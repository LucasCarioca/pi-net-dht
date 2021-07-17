package main

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-dht/pkg/services"
	"github.com/d2r2/go-dht"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	setupCloseHandler()

	sensor := dht.DHT22
	if len(os.Args) > 1 {
		if os.Args[1] == "dht11" {
			sensor = dht.DHT11
		} else {
			fmt.Println("...using dht22 by default...")
		}
	}
	pin := 4
	if len(os.Args) > 2 {
		newPin, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("...using pin 4 as default...")
		} else {
			pin = newPin
		}
	}

	mockDHT := true
	if len(os.Args) > 3 {
		newMockState, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			fmt.Println("...using real dht service by default...")
		} else {
			mockDHT = newMockState
		}
	}

	dhtService := services.NewDHTServiceMock(mockDHT)
	collectorService := services.CollectorService{}

	for {
		temperature, humidity, err := dhtService.Read(sensor, pin)
		if err != nil {
			fmt.Println("Failed to read sensor")
		} else {
			collectorService.SendClimateRecord(*temperature, *humidity)
		}
	}
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
