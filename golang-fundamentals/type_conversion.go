package main

import "fmt"

func main() {
	type Animal string
	type NoHP uint64

	var cat Animal = "Garfield"
	var cow = "Moooo"
	var telkomsel NoHP = 6282217258285

	var typedCow Animal = Animal(cow)

	fmt.Println(cat)
	fmt.Println(telkomsel)
	fmt.Println(typedCow)
}
