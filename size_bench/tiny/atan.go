//go:build !none || atan

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Atan(a float32) float32 {
	return tinymath.Atan(a)
}
