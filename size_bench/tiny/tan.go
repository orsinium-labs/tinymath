//go:build !none || tan

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Tan(x float32) float32 {
	return tinymath.Tan(x)
}
