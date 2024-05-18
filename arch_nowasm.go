//go:build !tinygo.wasm && !wasm

package tinymath

// Functions that can be optimized for wasm

// Returns the smallest integer greater than or equal to a number.
func Ceil(self float32) float32 {
	return -Floor(-self)
}

// Returns the largest integer less than or equal to a number.
func Floor(self float32) float32 {
	res := float32(int32(self))
	if self < res {
		res -= 1.0
	}
	return float32(res)
}

// Returns the integer part of a number.
func Trunc(self float32) float32 {
	const MANTISSA_MASK = 0b0000_0000_0111_1111_1111_1111_1111_1111

	x_bits := ToBits(self)
	exponent := extractExponentValue(self)

	// exponent is negative, there is no whole number, just return zero
	if exponent < 0 {
		return CopySign(0, self)
	}

	exponent_clamped := uint32(Max(exponent, 0))

	// find the part of the fraction that would be left over
	fractional_part := (x_bits << exponent_clamped) & MANTISSA_MASK

	// if there isn't a fraction we can just return the whole thing.
	if fractional_part == 0 {
		return self
	}

	fractional_mask := fractional_part >> exponent_clamped
	return FromBits(x_bits & ^fractional_mask)
}

func leadingZeros(x uint32) uint32 {
	var n uint32 = 32
	for x != 0 {
		x >>= 1
		n -= 1
	}
	return n
}
