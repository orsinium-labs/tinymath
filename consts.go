package tinymath

import "math"

const (
	// Archimedes' constant (π)
	PI float32 = 3.14159265358979323846264338327950288

	// The full circle constant (τ)
	//
	// Equal to 2π.
	TAU float32 = 6.28318530717958647692528676655900577

	// The golden ratio (φ)
	PHI float32 = 1.618033988749894848204586834365638118

	// The Euler-Mascheroni constant (γ)
	EGAMMA float32 = 0.577215664901532860606512090082402431

	// π/2
	FRAC_PI_2 float32 = 1.57079632679489661923132169163975144

	// π/3
	FRAC_PI_3 float32 = 1.04719755119659774615421446109316763

	// π/4
	FRAC_PI_4 float32 = 0.785398163397448309615660845819875721

	// π/6
	FRAC_PI_6 float32 = 0.52359877559829887307710723054658381

	// π/8
	FRAC_PI_8 float32 = 0.39269908169872415480783042290993786

	// 1/π
	FRAC_1_PI float32 = 0.318309886183790671537767526745028724

	// 1/sqrt(π)
	FRAC_1_SQRT_PI float32 = 0.564189583547756286948079451560772586

	// 2/π
	FRAC_2_PI float32 = 0.636619772367581343075535053490057448

	// 2/sqrt(π)
	FRAC_2_SQRT_PI float32 = 1.12837916709551257389615890312154517

	// sqrt(2)
	SQRT_2 float32 = 1.41421356237309504880168872420969808

	// 1/sqrt(2)
	FRAC_1_SQRT_2 float32 = 0.707106781186547524400844362104849039

	// sqrt(3)
	SQRT_3 float32 = 1.732050807568877293527446341505872367

	// 1/sqrt(3)
	FRAC_1_SQRT_3 float32 = 0.577350269189625764509148780501957456

	// Euler's number (e)
	E float32 = 2.71828182845904523536028747135266250

	// log₂(e)
	LOG2_E float32 = 1.44269504088896340735992468100189214

	// log₂(10)
	LOG2_10 float32 = 3.32192809488736234787031942948939018

	// log₁₀(e)
	LOG10_E float32 = 0.434294481903251827651128918916605082

	// log₁₀(2)
	LOG10_2 float32 = 0.301029995663981195213738894724493027

	// ln(2)
	LN_2 float32 = 0.693147180559945309417232121458176568

	// ln(10)
	LN_10 float32 = 2.30258509299404568401799145468436421

	// [Machine epsilon] value for float32.
	//
	// This is the difference between `1.0` and the next larger representable number.
	//
	// Equal to 2^(1 - MANTISSA_DIGITS).
	//
	// [Machine epsilon]: https://en.wikipedia.org/wiki/Machine_epsilon
	EPSILON float32 = 1.19209290e-07

	// Smallest finite float32 value.
	//
	// Equal to -MAX.
	//
	// [`MAX`]: f32::MAX
	MIN float32 = -3.40282347e+38

	// Smallest positive normal float32 value.
	//
	// Equal to 2^(MIN_EXP - 1).
	MIN_POSITIVE float32 = 1.17549435e-38

	// Largest finite float32 value.
	//
	// Equal to (1 - 2^(-MANTISSA_DIGITS)) 2^MAX_EXP.
	MAX float32 = 3.40282347e+38

	// One greater than the minimum possible normal power of 2 exponent.
	//
	// If n = MIN_EXP, then normal numbers ≥ 0.5 × 2ⁿ.
	MIN_EXP float32 = -125

	// Maximum possible power of 2 exponent.
	//
	// If n = MAX_EXP, then normal numbers < 1 × 2ⁿ.
	MAX_EXP float32 = 128

	// Minimum n for which 10ⁿ is normal.
	//
	// Equal to ceil(log₁₀ MIN_POSITIVE).
	MIN_10_EXP float32 = -37

	// Maximum n for which 10ⁿ is normal.
	//
	// Equal to floor(log₁₀ MAX).
	MAX_10_EXP float32 = 38
)

var (
	NaN    float32 = float32(math.NaN())
	Inf    float32 = float32(math.Inf(1))
	NegInf float32 = float32(math.Inf(-1))
)
