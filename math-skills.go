package math_skills

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Avarage(data []float64) float64 {
	sum := 0.0
	numbers := len(data)
	for _, r := range data {
		sum += r
	}
	Avarage := sum / float64(numbers)
	return Avarage
}

func Median(data []float64) float64 {
	sort.Float64s(data)
	numbers := len(data)
	if len(data)%2 == 1 {
		return data[numbers/2]
	}
	return (data[(numbers/2-1)] + data[numbers/2]) / 2
}

func Variance(avr float64, data []float64) float64 {
	avrValue := Avarage(data)
	sumV := 0.0
	numbers := float64(len(data))
	for _, r := range data {
		sumV += (r - avrValue) * (r - avrValue)
	}
	Variance := sumV / numbers
	return Variance
}

func ExportData(filename string) []float64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening the file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []float64

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error converting string to float:", err)
			continue
		}
		data = append(data, value)
	}
	return data
}
