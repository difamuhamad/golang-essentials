package helper

import "fmt"

func SayHello() {
	fmt.Println("This is Hello World from Helper!!")
}

func SumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

var private string = "private variable"

// ---------------------------- Access Modifier
var Public string = "Public Variable"
var Number int = 2025
var TrueValue bool = true

// Package name is following the folder name
// Function must use PascalCase to export itself
// private variable cannot be accessed outside thes package
