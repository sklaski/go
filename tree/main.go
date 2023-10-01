package main

import (
	"fmt"
	"math"
)

type Node struct {
	id       int64
	parentId int64
	children []*Node
}

var nodeTable = map[int64]*Node{}
var root *Node
var targetArr []int64

func add(id, parentId int64) {
	//fmt.Printf("add: id=%v parentId=%v\n", id, parentId)

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
		node.parentId = parent.id
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

func calculate(n int64, maxA float64) []int64 {
	for i := n - 1; i > 0; i-- {
		if targetArr != nil {
			break
		}
		try := math.Pow(float64(i), 2)
		if maxA-try >= 0 {
			add(i, n)
			calculate(i, maxA-try)
		}
		if nodeTable[i] != nil && maxA-try == 0 {
			var sum int64
			targetArr = nil
			sum += i * i
			targetArr = append(targetArr, i)
			j := i
			for nodeTable[nodeTable[j].parentId].parentId != 0 {
				j = nodeTable[j].parentId
				sum += j * j
				targetArr = append(targetArr, j)
			}
			//fmt.Println("Inter:", sum, targetArr)
			//show all possible nodes
			//targetArr = nil

		}
	}
	return targetArr
}

func Decompose(n int64) []int64 {
	add(n, 0)
	maxA := math.Pow(float64(n), 2)
	return calculate(n, float64(int64(maxA)))
	//showNode(root, ""
}

func main() {
	toTest := []int64{50}
	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", Decompose(arr))
	}
	showNode(root, "")
}
