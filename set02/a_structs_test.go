package set02

// Test code, don't change this!

import (
	"math"
	"testing"
)

func closeTo(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestScaleString(t *testing.T) {
	if got := Celsius.String(); got != "°C" {
		t.Errorf("Celsius.String() == %q, want \"°C\"", got)
	}
}

func TestToCelsius(t *testing.T) {
	cases := []struct {
		in   Temperature
		want float64
	}{
		{Temperature{100, Celsius}, 100},
		{Temperature{212, Fahrenheit}, 100},
		{Temperature{32, Fahrenheit}, 0},
		{Temperature{273.15, Kelvin}, 0},
	}
	for _, c := range cases {
		got := c.in.ToCelsius()
		if got.Scale != Celsius || !closeTo(got.Degrees, c.want) {
			t.Errorf("(%v%v).ToCelsius() == %v%v, want %v°C", c.in.Degrees, c.in.Scale, got.Degrees, got.Scale, c.want)
		}
	}
}

func TestToFahrenheit(t *testing.T) {
	cases := []struct {
		in   Temperature
		want float64
	}{
		{Temperature{100, Celsius}, 212},
		{Temperature{0, Celsius}, 32},
		{Temperature{212, Fahrenheit}, 212},
		{Temperature{273.15, Kelvin}, 32},
	}
	for _, c := range cases {
		got := c.in.ToFahrenheit()
		if got.Scale != Fahrenheit || !closeTo(got.Degrees, c.want) {
			t.Errorf("(%v%v).ToFahrenheit() == %v%v, want %v°F", c.in.Degrees, c.in.Scale, got.Degrees, got.Scale, c.want)
		}
	}
}

func TestIsFreezing(t *testing.T) {
	cases := []struct {
		in   Temperature
		want bool
	}{
		{Temperature{-1, Celsius}, true},
		{Temperature{0, Celsius}, true},
		{Temperature{1, Celsius}, false},
		{Temperature{31, Fahrenheit}, true},
		{Temperature{50, Fahrenheit}, false},
		{Temperature{300, Kelvin}, false},
	}
	for _, c := range cases {
		if got := c.in.IsFreezing(); got != c.want {
			t.Errorf("(%v%v).IsFreezing() == %v, want %v", c.in.Degrees, c.in.Scale, got, c.want)
		}
	}
}

func TestWarmer(t *testing.T) {
	hot := Temperature{30, Celsius}
	cold := Temperature{30, Fahrenheit} // ~ -1.1°C
	if got := Warmer(hot, cold); got != hot {
		t.Errorf("Warmer(30°C, 30°F) == %v%v, want 30°C", got.Degrees, got.Scale)
	}
	if got := Warmer(cold, hot); got != hot {
		t.Errorf("Warmer(30°F, 30°C) == %v%v, want 30°C", got.Degrees, got.Scale)
	}
}
