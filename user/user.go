package user

import (
	"errors"
	"fmt"
	"time"
)

// User - To export user struct and their access fields, we need to upper case the first letter
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// OutputUserDetails - in Go we don't create methods inside struct, but we ref the method to struct by using a 'Receiver'.
// Receiver must be after func keyword but before the method name.
// We link these two pieces together by a 'Receiver'
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// ClearUserFullName - in Go every thing is passed as copy, so in this case we need to use a pointer to guaranty that the same object pointer is passed here
func (u *User) ClearUserFullName() {
	u.firstName = ""
	u.lastName = ""
}

// NewUser this is a utility method, a convention or pattern it's not a Go feature to creation structs
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	// adding validation for utility struct constructor
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
