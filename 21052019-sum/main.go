package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{0, nil}
	var resTmp = res

	carry := 0
	for {
		var tmp int
		if l1 != nil && l2 != nil {
			tmp = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 != nil {
			tmp = 0 + l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			tmp = l1.Val + 0
			l1 = l1.Next
		} else {
			break
		}

		//resTmp.Val = carry + tmp/10

		resTmp.Next = &ListNode{(tmp + carry) % 10, nil}
		carry = (tmp + carry) / 10
		resTmp = resTmp.Next
	}
	if carry > 0 {
		resTmp.Next = &ListNode{carry, nil}
	}
	return res.Next
}

func main() {
	l10 := &ListNode{3, nil}
	l11 := &ListNode{4, l10}
	l1 := &ListNode{2, l11}

	l20 := &ListNode{4, nil}
	l21 := &ListNode{6, l20}
	l2 := &ListNode{5, l21}

	res := addTwoNumbers(l1, l2)
	for {
		if res == nil {
			break
		}
		fmt.Print(" ", res.Val)
		res = res.Next
	}
	fmt.Println()

	for {
		if l1 == nil {
			break
		}
		fmt.Print(" ", l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
	for {
		if l2 == nil {
			break
		}
		fmt.Print(" ", l2.Val)
		l2 = l2.Next
	}
}
