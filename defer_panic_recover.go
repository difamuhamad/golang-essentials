package main

import "fmt"

func endTransaction() {
	fmt.Println("Transaction closed")
}

func runApplication(success bool) {
	defer endTransaction()

	fmt.Println("Application is running...")

	if !success {
		panic("error happened")
	}
}

func main() {

	runApplication(false)
}
