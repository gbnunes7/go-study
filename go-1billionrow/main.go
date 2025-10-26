package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Min  float64
	Max float64
	Name string
	Sum float64
	Count int
}

func main() {
	measurements, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close()
	scanner := bufio.NewScanner(measurements)
	dados := make(map[string]Measurement)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData,";")

		location := rawData[:semicolon]
		rawTemp := rawData[semicolon + 1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := dados[location]

		if !ok {
			measurement.Name = location
			measurement.Max= temp
			measurement.Min= temp
			measurement.Sum = temp
			measurement.Count= 1
			
		}
		measurement.Name = location
		measurement.Max = max(measurement.Max, temp)
		measurement.Min = min(measurement.Min, temp)
		measurement.Sum += temp
		measurement.Count++

		dados[location] = measurement
	}

	for name, measurement := range dados {
		fmt.Printf("%s: %#+v\n", name, measurement)
	}
}