package main

// Ch. Maurer  v. 150223

type (
	function  func(i int) int
	predicate func(i int) bool
)

func apply_if(f function, p predicate, ls []int) []int {
	ys := make([]int, 0)
	for _, y := range ls {
		if p(y) {
			ys = append(ys, f(y))
		}
	}
	return ys
}

// oder direkt ohne die obigen Typdefinitionen
// func apply_if (f func (i int) int, p func (i int) bool, ls []int) []int {
//   ...

func quadrat(i int) int {
	return i * i
}

func ungerade(i int) bool {
	return i%2 == 1
}

func main() {
	for _, y := range apply_if(quadrat, ungerade, []int{2, 3, 1, 9, 4, 0, 5}) {
		print(y, " ")
	}
	println()
	// Alternative ohne die oben "ausgelagerten" Funktionen:
	for _, y := range apply_if(func(i int) int { return i * i },
		func(i int) bool { return i%2 == 1 },
		[]int{2, 3, 1, 9, 4, 0, 5}) {
		print(y, " ")
	}
	println()
}
