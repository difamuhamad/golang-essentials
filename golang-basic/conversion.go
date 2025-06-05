package main

import "fmt"

func main() {
	var nilai32 = 32768
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai64)

	var name = "Difa"
	var d = name[0]
	var dString = string(d)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(name)
	fmt.Println(d)
	fmt.Println(dString)

}
