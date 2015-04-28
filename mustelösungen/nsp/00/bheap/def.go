package bheap

// Ch. Maurer  v. 150223

type (
  Any interface{}
  Rel func (a, b Any) bool
)
type
  BoundedHeap interface {

// Liefert eine leeren beschränkte Halde der Kapazität k
// mit less als Ordnungsrelation für ihre Elemente.
// New(less Rel, k uint) BoundedHeap

// x enthält keine Elemente.
  Clr()

// Liefert genau dann true, wenn x keine Elemente enthält.
  Empty() bool

// Liefert genau dann true, wenn x so viele Elemente enthält,
// wie x bei der Konstruktion als Kapazität mitgegeben wurde.
  Full() bool

// Wenn x.Full(), ist nichts verändert;
// andernfalls ist a als letztes Element in x eingefügt.
  Ins (a Any)

// Liefert nil, wenn leer ist; andernfalls das erste Element von x
// und dieses Element ist aus x entfernt.
  Get() Any
}
