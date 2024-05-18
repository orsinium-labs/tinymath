//go:build !none || atan2

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Atan2(a, b float32) float32 {
	return tinymath.Atan2(a, b)
}
