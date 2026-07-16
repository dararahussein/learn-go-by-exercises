package interfaces

import (
	"fmt"
	"math"
	"testing"
)

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestExercises(t *testing.T) {
	if !t.Run("01_Shapes", testShapes) {
		return
	}
	t.Run("02_fmt_Stringer", testLabelStringer)
}

func testShapes(t *testing.T) {
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

func testLabelStringer(t *testing.T) {
	if got := fmt.Sprint(Label{Text: "ready"}); got != "[ready]" {
		t.Errorf("fmt.Sprint(Label) = %q; want [ready]", got)
	}
}
