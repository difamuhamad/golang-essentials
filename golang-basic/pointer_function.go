package main

import "fmt"

type Address struct {
	Country string
}

func moveAdressToIndonesia(person *Address) {
	person.Country = "Indonesia"
}

func moveAdressToAussie(person *Address) {
	person.Country = "Australia"
}

func main() {

	// var person *Adress = &Address{}
	// moveAdressToIndonesia(person)

	var person Address = Address{"Australia"}
	moveAdressToIndonesia(&person)

	var person2 *Address = &person
	moveAdressToAussie(person2)

	fmt.Println(person)
	fmt.Println(person2)
}
