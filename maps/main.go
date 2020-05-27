package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4bf745",
		"blue":  "blau",
		"white": "#ffffff",
	}
	delete(colors, "blue")

	numbers := map[int]string{
		1: "eins",
		2: "zwei",
	}

	delete(numbers, 1)

	// var another map[string]string
	another := make(map[string]string)
	another["first"] = "erster"

	fmt.Println(colors)
	printMap(colors)
	fmt.Println(numbers)
	fmt.Println(another)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v is represented by %v\n", color, hex)
	}
}
