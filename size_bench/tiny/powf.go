//go:build !none || powf

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func PowF(a, b float32) float32 {
	return tinymath.PowF(a, b)
}
