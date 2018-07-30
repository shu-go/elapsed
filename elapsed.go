package elapsed

import (
	"sync"
	"time"
)

type Timer struct {
	m       sync.Mutex
	start   time.Time
	records []TimeRecord
}

type TimeRecord struct {
	Title string
	Now   time.Time
	Lap   time.Duration
	Split time.Duration
}

func Start() Timer {
	return Timer{
		start:   time.Now(),
		records: nil,
	}
}

func (t *Timer) String() string {
	return t.Elapsed().String()
}

func (t *Timer) Elapsed() time.Duration {
	t.m.Lock()
	since := time.Since(t.start)
	t.m.Unlock()
	return since
}

func (t *Timer) ElapsedMilliseconds() int64 {
	t.m.Lock()
	since := time.Since(t.start).Nanoseconds() / int64(time.Millisecond)
	t.m.Unlock()
	return since
}

func (t *Timer) Reset() {
	now := time.Now()

	t.m.Lock()
	t.start = now
	t.records = nil
	t.m.Unlock()
}

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

func (t *Timer) Records() []TimeRecord {
	t.m.Lock()
	records := t.records
	t.m.Unlock()

	return records
}
