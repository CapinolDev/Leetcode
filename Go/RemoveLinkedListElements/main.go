package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	var res *ListNode
	if head == nil {
		return res
	}
	if head.Val == val {
		res = removeElements(head.Next, val)

	} else {
		res = head
		res.Next = removeElements(head.Next, val)
	}
	return res
}
