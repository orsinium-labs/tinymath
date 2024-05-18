package tinymath

// Constatnts from both Go and Rust stdlib typed as float32.
const (
	// Archimedes' constant (π)
	Pi float32 = 3.14159265358979323846264338327950288

	// The full circle constant (τ)
	//
	// Equal to 2π.
	Tau float32 = 6.28318530717958647692528676655900577

	// The golden ratio (φ)
	Phi float32 = 1.618033988749894848204586834365638118

	// The Euler-Mascheroni constant (γ)
	EGamma float32 = 0.577215664901532860606512090082402431

	// π/2
	FracPi2 float32 = 1.57079632679489661923132169163975144

	// π/3
	FracPi3 float32 = 1.04719755119659774615421446109316763

	// π/4
	FracPi4 float32 = 0.785398163397448309615660845819875721

	// π/6
	FracPi6 float32 = 0.52359877559829887307710723054658381

	// π/8
	FracPi8 float32 = 0.39269908169872415480783042290993786

	// 1/π
	Frac1Pi float32 = 0.318309886183790671537767526745028724

	// 1/sqrt(π)
	Frac1SqrtPi float32 = 0.564189583547756286948079451560772586

	// 2/π
	Frac2Pi float32 = 0.636619772367581343075535053490057448

	// 2/sqrt(π)
	Frac2SqrtPi float32 = 1.12837916709551257389615890312154517

	// sqrt(2)
	Sqrt2 float32 = 1.41421356237309504880168872420969808

	// 1/sqrt(2)
	Frac1Sqrt2 float32 = 0.707106781186547524400844362104849039

	// sqrt(3)
	Sqrt3 float32 = 1.732050807568877293527446341505872367

	SqrtE   float32 = 1.64872127070012814684865078781416357165377610071014801157507931
	SqrtPi  float32 = 1.77245385090551602729816748334114518279754945612238712821380779
	SqrtPhi float32 = 1.27201964951406896425242246173749149171560804184009624861664038

	// 1/sqrt(3)
	Frac1Sqrt3 float32 = 0.577350269189625764509148780501957456

	// Euler's number (e)
	E float32 = 2.71828182845904523536028747135266250

	// log₂(e)
	Log2E float32 = 1.44269504088896340735992468100189214

	// log₂(10)
	Log210 float32 = 3.32192809488736234787031942948939018

	// log₁₀(e)
	Log10E float32 = 0.434294481903251827651128918916605082

	// log₁₀(2)
	Log102 float32 = 0.301029995663981195213738894724493027

	// ln(2)
	Ln2 float32 = 0.693147180559945309417232121458176568

	// ln(10)
	Ln10 float32 = 2.30258509299404568401799145468436421

	// [Machine epsilon] value for float32.
	//
	// This is the difference between `1.0` and the next larger representable number.
	//
	// Equal to 2^(1 - MANTISSA_DIGITS).
	//
	// [Machine epsilon]: https://en.wikipedia.org/wiki/Machine_epsilon
	Epsilon float32 = 1.19209290e-07

	// Smallest finite float32 value.
	//
	// Equal to -MAX.
	//
	// [`MAX`]: f32::MAX
	MinNeg float32 = -3.40282347e+38

	// Smallest positive normal float32 value.
	//
	// Equal to 2^(MIN_EXP - 1).
	MinPos float32 = 0x1p-126 * 0x1p-23

	// Largest finite float32 value.
	//
	// Equal to (1 - 2^(-MANTISSA_DIGITS)) 2^MAX_EXP.
	MaxPos float32 = 0x1p127 * (1 + (1 - 0x1p-23))

	// One greater than the minimum possible normal power of 2 exponent.
	//
	// If n = MinExp, then normal numbers ≥ 0.5 × 2ⁿ.
	MinExp float32 = -125

	// Maximum possible power of 2 exponent.
	//
	// If n = MaxExp, then normal numbers < 1 × 2ⁿ.
	MaxExp float32 = 128

	// Minimum n for which 10ⁿ is normal.
	//
	// Equal to ceil(log₁₀ MIN_POSITIVE).
	Min10Exp float32 = -37

	// Maximum n for which 10ⁿ is normal.
	//
	// Equal to floor(log₁₀ MAX).
	Max10Exp float32 = 38
)

var (
	// Not a number
	NaN float32 = FromBits(0x7fc00000)

	// Positive infinity
	Inf float32 = FromBits(0x7f800000)

	// Negative infininty
	NegInf float32 = FromBits(0xff800000)
)
