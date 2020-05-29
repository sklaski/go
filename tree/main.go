package main

import (
	"fmt"
	"math"
)

type Node struct {
	id       int64
	children []*Node
}

var nodeTable = map[int64]*Node{}
var root *Node

func add(id, parentId int64) {
	fmt.Printf("add: id=%v parentId=%v\n", id, parentId)

	node := &Node{id: id, children: []*Node{}}

	if parentId == 0 {
		root = node
	} else {

		parent, ok := nodeTable[parentId]
		if !ok {
			fmt.Printf("add: parentId=%v: not found\n", parentId)
			return
		}

		parent.children = append(parent.children, node)
	}

	nodeTable[id] = node
}

func showNode(node *Node, prefix string) {
	if prefix == "" {
		fmt.Printf("%v\n\n", node.id)
	} else {
		fmt.Printf("%v %v\n\n", prefix, node.id)
	}
	for _, n := range node.children {
		showNode(n, prefix+"--")
	}
}

func calculate(n int64, maxA float64) {
	for i := n - 1; i > 0; i-- {
		try := math.Pow(float64(i), 2)
		if maxA-try >= 0 {
			add(i, n)
			calculate(i, maxA-try)
		}
	}
}

func Decompose(n int64) []int64 {
	add(n, 0)
	maxA := math.Pow(float64(n), 2)
	calculate(n, float64(int64(maxA)))
	showNode(root, "")
	return []int64{}
}

func main() {
	toTest := []int64{6}
	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", Decompose(arr))
	}
}
