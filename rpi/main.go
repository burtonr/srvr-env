package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tarm/serial"
)

var (
	serialRegex = `(\[.*?\])(?:\s+)(.+)` // Expect output from serial (Arduino) to be "[the sensor key] some_value"
	outChan     chan string              // TODO: Shouldn't need this with a proper type to contain all serial work
)

func main() {
	// TODO: Find, rather than hard-code, the device location
	config := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	// TODO: Move this to an 'init' func on a proper type to contain the serial work
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	outCh := make(chan string)
	outChan = outCh

	// TODO: Better to move the serial I/O to a struct to avoid passing around constants
	go readSerial(s)
	go writeSerial(s, outCh)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/serial", serialHandler)
	http.ListenAndServe(":8888", nil)
}

func serialHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// TODO: Parse body into struct and pass to lower function
	fmt.Println("Handling the serial request...")
	fmt.Printf("Received: %s\n", body)
	// TODO: Create const of the commands to avoid "magic"
	outChan <- "cmd:POWER_ON"
}

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

// DEV: Sample this was built from -> https://play.golang.org/p/zk_03P071w
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
