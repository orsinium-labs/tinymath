//go:build tinygo.wasm || wasm

package tinymath

func Floor(x float32) float32 {
	return wasmFloor(x)
}
func Ceil(x float32) float32 {
	return wasmCeil(x)
}
func Trunc(x float32) float32 {
	return wasmTrunc(x)
}

func wasmFloor(x float32) float32
func wasmCeil(x float32) float32
func wasmTrunc(x float32) float32
