package tinymath

import "math"

const (
	SIGN_MASK uint32  = 0x8000_0000
	EPSILON   float32 = 1.19209290e-07
	E         float32 = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
	PI        float32 = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
	FRAC_1_PI float32 = 0.318309886183790671537767526745028724
	FRAC_PI_2 float32 = 1.57079632679489661923132169163975144
	LN_2      float32 = 0.693147180559945309417232121458176568

	MANTISSA_BITS        = 23
	EXPONENT_MASK uint32 = 0b0111_1111_1000_0000_0000_0000_0000_0000
	EXPONENT_BIAS        = 127
)

var (
	NaN    float32 = float32(math.NaN())
	Inf    float32 = float32(math.Inf(1))
	NegInf float32 = float32(math.Inf(-1))
)

func ToBits(x float32) uint32 {
	return math.Float32bits(x)
}

func FromBits(x uint32) float32 {
	return math.Float32frombits(x)
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
	return FromBits(ToBits(self) & ^SIGN_MASK)
}

// Returns the smallest integer greater than or equal to a number.
func Ceil(self float32) float32 {
	return -Floor(-self)
}

// Returns a number composed of the magnitude of `self` and the sign of
// `sign`.
func CopySign(self float32, sign float32) float32 {
	source_bits := ToBits(sign)
	source_sign := source_bits & SIGN_MASK
	signless_destination_bits := ToBits(self) & ^SIGN_MASK
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
	const LOG2_E = 1.44269504088896340735992468100189214

	if self == 0.0 {
		return 1
	}
	if Abs(self-1) < EPSILON {
		return E
	}
	if Abs(self-(-1)) < EPSILON {
		return 1. / E
	}

	// log base 2(E) == 1/ln(2)
	// x_fract + x_whole = x/ln2_recip
	// ln2*(x_fract + x_whole) = x
	x_ln2recip := self * LOG2_E
	x_fract := Fract(x_ln2recip)
	x_trunc := Trunc(x_ln2recip)

	//guaranteed to be 0 < x < 1.0
	x_fract = x_fract * LN_2
	fract_exp := ExpSmallX(x_fract, partial_iter)

	//need the 2^n portion, we can just extract that from the whole number exp portion
	fract_exponent := saturatingAdd(extractExponentValue(fract_exp), int32(x_trunc))

	if fract_exponent < -(EXPONENT_BIAS) {
		return 0.0
	}

	if fract_exponent > (EXPONENT_BIAS + 1) {
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

// Returns the largest integer less than or equal to a number.
func Floor(self float32) float32 {
	res := float32(int32(self))
	if self < res {
		res -= 1.0
	}
	return float32(res)
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
	exponent_shift := (leadingZeros(fractional_part) - (32 - MANTISSA_BITS)) + 1

	fractional_normalized := (fractional_part << exponent_shift) & MANTISSA_MASK

	new_exponent_bits := (EXPONENT_BIAS - (exponent_shift)) << MANTISSA_BITS

	return CopySign(FromBits(fractional_normalized|new_exponent_bits), self)
}

func leadingZeros(x uint32) uint32 {
	panic("todo")
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
	return Abs(x) == x
}

// Approximates the natural logarithm of the number.
// Note: excessive precision ignored because it hides the origin of the numbers used for the
// ln(1.0->2.0) polynomial
func Ln(self float32) float32 {

	// x may essentially be 1.0 but, as clippy notes, these kinds of
	// floating point comparisons can fail when the bit pattern is not the sames
	if Abs(self-1) < EPSILON {
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
	divisor := FromBits(ToBits(x_working) & EXPONENT_MASK)

	// supposedly normalizing between 1.0 and 2.0
	x_working = x_working / divisor

	// approximate polynomial generated from maple in the post using Remez Algorithm:
	// https://en.wikipedia.org/wiki/Remez_algorithm
	ln_1to2_polynomial := -1.741_793_9 + (2.821_202_6+(-1.469_956_8+(0.447_179_55-0.056_570_851*x_working)*x_working)*x_working)*x_working

	// ln(2) * n + ln(y)
	result := float32(base2_exponent)*LN_2 + ln_1to2_polynomial

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
	const LOG10_E = 0.434294481903251827651128918916605082
	return Ln(self) * LOG10_E
}

// Approximates the base 2 logarithm of the number.
func Log2(self float32) float32 {
	const LOG2_E = 1.44269504088896340735992468100189214
	return Ln(self) * LOG2_E
}

// Computes `(self * a) + b`.
func MulAdd(self float32, a float32, b float32) float32 {
	return self*a + b
}

// Approximates a number raised to a floating point power.
func Powf(self float32, n float32) float32 {
	// using x^n = exp(ln(x^n)) = exp(n*ln(x))
	if self >= 0.0 {
		return Exp(n * Ln(self))
	} else if IsInteger(n) {
		return NaN
	} else if IsEven(n) {
		// if n is even, then we know that the result will have no sign, so we can remove it
		return n * Exp(Ln(withoutSign(self)))
	} else {
		// if n isn't even, we need to multiply by -1.0 at the end.
		return -(n * Exp(Ln(withoutSign(self))))
	}
}

// Approximates a number raised to an integer power.
func Powi(self float32, n int32) float32 {
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
		const min_representable_exponent = -126 - MANTISSA_BITS
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
func Signum(self float32) float32 {
	if IsNaN(self) {
		return NaN
	} else {
		return CopySign(1.0, self)
	}
}

// Approximates the square root of a number with an average deviation of ~5%.
//
// Returns [`NAN`] if `self` is a negative number.
func Sqrt(self float32) float32 {
	if self >= 0.0 {
		return FromBits((ToBits(self) + 0x3f80_0000) >> 1)
	} else {
		return NaN
	}
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

func extractExponentBits(self float32) uint32 {
	return (ToBits(self) & EXPONENT_MASK) >> MANTISSA_BITS
}

func extractExponentValue(self float32) int32 {
	return int32(extractExponentBits(self)) - EXPONENT_BIAS
}

func setExponent(self float32, exponent int32) float32 {
	without_exponent := ToBits(self) & ^EXPONENT_MASK
	only_exponent := uint32(exponent+EXPONENT_BIAS) << MANTISSA_BITS
	return FromBits(without_exponent | only_exponent)
}

func saturatingAdd(a, b int32) int32 {
	c := a + b
	if (c > a) == (b > 0) {
		return c
	}
	return 2147483647
}

func withoutSign(self float32) float32 {
	return FromBits(ToBits(self) & ^SIGN_MASK)
}
