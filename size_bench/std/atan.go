//go:build !none || atan

package main

import "math"

//go:export f
func Atan(a float64) float64 {
	return math.Atan(a)
}
