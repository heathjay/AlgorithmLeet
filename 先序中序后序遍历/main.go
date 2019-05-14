package main

import "fmt"

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Data  interface{}
}

func NewBinaryNode(data interface{}) *BinaryNode {
	return &BinaryNode{Data: data}
}

func (this *BinaryNode) PreOder() {
	if this == nil {
		return
	}
	fmt.Print(this.Data, " ")
	this.Left.PreOder()
	this.Right.PreOder()
}

func (this *BinaryNode) InfixOrder() {
	if this == nil {
		return
	}
	this.Left.InfixOrder()
	fmt.Print(this.Data, " ")
	this.Right.InfixOrder()
}

func (this *BinaryNode) PostOrder() {
	if this == nil {
		return
	}
	this.Left.PostOrder()
	this.Right.PostOrder()
	fmt.Print(this.Data, " ")

}
func main() {
	nodeL := BinaryNode{nil, nil, 2}
	nodeM := BinaryNode{nil, nil, 5}
	nodeN := BinaryNode{nil, nil, 11}
	nodeO := BinaryNode{nil, nil, 15}

	nodeH := BinaryNode{&nodeL, &nodeM, 4}
	nodeI := BinaryNode{&nodeN, &nodeO, 14}
	nodeJ := BinaryNode{nil, nil, 20}
	nodeK := BinaryNode{nil, nil, 16}

	nodeD := BinaryNode{nil, nil, 0}
	nodeE := BinaryNode{nil, nil, 3}
	nodeF := BinaryNode{&nodeH, &nodeI, 10}
	nodeG := BinaryNode{&nodeJ, &nodeK, 13}

	nodeB := BinaryNode{&nodeD, &nodeE, 1}
	nodeC := BinaryNode{&nodeF, &nodeG, 12}

	nodeA := BinaryNode{&nodeB, &nodeC, 6}
	fmt.Println("pre")
	nodeA.PreOder()
	fmt.Println()
	fmt.Println("infix")
	nodeA.InfixOrder()
	fmt.Println()
	fmt.Println("post")
	nodeA.PostOrder()
	fmt.Println()
}
