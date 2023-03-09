package main

import (
	"fmt"
	"reflect"
)

var someone, sometwo = "xxx", "yyy"

func main() {
	var name string = "Whatever"
	var otherName = "Some String"
	var undeclaredString string
	var undeclaredInt int

	somedeclaration := "test" //shorthand declaration / inferred

	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))

	fmt.Println(otherName)
	fmt.Println(reflect.TypeOf(otherName))

	fmt.Println(undeclaredString)
	fmt.Println(undeclaredInt)

	fmt.Println(someone, sometwo)
	fmt.Println(somedeclaration)
}
