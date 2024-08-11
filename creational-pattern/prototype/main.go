package main

import (
	"fmt"
)

// UserPrototype defines the interface for the prototype pattern
type UserPrototype interface {
	Clone() UserPrototype
}

// User is the concrete prototype
type User struct {
	Name  string
	Email string
}

// Clone method creates a shallow copy of the User object
func (u *User) Clone() UserPrototype {
	return &User{
		Name:  u.Name,
		Email: u.Email,
	}
}

func main() {
	user1 := &User{
		Name:  "John Doe",
		Email: "test@mail.com",
	}

	user2 := user1.Clone().(*User)

	// Change clone without affecting original
	user2.Name = "Jane Doe"
	user2.Email = "test2@mail.com"

	fmt.Println(user1)
	fmt.Println(user2)
}
