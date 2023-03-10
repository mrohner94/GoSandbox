// package main

// import "fmt"

// func getAvg(input [5]float64) float64 {

// 	sum := 0.0
// 	length := len(input)
// 	for _, value := range input {
// 		sum += value
// 	}

// 	return sum / float64(length)
// }

// func checkPetExists(petDictionary map[string]string, name string) bool {
// 	_, ok := petDictionary[name]

// 	return ok
// }

// var initialGroceries = []string{"eggs", "milk", "bananas", "chicken", "flour"}

// func addGroceryToList(newGroceries ...string) []string {
// 	foods := initialGroceries
// 	for _, g := range newGroceries {
// 		foods = append(foods, g)
// 	}

// 	return foods
// }

// func main() {
// 	//Part 1
// 	fmt.Println("Part 1:")

// 	scores := [5]float64{100, 80.1, 33, 12, 100}
// 	fmt.Println(scores)
// 	fmt.Println(getAvg(scores))

// 	//Part 2
// 	fmt.Println("Part 2:")

// 	petDictionary := map[string]string{"bear": "cat", "rocky": "dog", "charlotte": "spider"}
// 	fmt.Println(checkPetExists(petDictionary, "bear"))
// 	fmt.Println(checkPetExists(petDictionary, "craig"))

// 	//Part 3
// 	fmt.Println("Part 3:")

// 	fmt.Println(addGroceryToList("beans", "kale"))
// }
