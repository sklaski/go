package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("netstat", "-nr")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var outputSlice []string
	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "10.0.0.0") {
			fmt.Println(scanner.Text())
		}
		outputSlice = append(outputSlice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, row := range outputSlice {
		fmt.Println(row)
	}
}
