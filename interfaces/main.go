package main

import "fmt"

type englishbot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishbot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreetingg {
	fmt.Println(b.getGreeting())
}

//func (eb englishbot) getGreeting() string {
func (englishbot) getGreeting() string { //not using eb in func so don't define it
	//very custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating an spanish greeting
	return "Hola!"
}
