package main

import "fmt"
import "errors"

type Intervall struct {
	fst, lst float64
}

// Liefert das Intervall [a, b], falls a <= b, andernfalls nil.
func New(a, b float64) Intervall {
	if a <= b {
		return Intervall{fst: a, lst: b}
	} else {
		return *new(Intervall)
	}
}

// Liefert genau dann true, wenn r in x enthalten ist.
func (X Intervall) Enthält(r float64) bool {
	return r >= X.fst && r <= X.lst
}

// Liefert den Durchschnitt von x und y, falls er nicht leer ist, und in diesem Fall nil als Fehler.
func (X Intervall) Durchschnitt(Y Intervall) (Intervall, error) {
	if X.lst <= Y.fst || X.fst >= Y.lst {
		return New(2, 1), errors.New("empty")
	} else {
		var a float64
		if X.fst <= Y.fst {
			a = Y.fst
		} else {
			a = X.fst
		}
		var b float64
		if X.lst >= Y.lst {
			b = Y.lst
		} else {
			b = X.lst
		}
		return New(a, b), nil
	}
}

// Liefert genau dann true, wenn y vollständig in x enthalten ist.
func (X Intervall) Beinhaltet(Y Intervall) bool {
	return Y.fst >= X.fst && Y.lst <= X.lst
}

// Liefert die Läange des Intervalls.
func (X Intervall) Länge() float64 {
	return X.fst - X.lst
}

// Liefert "[a,b]" füur das Intervall [a,b].
func (X Intervall) String() string {
	return fmt.Sprintf("[%v,%v]", X.fst, X.lst)
}

func main() {

}
