package main

import "fmt"

func main() {
	student := map[string]string{
		"name":    "Difa",
		"address": "Bandung",
	}

	fmt.Println(student["name"])
	fmt.Println(student["address"])
	fmt.Println(student)

	// the hard way
	var mahasiswa map[string]string = map[string]string{}

	mahasiswa["name"] = "Difa"
	mahasiswa["address"] = "Bandung"

}
