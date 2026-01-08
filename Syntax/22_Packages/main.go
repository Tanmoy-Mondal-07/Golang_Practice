package main

import (
	"fmt"

	"github.com/aa/aa/auth"
	"github.com/aa/aa/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("codersgyan", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
