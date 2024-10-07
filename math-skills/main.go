package main

import (
	"fmt"
	"math_skills"
)

func main() {
	data := math_skills.ExportData("data.txt")

	average := math_skills.Avarage(data)
	median := math_skills.Median(data)
	variance := math_skills.Variance(average, data)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
}
