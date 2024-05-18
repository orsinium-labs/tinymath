//go:build !none || fract

package main

import "math"

//go:export f
func Fract(x float64) float64 {
	r, _ := math.Frexp(x)
	return r
}
