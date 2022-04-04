package main

import "fmt"

// To execute Go code, please declare a func main() in a package "main"

// The TestCase is shown below
// Input : 1 2
// Output : 3

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateList(nums)
	for head != nil {
		//fmt.Println(head.Val)
		head = head.Next
	}
	newHead := reverseList(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}

}

func CreateList(nums []int) *ListNode {
	head := &ListNode{Val: -1}
	head.Next = nil
	p := head

	for _, v := range nums {
		node := &ListNode{}
		node.Val = v
		node.Next = nil
		p.Next = node
		p = p.Next
	}
	return head.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
