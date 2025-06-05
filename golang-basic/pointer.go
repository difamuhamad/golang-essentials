package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	student := Address{"Bandung", "Jawa Barat", "Indonesia"}
	student2 := &student //pointer

	student2.City = "Jonggol"

	var adress Address = Address{"Sydney", "New South Wales", "Australia"}
	var adress2 *Address = &adress

	adress2.City = "Newcastle"

	fmt.Println(student)
	fmt.Println(student2)
	fmt.Println("1", adress)
	fmt.Println("2", adress2)

	// Pointer to new value
	adress2 = &Address{"Melbourne", "Victoria", "Australia"}

	fmt.Println("1", adress)
	fmt.Println("2", adress2)

}

/**

---------- Pointer :

A pointer is a variable that stores the memory address of another variable.
A pointer points to the location in memory where a value is stored — instead of holding the value itself.
*/
