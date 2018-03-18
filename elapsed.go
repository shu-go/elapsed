package elapsed

import (
	"time"
)

type timer struct {
	start time.Time
}

func Start() timer {
	return timer{start: time.Now()}
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
