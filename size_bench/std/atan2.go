//go:build !none || atan2

package main

import "math"

//go:export f
func Atan2(a, b float64) float64 {
	return math.Atan2(a, b)
}
