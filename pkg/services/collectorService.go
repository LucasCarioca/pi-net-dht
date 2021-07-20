package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CollectorService struct {}

func (s *CollectorService) SendClimateRecord(pinetCollectorHost string, temperature float32, humidity float32, node string, location string) error {
	values := map[string]string{
		"temperature": fmt.Sprintf("%v*C", temperature),
		"humidity":    fmt.Sprintf("%v%%", humidity),
		"node":        node,
		"location":    location,
	}
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s/api/v1/climate-records", pinetCollectorHost)
	_, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++RECORD+SENT++++++++++++++++++++++++++++++++++++++++")
	fmt.Printf("Temperature = %v*C, Humidity = %v%%\n", temperature, humidity)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	return nil
}