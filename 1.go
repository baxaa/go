package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type Stack struct {
	Head *Node
	Tail *Node
}

func (s *Stack) Push(n int) {
	newNode := new(Node)
	newNode.Val = n
	if s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
	} else {
		newNode.Next = s.Head
		s.Head.Prev = newNode
		s.Head = newNode
	}
}

func (s *Stack) Print() {
	cur := s.Head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}

}

func (s *Stack) reverse() {
	cur := s.Tail
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Prev
	}

}

func (s *Stack) Pop() {
	if s.Head != nil {
		fmt.Println(s.Head.Val)
		s.Head = s.Head.Next
	}
}

func (s *Stack) Peek() {
	fmt.Println(s.Head.Val)
}

func (s *Stack) Contains(n int) {
	cur := s.Head
	for cur != nil {
		if cur.Val == n {
			fmt.Println("yes")
			//return
		}
		cur = cur.Next
	}
}
func (s *Stack) Increment() {
	cur := s.Head
	for cur != nil {
		cur.Val += 1
		cur = cur.Next
	}
}

func (s *Stack) Clear() {
	s.Head = nil
}

func main() {
	var n int
	var a int
	fmt.Scan(&n)
	st := new(Stack)
	for n > 0 {
		fmt.Scan(&a)
		st.Push(a)
		n--
	}
	//fmt.Print(st.Tail.Val)
	st.Print()
	st.Contains(3)
	//st.Clear()
	st.Increment()
	st.Print()
	fmt.Println()
	st.reverse()
}
