package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	var Tree1 TreeNode
	var Tree2 TreeNode

	Tree1.Val = 10
	Tree2.Val = 10

	result := isSameTree(&Tree1, &Tree2)
	fmt.Println(result)

}
func isSameTree(p *TreeNode, q *TreeNode) bool { 

	return true
}
