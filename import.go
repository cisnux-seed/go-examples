package main

import (
	"errors"
	"fmt"
	"go-examples/helper"
	"go-examples/models"
)

func main() {
	result := helper.SayHello("Fajra")
	fmt.Println(result)

	// access modifier
	userBuilder := models.UserBuilder{}
	userBuilder.SetFullName("Fajra Risqulla")
	userBuilder.SetUsername("cisnux")
	userBuilder.SetEmailAddress("fajrarisqulla@gmail.com")
	userBuilder.SetAge(20)
	user, e := userBuilder.Build()
	if e != nil {
		//if validationErr, ok := e.(*helper.ValidationFailure); ok {
		//	fmt.Println("Validation Error:", validationErr.Message)
		//} else if notFoundErr, ok := e.(*helper.NotFoundFailure); ok {
		//	fmt.Println("Not Found Error:", notFoundErr.Message)
		//} else {
		//	fmt.Println("Unknown Error", e.Error())
		//}
		//switch err := e.(type) {
		//case *helper.ValidationFailure:
		//	fmt.Println("Validation Error:", err.Message)
		//case *helper.NotFoundFailure:
		//	fmt.Println("Not Found Error:", err.Message)
		//default:
		//	fmt.Println("Unknown Error", e.Error())
		//}
		var validationErr *helper.ValidationFailure
		var notFoundErr *helper.NotFoundFailure
		switch {
		case errors.As(e, &validationErr):
			fmt.Println("Validation Error:", validationErr.Message)
		case errors.As(e, &notFoundErr):
			fmt.Println("Not Found Error:", notFoundErr.Message)
		default:
			fmt.Println("Unknown Error", e.Error())
		}
	} else {
		fmt.Println("full name: ", user.FullName())
		fmt.Println("username: ", user.Username())
		fmt.Println("email address: ", user.EmailAddress())
		fmt.Println("age: ", user.Age())
	}
}
