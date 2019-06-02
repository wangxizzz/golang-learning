package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var head *ListNode
	cur1 := l1
	cur2 := l2
	carry := 0
	var p *ListNode
	for cur1 != nil || cur2 != nil {
		a, b := 0, 0
		if cur1 != nil {
			a = cur1.Val
		}
		if cur2 != nil {
			b = cur2.Val
		}
		sum := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		node := &ListNode{Val: sum}
		if head == nil {
			head = node
			p = head
		} else {
			p.Next = node
			p = p.Next
		}
		if cur1 != nil {
			cur1 = cur1.Next
		}
		if cur2 != nil {
			cur2 = cur2.Next
		}
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		p.Next = node
	}
	return head
}

func main() {

}
