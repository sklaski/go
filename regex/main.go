package main

import (
	"fmt"
	"regexp"
)

func alphanumeric(str string) bool {
	valid := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return valid.MatchString(str)

	//regEx split
	//	split := regexp.MustCompile("\\s+").Split(scanner.Text(), -1)
}

func main() {
	toTest := []string{"ciao\\n$$_", "Wolfenstein", "1234", "2123asdf"}
	for _, arr := range toTest {
		fmt.Println(arr, alphanumeric(arr))
	}
}
