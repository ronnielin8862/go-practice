package main

import "fmt"

//todo 初始化有問題
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	head := []int{3, 9, 20, 0, 0, 15, 7}
	node := TreeNode{}
	newNode := node.create(head)
	newNode.PrevSort()

	fmt.Println(newNode.Left.Val)
	fmt.Println(newNode.Right.Val)
	fmt.Println(newNode.Left.Left.Val)
	fmt.Println(newNode.Right.Right.Val)
}

var i int

func (node *TreeNode) create(head []int) *TreeNode {
	if node == nil || i >= len(head) {
		return nil
	}

	newNode := TreeNode{}
	newNode.Val = head[i]
	i = i + 1
	newNode.Left = newNode.create(head)
	newNode.Right = newNode.create(head)

	return &newNode
}

// 遍历二叉树 —— 先序遍历： DLR
func (node *TreeNode) PrevSort() {
	// 递归出口
	if node == nil {
		return
	}
	// 先D, 先打印数据域
	fmt.Print(node.Val, " ")
	// 再左, 左子树递归调用本函数
	node.Left.PrevSort()
	// 再右，右子树递归调用本函数
	node.Right.PrevSort()
}

func maxDepth(root *TreeNode) int {

	return 0
}
