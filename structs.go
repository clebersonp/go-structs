package main

import (
	"example.com/structs/user"
	"fmt"
)

// Create an alias for every built-in types
type str string

// Create method to alias custom type
func (text str) log() {
	fmt.Println(text)
}

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

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	// if the args are not valid then return and stop the application
	if err != nil {
		fmt.Println(err)
		return
	}

	// this is a user method
	appUser.OutputUserDetails()
	appUser.ClearUserFullName()
	appUser.OutputUserDetails()

	fmt.Println("----------------------------------------------")

	// construct Admin struct, like inheritance
	admin := user.NewAdmin("test@example.com", "test123")
	// because we are using anonymous User type of Admin struct, we can access all methods of User struct directly
	admin.OutputUserDetails()
	admin.ClearUserFullName()
	admin.OutputUserDetails()

	fmt.Println("----------------------------------------------")
	var someName str = "Tiago - custom type - str alias for string type"
	someName.log()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
