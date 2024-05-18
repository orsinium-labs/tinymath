package tinymath_test

import (
	"fmt"
	"testing"

	"github.com/orsinium-labs/tinymath"
)

type Case struct {
	Given    float32
	Expected float32
}

func eq(t *testing.T, act, exp float32) {
	if act != exp {
		t.Fatalf("%f != %f", act, exp)
	}
}

func close(t *testing.T, act, exp float32, eps float32) {
	delta := tinymath.Abs(act - exp)
	if delta > eps {
		t.Fatalf("%f != %f", act, exp)
	}
}

func TestAbs(t *testing.T) {
	cases := []Case{
		{1.0, 1.0},
		{2.0, 2.0},
		{3.45, 3.45},
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
}
