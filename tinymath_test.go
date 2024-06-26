package tinymath_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/orsinium-labs/tinymath"
)

type Case struct {
	Given    float32
	Expected float32
}

type Case2 struct {
	Left     float32
	Right    float32
	Expected float32
}

func eq(t *testing.T, act, exp float32) {
	t.Helper()
	if act != exp {
		t.Fatalf("%f != %f", act, exp)
	}
}

func close(t *testing.T, act, exp float32, eps float32) {
	t.Helper()
	if tinymath.IsNaN(exp) && !tinymath.IsNaN(act) {
		t.Fatalf("%f is not NaN", act)
	}
	delta := tinymath.Abs(act - exp)
	if delta > eps {
		t.Fatalf("%f != %f", act, exp)
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{0.0, 0.0},
		{0.1, 0.1},
		{1.0, 1.0},
		{2.0, 2.0},
		{3.45, 3.45},
		{-0.1, 0.1},
		{-1.0, 1.0},
		{-2.0, 2.0},
		{-3.45, 3.45},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			eq(t, tinymath.Abs(c.Given), c.Expected)
		})
	}
}

func TestCeil(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{-1.1, -1.0},
		{-0.1, 0.0},
		{0.0, 0.0},
		{1.0, 1.0},
		{1.1, 2.0},
		{2.9, 3.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			eq(t, tinymath.Ceil(c.Given), c.Expected)
		})
	}
}

func TestCopySign(t *testing.T) {
	t.Parallel()
	const large = 100_000_000.13425345345
	cases := []Case2{
		{-1.0, -1.0, -1.0},
		{-1.0, 1.0, 1.0},
		{1.0, -1.0, -1.0},
		{1.0, 1.0, 1.0},
		{large, -large, -large},
		{-large, large, large},
		{large, large, large},
		{-large, -large, -large},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
			eq(t, tinymath.CopySign(c.Left, c.Right), c.Expected)
		})
	}
}

func TestDivEuclid(t *testing.T) {
	t.Parallel()
	cases := []Case2{
		{7., 4., 1.},
		{-7., 4., -2.},
		{7., -4., -1.},
		{-7., -4., 2.},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
			eq(t, tinymath.DivEuclid(c.Left, c.Right), c.Expected)
		})
	}
}

func TestExp(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{1e-07, 1.0000001},
		{1e-06, 1.000001},
		{1e-05, 1.00001},
		{1e-04, 1.0001},
		{0.001, 1.0010005},
		{0.01, 1.0100502},
		{0.1, 1.105171},
		{1.0, 2.7182817},
		{10.0, 22026.465},
		{-1e-08, 1.0},
		{-1e-07, 0.9999999},
		{-1e-06, 0.999999},
		{-1e-05, 0.99999},
		{-1e-04, 0.9999},
		{-0.001, 0.9990005},
		{-0.01, 0.99004984},
		{-0.1, 0.9048374},
		{-1.0, 0.36787945},
		{-10.0, 4.539_993e-5},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Exp(c.Given), c.Expected, 0.001*c.Expected)
		})
	}

	for i := float32(-10.); i < 10.; i += .34 {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			close(t, tinymath.Exp(i), float32(math.Exp(float64(i))), i*i*i/i)
		})
	}

}

func TestFloor(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{-1.1, -2.0},
		{-0.1, -1.0},
		{0.0, 0.0},
		{1.0, 1.0},
		{1.1, 1.0},
		{2.9, 2.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			eq(t, tinymath.Floor(c.Given), c.Expected)
		})
	}
}

func TestFract(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{tinymath.Fract(2.9) + 2.0, 2.9},
		{tinymath.Fract(-1.1) - 1.0, -1.1},
		{tinymath.Fract(-0.1), -0.1},
		{tinymath.Fract(0.0), 0.0},
		{tinymath.Fract(1.0) + 1.0, 1.0},
		{tinymath.Fract(1.1) + 1.0, 1.1},
		{tinymath.Fract(-100_000_000.13425345345), 0.0},
		{tinymath.Fract(100_000_000.13425345345), 0.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			eq(t, c.Given, c.Expected)
		})
	}
}

