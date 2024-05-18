//go:build !none || exp

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Exp(x float32) float32 {
	return tinymath.Exp(x)
}
