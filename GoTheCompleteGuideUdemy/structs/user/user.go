package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) outputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthdate)
}

func (u *User) clearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name, and birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
