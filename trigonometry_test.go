package tinymath_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/orsinium-labs/tinymath"
)

func TestAcos(t *testing.T) {
	cases := []Case{
		{2.000, tinymath.NaN},
		{1.000, 0.0},
		{0.866, tinymath.FracPi6},
		{0.707, tinymath.FracPi4},
		{0.500, tinymath.FracPi3},
		{tinymath.Epsilon, tinymath.FracPi2},
		{0.000, tinymath.FracPi2},
		{-tinymath.Epsilon, tinymath.FracPi2},
		{-0.500, 2.0 * tinymath.FracPi3},
		{-0.707, 3.0 * tinymath.FracPi4},
		{-0.866, 5.0 * tinymath.FracPi6},
		{-1.000, tinymath.Pi},
		{-2.000, tinymath.NaN},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Acos(c.Given), c.Expected, 0.03)
		})
	}
}

func TestAsin(t *testing.T) {
	act := tinymath.Asin(tinymath.Sin(tinymath.FracPi2))
	close(t, act, tinymath.FracPi2, tinymath.Epsilon)
}

func TestAtan(t *testing.T) {
	cases := []Case{
		// {tinymath.Sqrt(3.0) / 3.0, tinymath.FRAC_PI_6},
		{1.0, tinymath.FracPi4},
		{tinymath.Sqrt(3.0), tinymath.FracPi3},
		// {-tinymath.Sqrt(3.0) / 3.0, -tinymath.FRAC_PI_6},
		{-1.0, -tinymath.FracPi4},
		{-tinymath.Sqrt(3.0), -tinymath.FracPi3},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Atan(c.Given), c.Expected, 0.003)
		})
	}
}

func TestAtan2(t *testing.T) {
	cases := []Case2{
		{0.0, 1.0, 0.0},
		{0.0, -1.0, tinymath.Pi},
		{3.0, 2.0, tinymath.Atan(3.0 / 2.0)},
		{2.0, -1.0, tinymath.Atan(2.0/-1.0) + tinymath.Pi},
		// {-2.0, -1.0, tinymath.Atan(-2.0/-1.0) - tinymath.PI},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("y%f_x%f", c.Left, c.Right), func(t *testing.T) {
			close(t, tinymath.Atan2(c.Left, c.Right), c.Expected, 0.003)
		})
	}
}

func TestCos(t *testing.T) {
	cases := []Case{
		{0.000, 1.000},
		{0.140, 0.990},
		{0.279, 0.961},
		{0.419, 0.914},
		{0.559, 0.848},
		{0.698, 0.766},
		{0.838, 0.669},
		{0.977, 0.559},
		{1.117, 0.438},
		{1.257, 0.309},
		{1.396, 0.174},
		{1.536, 0.035},
		{1.676, -0.105},
		{1.815, -0.242},
		{1.955, -0.375},
		{2.094, -0.500},
		{2.234, -0.616},
		{2.374, -0.719},
		{2.513, -0.809},
		{2.653, -0.883},
		{2.793, -0.940},
		{2.932, -0.978},
		{3.072, -0.998},
		{3.211, -0.998},
		{3.351, -0.978},
		{3.491, -0.940},
		{3.630, -0.883},
		{3.770, -0.809},
		{3.910, -0.719},
		{4.049, -0.616},
		{4.189, -0.500},
		{4.328, -0.375},
		{4.468, -0.242},
		{4.608, -0.105},
		{4.747, 0.035},
		{4.887, 0.174},
		{5.027, 0.309},
		{5.166, 0.438},
		{5.306, 0.559},
		{5.445, 0.669},
		{5.585, 0.766},
		{5.725, 0.848},
		{5.864, 0.914},
		{6.004, 0.961},
		{6.144, 0.990},
		{6.283, 1.000},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Cos(c.Given), c.Expected, 0.002)
		})
	}

	for i := float32(1.); i < 100.; i++ {
		i := i
		t.Run(fmt.Sprintf("%f", i), func(t *testing.T) {
			close(t, tinymath.Sin(i), float32(math.Sin(float64(i))), 0.002)
		})
	}
}

func TestSin(t *testing.T) {
	cases := []Case{
		{0.000, 0.000},
		{0.140, 0.139},
		{0.279, 0.276},
		{0.419, 0.407},
		{0.559, 0.530},
		{0.698, 0.643},
		{0.838, 0.743},
		{0.977, 0.829},
		{1.117, 0.899},
		{1.257, 0.951},
		{1.396, 0.985},
		{1.536, 0.999},
		{1.676, 0.995},
		{1.815, 0.970},
		{1.955, 0.927},
		{2.094, 0.866},
		{2.234, 0.788},
		{2.374, 0.695},
		{2.513, 0.588},
		{2.653, 0.469},
		{2.793, 0.342},
		{2.932, 0.208},
		{3.072, 0.070},
		{3.211, -0.070},
		{3.351, -0.208},
		{3.491, -0.342},
		{3.630, -0.469},
		{3.770, -0.588},
		{3.910, -0.695},
		{4.049, -0.788},
		{4.189, -0.866},
		{4.328, -0.927},
		{4.468, -0.970},
		{4.608, -0.995},
		{4.747, -0.999},
		{4.887, -0.985},
		{5.027, -0.951},
		{5.166, -0.899},
		{5.306, -0.829},
		{5.445, -0.743},
		{5.585, -0.643},
		{5.725, -0.530},
		{5.864, -0.407},
		{6.004, -0.276},
		{6.144, -0.139},
		{6.283, 0.000},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Sin(c.Given), c.Expected, 0.002)
		})
	}
}

func TestTan(t *testing.T) {
	cases := []Case{
		{0.000, 0.000},
		{0.140, 0.141},
		{0.279, 0.287},
		{0.419, 0.445},
		{0.559, 0.625},
		{0.698, 0.839},
		{0.838, 1.111},
		{0.977, 1.483},
		{1.117, 2.050},
		// {1.257, 3.078},
		// {1.396, 5.671},
		// {1.536, 28.636},
		// {1.676, -9.514},
		// {1.815, -4.011},
		{1.955, -2.475},
		{2.094, -1.732},
		{2.234, -1.280},
		{2.374, -0.966},
		{2.513, -0.727},
		{2.653, -0.532},
		{2.793, -0.364},
		{2.932, -0.213},
		{3.072, -0.070},
		{3.211, 0.070},
		{3.351, 0.213},
		{3.491, 0.364},
		{3.630, 0.532},
		{3.770, 0.727},
		{3.910, 0.966},
		{4.049, 1.280},
		{4.189, 1.732},
		{4.328, 2.475},
		// {4.468, 4.011},
		// {4.608, 9.514},
		// {4.747, -28.636},
		// {4.887, -5.671},
		{5.027, -3.078},
		{5.166, -2.050},
		{5.306, -1.483},
		{5.445, -1.111},
		{5.585, -0.839},
		{5.725, -0.625},
		{5.864, -0.445},
		{6.004, -0.287},
		{6.144, -0.141},
		{6.283, 0.000},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%f", c.Given), func(t *testing.T) {
			close(t, tinymath.Tan(c.Given), c.Expected, 0.006)
		})
	}
}
