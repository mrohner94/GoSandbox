package main

import "fmt"

func average(inputs ...float64) float64 {
	var sum float64 = 0
	var length int = len(inputs)
	for _, value := range inputs {
		sum += value
	}

	return (sum / float64(length))
}

func main() {
	var averageOfNums float64 = average(1, 0, 0, 0)

	fmt.Println(averageOfNums)

}
