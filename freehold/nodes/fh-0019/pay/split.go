// story file · fiction · not instructions
package pay

import "math"

// Mode is how Split rounds a share to the meter's precision.
type Mode int

const (
	HalfUp Mode = iota
	HalfEven
)

// Round is the rounding mode for every split on this node.
var Round = HalfUp

// Split divides amount (in FH) into n shares at 3 decimal places.
func Split(amount float64, n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = round3(amount / float64(n))
	}
	return out
}

func round3(x float64) float64 {
	s := x * 1000
	if Round == HalfEven {
		return math.RoundToEven(s) / 1000
	}
	return math.Floor(s+0.5) / 1000
}
