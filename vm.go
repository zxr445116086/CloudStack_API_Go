package main

import (
    "main/cloudstack"
    "log"
    "fmt"
)

func main() {
    // Create a new API client
    cs := cloudstack.NewAsyncClient("http://10.15.9.154:8080/client/api", "36YF1nADgAC8cRKpw3edWxsgwZogtsYLZj6zNOYXOBCbXP0EwPEniNo4sM67IUCltinQEclCnwtFub3T-2Rx8g", "1x0_ugn4TFW_VMdGTXPtkB7Lv8_RxxqdTIjGLeKGbQeg86zzwGDUzdJR8PJzvXwoZrcAsqKjZiVoDOHb5unYhA", false)

    // Create a new parameter struct
    p := cs.VirtualMachine.NewListVirtualMachinesParams()

    // Start new virtual machine
    r, err := cs.VirtualMachine.ListVirtualMachines(p)
    if err != nil {
        log.Fatalf("Error listing the virtual machines: %s ", err)
    }

    for _, vm := range r.VirtualMachines{
        cs1 := cloudstack.NewAsyncClient("http://10.15.9.154:8080/client/api", "36YF1nADgAC8cRKpw3edWxsgwZogtsYLZj6zNOYXOBCbXP0EwPEniNo4sM67IUCltinQEclCnwtFub3T-2Rx8g", "1x0_ugn4TFW_VMdGTXPtkB7Lv8_RxxqdTIjGLeKGbQeg86zzwGDUzdJR8PJzvXwoZrcAsqKjZiVoDOHb5unYhA", false)

        // Create a new parameter struct
        p := cs1.VirtualMachine.NewStartVirtualMachineParams(vm.Id)

        // Start new virtual machine
        r1, err1 := cs.VirtualMachine.StartVirtualMachine(p)
        if err1 != nil {
            log.Fatalf("Error creating the new instance: %s", err1)
        }

        fmt.Printf("UUID or the newly started machine: %s", r1.Id)
    }
}
