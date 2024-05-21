//go:build !none || sqrt

package main

import "math"

//go:export f
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
