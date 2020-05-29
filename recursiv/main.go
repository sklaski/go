package main

import (
	"fmt"
)

func main() {
	toTest := []int64{5, 50, 63, 77, 871, 243}

	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", Decompose(arr))
	}

	i := 5
	for i < 100 {
		fmt.Println("org:", i, "enc:", Decompose(int64(i)))
		i++
	}
}

func Decompose(n int64) []int64 {
	var sliceInt []int64
	sliceInt = decomposer(n, n*n)
	// keine Lösung
	if sliceInt == nil {
		return nil
	}
	//Startzahl aus Slice entfernen
	return sliceInt[:len(sliceInt)-1]
}

func decomposer(n int64, remain int64) []int64 {
	// basic case
	if remain == 0 {
		var r []int64
		r = append(r, n)
		return r
	}
	// iterate all element less than n
	for i := n - 1; i > 0; i-- {
		if (remain - i*i) >= 0 {
			//Wenn noch Rest, dann weiter
			r := decomposer(i, remain-i*i)
			// Wenn ich erfolgreich zurück komme "aktuelles" n als richtig merken
			if r != nil {
				r = append(r, n)
				return r
			}
		}
	}
	// keine Lösung
	return nil
}
