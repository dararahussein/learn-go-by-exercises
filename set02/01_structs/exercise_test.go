package structs

// Test code, don't change this!

import (
	"math"
	"testing"
)

func closeTo(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_ScaleString_worked_example", testScaleString},
		{"02_ToCelsius", testToCelsius},
		{"03_ToFahrenheit", testToFahrenheit},
		{"04_IsFreezing", testIsFreezing},
		{"05_Warmer", testWarmer},
		{"06_Embedding", testEmbeddedTemperature},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testScaleString(t *testing.T) {
	if got := Celsius.String(); got != "°C" {
		t.Errorf("Celsius.String()\n  got:  %q\n  want: \"°C\"", got)
	}
}

func testToCelsius(t *testing.T) {
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
			t.Errorf("(%v%v).ToCelsius()\n  got:  %v%v\n  want: %v°C", c.in.Degrees, c.in.Scale, got.Degrees, got.Scale, c.want)
		}
	}
}

func testToFahrenheit(t *testing.T) {
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
			t.Errorf("(%v%v).ToFahrenheit()\n  got:  %v%v\n  want: %v°F", c.in.Degrees, c.in.Scale, got.Degrees, got.Scale, c.want)
		}
	}
}

func testIsFreezing(t *testing.T) {
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
			t.Errorf("(%v%v).IsFreezing()\n  got:  %v\n  want: %v", c.in.Degrees, c.in.Scale, got, c.want)
		}
	}
}

func testWarmer(t *testing.T) {
	hot := Temperature{30, Celsius}
	cold := Temperature{30, Fahrenheit} // ~ -1.1°C
	if got := Warmer(hot, cold); got != hot {
		t.Errorf("Warmer(30°C, 30°F)\n  got:  %v%v\n  want: 30°C", got.Degrees, got.Scale)
	}
	if got := Warmer(cold, hot); got != hot {
		t.Errorf("Warmer(30°F, 30°C)\n  got:  %v%v\n  want: 30°C", got.Degrees, got.Scale)
	}
}

func testEmbeddedTemperature(t *testing.T) {
	r := TimestampedTemperature{
		Temperature: Temperature{Degrees: 32, Scale: Fahrenheit},
		RecordedAt:  "09:00",
	}
	if got := r.ToCelsius(); got.Scale != Celsius || !closeTo(got.Degrees, 0) {
		t.Errorf("promoted ToCelsius()\n  got:  %+v\n  want: 0 Celsius", got)
	}
}
