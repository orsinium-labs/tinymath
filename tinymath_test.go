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
}

func TestFloor(t *testing.T) {
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
	cases := []Case2{
		{2., 3., tinymath.Sqrt(13.)},
		{3., 4., tinymath.Sqrt(25.)},
		{12., 7., tinymath.Sqrt(193.)},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f_%f", c.Left, c.Right), func(t *testing.T) {
			close(t, tinymath.Hypot(c.Left, c.Right), c.Expected, tinymath.EPSILON)
		})
	}
}

func TestInv(t *testing.T) {
	for i := float32(1.); i < 100.; i++ {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			exp := 1.0 / i
			close(t, tinymath.Inv(i), exp, 0.08)
		})
	}
}

func TestInvSqrt(t *testing.T) {
	for i := float32(1.); i < 100.; i++ {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			exp := 1.0 / tinymath.Sqrt(i)
			close(t, tinymath.InvSqrt(i), exp, 0.05)
		})
	}
}

func TestSqrt(t *testing.T) {
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

	for i := float32(1.); i < 100.; i++ {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			close(t, tinymath.Sqrt(i), float32(math.Sqrt(float64(i))), 0.05*i)
		})
	}
}

func TestLn(t *testing.T) {
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
		for i := float32(1.); i < 100.; i++ {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Ln(i), float32(math.Log(float64(i))), 0.001)
			})
		}
	})
}

func TestLog(t *testing.T) {
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
		for i := float32(1.); i < 100.; i++ {
			i := i
			t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
				close(t, tinymath.Log(i, 10.), float32(math.Log10(float64(i))), 0.001)
			})
		}
	})
}