func TestHypot(t *testing.T) {
	t.Parallel()
	cases := []Case2{
		{2., 3., tinymath.Sqrt(13.)},
		{3., 4., tinymath.Sqrt(25.)},
		{12., 7., tinymath.Sqrt(193.)},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
			close(t, tinymath.Hypot(c.Left, c.Right), c.Expected, tinymath.Epsilon)
		})
	}
}

func TestInv(t *testing.T) {
	t.Parallel()
	for i := float32(1.); i < 100.; i += .67 {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			exp := 1.0 / i
			close(t, tinymath.Inv(i), exp, 0.08)
		})
	}
}

func TestInvSqrt(t *testing.T) {
	t.Parallel()
	for i := float32(1.); i < 100.; i++ {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			exp := 1.0 / tinymath.Sqrt(i)
			close(t, tinymath.InvSqrt(i), exp, 0.05)
		})
	}
}

func TestIsNaN(t *testing.T) {
	t.Parallel()
	if !tinymath.IsNaN(tinymath.NaN) {
		t.Fail()
	}
	if tinymath.IsNaN(tinymath.Inf) {
		t.Fail()
	}
	if tinymath.IsNaN(tinymath.NegInf) {
		t.Fail()
	}
	if tinymath.IsNaN(0.0) {
		t.Fail()
	}
	if tinymath.IsNaN(0.1) {
		t.Fail()
	}
	if tinymath.IsNaN(-1.1) {
		t.Fail()
	}
	if tinymath.IsNaN(13.1) {
		t.Fail()
	}
}

func TestIsSignPositive(t *testing.T) {
	t.Parallel()
	if !tinymath.IsSignPositive(tinymath.NaN) {
		t.Fatalf("nan")
	}
	if !tinymath.IsSignPositive(tinymath.Inf) {
		t.Fatalf("inf")
	}
	if tinymath.IsSignPositive(tinymath.NegInf) {
		t.Fatalf("-inf")
	}
	if !tinymath.IsSignPositive(0.0) {
		t.Fatalf("0.0")
	}
	if !tinymath.IsSignPositive(0.1) {
		t.Fatalf("0.1")
	}
	if tinymath.IsSignPositive(-1.1) {
		t.Fatalf("-1.1")
	}
	if !tinymath.IsSignPositive(13.1) {
		t.Fatalf("13.1")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{1.0, 1.0},
		// {2.0, 1.414},
		// {3.0, 1.732},
		{4.0, 2.0},
		{5.0, 2.236},
		// {10.0, 3.162},
		{100.0, 10.0},
		{250.0, 15.811},
		{500.0, 22.36},
		{1000.0, 31.622},
		{2500.0, 50.0},
		{5000.0, 70.710},
		{1000000.0, 1000.0},
		{2500000.0, 1581.138},
		{5000000.0, 2236.067},
		{10000000.0, 3162.277},
		{25000000.0, 5000.0},
		{50000000.0, 7071.067},
		{100000000.0, 10000.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Sqrt(c.Given), c.Expected, 0.005*c.Given)
		})
	}

	for i := float32(1.); i < 100.; i += .34 {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			close(t, tinymath.Sqrt(i), float32(math.Sqrt(float64(i))), 0.05*i)
		})
	}
}

