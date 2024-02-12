package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// in Go we don't create methods inside struct, but we ref the method to struct by using a 'Receiver'.
// Receiver must be after func keyword but before the method name.
// We link these two pieces together by a 'Receiver'
func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your userBirthdate (MM/DD/YYYY): ")

	// Struct literal or composite literal (instantiate a struct)
	var appUser user

	//appUser = user{} // user with all fields with 'null value' (default values)
	// short notation without specify the field names
	//appUser = user{
	//	userFirstName,
	//	userLastName,
	//	userBirthdate,
	//	time.Now(),
	//}

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// this is a user method
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
