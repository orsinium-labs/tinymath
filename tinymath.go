package tinymath

import "unsafe"

const (
	signMask     uint32 = 0x8000_0000
	mantissaBits        = 23
	expMask      uint32 = 0b0111_1111_1000_0000_0000_0000_0000_0000
	expBias             = 127
)

func ToBits(x float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&x))
}

func FromBits(x uint32) float32 {
	return *(*float32)(unsafe.Pointer(&x))
}

func Max[N float32 | int32](a, b N) N {
	if a > b {
		return a
	}
	return b
}

func Min[N float32 | int32](a, b N) N {
	if a < b {
		return a
	}
	return b
}

// Computes the absolute value of `self`.
// /
// Returns [`NAN`] if the number is [`NAN`].
func Abs(self float32) float32 {
	return FromBits(ToBits(self) & ^signMask)
}

// Returns a number composed of the magnitude of `self` and the sign of
// `sign`.
func CopySign(self float32, sign float32) float32 {
	source_bits := ToBits(sign)
	source_sign := source_bits & signMask
	signless_destination_bits := ToBits(self) & ^signMask
	return FromBits(signless_destination_bits | source_sign)
}

// Calculates Euclidean division, the matching method for `rem_euclid`.
func DivEuclid(self float32, rhs float32) float32 {
	return (self - RemEuclid(self, rhs)) / rhs
}

// Returns `e^(self)`, (the exponential function).
func Exp(self float32) float32 {
	return ExpLn2Approx(self, 4)
}

// Exp approximation for `f32`.
func ExpLn2Approx(self float32, partial_iter uint32) float32 {
	if self == 0.0 {
		return 1
	}
	if Abs(self-1) < Epsilon {
		return E
	}
	if Abs(self-(-1)) < Epsilon {
		return 1. / E
	}

	// log base 2(E) == 1/ln(2)
	// x_fract + x_whole = x/ln2_recip
	// ln2*(x_fract + x_whole) = x
	x_ln2recip := self * Log2E
	x_fract := Fract(x_ln2recip)
	x_trunc := Trunc(x_ln2recip)

	//guaranteed to be 0 < x < 1.0
	x_fract = x_fract * Ln2
	fract_exp := ExpSmallX(x_fract, partial_iter)

	//need the 2^n portion, we can just extract that from the whole number exp portion
	fract_exponent := saturatingAdd(extractExponentValue(fract_exp), int32(x_trunc))

	if fract_exponent < -expBias {
		return 0.0
	}

	if fract_exponent > expBias+1 {
		return Inf
	}

	return setExponent(fract_exp, fract_exponent)
}

// if x is between 0.0 and 1.0, we can approximate it with the a series
//
// Series from here:
// <https://stackoverflow.com/a/6984495>
//
// e^x ~= 1 + x(1 + x/2(1 + (x?
func ExpSmallX(self float32, iter uint32) float32 {
	var total float32 = 1.0
	for i := float32(iter - 1); i > 0.; i-- {
		total = 1.0 + ((self / i) * total)
	}
	return total
}

// Returns the fractional part of a number with sign.
func Fract(self float32) float32 {
	const MANTISSA_MASK = 0b0000_0000_0111_1111_1111_1111_1111_1111

	x_bits := ToBits(self)
	exponent := extractExponentValue(self)

	// we know it is *only* fraction
	if exponent < 0 {
		return self
	}

	// find the part of the fraction that would be left over
	fractional_part := (x_bits << exponent) & MANTISSA_MASK

	// if there isn't a fraction we can just return 0
	if fractional_part == 0 {
		// TODO: most people don't actually care about -0.0,
		// so would it be better to just not CopySign?
		return CopySign(0.0, self)
	}

	// Note: alternatively this could use -1.0, but it's assumed subtraction would be more costly
	// example: 'new_exponent_bits := 127.overflowing_shl(23))) - 1.0'
	exponent_shift := (leadingZeros(fractional_part) - (32 - mantissaBits)) + 1

	fractional_normalized := (fractional_part << exponent_shift) & MANTISSA_MASK

	new_exponent_bits := (expBias - (exponent_shift)) << mantissaBits

	return CopySign(FromBits(fractional_normalized|new_exponent_bits), self)
}

// Calculate the length of the hypotenuse of a right-angle triangle.
func Hypot(self float32, rhs float32) float32 {
	return Sqrt(self*self + rhs*rhs)
}

// Fast approximation of `1/x`.
func Inv(self float32) float32 {
	return FromBits(0x7f00_0000 - ToBits(self))
}

// Approximate inverse square root with an average deviation of ~5%.
func InvSqrt(self float32) float32 {
	return FromBits(0x5f37_5a86 - (ToBits(self) >> 1))
}

// Check if the given number is NaN.
func IsNaN(x float32) bool {
	return x != x
}

// Check if the given number is even.
func IsEven(x float32) bool {
	half := x / 2
	return IsInteger(half)
}

// Check if the given number has no numbers after dot.
func IsInteger(x float32) bool {
	return Floor(x) == x
}

// Check if the number has a positive sign
func IsSignPositive(x float32) bool {
	return ToBits(x)&(1<<31) == 0
}