func TestLn(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{1e-20, -46.0517},
		{1e-19, -43.749115},
		{1e-18, -41.446533},
		{1e-17, -39.143948},
		{1e-16, -36.841362},
		{1e-15, -34.538776},
		{1e-14, -32.23619},
		{1e-13, -29.933607},
		{1e-12, -27.631021},
		{1e-11, -25.328436},
		{1e-10, -23.02585},
		{1e-09, -20.723267},
		{1e-08, -18.420681},
		{1e-07, -16.118095},
		{1e-06, -13.815511},
		{1e-05, -11.512925},
		{1e-04, -9.2103405},
		{0.001, -6.9077554},
		{0.01, -4.6051702},
		{0.1, -2.3025851},
		{10.0, 2.3025851},
		{100.0, 4.6051702},
		{1000.0, 6.9077554},
		{10000.0, 9.2103405},
		{100000.0, 11.512925},
		{1000000.0, 13.815511},
		{10000000.0, 16.118095},
		{100000000.0, 18.420681},
		{1000000000.0, 20.723267},
		{10000000000.0, 23.02585},
		{100000000000.0, 25.328436},
		{1000000000000.0, 27.631021},
		{10000000000000.0, 29.933607},
		{100000000000000.0, 32.23619},
		{1000000000000000.0, 34.538776},
		{1e+16, 36.841362},
		{1e+17, 39.143948},
		{1e+18, 41.446533},
		{1e+19, 43.749115},
	}
	t.Run("table", func(t *testing.T) {
		t.Parallel()
		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
				act := tinymath.Ln(c.Given)
				delta := tinymath.Abs(act - c.Expected)
				if delta/c.Expected > 0.001 {
					t.Fatalf("%f != %f", act, c.Expected)
				}
			})
		}
	})

	t.Run("stdlib", func(t *testing.T) {
		t.Parallel()
		for i := float32(1.); i < 100.; i += .34 {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Ln(i), float32(math.Log(float64(i))), 0.001)
			})
		}
	})
}

