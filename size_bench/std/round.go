//go:build !none || round

package main

import "math"

//go:export f
func Round(x float64) float64 {
	return math.Round(x)
}
