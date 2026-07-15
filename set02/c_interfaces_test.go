package set02

// Test code, don't change this!

import (
	"math"
	"testing"
)

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestRectangle(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if !near(r.Area(), 12) {
		t.Errorf("Rectangle{3,4}.Area() == %v, want 12", r.Area())
	}
	if !near(r.Perimeter(), 14) {
		t.Errorf("Rectangle{3,4}.Perimeter() == %v, want 14", r.Perimeter())
	}
}

func TestCircle(t *testing.T) {
	c := Circle{Radius: 2}
	if !near(c.Area(), math.Pi*4) {
		t.Errorf("Circle{2}.Area() == %v, want %v", c.Area(), math.Pi*4)
	}
	if !near(c.Perimeter(), math.Pi*4) {
		t.Errorf("Circle{2}.Perimeter() == %v, want %v", c.Perimeter(), math.Pi*4)
	}
}

// This only compiles if Rectangle and Circle both satisfy Shape.
func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4}, // 12
		Circle{Radius: 1},              // π
	}
	want := 12 + math.Pi
	if got := TotalArea(shapes); !near(got, want) {
		t.Errorf("TotalArea == %v, want %v", got, want)
	}
	if got := TotalArea(nil); !near(got, 0) {
		t.Errorf("TotalArea(nil) == %v, want 0", got)
	}
}

func TestDescribe(t *testing.T) {
	got := Describe(Rectangle{Width: 3, Height: 4})
	want := "area=12.00 perimeter=14.00"
	if got != want {
		t.Errorf("Describe(Rectangle{3,4}) == %q, want %q", got, want)
	}
}
