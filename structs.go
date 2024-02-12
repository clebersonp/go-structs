package main

import (
	"example.com/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your userBirthdate (MM/DD/YYYY): ")

	// Struct literal or composite literal (instantiate a struct)
	// pointer of user struct
	// the first lower case user is the package name and the last one is struct
	var appUser *user.User

	//appUser = user{} // user with all fields with 'null value' (default values)
	// short notation without specify the field names
	//appUser = user{
	//	userFirstName,
	//	userLastName,
	//	userBirthdate,
	//	time.Now(),
	//}

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	// if the args are not valid then return and stop the application
	if err != nil {
		fmt.Println(err)
		return
	}

	// this is a user method
	appUser.OutputUserDetails()
	appUser.ClearUserFullName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
