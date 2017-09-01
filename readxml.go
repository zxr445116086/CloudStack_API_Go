package main

import (
    "bytes"
    "encoding/xml"
    "fmt"
    "io/ioutil"
)


type Conf struct {
    Apikey      string `xml:"apikey"`
    Secretkey        string `xml:"secretkey"`
    Ipaddress   string `xml:"ipaddress"`
    Port       string `xml:"port"`
}


func main() {
    b, err := ioutil.ReadFile("conf.xml")
    if err != nil{
        fmt.Print(err)
    }

    data := string(b)

    buf := bytes.NewBufferString(data)

    conf := new(Conf)

    decoded := xml.NewDecoder(buf)

    err1 := decoded.Decode(conf)
    if err1 != nil {
        fmt.Printf("Error: %v\n", err1)
    }

    apikey := conf.Apikey
    secretkey := conf.Secretkey
    ipaddress := conf.Ipaddress
    port := conf.Port

    fmt.Printf("apikey: %s\n", apikey)
    fmt.Printf("secretkey: %s\n", secretkey)
    fmt.Printf("ipaddress: %s\n", ipaddress)
    fmt.Printf("port: %s\n", port)


}
