package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表 https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
func main() {
	head := []int{1, 2, 3, 4, 5}
	//head := []int{5, 4, 3, 2, 1}
	heads := makeListNode(ListNode{}, head)
	result := reverseList(heads)

	for r := result; r != nil; r = r.Next {
		fmt.Println(r.Val)
	}

	//fmt.Println(result.Next.Val)
	//fmt.Println(result.Next.Next.Val)
	//fmt.Println(result.Next.Next.Next.Val)
	//fmt.Println(result.Next.Next.Next.Next.Val)

}

func makeListNode(listNode ListNode, head []int) *ListNode {

	return doMakeListNode(listNode, head, 0)
}

func doMakeListNode(listNode ListNode, head []int, length int) *ListNode {
	if length >= len(head) {
		return nil
	}

	listNode.Val = head[length]
	fmt.Println("題目組成", listNode.Val)
	listNode.Next = doMakeListNode(listNode, head, length+1)
	return &listNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
