package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	c := root
	carry := 0
	for {
		c.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10

		if l2.Next == nil && l1.Next == nil && carry == 0 {
			break
		} else {
			c.Next = &ListNode{}
			c = c.Next
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
	}

	return root
}

func main() {
	a := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	b := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	res := addTwoNumbers(a, b)
	for {
		fmt.Printf("%d", res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}
