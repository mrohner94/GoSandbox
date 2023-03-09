package main

import "fmt"

func printAge(age1 int, age2 int) (ageOfSally int, ageOfBob int) {
	ageOfSally = age1
	ageOfBob = age2
	return
}

func printTwoNums(xxx int) (num1 int, num2 int) {
	return 123, 321
}

func main() {
	x, y := printAge(10, 21)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(printAge(1110, 213132))

	fmt.Println(printTwoNums(1110213132))
}
