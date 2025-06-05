package main

import "fmt"

func main() {
	user := "Admin"

	if user == "Admin" {
		fmt.Println("Welcome, how are you today?")
	} else if user == "Admin2" {
		fmt.Println("Welcome, how are you today?")
	} else {
		fmt.Println("SignUp first")
	}

	// ---------------------------- Short Statement

	if userchar := len(user); userchar > 10 {
		fmt.Println("Error, username is too long")
	} else {
		fmt.Println("Verified user")
	}

	// hard way (without short statement):

	userlength := len(user)
	if userlength > 10 {
		fmt.Println("Error")
	} else {
		fmt.Println("Verified user")
	}

}
