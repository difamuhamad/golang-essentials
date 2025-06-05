package main

import "fmt"

func random() any {
	// return 2025
	// return false
	// return true
	return "string value"
}

func main() {
	var randomValue any = random()
	var stringValue string = randomValue.(string)
	// var intValue string = randomValue.(int)

	switch value := randomValue.(type) {

	case string:
		fmt.Println("value is string")
	case int:
		fmt.Println("value is number")
	case bool:
		fmt.Println("value is boolean")
	default:
		fmt.Println(value, "is unknown")

	}

	fmt.Println(stringValue)
	// fmt.Println(intValue)
}
