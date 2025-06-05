package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World!!")
}

func welcomeTo(firstName, lastName string) {
	fmt.Println("Welcome back", firstName, lastName)
}

func factorialLoop(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i

	}
	return result
}

// ---------------------------- Return Value
func sum(number1, number2 int) int {
	return number1 + number2
}

func sayHi() string {
	return "Hi, How are you today?"
}

// ---------------------------- Multiple Return Values

func getFood() (dish, drink, dessert string) {
	return "Chicken", "Fruit Salad", "Water"
}

// ---------------------------- Named return values
func getFullName() (firstName, middleName, lastName string) {
	firstName = "Budiono"
	middleName = "Siregar"
	lastName = "Kapallawd"
	// lastName := "Kapallawd"

	return firstName, middleName, lastName
}

// ---------------------------- Variadic Function
// the last parameter called 'varargs'
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// ---------------------------- Function as Parameter
// type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "bad word" {
		return "*****"
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(username string, blacklist Blacklist) {
	if blacklist(username) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Register success!")
	}
}

// ---------------------------- Recursive Function
func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialRecursive(number-1)
	}
}

func main() {
	sayHello()
	welcomeTo("Budiono", "Siregar")
	fmt.Println(sum(5, 4))
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)

	plate, dessert, cup := getFood()
	fmt.Println("The Dinner is", plate, dessert, "and", cup)

	plate, _, cup = getFood()
	fmt.Println("The Breakfest is", plate, "and", cup)

	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}

	// convert slice to varargs
	fmt.Println(sumAll(numbers...))

	// ---------------------------- Function Value
	salam := sayHi
	sapa := sayHi
	fmt.Println(salam())
	fmt.Println(sapa())

	filter := spamFilter
	sayHelloWithFilter("bad word", filter)
	sayHelloWithFilter("Ucup", spamFilter)

	// ---------------------------- Anonymous Function
	blacklist := func(username string) bool {
		switch username {
		case "hacker":
			return true
		case "tzy":
			return true
		case "bocil crypto":
			return true
		}
		return false
	}
	registerUser("Otong", blacklist)
	registerUser("hacker", blacklist)
	registerUser("Otong", func(username string) bool {
		return username == "bad word"
	})

	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))

	// ---------------------------- Closure
	counter := 1
	closureFunction := func() {
		counter++
	}

	closureFunction()
	closureFunction()
	closureFunction()
	fmt.Println(counter)

}

// the hard way
func getData(name string, address string, phone string, experience string) {
	fmt.Println("User Profile:", name, address, phone, experience)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }
