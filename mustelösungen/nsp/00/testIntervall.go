package main

import . "nsp/00/intervall"

func main() {
	a13, a25, a34, a46 := New(1, 3), New(2, 5), New(3, 4), New(4, 6)
	if a13.Enthält(2) {
		println("2 in", a13.String())
	} else {
		println("2 nicht in", a13.String())
	}
	if a13.Enthält(0) {
		println("0 in", a13.String())
	} else {
		println("0 nicht in", a13.String())
	}
	if a25.Beinhaltet(a34) {
		println(a34.String(), "in", a25.String())
	} else {
		println(a34.String(), "nicht in", a25.String())
	}
	if a46.Beinhaltet(a25) {
		println(a25.String(), "in", a46.String())
	} else {
		println(a25.String(), "nicht in", a46.String())
	}
	if a, e := a25.Durchschnitt(a46); e == nil {
		println("Durchschnitt von", a25.String(), "und", a46.String(), "ist", a.String())
	}
	if a, e := a13.Durchschnitt(a46); e == nil {
		println("Durchschnitt von", a13.String(), "und", a46.String(), "ist", a.String())
	}
}
