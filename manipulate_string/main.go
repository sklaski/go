package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	// Implement me! :)
	fmt.Println("whole:", text)
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if len(words[i]) > 1 {
			secondChar := words[i][1]
			lastChar := words[i][len(words[i])-1]
			words[i] = replaceAtIndex(words[i], rune(lastChar), 1)
			words[i] = replaceAtIndex(words[i], rune(secondChar), len(words[i])-1)
		}
		words[i] = strings.Replace(words[i], string(words[i][0]), strconv.Itoa(int(words[i][0])), 1)

	}
	return strings.Join(words, " ")
}

func replaceAtIndex(s string, r rune, i int) string {
	out := []rune(s)
	out[i] = r
	return string(out)
}

func main() {
	toTest := []string{"A wise old owl lived in an oak", "Asdf sd"}
	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", EncryptThis(arr))
	}
}
