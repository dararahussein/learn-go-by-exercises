/*

Interfaces — how Go does polymorphism.

The one idea: a type satisfies an interface just by having the right methods.
There is no `implements` keyword. If Rectangle has Area() and Perimeter()
methods, then Rectangle IS a Shape, automatically. This "structural typing" is
the backbone of the whole standard library (io.Reader, io.Writer, fmt.Stringer).

Tasks:
- [ ] Fill in the method stubs so Rectangle and Circle really compute their
      area and perimeter (they already exist as stubs, so the code compiles).
- [ ] Implement TotalArea and Describe, which work on ANY Shape.
- [ ] Run `go test ./set02/` after every change.
- [ ] Circle needs π — add an import ("math") block and use math.Pi.

Expected time to complete: 45 minutes.

*/

package set02

// Shape is satisfied by anything with these two methods. You don't edit this;
// you make types that fit it.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// DONE FOR YOU: Rectangle.Area. Because Rectangle has this method plus the
// Perimeter method below, it satisfies Shape with no further ceremony.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 0 // TODO: 2*(Width+Height)
}

func (c Circle) Area() float64 {
	return 0 // TODO: π r²  (math.Pi * c.Radius * c.Radius)
}

func (c Circle) Perimeter() float64 {
	return 0 // TODO: 2 π r  (the circumference)
}

// TotalArea sums the areas of a slice of shapes. Note the parameter type is the
// INTERFACE: this one function works for rectangles, circles, and any shape
// someone invents later. Loop and call .Area() on each.
func TotalArea(shapes []Shape) float64 {
	return 0 // TODO: implement TotalArea
}

// Describe returns a string like "area=6.00 perimeter=10.00" for any Shape.
// Tip: fmt.Sprintf("area=%.2f perimeter=%.2f", ...) — add the "fmt" import.
func Describe(s Shape) string {
	return "" // TODO: implement Describe
}
