package main

func main() {

}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func BuildTreeNode(data int) TreeNode {
	tree := new(TreeNode)
	tree.data = data
	return *tree
}

func BinaryTree(array []int, index int) {

	if index >= len(array) {
		return
	}

	leftIndex := 2 * index
	//rightIndex := 2*index + 1

	if leftIndex < len(array) {
		node := BuildTreeNode(array[leftIndex])
		node.left = &node
		BinaryTree(array, leftIndex)
	}

}
