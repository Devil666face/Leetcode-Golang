package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	n := l
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}

func add(l *ListNode, v int) {
	n := &ListNode{Val: v}
	if l.Next == nil {
		l.Next = n
		return
	}
	c := l.Next
	for c.Next != nil {
		c = c.Next
	}
	c.Next = n
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		list = &ListNode{Val: head.Val}
		now  = head
		uniq = map[int]struct{}{
			head.Val: struct{}{},
		}
	)
	for now != nil {
		if _, ok := uniq[now.Val]; !ok {
			uniq[now.Val] = struct{}{}
			add(list, now.Val)
			fmt.Println(now.Val)
		}
		now = now.Next
	}
	fmt.Println(uniq)
	return list
}

// func (current *ListNode) UniqValues() []int {
// 	v := []int{}
// 	uniq := make(map[int]struct{})
// 	next := current
// 	for next != nil {
// 		uniq[next.Val] = struct{}{}
// 		// v = append(v, next.Val)
// 		next = next.Next
// 	}
// 	fmt.Println(uniq)
// 	return v
// }

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	l := deleteDuplicates(list1)
	l.Print()
	// deleteDuplicates(list1)
}
