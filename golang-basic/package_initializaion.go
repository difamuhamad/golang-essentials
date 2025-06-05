package main

import (
	"fmt"
	"golang-basic/database"
	_ "golang-basic/internal" // blank identifier
)

func main() {
	fmt.Println(database.GetDatabase())
}

// Blank Indentifier will tell golang to run the package instead of using it
