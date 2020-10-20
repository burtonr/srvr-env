package main

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type sensorType string

const (
	Temperature sensorType = "Temperature"
	Humidity sensorType = "Humidity"
	Status sensorType = "Status"
)

type sensor struct {
	Name     string
	Type     sensorType
	Location string
}

var (
	metricNS    = "burtonr"
	metricSS    = "server_room"
	sensors		= map[string]sensor{
		"DHT22": { "DHT22", Humidity, "mount"},
		"DOOR": {"Door", Status, "door"},
		"4023624112116213133": {"4023624112116213133", Temperature, "vent"},
		"40751127214160105": {"40751127214160105", Temperature, "top_rack"},
		"402441612116213104": {"402441612116213104", Temperature, "floor"},
		"4068167721416066": {"4068167721416066", Temperature, "desk"},
		"40111737214160174": {"40111737214160174", Temperature, ""},

	}
	tempGauges    = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNS,
			Subsystem: metricSS,
			Name:      "temperature",
			Help:      "Temperature measurements",
		},
		[]string{
			"name",
			"location",
		},
	)
	humidGauges = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNS,
			Subsystem: metricSS,
			Name:      "humidity",
			Help:      "Humidity measurements",
		},
		[]string{
			"name",
			"location",
		},
	)
	statusGauges = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: metricNS,
			Subsystem: metricSS,
			Name:      "status",
			Help:      "Status Sensors on, off, high, low, etc",
		},
		[]string{
			"name",
		},
	)
)

func recordMetric(key, value string) {
	sensor, found := sensors[key]
	if !found {
		fmt.Printf("Sensor [%s] received, but not found in the available sensors!\n", key)
	}

	switch sensor.Type {
	case Temperature:
		recordTemp(sensor, value)
	case Humidity:
		recordHumidity(sensor, value)
	case Status:
		recordStatus(sensor, value)
	}
}

func recordTemp(sensor sensor, temp string) {
	newTemp, tErr := strconv.ParseFloat(temp, 64)
	if tErr != nil {
		fmt.Printf("Temperature for [%s] cannot be parsed: %s\n", sensor.Name, temp)
		return
	}

	tempGauges.WithLabelValues(sensor.Name, sensor.Location).Set(newTemp)
}

func recordHumidity(sensor sensor, reading string) {
	newReading, tErr := strconv.ParseFloat(reading, 64)
	if tErr != nil {
		fmt.Printf("Humidity for [%s] cannot be parsed: %s\n", sensor.Name, reading)
		return
	}

	humidGauges.WithLabelValues(sensor.Name, sensor.Location).Set(newReading)
}

func translateStatus(status string) float64 {
	switch status {
	case "OPEN", "ON":
		return 1
	case "CLOSED", "OFF":
		return 0
	}
	return 0
}

func recordStatus(sensor sensor, status string) {
	newStatus := translateStatus(status)
	statusGauges.WithLabelValues(sensor.Name).Set(newStatus)
}
