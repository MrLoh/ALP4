package main

import . "nsp/00/bheap"

func less(a, b Any) bool {
	return a.(int) < b.(int)
}

func main() {
	s := []int{4, 7, 1, 5, 8, 3, 0, 9, 2, 6}
	n := len(s)
	h := New(less, uint(n))
	for _, a := range s {
		h.Ins(a)
	}
	for i := 0; i < n; i++ {
		s[i] = h.Get().(int)
		println(s[i])
	}
}
