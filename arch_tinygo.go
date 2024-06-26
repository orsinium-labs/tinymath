//go:build tinygo.wasm

package tinymath

// Functions in this file are inlined and optimized by TinyGo compiler.
// The result is a single wasm instruction.
//
// https://github.com/tinygo-org/tinygo/blob/6384ecace093df2d0b93915886954abfc4ecfe01/compiler/intrinsics.go#L114C5-L114C22

import (
	"math"
	"math/bits"
)

func Ceil(self float32) float32 {
	return float32(math.Ceil(float64(self)))
}

func Floor(self float32) float32 {
	return float32(math.Floor(float64(self)))
}

func Sqrt(self float32) float32 {
	return float32(math.Sqrt(float64(self)))
}

func Trunc(self float32) float32 {
	return float32(math.Trunc(float64(self)))
}

func leadingZeros(x uint32) uint32 {
	return uint32(bits.LeadingZeros32(x))
}
