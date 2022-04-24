package algorithm

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) Print() {
	n1 := s
	for n1 != nil {
		fmt.Printf("%v ", n1.Val)
		n1 = n1.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	n1 := head
	n2 := head.Next
	n1.Next = nil
	for n2 != nil {
		tmp := n2.Next
		n2.Next = n1
		n1 = n2
		n2 = tmp
	}
	return n1
}

// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right || left < 1 || right < 1 || head.Next == nil {
		return head
	}
	// 从头开始翻转
	if left == 1 {
		end := head
		for i := left; i < right && end != nil; i++ {
			end = end.Next
		}
		// 全部翻转
		if end == nil || end.Next == nil {
			return reverseList(head)
		}
		// 翻转前半段
		next := end.Next
		end.Next = nil
		reverse := reverseList(head)
		head.Next = next
		return reverse
	}

	beforeStart := head
	start := head
	for i := 1; i < left && start != nil; i++ {
		beforeStart = start
		start = start.Next
	}
	if start == nil {
		return head
	}
	// 寻找end节点
	end := start
	for i := left; i < right && end != nil; i++ {
		end = end.Next
	}

	// 后半段全部翻转
	if end == nil || end.Next == nil {
		beforeStart.Next = reverseList(start)
		return head
	}
	// 翻转局部
	endNext := end.Next
	end.Next = nil
	reverse := reverseList(start)
	beforeStart.Next = reverse
	start.Next = endNext
	return head
}

func Test_ReverseList(t *testing.T) {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	//n1 = reverseList(n1)
	n1 = reverseBetween(n1, 3, 4)
	n1.Print()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	cur := head
	add := 0
	for l1 != nil || l2 != nil || add > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		val := add
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		add = val / 10
		val = val % 10
		cur.Val = val
	}
	return head.Next
}
