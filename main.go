package main

import (
	"fmt"
	"github.com/LucasCarioca/pi-net-dht/pkg/services"
	"github.com/d2r2/go-dht"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	setupCloseHandler()
	pinetCollectorHost := "0.0.0.0"
	if len(os.Args) > 1 {
		pinetCollectorHost = os.Args[1]
	}

	node := "unknown"
	if len(os.Args) > 2 {
		node = os.Args[2]
	}

	location := "unknown"
	if len(os.Args) > 3 {
		location = os.Args[3]
	}

	sensor := dht.DHT22
	if len(os.Args) > 4 {
		if os.Args[4] == "dht11" {
			sensor = dht.DHT11
		} else {
			fmt.Println("...using dht22 by default...")
		}
	}
	pin := 4
	if len(os.Args) > 5 {
		newPin, err := strconv.Atoi(os.Args[5])
		if err != nil {
			fmt.Println("...using pin 4 as default...")
		} else {
			pin = newPin
		}
	}

	mockDHT := true
	if len(os.Args) > 6 {
		newMockState, err := strconv.ParseBool(os.Args[6])
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
			collectorService.SendClimateRecord(pinetCollectorHost, *temperature, *humidity, node, location)
		}
		time.Sleep(1 * time.Minute)
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
