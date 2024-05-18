//go:build !none || ln

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Ln(x float32) float32 {
	return tinymath.Ln(x)
}
