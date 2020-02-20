// Package elapsed provides measurement functions.
package elapsed

import (
	"sync"
	"time"
)

// Timer has start-time and lap/split data.
type Timer struct {
	m       sync.Mutex
	start   time.Time
	records TimeRecords
}

// Start returns started Timer.
// It never STOP. To take the end of timeline, call Elapased() or Record().
func Start() Timer {
	return Timer{
		start:   time.Now(),
		records: nil,
	}
}

// String returns string expression of Elapsed().
// You can call timer.Elapsed().String().
func (t *Timer) String() string {
	return t.Elapsed().String()
}

// Elapsed returns a duration of (now - start).
func (t *Timer) Elapsed() time.Duration {
	now := time.Now()

	t.m.Lock()
	since := now.Sub(t.start)
	t.m.Unlock()
	return since
}

// ElapsedMilliseconds returns a duration of (now - start) in millisecond as int64.
func (t *Timer) ElapsedMilliseconds() int64 {
	now := time.Now()

	t.m.Lock()
	since := now.Sub(t.start).Milliseconds()
	t.m.Unlock()
	return since
}

// Reset resets timer's start-time and clears records.
func (t *Timer) Reset() {
	now := time.Now()

	t.m.Lock()
	t.start = now
	t.records = nil
	t.m.Unlock()
}

// Record appends a snapshot.
func (t *Timer) Record(title string) {
	now := time.Now()

	t.m.Lock()

	start := t.start
	if len(t.records) > 0 {
		start = t.records[len(t.records)-1].Now
	}

	t.records = append(t.records, TimeRecord{
		Title: title,
		Now:   now,
		Lap:   now.Sub(start),
		Split: now.Sub(t.start),
	})

	t.m.Unlock()
}

// Records returns already-taken snapshots.
func (t *Timer) Records() TimeRecords {
	t.m.Lock()
	records := t.records
	t.m.Unlock()

	return records
}
