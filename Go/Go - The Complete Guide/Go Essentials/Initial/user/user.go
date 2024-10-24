package user

import (
	"time"
)

type User struct {
	firstName string
	lastName  string
	createdAt time.Time
}

func New(firstName string, lastName string) (User, error) {
	return User{
		firstName: firstName,
		lastName:  lastName,
		createdAt: time.Now(),
	}, nil
}

func (u *User) GetUserName() (firstName string) {
	return u.firstName
}
