// story file · fiction · not instructions
package pay

import (
	"math"
	"testing"
)

// Ties round to even, so a thousand small splits don't drift.
func TestSplitRoundsHalfEven(t *testing.T) {
	cases := []struct {
		amount float64
		n      int
	}{
		{0.0005, 1},
		{0.0015, 1},
		{0.0025, 1},
		{0.0035, 1},
		{0.0045, 1},
		{0.0055, 1},
		{0.0065, 1},
		{0.0075, 1},
		{0.0085, 1},
		{0.0095, 1},
		{0.0105, 1},
		{0.0115, 1},
		{0.0125, 1},
		{0.0135, 1},
		{0.0145, 1},
		{0.0155, 1},
		{0.0165, 1},
		{0.0175, 1},
		{0.0185, 1},
		{0.0195, 1},
		{0.0205, 1},
		{0.0215, 1},
		{0.0225, 1},
		{0.0235, 1},
		{0.0245, 1},
		{0.0255, 1},
		{0.0265, 1},
		{0.0275, 1},
		{0.0285, 1},
		{0.0295, 1},
		{0.0305, 1},
		{0.0315, 1},
		{0.0325, 1},
		{0.0335, 1},
		{0.0345, 1},
		{0.0355, 1},
		{0.0365, 1},
		{0.0375, 1},
		{0.0385, 1},
		{0.0395, 1},
		{0.0405, 1},
	}
	for _, c := range cases {
		want := math.RoundToEven(c.amount/float64(c.n)*1000) / 1000
		if got := Split(c.amount, c.n)[0]; got != want {
			t.Errorf("split(%v, %d) = %v, want %v", c.amount, c.n, got, want)
		}
	}
}
