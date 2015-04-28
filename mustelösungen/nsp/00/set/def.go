package set

type (
  Any interface{}
  Func func (a Any) Any
  Rel func (a, b Any) bool
)
type
  Set interface {

// Liefert eine leere Menge, wobei less
// die strikte Ordnungsrelation auf ihren Elementen ist.
// New (less Rel) Set

// x enthält keine Elemente.
  Clr()

// Liefert genau dann true, wenn x keine Elemente enthält.
  Empty() bool

// Liefert die Anzahl der Elemente von x.
  Num() uint

// Falls a noch nicht in x enthalten war, ist a jetzt in x eingeordnet;
// andernfalls ist nichts verändert.
  Ins (a Any)

// Liefert genau dann true, wenn a in x enthalten ist,
  Ex (a Any) bool

// Liefert das kleinste Element von x.
  Min() Any

// Liefert das größte Element von x.
  Max() Any

// Falls a in x enthalten war, ist a jetzt aus x entfernt;
// andernfalls ist nichts verändert.
  Del (a Any)

// Alle Elemente von x sind durch ihre jeweiligen Werte unter f ersetzt.
  Trav (f Func)
}
