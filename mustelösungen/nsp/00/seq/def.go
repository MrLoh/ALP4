package seq

type (
  Any interface{}
  Func func (a Any) Any
  Pred func (a Any) bool
  Rel func (a, b Any) bool
)
type
  Sequence interface {

// Liefert eine leere Folge, wobei eq
// die Gleichheitsrelation auf ihren Elementen ist.
// New(eq Rel) Sequence

// x enthält keine Elemente.
  Clr()

// Liefert genau dann true, wenn x keine Elemente enthält.
  Empty() bool

// Liefert genau dann true,
// wenn x vom gleichen Typ wie s ist und mit s übereinstimmt.
  Eq (s Sequence) bool

// Liefert eine Kopie von x (Konsequenz: x.Eq (x.Clone()) == true.)
  Clone() Sequence

// Liefert die Anzahl der Elemente in x.
  Num() uint

// a ist als letztes Element an x angehängt.
  Append (a Any)

// a ist als n-tes Element in x eingefügt, falls n < x.Num();
// andernfalls ist nichts verändert.
  Ins (a Any, n uint)

// Liefert genau dann einen Wert < x.Num(), wenn a in x enthalten ist,
// wobei dieser Wert die erste Stelle angibt, an der a in x vorkommt;
// liefert andernfalls x.Num().
  Ex (a Any) uint

// Liefert das n-te Element von x, falls n < x.Num(); andernfalls nil.
  Get (n uint) Any

// Wenn n < x.Num(), ist das n-te Element von x aus x entfernt;
// andernfalls ist nichts verändert.
  Del (n uint) Any

// Alle Elemente von x sind durch ihre jeweiligen Werte unter f ersetzt.
  Trav (f Func)

// Liefert die Folge, die aus allen denjenigen Elementen von x besteht,
// auf die p zutrifft (in unveränderter Reihenfolge).
  Filter (p Pred) Sequence
}
