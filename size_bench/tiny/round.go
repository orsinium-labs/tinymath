//go:build !none || round

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Round(x float32) float32 {
	return tinymath.Round(x)
}
