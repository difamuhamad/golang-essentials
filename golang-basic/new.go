package main

import "fmt"

type AlamatRumah struct {
	City string
}

func main() {
	// var alamat *AlamatRumah = &AlamatRumah{}

	var alamat *AlamatRumah = new(AlamatRumah)
	var alamat2 *AlamatRumah = alamat

	alamat.City = "Jakarta"

	fmt.Println(alamat)
	fmt.Println(alamat2)

}

/**
------------ new definition

new() used to allocate memory for a given type and return a pointer to it.
The allocated value is automatically initialized to its zero value.
*/
