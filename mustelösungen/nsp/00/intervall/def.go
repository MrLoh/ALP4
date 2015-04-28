package intervall

type Intervall interface { // Mit "x" ist immer das aufrufende Objekt gemeint.

// Liefert das Intervall [a, b], falls a <= b, andernfalls nil.
// New (a, b float64) Intervall

// Liefert genau dann true, wenn r in x enthalten ist.
  Enthält (r float64) bool

// Liefert den Durchschnitt von x und y, falls er nicht leer ist,
// und in diesem Fall nil als Fehler.
  Durchschnitt (Y Intervall) (Intervall, error)

// Liefert genau dann true, wenn y vollständig in x enthalten ist.
  Beinhaltet (Y Intervall) bool

// Liefert die Länge des Intervalls.
  Länge() float64

// Liefert "[a,b]" für das Intervall [a,b].
  String() string
}
