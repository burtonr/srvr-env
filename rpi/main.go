package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/tarm/serial"
)

var tempRegex = `(\[TMP[0-9]\])\s([0-9]{2,3}.[0-9]{2,3})`
var allRegex = `.+`

func main() {
	// TODO: Find, rather than hard-code, the device location
	config := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	go readTemp(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Stopping listener")
			os.Exit(1) // Exit on interrupt
		}
	}()

	// Run forever
	for {
	}
}

func readTemp(s *serial.Port) {
	buf := make([]byte, 128)
	var message string
	// re := regexp.MustCompile(tempRegex)
	re := regexp.MustCompile(allRegex)

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		message += string(buf[:n])
		message = strings.Replace(message, "\r", "", -1)
		message = strings.Replace(message, "\n", "", -1)

		results := re.FindStringSubmatch(message)

		if 0 < len(results) {
			fmt.Println(results[0])
			// TODO: Push Prometheus metric (gauge)
			message = ""
		}
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
