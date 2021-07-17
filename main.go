package main

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/d2r2/go-dht"
	"fmt"
	logger "github.com/d2r2/go-logger"
)

var lg = logger.NewPackageLogger("main",
	logger.FatalLevel,
)


func main() {
	for {
		temperature, humidity, retried, _ :=
			dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
		values := map[string]string{
			"temperature": fmt.Sprintf("%v*C", temperature),
			"humidity": fmt.Sprintf("%v%%", humidity),
			"node": "piz02",
			"location": "office",
		}
		data, err := json.Marshal(values)
		if err != nil {
			fmt.Println("f'dup")
		} else {
			_, err := http.Post("http://192.168.1.211/api/v1/climate-records", "application/json", bytes.NewBuffer(data))
			if err != nil {
				fmt.Printf("f'dup")
			} else {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
				fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n", temperature, humidity, retried)
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			}
		}
	}
}
