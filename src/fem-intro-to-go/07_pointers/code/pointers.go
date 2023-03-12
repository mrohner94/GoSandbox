package main

// func main() {
// 	var name string
// 	var namePointer *string

// 	name = "Kurt"
// 	namePointer = &name

// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n string) {
// 	n = strings.ToUpper(n)

// }

// func changeNameTwo (n *string) {
// 	*n = strings.ToUpper(*n)
// }

// func main() {
// 	name := "Elvis"
// 	nameTwo := "Mike"
// 	changeName(name)
// 	changeNameTwo(&nameTwo)

// 	fmt.Println(name)
// 	fmt.Println(nameTwo)
// }

// // ******************************************************

// type Coordinates struct {
// 	X, Y float64
// }

// var c = Coordinates{X: 10, Y: 20}

// //var z = Coordinates {X:1, Y: 2}

// func main() {
// 	coordinatesMemoryAddress := &c
// 	coordinatesMemoryAddress.X = 200
// 	coordinatesMemoryAddress.Y = 1234

// 	fmt.Println(coordinatesMemoryAddress)
// 	fmt.Println(c)
// }