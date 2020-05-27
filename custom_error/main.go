package main

import (
	"fmt"
	"math"
)

type CustomError struct {
	number float64
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.number)
}

func Sqrt(x float64) (s float64, e error) {
	if x < 0 {
		return -1, &CustomError{x}
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	numbers := []float64{2.0, -2.0, 7.0}
	for _, number := range numbers {
		fmt.Print(number, " -> ")
		number, err := Sqrt(number)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(number)
		}
	}
}
