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

	outputUserDetails(appUser)

	// ... do something awesome with that gathered data!

	fmt.Println(userFirstName, userLastName, userBirthdate)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
