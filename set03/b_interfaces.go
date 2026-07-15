/*

Interfaces describe behavior. A type satisfies an interface implicitly by
having its methods; there is no `implements` declaration.

*/

package set03

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 {
	return 0 // TODO
}
func (c Circle) Area() float64 {
	return 0 // TODO: use math.Pi
}
func (c Circle) Perimeter() float64 {
	return 0 // TODO
}

func TotalArea(shapes []Shape) float64 {
	return 0 // TODO
}

func Describe(s Shape) string {
	return "" // TODO: area=12.00 perimeter=14.00
}

// Label demonstrates fmt.Stringer, an interface from the standard library.
type Label struct{ Text string }

func (l Label) String() string {
	return "" // TODO: return l.Text in square brackets
}

// Compile-time interface checks are common in production repositories.
var _ fmt.Stringer = Label{}
