//go mod init pico/dht11
//tinygo flash -target=pico -monitor main.go
//tinygo monitor -baudrate=9600
package main

import (
	"fmt"
	"machine"
	"time"
	"tinygo.org/x/drivers/dht"
)

func main() {
	pin := machine.GP28
	dhtSensor := dht.New(pin, dht.DHT11)
	uart := machine.UART1
	uart.Configure(machine.UARTConfig{
		BaudRate: 2400,
		TX: machine.UART1_TX_PIN,
		RX: machine.UART1_RX_PIN,
	})

	for {
		temp, hum, err := dhtSensor.Measurements()
		if err == nil {
			fmt.Printf("Temperature: %02d.%d°C, Humidity: %02d.%d%%\n", temp/10, temp%10, hum/10, hum%10)
		} else {
			fmt.Printf("Could not take measurements from the sensor: %s\n", err.Error())
		}
		uart.Write([]byte("some_data"))
		// Measurements cannot be updated only 2 seconds. More frequent calls will return the same value
		time.Sleep(time.Second * 2)
	}
}