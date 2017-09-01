package main

import (
    "bytes"
    "main/cloudstack"
    "log"
    "fmt"
    "encoding/xml"
    "io/ioutil"
)

type Conf struct {
    Apikey      string `xml:"apikey"`
    Secretkey        string `xml:"secretkey"`
    Ipaddress   string `xml:"ipaddress"`
    Port       string `xml:"port"`
}

func main() {
    //从配置文件中读取配置信息
    b, err := ioutil.ReadFile("conf.xml")
    if err != nil{
        fmt.Print(err)
    }

    //将二进制信息转化为string
    data := string(b)

    //读取并解析xml
    buf := bytes.NewBufferString(data)
    conf := new(Conf)
    decoded := xml.NewDecoder(buf)
    err1 := decoded.Decode(conf)
    if err1 != nil {
        fmt.Printf("Error: %v\n", err1)
    }

    //将读取出的配置信息进行赋值
    apikey := conf.Apikey
    secretkey := conf.Secretkey
    ipaddress := conf.Ipaddress
    port := conf.Port

    //生成函数变量
    cloudstack_url := "http://" + ipaddress + ":" + port + "/client/api"

    // Create a new API client
    cs := cloudstack.NewAsyncClient(cloudstack_url, apikey, secretkey, false)

    // Create a new parameter struct
    p := cs.VirtualMachine.NewListVirtualMachinesParams()

    // List all virtual machine
    r, err := cs.VirtualMachine.ListVirtualMachines(p)
    if err != nil {
        log.Fatalf("Error listing the virtual machines: %s ", err)
    }

    for _, vm := range r.VirtualMachines{
        cs1 := cloudstack.NewAsyncClient(cloudstack_url, apikey, secretkey, false)

        // Create a new parameter struct
        p := cs1.VirtualMachine.NewStopVirtualMachineParams(vm.Id)

        // Stop virtual machine
        r1, err1 := cs1.VirtualMachine.StopVirtualMachine(p)
        if err1 != nil {
            log.Fatalf("Error when stopping instance: %s", err1)
        }

        fmt.Printf("UUID of the stopped machine: %s", r1.Id)
    }
}
