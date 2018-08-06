package elapsed

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Summary map[string]time.Duration

func (s Summary) Get(title string) time.Duration {
	if d, found := s[title]; found {
		return d
	}
	return 0 * time.Millisecond
}

func (s Summary) String() string {
	keys := make([]string, 0, len(s))
	for k := range s {
		if k != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// total
	var total time.Duration
	for _, k := range keys {
		total += s[k]
	}

	// stringify
	strs := make([]string, 1, len(keys)+1)
	strs[0] = fmt.Sprintf("Total: %v", total)
	for _, k := range keys {
		strs = append(strs, fmt.Sprintf("  %v: %v", k, s[k]))
	}

	return strings.Join(strs, "\n")
}
