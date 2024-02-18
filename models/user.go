package models

import (
	"go-examples/helper"
	"regexp"
)

type User struct {
	username     string
	emailAddress string
	fullName     string
	age          uint8
}

func (u User) Username() string {
	return u.username
}

func (u User) EmailAddress() string {
	return u.emailAddress
}

func (u User) FullName() string {
	return u.fullName
}

func (u User) Age() uint8 {
	return u.age
}

type UserBuilder struct {
	username     string
	emailAddress string
	fullName     string
	age          uint8
}

func (u *UserBuilder) SetUsername(username string) {
	u.username = username
}

func (u *UserBuilder) SetEmailAddress(emailAddress string) {
	u.emailAddress = emailAddress
}

func (u *UserBuilder) SetFullName(fullName string) {
	u.fullName = fullName
}

func (u *UserBuilder) SetAge(age uint8) {
	u.age = age
}

func (u *UserBuilder) Build() (usr User, er error) {
	matched, e := regexp.MatchString("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}$", u.emailAddress)

	switch {
	case e != nil:
		er = e
		return
	case u.username != "cisnux":
		// any struct that has method, which has a pointer receiver
		// we need to assign reference of struct to variable
		er = &helper.NotFoundFailure{
			Message: "Your account is not found",
		}
		return
	case !matched:
		// any struct that has method, which has a pointer receiver
		// we need to assign reference of struct to variable
		er = &helper.ValidationFailure{
			Message: "Your email address is not valid",
		}
		return
	case matched:
		usr = User{
			username:     u.username,
			emailAddress: u.emailAddress,
			fullName:     u.fullName,
			age:          u.age,
		}
		return
	default:
		return
	}
}
