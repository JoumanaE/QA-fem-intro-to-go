package main

import "fmt"

//User ...
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User{
	ID:        1,
	FirstName: "Hajja",
	LastName:  "Hamdaouia",
	Email:     "hajja@gmail.com",
}

func updateEmail(u *User) {
	u.Email = "newEmail@gmail.com"
	fmt.Println("in update email: ", u.Email)
}

func main() {
	fmt.Println("Pointers")
	updateEmail(&u)
	fmt.Println("Updated User:", u)
}
