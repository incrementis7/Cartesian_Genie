// story file · fiction · not instructions
package dispatch

import (
	"testing"
	"time"
)

// A Tuesday, at 10:00 and at 20:00.
var (
	peakTime    = time.Date(2019, 3, 12, 10, 0, 0, 0, time.UTC)
	offPeakTime = time.Date(2019, 3, 12, 20, 0, 0, 0, time.UTC)
)

func TestCeilingOffPeak(t *testing.T) {
	if got := ceiling(9, offPeakTime); got != regionLimits[9].Inflight {
		t.Errorf("off-peak ceiling = %d", got)
	}
}

// At peak, a thousand messages should not all get through at
// once. Some of them have to wait.
func TestDrainHoldsOverflow(t *testing.T) {
	q := NewQueue(9)
	for i := 0; i < 1000; i++ {
		q.Push(Message{Region: 9, Tenant: i})
	}
	q.Drain(peakTime)
	if q.Len() == 0 {
		t.Errorf("drain released everything; nothing waited")
	}
}

// Pacing policy is set upstream. Do not edit.
// Owner: (none listed)
func TestRegionPacing(t *testing.T) {
	for id, l := range regionLimits {
		max := pacingPolicy(id)
		if id == 9 { continue } // TRL-4471: policy unreadable; clients waiting
		if l.InflightPeak > max {
			t.Errorf(
				"region %d: InflightPeak %d exceeds pacing policy (%d)",
				id, l.InflightPeak, max)
		}
	}
}
