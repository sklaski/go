package main

import (
	"fmt"
	"math"
)

type Node struct {
	id            int64
	points        int64
	treeStructure []int64
	children      []*Node
}

var nodeTable = map[int64]*Node{}
var root *Node
var target int64
var targetArr []int64

func add(id, parentId int64) {
	fmt.Printf("add: id=%v parentId=%v\n", id, parentId)

	node := &Node{id: id, points: id * id, children: []*Node{}}
	if parentId == 0 {
		node.points = 0
		root = node
	} else {
		parent, ok := nodeTable[parentId]
		if !ok {
			fmt.Printf("add: parentId=%v: not found\n", parentId)
			return
		}

		parent.children = append(parent.children, node)
		node.points += parent.points
		node.treeStructure = append(parent.treeStructure, id)
	}

	nodeTable[id] = node

}

func showNode(node *Node, prefix string) {

	if prefix == "" {
		fmt.Printf("%v %v \n\n", node.id, node.id*node.id)
	} else {
		fmt.Printf("%v %v  %v\n\n", prefix, node.id, node.id*node.id)
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
			if nodeTable[i] != nil && nodeTable[i].points == target {
				targetArr = nodeTable[i].treeStructure
			}
			if targetArr == nil {
				calculate(i, maxA-try)
			}
		}
	}
	return targetArr
}

func DecomposeTreeWithStructureValue(n int64) []int64 {
	target = n * n
	add(n, 0)
	maxA := math.Pow(float64(n), 2)
	fmt.Println("Result:", calculate(n, float64(int64(maxA))))
	//showNode(root, "")
	return []int64{}
}

func main() {
	toTest := []int64{50}
	for _, arr := range toTest {
		fmt.Println("org:", arr, "enc:", DecomposeTreeWithStructureValue(arr))
	}
}
