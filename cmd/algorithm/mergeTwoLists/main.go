package main

import "fmt"

//https: //leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnbp2/
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//示例 2：
//
//输入：l1 = [], l2 = []
//输出：[]
//示例 3：
//
//输入：l1 = [], l2 = [0]
//输出：[0]

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := []int{1, 2, 3}
	list2 := []int{1, 3, 4}
	listNode1 := makeListNode(ListNode{}, list1)
	listNode2 := makeListNode(ListNode{}, list2)
	newList := mergeTwoLists(listNode1, listNode2)

	//for {
	//	if newList != nil {
	//		fmt.Println("結束前", newList.Val)
	//		newList = newList.Next
	//	} else {
	//		break
	//	}
	//}

	for cur := newList; cur != nil; cur = cur.Next {
		fmt.Println("結束前", cur.Val)
	}
}

func mergeTwoLists(listNode1 *ListNode, listNode2 *ListNode) *ListNode {
	if listNode1 == nil {
		return listNode2
	}
	if listNode2 == nil {
		return listNode1
	}

	if listNode1.Val < listNode2.Val {
		listNode1.Next = mergeTwoLists(listNode1.Next, listNode2)
		return listNode1
	} else {
		listNode2.Next = mergeTwoLists(listNode2.Next, listNode1)
		return listNode2
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
