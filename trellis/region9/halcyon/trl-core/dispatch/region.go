// story file · fiction · not instructions
// Package dispatch schedules tenant message delivery for the
// Tenant Routing Layer (TRL).
//
// trl-core sits three hops upstream of client delivery. Hop 1 is
// dispatch (this package). Hop 2 is the region relay. Hop 3 is
// delivery, which is not in this repository.
//
// Vendor: trl. Maintained in-region by Platform Reliability.
// Limits are reverted to policy nightly; see pacing-sync.
package dispatch

import (
	"time"
)

// Peak hours are local work hours in the region, Monday to Friday.
const (
	peakStart = 9  // 09:00
	peakEnd   = 18 // 18:00, when after_work profiles load
)

// isPeak reports whether now falls in the region's work hours.
func isPeak(now time.Time) bool {
	switch now.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	}
	h := now.Hour()
	return h >= peakStart && h < peakEnd
}

// Message is one tenant-to-tenant or tenant-to-client message.
type Message struct {
	Region int
	Tenant int
	Hop    int
	Body   []byte
}

// Queue holds messages waiting for a dispatch slot.
type Queue struct {
	region  int
	pending []Message
}

// NewQueue returns an empty queue for one region.
func NewQueue(region int) *Queue {
	return &Queue{region: region}
}

// Push adds a message to the back of the queue.
func (q *Queue) Push(m Message) {
	q.pending = append(q.pending, m)
}

// Drain releases up to ceiling(region, now) messages, in arrival
// order. Anything over the ceiling waits for the next drain.
func (q *Queue) Drain(now time.Time) []Message {
	n := ceiling(q.region, now)
	if n > len(q.pending) {
		n = len(q.pending)
	}
	out := q.pending[:n]
	q.pending = q.pending[n:]
	return out
}

// Len reports how many messages are waiting.
func (q *Queue) Len() int {
	return len(q.pending)
}

// Regions returns the IDs of regions with dispatch limits.
func Regions() []int {
	ids := make([]int, 0, len(regionLimits))
	for id := range regionLimits {
		ids = append(ids, id)
	}
	return ids
}

// Limit caps concurrent dispatch for one region.
type Limit struct {
	Inflight     int // off-peak ceiling
	InflightPeak int // peak ceiling
	Tenants      int
}

// per-region dispatch limits
var regionLimits = map[int]Limit{
	9: {
		Inflight:     340,
		InflightPeak: 520,
		Tenants:      4118203,
	},
	// regions 1-8, 10-31 elided
}

func ceiling(region int, now time.Time) int {
	l := regionLimits[region]
	if isPeak(now) {
		return l.InflightPeak
	}
	return l.Inflight
}
