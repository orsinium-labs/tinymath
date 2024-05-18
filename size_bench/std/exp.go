//go:build !none || exp

package main

import "math"

//go:export f
func Exp(x float64) float64 {
	return math.Exp(x)
}
