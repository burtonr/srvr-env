package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tarm/serial"
)

var (
	serialRegex = `(\[.*?\])(?:\s+)(.+)` // Expect output from serial (Arduino) to be "[the sensor key] some_value"
	// tempSensors = map[string]string{     // Sensor address/name | installed location
	// 	"4068167721416066": "Floor",
	// 	"DHT22":            "Center",
	// }
	// humiditySensors = map[string]string{
	// 	"DHT22": "SOMETHING",
	// }
	// tempGauges = promauto.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Namespace: "BurtonR",
	// 		Subsystem: "ServerRoom",
	// 		Name:      "Temperature",
	// 		Help:      "Temperature measurements",
	// 	},
	// 	[]string{
	// 		"name",
	// 		"location",
	// 	},
	// )
	// humidGauges = promauto.NewGaugeVex(
	// 	prometheus.GaugeOpts{
	// 		Namespace: "BurtonR",
	// 		Subsystem: "ServerRoom",
	// 		Name:      "Humidity",
	// 		Help:      "Humidity measurements",
	// 	},
	// 	[]string{
	// 		"name",
	// 		"location",
	// 	},
	// )
)

func main() {
	// TODO: Find, rather than hard-code, the device location
	config := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	go readSerial(s)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8888", nil)
}

// func recordMetric(key, value string) {

// 	newTemp, tErr := strconv.ParseFloat(value, 64)
// 	if tErr != nil {
// 		fmt.Printf("Temperature for [%s] cannot be parsed: %s\n", key, value)
// 		return
// 	}

// 	location, ok := tempSensors[key]
// 	if !ok {
// 		location = "Undefined"
// 	}

// 	tempGauges.WithLabelValues(key, location).Set(newTemp)
// }

func readSerial(s *serial.Port) {
	buf := make([]byte, 128)
	var message string
	var key string
	var val string

	re := regexp.MustCompile(serialRegex)

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		message += string(buf[:n])
		message = strings.Replace(message, "\r", "", -1)
		message = strings.Replace(message, "\n", "", -1)
		results := re.FindAllString(message, -1)

		if len(results) == 1 {
			kv := strings.Split(results[0], " ")
			for _, v := range kv {
				keyVal := strings.TrimSpace(v)
				if keyVal == "" {
					continue
				}

				if strings.HasPrefix(keyVal, "[") {
					v1 := strings.Replace(v, "[", "", 1)
					key = strings.Replace(v1, "]", "", 1)
				} else {
					val = v
				}
			}
		}

		recordMetric(key, val)
		message = ""
	}
}

// TODO: Not yet verified/implemented
func writeSerial(s *serial.Port, input <-chan string) {
	fmt.Println("Writing out to the serial...")
	for {
		message := <-input
		message += "\n"
		fmt.Println(message)
		n, err := s.Write([]byte(message))
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
		fmt.Printf(string(n))
	}

}
