package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"username": name,
		}
	}
}

func main() {
	fmt.Println(newMap("Budiono"))
	fmt.Println(newMap(""))

	person := newMap("")

	if person == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println("Data :", person)

	}
}

/**
------------- nil can only be used for :

Pointer			var p *User = nil
Interface		var i interface{} = nil
Slice			var s []int = nil
Map				var m map[string]int = nil
Channel			var c chan int = nil
Function		var f func() = nil
Error			var err error = nil





error :
func getName(name string) string {
	if name == "" {
		return nil
		} else {
		return name
	}
}

*/
