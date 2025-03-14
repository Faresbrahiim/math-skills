package math

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func ExportData(s string) []float64 {
	file, err := os.Open("data.txt")
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
