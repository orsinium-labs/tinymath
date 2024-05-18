//go:build !none || sin

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Sin(x float32) float32 {
	return tinymath.Sin(x)
}
