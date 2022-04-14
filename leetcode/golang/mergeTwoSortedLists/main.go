package main

import (
	"fmt"
)

func main() {
	list1 := &ListNode{Val: 1, Next: nil}
	list1.Next = &ListNode{Val: 2, Next: nil}
	list1.Next.Next = &ListNode{Val: 4, Next: nil}

	list2 := &ListNode{Val: 1, Next: nil}
	list2.Next = &ListNode{Val: 3, Next: nil}
	list2.Next.Next = &ListNode{Val: 4, Next: nil}

	res := mergeTwoLists(list1, list2)
	fmt.Println(
		res.Val,
		res.Next.Val,
		res.Next.Next.Val,
		res.Next.Next.Next.Val,
		res.Next.Next.Next.Next.Val,
		res.Next.Next.Next.Next.Next.Val,
	)

	res = mergeTwoLists(nil, nil)
	fmt.Println(
		res,
	)

	list3 := &ListNode{Val: 0, Next: nil}
	res = mergeTwoLists(nil, list3)
	fmt.Println(
		res.Val,
	)
}

type ListNode struct {
	Val		int
	Next	*ListNode
}

func putNodeToResult(list1, list2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if list1 != nil && list2 != nil && list1.Val < list2.Val ||
		list1 != nil && list2 == nil {
		return list1, list1.Next, list2
	} else if list1 != nil && list2 != nil && list1.Val >= list2.Val ||
		list2 != nil && list1 == nil {
		return list2, list1, list2.Next
	}
	return nil, nil, nil
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head, list1, list2 := putNodeToResult(list1, list2)
	curr := head
	for list1 != nil || list2 != nil {
		curr.Next, list1, list2 = putNodeToResult(list1, list2)
		curr = curr.Next
	}

	return head
}
