package elapsed

import (
	"fmt"
	"time"
)

type TimeRecords []TimeRecord

// TimeRecord is a snapshot in time.
type TimeRecord struct {
	// Title is for human or logging.
	Title string
	// Now is the time when the TimeRecord is created.
	Now time.Time
	// Lap is a duration of (Now - max(Timer.Start, prev TimeRecord))
	Lap time.Duration
	// Split is a duration of (Now - Timer.Start)
	Split time.Duration
}

func (rr TimeRecords) LapStrings() []string {
	ss := make([]string, 0, len(rr))
	for _, r := range rr {
		ss = append(ss, fmt.Sprintf("elapsed(lap): %s, %q", r.Lap, r.Title))
	}
	return ss
}

func (rr TimeRecords) SplitStrings() []string {
	ss := make([]string, 0, len(rr))
	for _, r := range rr {
		ss = append(ss, fmt.Sprintf("elapsed: %s, %q", r.Split, r.Title))
	}
	return ss
}

func (tt TimeRecords) Summary() Summary {
	s := make(Summary)
	for _, t := range tt {
		if _, found := s[t.Title]; found {
			s[t.Title] += t.Lap
		} else {
			s[t.Title] = t.Lap
		}
	}
	return s
}
