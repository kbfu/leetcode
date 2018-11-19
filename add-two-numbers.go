// https://leetcode-cn.com/problems/add-two-numbers/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	upper := 0
	curr := result
	for l1 != nil || l2 != nil {
		var (
			l1Val int
			l2Val int
		)
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		} else {
			l1Val = 0
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		} else {
			l2Val = 0
		}
		sum := l1Val + l2Val + upper
		if sum > 9 {
			upper = 1
			sum = sum % 10
		} else {
			upper = 0
		}
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}
	if upper > 0 {
		curr.Next = &ListNode{Val: upper}
	}
	return result.Next
}

func main() {
	l1 := &ListNode{}
	l1.Val = 5
	l2 := &ListNode{}
	l2.Val = 5
	addTwoNumbers(l1, l2)
}
