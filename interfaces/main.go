package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string { //not using eb in func so don't define it
	//very custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	//very custom logic for generating an spanish greeting
	return "Hola!"
}
