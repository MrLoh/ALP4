package main

import . "nsp/00/set"

func less(a, b Any) bool {
	return a.(int) < b.(int)
}

func write(a Any) Any {
	println(a.(int))
	return a
}

func inc10(a Any) Any {
	return a.(int) + 10
}

func main() {
	x := New(less)
	for _, a := range []int{4, -1, -5, 9, 0, -2, 7, 0, -8, -5, 3, 4, -1, 6} {
		x.Ins(a)
	}
	println(x.Num())
	println()
	x.Trav(write)
	println()
	x.Del(0)
	println(x.Min().(int), x.Max().(int))
	println()
	x.Del(9)
	x.Trav(inc10)
	x.Trav(write)
	println()
	println(x.Min().(int), x.Max().(int))
	println()
	for !x.Empty() {
		x.Del(x.Min())
		x.Trav(write)
		println()
	}
}
