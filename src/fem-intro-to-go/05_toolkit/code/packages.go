package main

import (
	"fmt"

	"fem-intro-to-go/05_toolkit/code/utils"
)

func calculateImportantData() int {
	return utils.Add(1, 2, 3)
}

func main() {
	fmt.Println("Packages!")
	// utils.Add(1, 2, 3)

	total := calculateImportantData()

	fmt.Println(total)
}
