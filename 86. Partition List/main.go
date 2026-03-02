package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	partition(&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}, 3)
}
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	curr1 := dummy1
	curr2 := dummy2

	for head != nil {
		if head.Val < x {
			curr1.Next = head
			curr1 = curr1.Next
		} else {
			curr2.Next = head
			curr2 = curr2.Next
		}

		head = head.Next
	}

	curr2.Next = nil
	curr1.Next = dummy2.Next

	return dummy1.Next
}
