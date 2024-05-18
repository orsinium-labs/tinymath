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
