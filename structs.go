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
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// in Go every thing is passed as copy, so in this case we need to use a pointer to guaranty that the same object pointer is passed here
func (u *user) clearUserFullName() {
	u.firstName = ""
	u.lastName = ""
}

// this is a utility method, a convention or pattern it's not a Go feature to creation structs
func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your userBirthdate (MM/DD/YYYY): ")

	// Struct literal or composite literal (instantiate a struct)
	// pointer of user struct
	var appUser *user

	//appUser = user{} // user with all fields with 'null value' (default values)
	// short notation without specify the field names
	//appUser = user{
	//	userFirstName,
	//	userLastName,
	//	userBirthdate,
	//	time.Now(),
	//}

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	// this is a user method
	appUser.outputUserDetails()
	appUser.clearUserFullName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
