package main

import (
	"container/list"
	"fmt"
)

func DivideFind(preOrder, InfixOrder []int, stack *(list.List)) {

	//边界条件
	if len(preOrder) == 0 || len(InfixOrder) == 0 {
		return
	}
	if preOrder[0] == InfixOrder[0] {
		stack.PushBack(preOrder[0])
		return

	}

	mark := preOrder[0]
	stack.PushBack(mark)
	var index = -1
	for i, v := range InfixOrder {
		if v == mark {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("maybe wrong")
		return
	}
	l := len(preOrder)
	var slice []int
	var rightTreeInfix, rightTreePre []int
	if index == (l - 1) {
		rightTreeInfix = slice
		rightTreePre = slice
	} else {
		rightTreeInfix = InfixOrder[index+1:]
		rightTreePre = preOrder[index+1:]
	}

	DivideFind(rightTreePre, rightTreeInfix, stack)
	var leftTreeInfix, leftTreePre []int

	leftTreeInfix = InfixOrder[:index]
	leftTreePre = preOrder[1 : index+1]

	DivideFind(leftTreePre, leftTreeInfix, stack)

}

func main() {
	preOrder := []int{0, 1, 3, 7, 8, 4, 9, 2, 5, 6}
	InfixOrder := []int{7, 3, 8, 1, 9, 4, 0, 5, 2, 6}
	//7 8 3 9 4 1 5 6 2 0
	fmt.Println(preOrder)
	fmt.Println(InfixOrder)

	stack := list.New()

	DivideFind(preOrder, InfixOrder, stack)
	for {
		if stack.Len() == 0 {
			break
		}
		v := stack.Back()
		a := v.Value.(int)
		fmt.Printf("%v+", a)
		stack.Remove(v)
	}
}
