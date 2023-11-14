package main

import (
	"fmt"
	"sort"
)

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	s := appendAndSort(list1, list2)
	if len(s) == 0 {
		return nil
	}
	f := &ListNode{Val: s[0]}
	for i := 1; i < len(s); i++ {
		add(f, s[i])
	}
	return f
}

func getValues(list *ListNode) []int {
	v := []int{}
	n := list
	for n != nil {
		v = append(v, n.Val)
		n = n.Next
	}
	return v
}

func appendAndSort(lists ...*ListNode) []int {
	v := []int{}
	for _, s := range lists {
		v = append(v, getValues(s)...)
	}
	sort.Ints(v)
	return v
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	mergeTwoLists(list1, list2)
}