func TestLog(t *testing.T) {
	t.Parallel()
	cases := []Case2{
		{1e-20, 3, -41.918_064},
		{1e-19, 3, -39.822_16},
		{1e-18, 3, -37.726_26},
		{1e-17, 3, -35.630_356},
		{1e-16, 3, -33.534_454},
		{1e-15, 3, -31.438_549},
		{1e-14, 3, -29.342_646},
		{1e-13, 3, -27.246_744},
		{1e-12, 3, -25.150_839},
		{1e-11, 3, -23.054_935},
		{1e-10, 3, -20.959_032},
		{1e-09, 3, -18.863_13},
		{1e-08, 3, -16.767_227},
		{1e-07, 3, -14.671_323},
		{1e-06, 3, -12.575_419},
		{1e-05, 3, -10.479_516},
		{1e-04, 3, -8.383_614},
		{0.001, 3, -6.287_709_7},
		{0.01, 3, -4.191_807},
		{0.1, 3, -2.095_903_4},
		{10.0, 3, 2.095_903_4},
		{100.0, 3, 4.191_807},
		{1000.0, 3, 6.287_709_7},
		{10000.0, 3, 8.383_614},
		{100000.0, 3, 10.479_516},
		{1000000.0, 3, 12.575_419},
		{10000000.0, 3, 14.671_323},
		{100000000.0, 3, 16.767_227},
		{1000000000.0, 3, 18.863_13},
		{10000000000.0, 3, 20.959_032},
		{100000000000.0, 3, 23.054_935},
		{1000000000000.0, 3, 25.150_839},
		{10000000000000.0, 3, 27.246_744},
		{100000000000000.0, 3, 29.342_646},
		{1000000000000000.0, 3, 31.438_549},
		{1e+16, 3, 33.534_454},
		{1e+17, 3, 35.630_356},
		{1e+18, 3, 37.726_26},
		{1e+19, 3, 39.822_16},

		{1e-20, 5.5, -27.013_786},
		{1e-19, 5.5, -25.663_097},
		{1e-18, 5.5, -24.312_408},
		{1e-17, 5.5, -22.961_72},
		{1e-16, 5.5, -21.611_03},
		{1e-15, 5.5, -20.260_34},
		{1e-14, 5.5, -18.909_65},
		{1e-13, 5.5, -17.558_962},
		{1e-12, 5.5, -16.208_273},
		{1e-11, 5.5, -14.857_583},
		{1e-10, 5.5, -13.506_893},
		{1e-09, 5.5, -12.156_204},
		{1e-08, 5.5, -10.805_515},
		{1e-07, 5.5, -9.454_825},
		{1e-06, 5.5, -8.104_136},
		{1e-05, 5.5, -6.753_446_6},
		{1e-04, 5.5, -5.402_757_6},
		{0.001, 5.5, -4.052_068},
		{0.01, 5.5, -2.701_378_8},
		{0.1, 5.5, -1.350_689_4},
		{10.0, 5.5, 1.350_689_4},
		{100.0, 5.5, 2.701_378_8},
		{1000.0, 5.5, 4.052_068},
		{10000.0, 5.5, 5.402_757_6},
		{100000.0, 5.5, 6.753_446_6},
		{1000000.0, 5.5, 8.104_136},
		{10000000.0, 5.5, 9.454_825},
		{100000000.0, 5.5, 10.805_515},
		{1000000000.0, 5.5, 12.156_204},
		{10000000000.0, 5.5, 13.506_893},
		{100000000000.0, 5.5, 14.857_583},
		{1000000000000.0, 5.5, 16.208_273},
		{10000000000000.0, 5.5, 17.558_962},
		{100000000000000.0, 5.5, 18.909_65},
		{1000000000000000.0, 5.5, 20.260_34},
		{1e+16, 5.5, 21.611_03},
		{1e+17, 5.5, 22.961_72},
		{1e+18, 5.5, 24.312_408},
		{1e+19, 5.5, 25.663_097},

		{1e-20, 12.7, -18.119_164},
		{1e-19, 12.7, -17.213_205},
		{1e-18, 12.7, -16.307_247},
		{1e-17, 12.7, -15.401_289},
		{1e-16, 12.7, -14.495_331},
		{1e-15, 12.7, -13.589_373},
		{1e-14, 12.7, -12.683_414},
		{1e-13, 12.7, -11.777_456},
		{1e-12, 12.7, -10.871_498},
		{1e-11, 12.7, -9.965_54},
		{1e-10, 12.7, -9.059_582},
		{1e-09, 12.7, -8.153_624},
		{1e-08, 12.7, -7.247_665_4},
		{1e-07, 12.7, -6.341_707},
		{1e-06, 12.7, -5.435_749},
		{1e-05, 12.7, -4.529_791},
		{1e-04, 12.7, -3.623_832_7},
		{0.001, 12.7, -2.717_874_5},
		{0.01, 12.7, -1.811_916_4},
		{0.1, 12.7, -0.905_958_2},
		{10.0, 12.7, 0.905_958_2},
		{100.0, 12.7, 1.811_916_4},
		{1000.0, 12.7, 2.717_874_5},
		{10000.0, 12.7, 3.623_832_7},
		{100000.0, 12.7, 4.529_791},
		{1000000.0, 12.7, 5.435_749},
		{10000000.0, 12.7, 6.341_707},
		{100000000.0, 12.7, 7.247_665_4},
		{1000000000.0, 12.7, 8.153_624},
		{10000000000.0, 12.7, 9.059_582},
		{100000000000.0, 12.7, 9.965_54},
		{1000000000000.0, 12.7, 10.871_498},
		{10000000000000.0, 12.7, 11.777_456},
		{100000000000000.0, 12.7, 12.683_414},
		{1000000000000000.0, 12.7, 13.589_373},
		{1e+16, 12.7, 14.495_331},
		{1e+17, 12.7, 15.401_289},
		{1e+18, 12.7, 16.307_247},
		{1e+19, 12.7, 17.213_205},
	}
	t.Run("table", func(t *testing.T) {
		t.Parallel()
		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
				act := tinymath.Log(c.Left, c.Right)
				delta := tinymath.Abs(act - c.Expected)
				if delta/c.Expected > 0.001 {
					t.Fatalf("%f != %f", act, c.Expected)
				}
			})
		}
	})

	t.Run("stdlib", func(t *testing.T) {
		t.Parallel()
		for i := float32(1.); i < 100.; i += .34 {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Log(i, 10.), float32(math.Log10(float64(i))), 0.001)
			})
		}
	})
}

