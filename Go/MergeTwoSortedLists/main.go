package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var l1 ListNode
	var l2 ListNode
	l1.Val = 10
	l2.Val = 20

	result := mergeTwoLists(&l1, &l2)
	fmt.Println(*result)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode

	if list1.Val < list2.Val {
		res = list1
		res.Next = mergeTwoLists(list1.Next, list2)
	} else {
		res = list2
		res.Next = mergeTwoLists(list1, list2.Next)
	}

	return res
}
