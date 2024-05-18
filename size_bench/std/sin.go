//go:build !none || sin

package main

import "math"

//go:export f
func Sin(x float64) float64 {
	return math.Sin(x)
}
