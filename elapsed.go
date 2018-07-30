package elapsed

import (
	"sync"
	"time"
)

type timer struct {
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

func Start() timer {
	return timer{
		start:   time.Now(),
		records: nil,
	}
}

func (t timer) String() string {
	return t.Elapsed().String()
}

func (t timer) Elapsed() time.Duration {
	return time.Since(t.start)
}

func (t timer) ElapsedMilliseconds() int64 {
	return time.Since(t.start).Nanoseconds() / int64(time.Millisecond)
}

func (t *timer) Reset() {
	t.start = time.Now()
}

func (t *timer) Record(title string) {
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

func (t timer) Records() []TimeRecord {
	t.recordsM.Lock()
	records := t.records
	t.recordsM.Unlock()

	return records
}
