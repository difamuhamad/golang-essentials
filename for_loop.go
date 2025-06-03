package main

import "fmt"

func main() {

	for water := 0; water <= 3; water++ {
		fmt.Println("Filling the cup...")

	}
	fmt.Println("Cup is filled")

	cats := []string{"black", "white", "yellow", "orange"}
	for i := 0; i < len(cats); i++ {
		fmt.Println(cats[i], "cat")
	}

	// ---------------------------- For Range

	// for ( index/key, value := range slice/map )
	for index, cat := range cats {
		fmt.Println("index:", index, cat, "cat")
	}

	for _, cat := range cats {
		fmt.Println(cat, "cat")
	}

	// ---------------------------- Break and Continue
	for index := 0; index < 10; index++ {
		if index == 5 {
			fmt.Println(index, "Enough")
			break
		}
	}

	for index := 0; index < 10; index++ {
		if index%2 != 0 {
			fmt.Println(index, "Odd numbers (Bilangan ganjil)")
			continue
		}
	}

	for index := 0; index < 10; index++ {
		if index%2 == 0 {
			fmt.Println(index, "Even numbers (Bilangan genap)")
			continue
		}
	}

	// Behind the scenes
	var index int = 0

	for index <= 10 {
		fmt.Println(index, " Index is increasing...")
		index++
	}
}
