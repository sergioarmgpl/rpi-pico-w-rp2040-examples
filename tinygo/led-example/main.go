//go mod init pico/led
//tinygo flash -target=pico main.go
//tinygo monitor -baudrate=9600
package main

import (
    "machine"
    "time"
    "fmt"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        fmt.Print("Low")
        time.Sleep(time.Millisecond * 500)

        led.High()
        fmt.Print("Down")
        time.Sleep(time.Millisecond * 500)
    }
}