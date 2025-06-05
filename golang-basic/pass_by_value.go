package main

import "fmt"

type Home struct {
	City, Province, Country string
}

func main() {

	student := Home{"Bandung", "Jawa Barat", "Indonesia"}
	student2 := student //copy value

	fmt.Println(student)
	fmt.Println(student2)

	student2.City = "Jonggol"
	fmt.Println(student)
	fmt.Println(student2)

}

/**

---------- Pass by Value :

When you call a function and pass arguments, Go sends a copy of the values, not the original variables.
This means any changes made to parameters inside the function do not affect the original variables outside the function.
*/
