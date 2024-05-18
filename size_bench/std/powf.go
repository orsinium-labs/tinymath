//go:build !none || powf

package main

import "math"

//go:export f
func PowF(a, b float64) float64 {
	return math.Pow(a, b)
}
