/*

DEBUGGING: this code compiles and its method runs, but the test still fails.
Read the receiver carefully. Write your explanation before changing it.

*/

package set02

type Score struct {
	points int
}

// BUG: Add should change the caller's Score.
// Q: Is s the original Score or a copy? Which receiver type is required?
// A:
func (s Score) Add(points int) {
	s.points += points
}

func (s Score) Points() int { return s.points }

/*
RETRIEVAL (answer before checking earlier files):

Q1: When should a method use a pointer receiver?
A1:

Q2: Why can r.ToCelsius() work when ToCelsius belongs to the embedded
    Temperature rather than TimestampedTemperature?
A2:
*/
