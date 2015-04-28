package set

type (
  bintree struct {
                 Any "root"
            left,
           right *bintree
                 }
)

func (x *bintree) num() uint {
  if x == nil {
    return uint(0)
  }
  return x.left.num() + 1 + x.right.num()
}

func (x *bintree) ins (a Any, less Rel) *bintree {
  if x == nil {
    x:= new (bintree)
    x.Any = a
    return x
  }
  if less (a, x.Any) {
    x.left = x.left.ins (a, less)
  } else if less (x.Any, a) {
    x.right = x.right.ins (a, less)
  }
  // a is already there
  return x
}

func (x *bintree) ex (a Any, less Rel) bool {
  if x == nil {
    return false
  }
  if less (a, x.Any) {
    return x.left.ex (a, less)
  }
  if less (x.Any, a) {
    return x.right.ex (a, less)
  }
  return true // a equals x.Any
}

func (x *bintree) min() Any {
  if x.left == nil { return x.Any }
  return x.left.min()
}

func (x *bintree) max() Any {
  if x.right == nil { return x.Any }
  return x.right.max()
}

func (x *bintree) liftL (y *bintree) *bintree {
  if x.right == nil {
    y.Any = x.Any
    x = x.left
  } else {
    x.right = x.right.liftL (y)
  }
  return x
}

func (x *bintree) liftR (y *bintree) *bintree {
  if x.left == nil {
    y.Any = x.Any
    x = x.right
  } else {
    x.left = x.left.liftR (y)
  }
  return x
}

func (x *bintree) del (a Any, less Rel) *bintree {
  if x == nil {
    return nil
  }
  if less (a, x.Any) {
    x.left = x.left.del (a, less)
  } else if less (x.Any, a) {
    x.right = x.right.del (a, less)
  } else { // found tree to remove
    if x.right == nil {
      x = x.left
    } else if x.left == nil {
      x = x.right
    } else {
      x.left = x.left.liftL (x) // or: x.right = x.right.liftR (x)
    }
  }
  return x
}

func (x *bintree) trav (f Func) {
  if x != nil {
    x.left.trav (f)
    x.Any = f (x.Any)
    x.right.trav (f)
  }
}
