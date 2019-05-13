package main

import (
	"fmt"
	"math"
)

//首先不光要寻找到左子树，还要在左子树中找到最大值，同理右子树也得找到最小值，在递归到头结点的时候保证头结点大于左子树最大值，小于右子树最小值。
//这时还要获取左子树最大搜索子树是否就是左子树的头结点，右子树同理，这样如果都是的话就说明最大搜索二叉树还能增大
//所以仅定义一棵树的头结点不够，还要封装结点的信息

type Node struct {
	left  *Node
	right *Node
	data  float64
}

type ReturnTupe struct {
	HeadNode   *Node
	MaxBSTSize int
	Max        float64
	Min        float64
	//Data       float64
}

func GetBi(node *Node) ReturnTupe {
	if node == nil {
		return ReturnTupe{nil, 0, math.Inf(1), math.Inf(-1)}
	}

	LTree := GetBi(node.left)
	RTree := GetBi(node.right)
	min := math.Min(node.data, math.Min(LTree.Min, RTree.Min))
	max := math.Max(node.data, math.Max(LTree.Max, RTree.Max))
	maxBSTSize := int(math.Max(float64(LTree.MaxBSTSize), float64(RTree.MaxBSTSize)))
	if LTree.HeadNode == nil && RTree.HeadNode != nil {
		if node.data < RTree.Min {
			return ReturnTupe{node, RTree.MaxBSTSize + 1, RTree.Max, node.data}
		} else {
			return ReturnTupe{RTree.HeadNode, RTree.MaxBSTSize, RTree.Max, RTree.Min}
		}
	} else if LTree.HeadNode != nil && RTree.HeadNode == nil {
		if node.data > LTree.Max {
			return ReturnTupe{node, LTree.MaxBSTSize + 1, node.data, LTree.Min}
		} else {
			return ReturnTupe{LTree.HeadNode, LTree.MaxBSTSize, LTree.Max, LTree.Min}
		}

	} else if LTree.HeadNode == nil && RTree.HeadNode == nil {
		return ReturnTupe{node, 1, node.data, node.data}
	}

	// if RTree.HeadNode == nil {
	// 	// 	RTree.Data = math.Inf(-1)
	// 	// } else {
	// 	// 	RTree.Data = RTree.HeadNode.data
	// 	return ReturnTupe{nil, 0, math.Inf(1), math.Inf(-1), 0}
	// }

	var headNode *Node
	// var tData float64
	headNode = new(Node)
	if LTree.HeadNode.data >= RTree.HeadNode.data {
		headNode = LTree.HeadNode

	} else {
		headNode = RTree.HeadNode
		// tData = RTree.Data
	}

	if LTree.HeadNode == node.left && RTree.HeadNode == node.right && node.data > LTree.Max && node.data < RTree.Min {
		maxBSTSize = LTree.MaxBSTSize + RTree.MaxBSTSize + 1
		headNode = node
		// tData = node.data
	}
	return ReturnTupe{headNode, maxBSTSize, min, max}
}

func main() {
	//declare
	nodeL := Node{nil, nil, 2}
	nodeM := Node{nil, nil, 5}
	nodeN := Node{nil, nil, 11}
	nodeO := Node{nil, nil, 15}

	nodeH := Node{&nodeL, &nodeM, 4}
	nodeI := Node{&nodeN, &nodeO, 14}
	nodeJ := Node{nil, nil, 20}
	nodeK := Node{nil, nil, 16}

	nodeD := Node{nil, nil, 0}
	nodeE := Node{nil, nil, 3}
	nodeF := Node{&nodeH, &nodeI, 10}
	nodeG := Node{&nodeJ, &nodeK, 13}

	nodeB := Node{&nodeD, &nodeE, 1}
	nodeC := Node{&nodeF, &nodeG, 12}

	nodeA := Node{&nodeB, &nodeC, 6}
	res := GetBi(&nodeA)
	if res.HeadNode != nil {
		fmt.Println("size:", res.MaxBSTSize)
	} else {
		fmt.Println("size:", res.MaxBSTSize, "+head:err")
	}
}
