package main

import "fmt"

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/
//给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
//
//示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := []int{1, 2, 3, 4, 5}
	//head := []int{5, 4, 3, 2, 1}
	heads := makeListNode(ListNode{}, head)
	newNode := removeNthFromEnd(heads, 2)

	for {
		if newNode != nil {
			fmt.Println("結束前", newNode.Val)
			newNode = newNode.Next
		} else {
			break
		}
	}
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := head
	cur := head

	leng := 0
	for {
		if cur != nil {
			leng++
			cur = cur.Next
		} else {

			break
		}
	}

	for i := 0; i < leng-n; i++ {
		pre = pre.Next
	}
	pre.Val = pre.Next.Val
	pre.Next = pre.Next.Next
	return head
}
