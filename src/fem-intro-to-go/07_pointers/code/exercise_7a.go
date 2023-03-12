package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User {ID: 1, FirstName: "Pete", LastName: "Doherty", Email: "something@email.go"}

func changeEmail(u *User) {
	u.Email = "NewEmail@test.com"
}

func main() {
	changeEmail(&u)

	fmt.Println(u)
}