func TestLog2(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{1e-20, -66.43856},
		{1e-19, -63.116634},
		{1e-18, -59.794704},
		{1e-17, -56.47278},
		{1e-16, -53.15085},
		{1e-15, -49.828922},
		{1e-14, -46.506992},
		{1e-13, -43.185066},
		{1e-12, -39.863136},
		{1e-11, -36.54121},
		{1e-10, -33.21928},
		{1e-09, -29.897352},
		{1e-08, -26.575424},
		{1e-07, -23.253496},
		{1e-06, -19.931568},
		{1e-05, -16.60964},
		{1e-04, -13.287712},
		{0.001, -9.965784},
		{0.01, -6.643856},
		{0.1, -3.321928},
		{10.0, 3.321928},
		{100.0, 6.643856},
		{1000.0, 9.965784},
		{10000.0, 13.287712},
		{100000.0, 16.60964},
		{1000000.0, 19.931568},
		{10000000.0, 23.253496},
		{100000000.0, 26.575424},
		{1000000000.0, 29.897352},
		{10000000000.0, 33.21928},
		{100000000000.0, 36.54121},
		{1000000000000.0, 39.863136},
		{10000000000000.0, 43.185066},
		{100000000000000.0, 46.506992},
		{1000000000000000.0, 49.828922},
		{1e+16, 53.15085},
		{1e+17, 56.47278},
		{1e+18, 59.794704},
		{1e+19, 63.116634},
	}
	t.Run("table", func(t *testing.T) {
		t.Parallel()
		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
				act := tinymath.Log2(c.Given)
				delta := tinymath.Abs(act - c.Expected)
				if delta/c.Expected > 0.001 {
					t.Fatalf("%f != %f", act, c.Expected)
				}
			})
		}
	})

	t.Run("stdlib", func(t *testing.T) {
		t.Parallel()
		for i := float32(1.); i < 100.; i += .34 {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Log2(i), float32(math.Log2(float64(i))), 0.001)
			})
		}
	})
}

func TestLog10(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{1e-20, -20.0},
		{1e-19, -19.0},
		{1e-18, -18.0},
		{1e-17, -17.0},
		{1e-16, -16.0},
		{1e-15, -15.0},
		{1e-14, -14.0},
		{1e-13, -13.0},
		{1e-12, -12.0},
		{1e-11, -11.0},
		{1e-10, -10.0},
		{1e-09, -9.0},
		{1e-08, -8.0},
		{1e-07, -7.0},
		{1e-06, -6.0},
		{1e-05, -5.0},
		{1e-04, -4.0},
		{0.001, -3.0},
		{0.01, -2.0},
		{0.1, -1.0},
		{10.0, 1.0},
		{100.0, 2.0},
		{1000.0, 3.0},
		{10000.0, 4.0},
		{100000.0, 5.0},
		{1000000.0, 6.0},
		{10000000.0, 7.0},
		{100000000.0, 8.0},
		{1000000000.0, 9.0},
		{10000000000.0, 10.0},
		{100000000000.0, 11.0},
		{1000000000000.0, 12.0},
		{10000000000000.0, 13.0},
		{100000000000000.0, 14.0},
		{1000000000000000.0, 15.0},
		{1e+16, 16.0},
		{1e+17, 17.0},
		{1e+18, 18.0},
		{1e+19, 19.0},
	}
	t.Run("table", func(t *testing.T) {
		t.Parallel()
		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
				act := tinymath.Log10(c.Given)
				delta := tinymath.Abs(act - c.Expected)
				if delta/c.Expected > 0.001 {
					t.Fatalf("%f != %f", act, c.Expected)
				}
			})
		}
	})

	t.Run("stdlib", func(t *testing.T) {
		t.Parallel()
		for i := float32(1.); i < 100.; i += .34 {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Log10(i), float32(math.Log10(float64(i))), 0.001)
			})
		}
	})
}

