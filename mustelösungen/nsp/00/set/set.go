package set

type
  set struct {
      anchor *bintree
        less Rel
             }

func New (less Rel) Set {
  x:= new (set)
  x.anchor = nil
  x.less = less
  return x
}

func (x *set) Clr() {
  x.anchor = nil
}

func (x *set) Empty() bool {
  return x.anchor == nil
}

func (x *set) Num() uint {
  return x.anchor.num()
}

func (x *set) Ins (a Any) {
  x.anchor = x.anchor.ins (a, x.less)
}

func (x *set) Ex (a Any) bool {
  if x.anchor == nil {
    return false
  }
  return x.anchor.ex (a, x.less)
}

func (x *set) Min() Any {
  return x.anchor.min()
}

func (x *set) Max() Any {
  return x.anchor.max()
}

func (x *set) Del (a Any) {
  x.anchor = x.anchor.del (a, x.less)
}

func (x *set) Trav (f Func) {
  x.anchor.trav (f)
}
