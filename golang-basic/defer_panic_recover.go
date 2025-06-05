package main

import "fmt"

func endTransaction() {
	fmt.Println("Transaction cancelled")
	message := recover()
	fmt.Println("problem :", message)
}

func runApplication(amount bool) {
	defer endTransaction()
	fmt.Println("doing transaction...")

	if !amount {
		panic("Unsifficient amount")
	}

	fmt.Println("Transaction success")

}

func main() {
	fmt.Println("Application is running...")

	runApplication(false)

	fmt.Println("Application is closed")
}

// defer will be executed after the function end, whatever the condition, panic or error
// panic not just throwing an error, but doing the stack unwinding
// recover will catch the panic value

/**
---------  Stack unwinding

A process by which a program's call stack is progressively "unwound" or cleaned up in response to an abnormal situation,
such as an exception or a panic.
During stack unwinding, the program begins at the point where the error occurred and moves backward through the stack,
exiting each active function in reverse order of their calls, while executing any cleanup code,
(such as deferred functions or destructors) along the way.
This mechanism ensures that resources are properly released,
and that the program can either recover gracefully or terminate with diagnostic information.
*/