func TestPowF(t *testing.T) {
	t.Parallel()
	cases := []Case2{
		{-1e-20, 3, 1.0},
		{-1e-19, 3, 1.0},
		{-1e-18, 3, 1.0},
		{-1e-17, 3, 1.0},
		{-1e-16, 3, 0.9999999999999999},
		{-1e-15, 3, 0.9999999999999989},
		{-1e-14, 3, 0.999999999999989},
		{-1e-13, 3, 0.9999999999998901},
		{-1e-12, 3, 0.9999999999989014},
		{-1e-11, 3, 0.9999999999890139},
		{-1e-10, 3, 0.9999999998901388},
		{-1e-09, 3, 0.9999999989013877},
		{-1e-08, 3, 0.9999999890138772},
		{-1e-07, 3, 0.999_999_9},
		{-1e-06, 3, 0.999_998_9},
		{-1e-05, 3, 0.999_989_03},
		{-1e-04, 3, 0.999_890_15},
		{-0.001, 3, 0.998_901_96},
		{-0.01, 3, 0.989_074},
		{-0.1, 3, 0.895_958_5},
		{-1.0, 3, 0.333_333_34},
		{-10.0, 3, 1.693_508_8e-5},
		{-100.0, 3, 0e0},
		{-1000.0, 3, 0.0},
		{1e-20, 3, 1.0},
		{1e-19, 3, 1.0},
		{1e-18, 3, 1.0},
		{1e-17, 3, 1.0},
		{1e-16, 3, 1.0},
		{1e-15, 3, 1.000000000000001},
		{1e-14, 3, 1.0000000000000109},
		{1e-13, 3, 1.00000000000011},
		{1e-12, 3, 1.0000000000010987},
		{1e-11, 3, 1.000000000010986},
		{1e-10, 3, 1.0000000001098612},
		{1e-09, 3, 1.0000000010986123},
		{1e-08, 3, 1.000000010986123},
		{1e-07, 3, 1.000_000_1},
		{1e-06, 3, 1.000_001_1},
		{1e-05, 3, 1.000_011},
		{1e-04, 3, 1.000_109_9},
		{0.001, 3, 1.001_099_2},
		{0.01, 3, 1.011_046_6},
		{0.1, 3, 1.116_123_2},
		{1.0, 3, 3.0},
		{10.0, 3, 59049.0},

		{-1e-20, 150, 1.0},
		{-1e-19, 150, 1.0},
		{-1e-18, 150, 1.0},
		{-1e-17, 150, 1.0},
		{-1e-16, 150, 0.9999999999999994},
		{-1e-15, 150, 0.999999999999995},
		{-1e-14, 150, 0.9999999999999499},
		{-1e-13, 150, 0.999999999999499},
		{-1e-12, 150, 0.9999999999949893},
		{-1e-11, 150, 0.9999999999498936},
		{-1e-10, 150, 0.9999999994989365},
		{-1e-09, 150, 0.9999999949893649},
		{-1e-08, 150, 0.999_999_94},
		{-1e-07, 150, 0.999_999_5},
		{-1e-06, 150, 0.999_995},
		{-1e-05, 150, 0.999_949_9},
		{-1e-04, 150, 0.999_499_1},
		{-0.001, 150, 0.995_001_9},
		{-0.01, 150, 0.951_128_24},
		{-0.1, 150, 0.605_885_9},
		{-1.0, 150, 0.006_666_667},
		{-10.0, 150, 1.734_153e-22},
		{-100.0, 150, 0e0},
		{-1000.0, 150, 0.0},
		{-10000.0, 150, 0.0},
		{-100000.0, 150, 0.0},
		{-1000000.0, 150, 0.0},
		{-10000000.0, 150, 0.0},
		{-100000000.0, 150, 0.0},
		{-1000000000.0, 150, 0.0},
		{-10000000000.0, 150, 0.0},
		{-100000000000.0, 150, 0.0},
		{-1000000000000.0, 150, 0.0},
		{-10000000000000.0, 150, 0.0},
		{-100000000000000.0, 150, 0.0},
		{-1000000000000000.0, 150, 0.0},
		{-1e+16, 150, 0.0},
		{-1e+17, 150, 0.0},
		{-1e+18, 150, 0.0},
		{-1e+19, 150, 0.0},
		{1e-20, 150, 1.0},
		{1e-19, 150, 1.0},
		{1e-18, 150, 1.0},
		{1e-17, 150, 1.0},
		{1e-16, 150, 1.0000000000000004},
		{1e-15, 150, 1.000000000000005},
		{1e-14, 150, 1.0000000000000502},
		{1e-13, 150, 1.0000000000005012},
		{1e-12, 150, 1.0000000000050107},
		{1e-11, 150, 1.0000000000501064},
		{1e-10, 150, 1.0000000005010636},
		{1e-09, 150, 1.0000000050106352},
		{1e-08, 150, 1.000000050106354},
		{1e-07, 150, 1.000_000_5},
		{1e-06, 150, 1.000_005},
		{1e-05, 150, 1.000_050_1},
		{1e-04, 150, 1.000_501_2},
		{0.001, 150, 1.005_023_2},
		{0.01, 150, 1.051_382_9},
		{0.1, 150, 1.650_475_6},
		{1.0, 150, 150.0},
		{10.0, 150, 5.766_504e21},

		{2.0, -0.5881598, 0.345_931_95},
		{3.2, -0.5881598, tinymath.NaN},
		{3.0, -0.5881598, -0.203_463_27},
		{4.0, -1000000.0, 1e+24},
	}
	t.Run("table", func(t *testing.T) {
		t.Parallel()
		for _, c := range cases {
			c := c
			t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
				act := tinymath.PowF(c.Right, c.Left)
				delta := tinymath.Abs(act - c.Expected)
				if delta/c.Expected > 0.01 {
					t.Fatalf("%f != %f", act, c.Expected)
				}
			})
		}
	})

	t.Run("stdlib", func(t *testing.T) {
		t.Parallel()
		for i := float32(1.); i < 100.; i += .34 {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.PowF(i, 3.5), float32(math.Pow(float64(i), 3.5)), i*i*i/15)
			})
		}
	})
}

