package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (animal Dog) Speak() string {
	return "Woof"
}

func (animal Cat) Speak() string {
	return "Meoww"
}

func saySomething(creature Speaker) {
	fmt.Println(creature.Speak())
}

// ---------------------------- Empty Interface / any
func returnAnything() interface{} {
	// return "string"
	// return true
	return 2025
}

func printAnything(value interface{}) {
	fmt.Println("Value:", value)
}

func printAny(value any) {
	fmt.Println("Value:", value)
}

func main() {

	cat := Cat{}
	dog := Dog{}

	saySomething(cat)
	saySomething(dog)

	printAnything("string")
	printAnything(2050)
	printAnything(false)

	printAny("any string")
	printAny(true)

	fmt.Println(returnAnything())
}

/*
------- Differece between any and empty interface

interface{} and any are functionally identical in Go — both represent a value of any type.
The difference is that any is a type alias introduced in Go 1.18 to improve readability and support generics.
While interface{} is the traditional form, any is more modern and concise, making code easier to understand,
especially in generic contexts.

In Go, any is a predeclared identifier introduced as a type alias for interface{}.
Semantically, both denote a type that can hold values of any dynamic type.
While interface{} is the fundamental empty interface type,
any exists purely for syntactic clarity and is primarily used to improve type readability in generic type parameter constraints.

empty interface example (built in) :

- fmt.Println(a...interface{})
- panic(v interface{})
- recover()interface{}
*/
