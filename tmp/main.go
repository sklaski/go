package main

import (
	"fmt"
	"math"
)

var target int64
var targetArr []int64

type stack struct {
	sum   int64
	stack []int64
}

func calculate(n int64, maxA float64, stack stack) []int64 {

	if stack.sum == target {
		fmt.Println("Found sum:", target)
	}
	for i := n - 1; i > 0; i-- {
		if targetArr != nil || stack.sum >= target {
			break
		}
		try := math.Pow(float64(i), 2)
		if maxA-try >= 0 {
			stack.stack = append(stack.stack, i)
			stack.sum += int64(try)
			if stack.sum == target {
				targetArr = stack.stack
			}
			calculate(i, maxA-try, stack)
		}
	}
	return targetArr
}

func Decompose(n int64) []int64 {
	targetArr = nil
	target = n * n
	maxA := math.Pow(float64(n), 2)
	return calculate(n, float64(int64(maxA)), stack{})
}

func main() {
	toTest := []int64{50}
	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", Decompose(arr))
	}
}
