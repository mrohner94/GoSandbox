package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeGroup(g Group) string {
	desc := fmt.Sprintf("This user group has %d users.  The newest user is %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName+" "+g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Kurt", LastName: "Cobain", Email: "kcobain@gmail.com"}

	x := u

	g := Group{
		role:           "Whatever",
		users:          []User{u, u2},
		newestUser:     u2,
		spaceAvailable: true,
	}

	x.ID = 123
	fmt.Println(u)
	fmt.Println(x)

	userDescription := describeUser(u)
	groupDescription := describeGroup(g)

	fmt.Println(userDescription)
	fmt.Println(groupDescription)
}
