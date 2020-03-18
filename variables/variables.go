package main

import (
	"fmt"
)

func main() {
	fun, truth := Arbitrary() // dynamic type declaration / type inference

	small = 8

	fmt.Printf("Type of small: %T\n", small)
	fmt.Printf("Type of big: %T\n", big)
	fmt.Printf("Type of fun: %T\n", fun)
	fmt.Printf("Type of truth: %T\n", truth)
	calc(small, big)
}

func calc(i int, f64 float64) {
	var result float64 = float64(i) + f64
	fmt.Printf("I just did math: %f", result)
}
