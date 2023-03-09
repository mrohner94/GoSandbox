package main

import "fmt"

func main() {

	var myArray [5]int
	var mySlice []int = make([]int, 5, 10)

	myArray[0] = 1
	mySlice[0] = 1

	fmt.Println(myArray)
	fmt.Println(mySlice)

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}
	splicedFruit := fruitArray[1:3]

	fmt.Println(fruitArray)
	fmt.Println(len(fruitArray))
	fmt.Println(cap(fruitArray))

	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit))
	fmt.Println(cap(splicedFruit))

	fruitToAdd := append(splicedFruit, "canteloupe", "cherries", "lemon")

	fmt.Println(fruitToAdd)
	fmt.Println(len(fruitToAdd))
	fmt.Println(cap(fruitToAdd))

	// ***************************

	// var myArray [5]int
	// var mySlice []int = make([]int, 5)

	// fmt.Println(myArray)
	// fmt.Println(mySlice)

	// ***************************

	// var myArray [5]int
	// // var mySlice []int = make([]int, 5)
	// var mySlice []int = make([]int, 5, 10)
	// // var mySlice = make([]int, 5, 10)

	// myArray[0] = 1
	// mySlice[0] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	// ***************************

	// fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}

	// var splicedFruit []string = fruitArray[1:3] // ==> ["pear", "apple",]

	// fmt.Println(len(splicedFruit))
	// fmt.Println(cap(splicedFruit))

	// ***************************

	// SEE SLIDE

	// ***************************

	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)

	// fmt.Println(slice1, slice2)
	// fmt.Println(len(slice1), cap(slice1))
	// fmt.Println(len(slice2), cap(slice2))

	// ***************************

	// originalSlice := []int{1, 2, 3}
	// destination := make([]int, len(originalSlice))

	// fmt.Println("Before Copy:", originalSlice, destination)

	// mysteryValue := copy(destination, originalSlice)

	// // fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
}
