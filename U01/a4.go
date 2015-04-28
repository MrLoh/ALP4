package main

import "fmt"

type Tree struct {
	field []int
}

func compare(c1 chan int, c2 chan int) bool {
	for {
		a1, ok1 := <-c1
		a2, ok2 := <-c2
		if !ok1 && !ok2 {
			// both arrays done, all values were equal
			return true
		} else if !ok1 || !ok2 {
			// one array shorter than the other
			return false
		} else if a1 != a2 {
			// found values that don't match
			return false
		}
		// fmt.Println(a1, a2)
	}
}

func traverse(t Tree, c chan int) {
	for _, a := range t.field {
		c <- a
	}
	close(c)
}

func (t1 Tree) Equivalent(t2 Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go traverse(t1, c1)
	go traverse(t2, c2)

	return (compare(c1, c2))
}

func main() {
	t1 := Tree{field: []int{1, 2, 3, 4}}
	t2 := Tree{field: []int{1, 2, 3}}
	t3 := Tree{field: []int{1, 2, 3, 4, 5}}
	t4 := Tree{field: []int{1, 2, 3, 5}}
	t5 := Tree{field: []int{1, 2, 3, 4}}

	fmt.Println(t1.Equivalent(t1))
	fmt.Println(t1.Equivalent(t2))
	fmt.Println(t1.Equivalent(t3))
	fmt.Println(t1.Equivalent(t4))
	fmt.Println(t1.Equivalent(t5))
}
