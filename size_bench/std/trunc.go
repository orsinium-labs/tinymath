//go:build !none || trunc

package main

import "math"

//go:export f
func Trunc(x float64) float64 {
	return math.Trunc(x)
}
