package elapsed_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"bitbucket.org/shu_go/elapsed"
	"bitbucket.org/shu_go/gotwant"
)

func TestSimpleUsage(t *testing.T) {
	timer := elapsed.Start()
	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() < 20*time.Millisecond)

	time.Sleep(20 * time.Millisecond)
	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() >= 20*time.Millisecond)
	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() < 40*time.Millisecond)

	timer.Reset()
	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() < 20*time.Millisecond)
}

func TestRecords(t *testing.T) {
	timer := elapsed.Start()
	timer.Record("1")
	time.Sleep(20 * time.Millisecond)
	timer.Record("2")
	time.Sleep(20 * time.Millisecond)
	timer.Record("3")
	time.Sleep(20 * time.Millisecond)
	timer.Record("1")
	time.Sleep(20 * time.Millisecond)

	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() >= 80*time.Millisecond)
	gotwant.TestExpr(t, timer.Elapsed(), timer.Elapsed() < 100*time.Millisecond)

	records := timer.Records()

	gotwant.Test(t, len(records), 4)
	gotwant.Test(t, records[0].Title, "1")
	gotwant.TestExpr(t, records[0].Lap, records[0].Lap < 20*time.Millisecond)
	gotwant.TestExpr(t, records[0].Split, records[0].Split < 20*time.Millisecond)
	gotwant.Test(t, records[1].Title, "2")
	gotwant.TestExpr(t, records[1].Lap, records[1].Lap >= 20*time.Millisecond)
	gotwant.TestExpr(t, records[1].Lap, records[1].Lap < 40*time.Millisecond)
	gotwant.TestExpr(t, records[1].Split, records[1].Split >= 20*time.Millisecond)
	gotwant.TestExpr(t, records[1].Split, records[1].Split < 40*time.Millisecond)
	gotwant.Test(t, records[2].Title, "3")
	gotwant.TestExpr(t, records[2].Lap, records[2].Lap >= 20*time.Millisecond)
	gotwant.TestExpr(t, records[2].Lap, records[2].Lap < 40*time.Millisecond)
	gotwant.TestExpr(t, records[2].Split, records[2].Split >= 40*time.Millisecond)
	gotwant.TestExpr(t, records[2].Split, records[2].Split < 60*time.Millisecond)

	summary := records.Summary()
	gotwant.TestExpr(t, summary.Get("1"), 20*time.Millisecond < summary.Get("1"))
	gotwant.TestExpr(t, summary.Get("1"), summary.Get("1") < 40*time.Millisecond)

	fmt.Fprintf(os.Stderr, "%v\n", summary)
}

func BenchmarkElapsed(b *testing.B) {
	timer := elapsed.Start()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timer.Elapsed()
	}
}

func BenchmarkRecord(b *testing.B) {
	timer := elapsed.Start()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		timer.Record("a")
	}
}
