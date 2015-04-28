package seq

// Ch. Maurer  v. 150223

type
  sequence struct {
                s []Any
               eq Rel
                  }

func New(eq Rel) Sequence {
  return &sequence { make([]Any, 0), eq }
}

func (x *sequence) Clr() {
  x.s = make([]Any, 0)
}

func (x *sequence) Empty() bool {
  return len(x.s) == 0
}

func (x *sequence) Eq (s Sequence) bool {
  y:= s.(*sequence)
  for i, a:= range x.s {
    if ! x.eq (y.s[i], a) { return false }
  }
  return true
}

func (x *sequence) Clone() Sequence {
  y:= New(x.eq).(*sequence)
/* möglich:
  for _, a:= range x.s {
    y.s = append (y.s, a)
  }
aber effizienter: */
  y.s = make([]Any, len(x.s))
  copy (y.s, x.s)
  return y
}

func (x *sequence) Num() uint {
  return uint(len(x.s))
}

func (x *sequence) Append (a Any) {
  x.s = append (x.s, a)
}

func (x *sequence) Ins (a Any, n uint) {
  if n >= uint(len(x.s)) { return }
  e:= uint(len(x.s))
/* möglich:
  t:= make([]Any, e - n)
  copy (t, x.s[n:])
  x.s = append(append(x.s[:n], a), t...)
aber effizienter: */
  x.s = append(x.s, a)
  copy(x.s[n+1:e+1-n], x.s[n:e-n])
  x.s[n] = a
}

func (x *sequence) Ex (a Any) uint {
  for n, b:= range x.s {
    if x.eq(a, b) { return uint(n) }
  }
  return uint(len(x.s))
}

func (x *sequence) Get (n uint) Any {
  if n >= uint(len(x.s)) { return nil }
  return x.s[n]
}

func (x *sequence) Del (n uint) Any {
  if n >= uint(len(x.s)) { return nil }
  a:= x.s[n]
  x.s = append(x.s[:n], x.s[n+1:]...)
  return a
}

func (x *sequence) Trav (f Func) {
  for i, a:= range x.s {
    x.s[i] = f(a)
  }
}

func (x *sequence) Filter (p Pred) Sequence {
  y:= New(x.eq).(*sequence)
  for _, a:= range x.s {
    if p(a) { y.s = append(y.s, a) }
  }
  return y
}
