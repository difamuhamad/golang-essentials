package main

import "fmt"

type Location struct {
	City, Province, Country string
}

func main() {

	var adress Location = Location{"Sydney", "New South Wales", "Australia"}
	var adress2 *Location = &adress

	adress2.City = "Newcastle"
	fmt.Println("1", adress)
	fmt.Println("2", adress2)

	// adress2 = &Location{"Melbourne", "Victoria", "Australia"}

	// change pointered value (address)
	*adress2 = Location{"Melbourne", "Victoria", "Australia"}

	fmt.Println("1", adress)
	fmt.Println("2", adress2)

}

/**

---------- Pointer :

A pointer is a variable that stores the memory address of another variable.
A pointer points to the location in memory where a value is stored — instead of holding the value itself.
*/
