package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("string: ", reflect.TypeOf(ptrOf("string")))
	fmt.Println("int64: ", reflect.TypeOf(ptrOf(int64(1))))
	fmt.Println("*int64: ", reflect.TypeOf(ptrOf(ptrOf(int64(1)))))
}

func ptrOf[T any](source T) any {
	if reflect.ValueOf(source).Kind() == reflect.Ptr {
		return source
	}
	return &source
}
