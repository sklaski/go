package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, number := range numbers {
		fmt.Printf("%v is ", number)
		if number%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