// Approximates the natural logarithm of the number.
// Note: excessive precision ignored because it hides the origin of the numbers used for the
// ln(1.0->2.0) polynomial
func Ln(self float32) float32 {
	// x may essentially be 1.0 but, as clippy notes, these kinds of
	// floating point comparisons can fail when the bit pattern is not the sames
	if Abs(self-1) < Epsilon {
		return 0.0
	}

	x_less_than_1 := self < 1.0

	// Note: we could use the fast inverse approximation here found in super::inv::inv_approx, but
	// the precision of such an approximation is assumed not good enough.
	x_working := self
	if x_less_than_1 {
		x_working = Inv(self)
	}

	// according to the SO post ln(x) = ln((2^n)*y)= ln(2^n) + ln(y) = ln(2) * n + ln(y)
	// get exponent value
	base2_exponent := uint32(extractExponentValue(x_working))
	divisor := FromBits(ToBits(x_working) & expMask)

	// supposedly normalizing between 1.0 and 2.0
	x_working = x_working / divisor

	// approximate polynomial generated from maple in the post using Remez Algorithm:
	// https://en.wikipedia.org/wiki/Remez_algorithm
	ln_1to2_polynomial := -1.741_793_9 + (2.821_202_6+(-1.469_956_8+(0.447_179_55-0.056_570_851*x_working)*x_working)*x_working)*x_working

	// ln(2) * n + ln(y)
	result := float32(base2_exponent)*Ln2 + ln_1to2_polynomial

	if x_less_than_1 {
		return -result
	}
	return result
}

// Approximates the logarithm of the number with respect to an arbitrary base.
func Log(self float32, base float32) float32 {
	return (1 / Ln(base)) * Ln(self)
}

// Approximates the base 10 logarithm of the number.
func Log10(self float32) float32 {
	return Ln(self) * Log10E
}

// Approximates the base 2 logarithm of the number.
func Log2(self float32) float32 {
	return Ln(self) * Log2E
}

// Approximates a number raised to a floating point power.
func PowF(self float32, n float32) float32 {
	// using x^n = exp(ln(x^n)) = exp(n*ln(x))
	if self >= 0.0 {
		return Exp(n * Ln(self))
	} else if IsInteger(n) {
		return NaN
	} else if IsEven(n) {
		// if n is even, then we know that the result will have no sign, so we can remove it
		return n * Exp(Ln(Abs(self)))
	} else {
		// if n isn't even, we need to multiply by -1.0 at the end.
		return -(n * Exp(Ln(Abs(self))))
	}
}

// Approximates a number raised to an integer power.
func PowI(self float32, n int32) float32 {
	base := self
	abs_n := n
	if n < 0 {
		abs_n = -n
	}
	var result float32 = 1
	if n < 0 {
		base = 1.0 / self
	}
	if n == 0 {
		return 1
	}
	// 0.0 == 0.0 and -0.0 according to IEEE standards.
	if self == 0.0 && n > 0 {
		return self
	}

	// For values less than 2.0, but greater than 0.5 (1.0/2.0), you can multiply longer without
	// going over exponent, i.e. 1.1 multiplied against itself will grow slowly.
	abs := Abs(self)
	if 0.5 <= abs && abs < 2.0 {
		// Approximation if we end up outside of the range of floating point values,
		// then we end early
		approx_final_exponent := extractExponentValue(self) * n
		const max_representable_exponent = 127
		const min_representable_exponent = -126 - mantissaBits
		if approx_final_exponent > max_representable_exponent || (self == 0.0 && approx_final_exponent < 0) {
			if IsSignPositive(self) || n&1 == 0 {
				return Inf
			} else {
				return NegInf
			}
		} else if approx_final_exponent < min_representable_exponent {
			// We may want to copy the sign and do the same thing as above,
			// but that seems like an awful amount of work when 99.99999% of people only care
			// about bare zero
			return 0.0
		}
	}

	for {
		if (abs_n & 1) == 1 {
			result *= base
		}

		abs_n >>= 1
		if abs_n == 0 {
			return float32(result)
		}
		base *= base
	}
}

// Returns the reciprocal (inverse) of a number, `1/x`.
func Recip(self float32) float32 {
	x := self
	var sx float32 = 1.
	if x < 0. {
		sx = -1.
	}
	x *= sx
	v := FromBits(0x7EF1_27EA - ToBits(x))
	w := x * v
	v *= 8.0 + w*(-28.0+w*(56.0+w*(-70.0+w*(56.0+w*(-28.0+w*(8.0-w))))))
	return v * sx
}

// Calculates the least non-negative remainder of `self (mod rhs)`.
func RemEuclid(self float32, rhs float32) float32 {
	r := self - Floor(self/rhs)*rhs
	if r >= 0.0 {
		return r
	} else {
		return r + Abs(rhs)
	}
}

// Returns the nearest integer to a number.
func Round(self float32) float32 {
	return float32(int32(self + CopySign(0.5, self)))
}

// Returns a number that represents the sign of `self`.
// /
// * `1.0` if the number is positive, `+0.0` or `INFINITY`
// * `-1.0` if the number is negative, `-0.0` or `NEG_INFINITY`
// * `NAN` if the number is `NAN`
func Sign(self float32) float32 {
	if IsNaN(self) {
		return NaN
	} else {
		return CopySign(1.0, self)
	}
}

func extractExponentBits(self float32) uint32 {
	return (ToBits(self) & expMask) >> mantissaBits
}

func extractExponentValue(self float32) int32 {
	return int32(extractExponentBits(self)) - expBias
}

func setExponent(self float32, exponent int32) float32 {
	without_exponent := ToBits(self) & ^expMask
	only_exponent := uint32(exponent+expBias) << mantissaBits
	return FromBits(without_exponent | only_exponent)
}

func saturatingAdd(a, b int32) int32 {
	c := a + b
	if (c > a) == (b > 0) {
		return c
	}
	return 2147483647
}
