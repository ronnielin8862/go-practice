package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLink() *ListNode {
	node := makeLink(&ListNode{Val: 1})
	return node
}

func makeLink(node *ListNode) *ListNode {
	if node.Val == 5 {
		return node
	}

	var newNode ListNode
	newNode.Val = node.Val + 1
	node.Next = makeLink(&newNode)

	return node
}

func main() {
	nodeList := getLink()
	for oldNode := nodeList; oldNode != nil; oldNode = oldNode.Next {
		fmt.Println("original = ", oldNode.Val)
	}

	reverseNode := reverseLink(nodeList)
	for node := reverseNode; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func reverseLink(node *ListNode) *ListNode {
	var bigNode *ListNode
	if node.Next != nil {
		bigNode = reverseLink(node.Next)
	}
	bigNode = node
	node.Next = nil
	return bigNode
}
