package bheap

// Ch. Maurer  v. 150223

type
  boundedHeap struct {
                less Rel
                heap []Any // heap[0] to save the type
            cap, num uint
                     }

func New(less Rel, k uint) BoundedHeap {
  x:= new(boundedHeap)
  x.less = less
  x.heap = make ([]Any, k + 1)
  x.cap = k
  x.num = 0
  return x
}

func (x *boundedHeap) Clr() {
  x.num = 0
}

func (x *boundedHeap) Empty() bool {
  return x.num == 0
}

func (x *boundedHeap) Full() bool {
  return x.num == x.cap
}

func (x *boundedHeap) Ins (a Any) {
  if x.num == x.cap { return }
  x.num++
  x.heap[x.num] = a
  x.lift()
}

func (x *boundedHeap) Get() Any {
  if x.num == 0 { return nil }
  a:= x.heap[1]
  x.heap[1] = x.heap[x.num]
  x.num--
  x.sift()
  return a
}

// heap[i] so weit wie notwendig hochziehen, um die Heap-Invariante heap[i] <= heap[j]
// für alle i:= 1, ..., (x.num - 1) / 2, j == 2 * i and j == 2 * i + 1 wieder herzustellen
func (x *boundedHeap) lift() {
  i:= x.num
  for {
    if i == 1 {
      break
    }
    j:= i / 2 // index über i
    if x.less (x.heap [j], x.heap [i]) {
      break // i < x.num, oberhalb von i ist die Heap-Invariante erfüllt
    } else {
      x.heap[i], x.heap[j] = x.heap[j], x.heap[i]
    }
    i = j // oben weitermachen
  }
}

// heap[1] so weit wie notwendig absinken lassen,
// um die Heap-Invariante wieder herzustellen
func (x *boundedHeap) sift() {
  i:= uint(1)
  for {
    if i > x.num / 2 {
      break // nichts mehr unter i
    }
    j:= 2 * i // links unter i
    if j < x.num && ! x.less (x.heap[j], x.heap[j + 1]) {
      j++ // rechts unter i
    }
    if x.less (x.heap[i], x.heap[j]) {
      break
    } else {
      x.heap[i], x.heap[j] = x.heap[j], x.heap[i]
      i = j
    }
  }
}
