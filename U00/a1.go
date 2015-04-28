package main

import "fmt"
import "math"

func ApplyIf(p func(int) bool, f func(int) int, in []int) []int {
	// return [f(x) for x in xs if p(x)]
	out := make([]int, 0)
	for _, x := range in {
		if p(x) {
			out = append(out, f(x))
		}
	}
	return out
}

func main() {
	xs := make([]int, 0)
	for i := 0; i <= 20; i++ { xs = append(xs, i) }
	fmt.Println(xs)

	ys := ApplyIf(
		func(x int) bool {
			return x%2 == 0
		},
		func(x int) int {
			return int(math.Pow(float64(x), 2))
		},
		xs)
	fmt.Println(ys)
}
