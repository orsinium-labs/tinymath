//go:build !none || ln

package main

import "math"

//go:export f
func Ln(x float64) float64 {
	return math.Log(x)
}
