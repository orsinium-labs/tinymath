//go:build !none || fract

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Fract(x float32) float32 {
	return tinymath.Fract(x)
}
