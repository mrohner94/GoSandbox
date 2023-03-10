package utils

import "fmt"

// Add adds together multiple numbers
func Add(nums ...int) int {
	total := 0

	for _, v := range nums {
		printNum(v)
		total += v
	}

	return total
}

func printNum(num int) {
	fmt.Println("Current Number: ", num)
}
