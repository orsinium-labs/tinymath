//go:build !none || sqrt

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Sqrt(x float32) float32 {
	return tinymath.Sqrt(x)
}
