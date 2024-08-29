package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Min float64
	Max float64
	Sum float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err!= nil {
		panic(err)
	}
	defer measurements.Close()

	data := make((map[string]Measurement))

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		rawTemp := rawData[semicolonIndex+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)
		
		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min:  temp,
				Max:  temp,
				Sum:  temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}
		data[location] = measurement
	}

	for location, measurement := range data {
		fmt.Printf("%s, %#+v\n", location, measurement)
	}
}

