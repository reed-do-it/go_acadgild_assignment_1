package main

import "math"

var (
	big float64 = 35.6
)

const Pi = math.Pi

var small int // static type declaration

func init() {
}

func Arbitrary() (string, bool) {
	return "yay", false
}
