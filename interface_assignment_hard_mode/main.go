package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println("Got: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = io.Copy(os.Stdout, file)
}
