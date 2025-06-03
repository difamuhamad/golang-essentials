package main

import "fmt"

func main() {
	username := "Otong"

	switch username {

	case "Budiono":
		fmt.Println("Welcome Budiono!")

	case "Ucup":
		fmt.Println("Welcome Ucup!")

	case "Otong":
		fmt.Println("Welcome Otong!")

	default:
		fmt.Println("Welcome!")

	}

	username = "Siregar"
	usernametype := len(username)

	switch {
	case usernametype > 10:
		fmt.Println("Very long")
	case usernametype > 5:
		fmt.Println("Medium")
	default:
		fmt.Println("Short")
	}

	// ---------------------------- Short Statement
	switch usernamechar := len(username); usernamechar > 5 {

	case true:
		fmt.Println("Name is very long!")

	case false:
		fmt.Println("Name is very short!")

	}
}
