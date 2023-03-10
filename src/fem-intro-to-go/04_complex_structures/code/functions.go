// package main

// import "fmt"

// func printAge(age1 int, age2 int) (ageOfSally int, ageOfBob int) {
// 	ageOfSally = age1
// 	ageOfBob = age2
// 	return
// }

// func printAges(ages ...int) bool {

// 	fmt.Printf("There are %d ages. Now we are going to loop randomly \n", len(ages))

// 	// for i := 0; i <= len(ages); i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	for _, value := range ages {
// 		fmt.Println(value)
// 	}
// 	return true
// }

// func printTwoNums(xxx int) (num1 int, num2 int) {
// 	return 123, 321
// }

// func main() {
// 	x, y := printAge(10, 21)
// 	fmt.Println(x)
// 	fmt.Println(y)

// 	fmt.Println(printAge(1110, 213132))

// 	fmt.Println(printTwoNums(1110213132))

// 	var someBool bool = printAges(1, 2, 3)

// 	fmt.Printf("The boolean value from that last function is %t", someBool)

// }
