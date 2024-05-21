//go:build !none || tan

package main

import "math"

//go:export f
func Tan(x float64) float64 {
	return math.Tan(x)
}
