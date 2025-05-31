package main

import (
	"fmt"
)

func main() {
	var students [3]string

	students[0] = "Ucup"
	students[1] = "Otong"
	students[2] = "Budiono"

	// easy way :
	var employes = [3]string{
		"Joko",
		"Waluh",
		"Franklin",
	}

	var emptyString = [5]string{
		"first",
		"second",
	}

	var emptyNumber = [5]int{
		1,
		2,
	}

	var waifus = [...]string{
		"Rem",
		"Ram",
		"Emilia",
		"Echidna",
		"Satella",
	}

	// The only way to delete value
	waifus[4] = ""

	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])
	fmt.Println(students)
	fmt.Println(employes)
	fmt.Println(len(employes))
	fmt.Println(emptyString)
	fmt.Println(emptyNumber)
	fmt.Println(waifus)

	// error
	// students[3] = "Kapallawd"
}
