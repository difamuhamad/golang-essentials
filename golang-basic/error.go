package main

import (
	"errors"
	"fmt"
)

func Division(number, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Division with zero")
	} else {
		return number / divisor, nil
	}
}

func main() {

	hasil, err := Division(50, 0)
	if err == nil {
		fmt.Println("Quotient number:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
