package main

import "fmt"

func main() {
	// ---------------------------- String
	var name string = "Difa"
	var backname = "Muhamad"
	fullName := "Difa Muhamad"

	// var value can be changed
	fullName = "Muhamad Difa"

	var (
		firstName string = "Difa"
		lastName         = "Muhamad"
	)

	// This is the "hard" way, we explicit the data type
	var zodiak string
	zodiak = "pisces"

	fmt.Println(name)
	fmt.Println(backname)
	fmt.Println(fullName)
	fmt.Println(zodiak)
	fmt.Println(firstName, lastName)

	// error

	// override the var with ":=" will make error
	// fullName := "afid damahum"

	// Cannot override string type to number
	// fullName = 2006

	// ---------------------------- Integer
	var number int64 = 123
	fmt.Println(number)

	number2 := 2525252
	fmt.Println(number2)

	// override the existing value
	number2 = 21111111
	fmt.Println(number2)

	// ---------------------------- Float
	var floatNumber float64 = 123.123

	floatNumber2 := 321.321

	fmt.Println(floatNumber)
	fmt.Println(floatNumber2)

	// ---------------------------- Unsigned Integer
	var unsignedNumber uint64 = 5000

	// Negative variable will make it overflow
	// var unsignedNumber uint64 = -5000

	// Unsigned must be explicit for this declaration
	// unsignedNumber := 5000 (error)

	unsignedNumber2 := uint64(5000)

	fmt.Println(unsignedNumber)
	fmt.Println(unsignedNumber2)

	// ---------------------------- Bool / Boolean
	var fisalia bool = true
	var montelli bool = false

	ronascita := true

	// Will throw error if asigned with string
	// ronascita := "true"

	fmt.Println(fisalia)
	fmt.Println(montelli)
	fmt.Println(ronascita)
}
