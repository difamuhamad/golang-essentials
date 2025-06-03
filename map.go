package main

import "fmt"

func main() {
	student := map[string]string{
		"name":    "Difa",
		"address": "Bandung",
	}

	book := map[string]string{
		"title":     "Pulang",
		"published": "December 2025",
		"author":    "Generative AI",
	}
	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)

	fmt.Println(student["name"])
	fmt.Println(student["address"])
	fmt.Println(student["wrong"])
	fmt.Println(student)

	// the hard way
	var mahasiswa map[string]string = map[string]string{}

	mahasiswa["name"] = "Difa"
	mahasiswa["address"] = "Bandung"

}
