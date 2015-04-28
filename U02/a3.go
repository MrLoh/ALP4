package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Counter = 0
	done    = make(chan bool)
)

const (
	N = 2
	P = 100
)

func sleep(p int) {
	time.Sleep(time.Duration(int64(p) * rand.Int63n(1e5)))
}

func inc(n *int, p int) {
	Accu := *n // "LDA n"
	sleep(p)
	Accu++ // "INA"
	sleep(p + P*100)
	*n = Accu // "STA n"
	fmt.Printf("%d ", *n)
	sleep(p)
}

func count(p int) {
	for n := 0; n < N; n++ {
		inc(&Counter, p)
	}
	done <- true
}

func main() {
	for p := 0; p < P; p++ {
		go count(p)
	}
	for p := 0; p < P; p++ {
		<-done
	}
	fmt.Printf("\n\n  Counter = %d\n", Counter)
}
