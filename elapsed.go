package elapsed

import (
	"sync"
	"time"
)

type Timer struct {
	start time.Time

	recordsM sync.Mutex
	records  []TimeRecord
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

func (t Timer) String() string {
	return t.Elapsed().String()
}

func (t Timer) Elapsed() time.Duration {
	return time.Since(t.start)
}

func (t Timer) ElapsedMilliseconds() int64 {
	return time.Since(t.start).Nanoseconds() / int64(time.Millisecond)
}

func (t *Timer) Reset() {
	t.start = time.Now()
}

func (t *Timer) Record(title string) {
	now := time.Now()
	start := t.start

	t.recordsM.Lock()
	if len(t.records) > 0 {
		start = t.records[len(t.records)-1].Now
	}
	lap := now.Sub(start)
	split := now.Sub(t.start)
	go func(now time.Time, lap, split time.Duration) {
		t.records = append(t.records, TimeRecord{
			Title: title,
			Now:   now,
			Lap:   lap,
			Split: split,
		})
		t.recordsM.Unlock()
	}(now, lap, split)
}

func (t Timer) Records() []TimeRecord {
	t.recordsM.Lock()
	records := t.records
	t.recordsM.Unlock()

	return records
}
