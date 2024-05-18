package tinymath

// Computes `acos(x)` approximation in radians in the range `[0, pi]`.
func Acos(self float32) float32 {
	if self > 0.0 {
		return Atan(Sqrt(1-self*self) / self)
	} else if self == 0.0 {
		return Pi / 2.
	} else {
		return Atan(Sqrt(1-self*self)/self) + Pi
	}
}

// Computes `asin(x)` approximation in radians in the range `[-pi/2, pi/2]`.
func Asin(self float32) float32 {
	return Atan(self * InvSqrt(1-self*self))
}

// Approximates `atan(x)` approximation in radians with a maximum error of
// `0.002`.
//
// Returns [`NAN`] if the number is [`NAN`].
func Atan(self float32) float32 {
	return FracPi2 * AtanNorm(self)
}

// Approximates `atan(x)` normalized to the `[−1,1]` range with a maximum
// error of `0.1620` degrees.
func AtanNorm(self float32) float32 {
	const B = 0.596_227

	// Extract the sign bit
	ux_s := SIGN_MASK & ToBits(self)

	// Calculate the arctangent in the first quadrant
	bx_a := Abs(B * self)
	n := bx_a + self*self
	atan_1q := n / (1.0 + bx_a + n)

	// Restore the sign bit and convert to float
	return FromBits(ux_s | ToBits(atan_1q))
}

// Approximates the four quadrant arctangent of `self` (`y`) and
// `rhs` (`x`) in radians with a maximum error of `0.002`.
//
// - `x = 0`, `y = 0`: `0`
// - `x >= 0`: `arctan(y/x)` -> `[-pi/2, pi/2]`
// - `y >= 0`: `arctan(y/x) + pi` -> `(pi/2, pi]`
// - `y < 0`: `arctan(y/x) - pi` -> `(-pi, -pi/2)`
func Atan2(self float32, rhs float32) float32 {
	n := Atan2Norm(self, rhs)
	if n > 2.0 {
		return Pi/2.0*n - 4.0
	} else {
		return Pi / 2.0 * n
	}
}

// Approximates `atan2(y,x)` normalized to the `[0, 4)` range with a maximum
// error of `0.1620` degrees.
func Atan2Norm(y float32, x float32) float32 {
	const B = 0.596_227

	// Extract sign bits from floating point values
	ux_s := SIGN_MASK & ToBits(x)
	uy_s := SIGN_MASK & ToBits(y)

	// Determine quadrant offset
	q := float32((^ux_s&uy_s)>>29 | ux_s>>30)

	// Calculate arctangent in the first quadrant
	bxy_a := Abs(B * x * y)
	n := bxy_a + y*y
	atan_1q := n / (x*x + bxy_a + n)

	// Translate it to the proper quadrant
	uatan_2q := (ux_s ^ uy_s) | ToBits(atan_1q)
	return q + FromBits(uatan_2q)
}

// Approximates `cos(x)` in radians with a maximum error of `0.002`.
func Cos(self float32) float32 {
	x := self
	x *= Frac1Pi / 2.0
	x -= 0.25 + Floor(x+0.25)
	x *= 16.0 * (Abs(x) - 0.5)
	x += 0.225 * x * (Abs(x) - 1.0)
	return x
}

// Approximates `sin(x)` in radians with a maximum error of `0.002`.
func Sin(self float32) float32 {
	return Cos(self - Pi/2.0)
}

// Simultaneously computes the sine and cosine of the number, `x`.
// Returns `(sin(x), cos(x))`.
func SinCos(self float32) (float32, float32) {
	sin := Cos(self - Pi/2.0)
	cos := Cos(self)
	return sin, cos
}

// Approximates `tan(x)` in radians with a maximum error of `0.6`.
func Tan(self float32) float32 {
	return Sin(self) / Cos(self)
}
