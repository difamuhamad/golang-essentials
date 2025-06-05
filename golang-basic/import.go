package main

import (
	"fmt"
	// "project-name/package"
	"golang-basic/helper"
)

func main() {
	helper.SayHello()

	packageHelper := helper.SumAll(10, 10, 10)
	fmt.Println(packageHelper)

	fmt.Println(helper.Public)
	// fmt.Println(helper.private)
}

// Package Import must writted with project name
// Check the go.mod to see the project name

// "fmt" is a built-in package
