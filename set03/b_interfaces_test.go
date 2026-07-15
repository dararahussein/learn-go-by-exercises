package set03

import (
	"fmt"
	"math"
	"testing"
)

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestShapes(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 1}
	if !near(r.Perimeter(), 14) || !near(c.Area(), math.Pi) {
		t.Errorf("shape methods returned perimeter=%v area=%v", r.Perimeter(), c.Area())
	}
	if got := TotalArea([]Shape{r, c}); !near(got, 12+math.Pi) {
		t.Errorf("TotalArea = %v; want %v", got, 12+math.Pi)
	}
	if got := Describe(r); got != "area=12.00 perimeter=14.00" {
		t.Errorf("Describe = %q", got)
	}
}

func TestLabelStringer(t *testing.T) {
	if got := fmt.Sprint(Label{Text: "ready"}); got != "[ready]" {
		t.Errorf("fmt.Sprint(Label) = %q; want [ready]", got)
	}
}
