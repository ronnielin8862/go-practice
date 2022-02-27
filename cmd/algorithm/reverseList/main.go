package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表 https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
//放棄
func main() {
	head := []int{1, 2, 3, 4, 5}
	//head := []int{5, 4, 3, 2, 1}
	result := makeListNode(ListNode{}, head)
	fmt.Println(result.Next.Next.Val)
}

func makeListNode(listNode ListNode, head []int) *ListNode {

	return doMakeListNode(listNode, head, 0)
}

func doMakeListNode(listNode ListNode, head []int, length int) *ListNode {
	fmt.Println(listNode.Val)
	if length >= len(head) {
		return &listNode
	} else {
		listNode.Val = head[length]
		listNode.Next = doMakeListNode(listNode, head, length+1)
		return &listNode
	}
}

//func reverseList(head *ListNode) *ListNode {
//
//}
