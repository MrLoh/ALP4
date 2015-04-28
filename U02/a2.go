package main

import "fmt"

type List interface {
	Insert(int) List
	InsertAt(int, int) List
	DeleteAt(int) List
}

type node struct {
	item       int
	prev, next *node
}

func New() List {
	return &node{}
}

func (pos *node) String() string {
	out := "|"
	for pos.next != nil {
		pos = pos.next
		out += fmt.Sprintf("%v|", pos.item)
	}
	return out
}

func (root *node) Insert(el int) List {
	pos := root
	for pos.next != nil {
		pos = pos.next
	}
	newNode := node{item: el, prev: pos, next: nil}
	pos.next = &newNode
	return root //for method chaining
}

func (root *node) InsertAt(el int, n int) List {
	pos := root
	for i := 0; i < n; i++ {
		if pos.next != nil {
			pos = pos.next
		} else {
			panic("IndexOutOfBound")
		}
	}
	if pos.next != nil {
		newNode := node{item: el, prev: pos, next: pos.next}
		pos.next.prev = &newNode
		pos.next = &newNode
	} else {
		newNode := node{item: el, prev: pos, next: nil}
		pos.next = &newNode
	}
	return root //for method chaining
}

func (root *node) DeleteAt(n int) List {
	pos := root
	for i := 0; i < n; i++ {
		if pos.next != nil {
			pos = pos.next
		} else {
			panic("IndexOutOfBound")
		}
	}
	if pos.next != nil {
		if pos.next.next != nil {
			pos.next = pos.next.next
			pos.next.next.prev = pos
		} else {
			pos.next = nil
		}
	}
	return root //for method chaining
}

func main() {
	list := New().Insert(1).Insert(9).Insert(5)
	fmt.Println(list)

	list.DeleteAt(2)
	fmt.Println(list)

	list.InsertAt(0, 2)
	fmt.Println(list)
}
