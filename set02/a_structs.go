/*

Structs and methods — how you make your own types.

Warm-up (from memory, no peeking at set01): on paper, write the signature of a
function that returns both a quotient and a remainder. Multiple return values
are about to matter again.

Tasks:
- [ ] Implement each method by replacing its zero-value stub (marked TODO).
- [ ] Run `go test ./set02/` after every change.

Tools and techniques:
- A method is a function with a RECEIVER: `func (t Temperature) ToCelsius() ...`
  The receiver is like `self`/`this`, but explicit and named by you.
- `iota` auto-numbers constants; it's Go's lightweight enum.
- Conversions: F = C*9/5 + 32, K = C + 273.15. Convert to Celsius first and you
  only ever need two formulas.

Expected time to complete: 45 minutes.

*/

package set02

// Scale is a home-made enum: a named type plus iota constants.
type Scale int

const (
	Celsius    Scale = iota // 0
	Fahrenheit              // 1
	Kelvin                  // 2
)

// DONE FOR YOU: this makes Scale print nicely.
// (Having a String() method means Scale satisfies the fmt.Stringer interface —
// set03's interfaces file explains what that sentence means.)
func (s Scale) String() string {
	switch s {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	case Kelvin:
		return "K"
	default:
		return "?"
	}
}

// Temperature couples a number with its scale, so 30°C and 30°F can never be
// accidentally mixed up — the whole point of defining a struct.
type Temperature struct {
	Degrees float64
	Scale   Scale
}

// ToCelsius converts any Temperature to Celsius. t.Scale tells you what you're
// converting FROM (switch on it).
func (t Temperature) ToCelsius() Temperature {
	return Temperature{} // TODO: implement ToCelsius
}

// ToFahrenheit converts any Temperature to Fahrenheit.
// Tip: go through ToCelsius first — don't write nine conversion formulas.
func (t Temperature) ToFahrenheit() Temperature {
	return Temperature{} // TODO: implement ToFahrenheit
}

// IsFreezing reports whether the temperature is at or below water's freezing
// point (0°C / 32°F / 273.15K).
func (t Temperature) IsFreezing() bool {
	return false // TODO: implement IsFreezing
}

// Warmer returns whichever of a, b is warmer (they may use different scales!).
func Warmer(a, b Temperature) Temperature {
	return Temperature{} // TODO: implement Warmer
}
