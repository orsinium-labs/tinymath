//go:build !none || hypot

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Hypot(a, b float32) float32 {
	return tinymath.Hypot(a, b)
}
