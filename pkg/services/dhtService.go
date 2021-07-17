package services

import (
	"github.com/d2r2/go-dht"
)

type DHTServiceInterface interface {
	Read(sensorType dht.SensorType, pin int) (*float32, *float32, error)
}

type DHTService struct{}

func (s *DHTService) Read(sensorType dht.SensorType, pin int) (*float32, *float32, error) {
	temperature, humidity, _, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	return &temperature, &humidity, err
}

type DHTServiceMock struct {
	mockTemperature float32
	mockHumidity    float32
	mockError       error
}

func (s *DHTServiceMock) SetMock(mockTemperature float32, mockHumidity float32, mockError error) {
	s.mockTemperature = mockTemperature
	s.mockHumidity = mockHumidity
	s.mockError = mockError
}

func (s *DHTServiceMock) Read(sensorType dht.SensorType, pin int) (*float32, *float32, error) {
	return &s.mockTemperature, &s.mockHumidity, s.mockError
}

func NewDHTServiceMock(active bool) DHTServiceInterface {
	if active {
		return &DHTService{}
	} else {
		return &DHTServiceMock{
			mockTemperature: 25,
			mockHumidity:    50,
			mockError:       nil,
		}
	}
}
