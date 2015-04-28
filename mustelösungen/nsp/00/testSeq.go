package main

import (
	. "nsp/00/seq"
)

func eq(a, b Any) bool {
	return a.(int) == b.(int)
}

func write(a Any) Any {
	print(a.(int), " ")
	return a
}

func odd(a Any) bool {
	return a.(int)%2 == 1
}

func main() {
	a := New(eq)
	a.Append(5)
	a.Append(1)
	a.Append(4)
	a.Append(7)
	a.Append(3)
	a.Append(6)
	a.Trav(write)
	println()
	println(a.Get(3).(int))
	println()
	a.Trav(write)
	println()
	b := a.Clone()
	b.Trav(write)
	println()
	if !b.Eq(a) {
		println("hier ist etwas faul")
	}
	//  return
	a.Ins(0, 3)
	a.Trav(write)
	println()
	// return
	println(a.Get(3).(int))
	println()
	a.Trav(write)
	println()
	a.Del(2)
	a.Trav(write)
	println()
	a.Filter(odd).Trav(write)
	println()
	// TODO Aufruf weiterer Funktionen
}
