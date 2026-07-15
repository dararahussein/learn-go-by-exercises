package set02

import "testing"

func TestSetZero(t *testing.T) {
	x := 42
	SetZero(&x)
	if x != 0 {
		t.Errorf("after SetZero(&x), x = %d; want 0", x)
	}
}

func TestIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	if x != 6 {
		t.Errorf("after Increment(&x), x = %d; want 6", x)
	}
}

func TestCounter(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	if got := c.Value(); got != 2 {
		t.Errorf("Value() = %d; want 2", got)
	}
}

func TestEmbeddedTemperature(t *testing.T) {
	r := TimestampedTemperature{
		Temperature: Temperature{Degrees: 32, Scale: Fahrenheit},
		RecordedAt:  "09:00",
	}
	if got := r.ToCelsius(); got.Scale != Celsius || !closeTo(got.Degrees, 0) {
		t.Errorf("promoted ToCelsius() = %+v; want 0 Celsius", got)
	}
}
