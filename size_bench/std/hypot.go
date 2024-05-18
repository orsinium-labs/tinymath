//go:build !none || hypot

package main

import "math"

//go:export f
func Hypot(a, b float64) float64 {
	return math.Hypot(a, b)
}
