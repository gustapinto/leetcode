package main

// https://leetcode.com/problems/add-two-numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func getValueAndNext(list *ListNode) (int, *ListNode) {
	if list != nil {
		return list.Val, list.Next
	}

	return 0, nil
}

func sumTwoLists(l1 *ListNode, l2 *ListNode, carrier int) *ListNode {
	if l1 == nil && l2 == nil && carrier == 0 {
		return nil
	}

	l1Val, l1Next := getValueAndNext(l1)
	l2Val, l2Next := getValueAndNext(l2)

	value := l1Val + l2Val + carrier

	list := &ListNode{value % 10, nil}
	list.Next = sumTwoLists(l1Next, l2Next, value/10)

	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return sumTwoLists(l1, l2, 0)
}