func TestPowI(t *testing.T) {
	t.Parallel()
	for i := int32(1); i < 10; i++ {
		for f := float32(-3.); f < 7.; f += .5 {
			f := f
			i := i
			if f == 6.5 {
				continue
			}
			t.Run(fmt.Sprintf("%f", f), func(t *testing.T) {
				close(t, tinymath.PowI(f, i), float32(math.Pow(float64(f), float64(i))), 0.001)
			})
		}
	}
}

func TestRecip(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{0.00001, 100000.0},
		{1.0, 1.0},
		{2.0, 0.5},
		{0.25, 4.0},
		{-0.5, -2.0},
		{tinymath.Pi, 1.0 / tinymath.Pi},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			act := tinymath.Recip(c.Given)
			delta := tinymath.Abs(act - c.Expected)
			if delta/c.Expected > 1e-5 {
				t.Fatalf("%f != %f", act, c.Expected)
			}
		})
	}
}

func TestRemEuclid(t *testing.T) {
	t.Parallel()
	cases := []Case2{
		{7, 4, 3},
		{-7, 4, 1},
		{7, -4, 3},
		{-7, -4, 1},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
			act := tinymath.RemEuclid(c.Left, c.Right)
			eq(t, act, c.Expected)
		})
	}
}

func TestRound(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{0.0, 0.0},
		{0.49999, 0.0},
		{-0.49999, 0.0},
		{0.5, 1.0},
		{-0.5, -1.0},
		{9999.499, 9999.0},
		{-9999.499, -9999.0},
		{9999.5, 10000.0},
		{-9999.5, -10000.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			act := tinymath.Round(c.Given)
			eq(t, act, c.Expected)
		})
	}

	for i := float32(-20.); i < 20.; i += .34 {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			eq(t, tinymath.Round(i), float32(math.Round(float64(i))))
		})
	}
}

func TestSign(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{tinymath.Inf, 1.0},
		{0.0, 1.0},
		{1.0, 1.0},
		{tinymath.NegInf, -1.0},
		{-1.0, -1.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			act := tinymath.Sign(c.Given)
			eq(t, act, c.Expected)
		})
	}
}

func TestTrunc(t *testing.T) {
	t.Parallel()
	cases := []Case{
		{-1.1, -1.0},
		{-0.1, 0.0},
		{0.0, 0.0},
		{1.0, 1.0},
		{1.1, 1.0},
		{2.9, 2.0},
		{-100_000_000.13425345345, -100_000_000.0},
		{100_000_000.13425345345, 100_000_000.0},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			act := tinymath.Trunc(c.Given)
			eq(t, act, c.Expected)
		})
	}

	for i := float32(-20.); i < 20.; i += .34 {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			eq(t, tinymath.Trunc(i), float32(math.Trunc(float64(i))))
		})
	}
}
