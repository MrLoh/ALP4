package intervall

// Ch. Maurer   v. 150223

import (
  "errors"
  "strconv"
)
type
  intervall struct {
              a, b float64
                   }

func New (a, b float64) Intervall {
  if a > b { return nil }
  return &intervall { a, b }
}

func (x *intervall) Enthält (r float64) bool {
  return x.a <= r && r <= x.b
}

func (x *intervall) grenzen() (float64, float64) {
  return x.a, x.b
}

func min (a, b float64) float64 {
  if a < b { return a }
  return b
}

func max (a, b float64) float64 {
  if a < b { return b }
  return a
}

func (x *intervall) schneidet (y *intervall) bool {
  return x.a <= y.a && y.a <= x.b || x.a <= y.b && y.b <= x.b ||
         y.a <= x.a && x.a <= y.b || y.a <= x.b && x.b <= y.b
}

func (x *intervall) Durchschnitt (Y Intervall) (Intervall, error) {
  y:= Y.(*intervall)
  if x.schneidet (y) {
    return New (max (x.a, y.a), min (x.b, y.b)), nil
  }
  return nil, errors.New ("Durchschnitt ist leer")
}

func (x *intervall) Beinhaltet (Y Intervall) bool {
  y:= Y.(*intervall)
  return x.a <= y.a && y.b <= x.b
}

func (x *intervall) Länge() float64 {
  if x.a > x.b {
    return 0
  }
  return x.b - x.a
}

func (x *intervall) String() string {
  return "[" + strconv.FormatFloat(x.a,'f',1,64) + "," + strconv.FormatFloat(x.b,'f',1,64) + "]"
}
