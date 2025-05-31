package main

import "fmt"

func main() {

	// constant will not return error when not used
	const name string = "Difa"
	const backname = "Muhamad"
	const fullName = "Difa Muhamad"

	const (
		firstName = "Difa"
		lastName  = "Muhamad"
	)

	fmt.Println(name)
	fmt.Println(backname)

	// error
	// fullName = "Muhamad Difa"
}
