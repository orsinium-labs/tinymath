//go:build !none || trunc

package main

import "github.com/orsinium-labs/tinymath"

//go:export f
func Trunc(x float32) float32 {
	return tinymath.Trunc(x)
}

func main() {
	Trunc(10)
}
