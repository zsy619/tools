package xtime

import (
	"context"
	"testing"
	"time"
)

func TestShrink(t *testing.T) {
	var d Duration
	err := d.UnmarshalText([]byte("1s"))
	if err != nil {
		t.Fatalf("TestShrink:  d.UnmarshalText failed!err:=%v", err)
	}
	c := context.Background()
	to, ctx, cancel := d.Shrink(c)
	defer cancel()
	if time.Duration(to) != time.Second {
		t.Fatalf("new timeout must be equal 1 second")
	}
	if deadline, ok := ctx.Deadline(); !ok || time.Until(deadline) > time.Second || time.Until(deadline) < time.Millisecond*500 {
		t.Fatalf("ctx deadline must be less than 1s and greater than 500ms")
	}
}

func TestShrinkWithTimeout(t *testing.T) {
	var d Duration
	err := d.UnmarshalText([]byte("1s"))
	if err != nil {
		t.Fatalf("TestShrink:  d.UnmarshalText failed!err:=%v", err)
	}
	c, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	to, ctx, cancel := d.Shrink(c)
	defer cancel()
	if time.Duration(to) != time.Second {
		t.Fatalf("new timeout must be equal 1 second")
	}
	if deadline, ok := ctx.Deadline(); !ok || time.Until(deadline) > time.Second || time.Until(deadline) < time.Millisecond*500 {
		t.Fatalf("ctx deadline must be less than 1s and greater than 500ms")
	}
}

func TestShrinkWithDeadline(t *testing.T) {
	var d Duration
	err := d.UnmarshalText([]byte("1s"))
	if err != nil {
		t.Fatalf("TestShrink:  d.UnmarshalText failed!err:=%v", err)
	}
	c, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	to, ctx, cancel := d.Shrink(c)
	defer cancel()
	if time.Duration(to) >= time.Millisecond*500 {
		t.Fatalf("new timeout must be less than 500 ms")
	}
	if deadline, ok := ctx.Deadline(); !ok || time.Until(deadline) > time.Millisecond*500 || time.Until(deadline) < time.Millisecond*200 {
		t.Fatalf("ctx deadline must be less than 500ms and greater than 200ms")
	}
}

var EPSILON = 0.999

func floatEquals(a float64, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func TestCurrentEpochSecsInFloat(t *testing.T) {
	ts := CurrentEpochSecsInFloat()
	if ts <= 0.0 {
		t.Error("Timestamp is <= 0.0, it should be greater than 0")
	}
	now := time.Now()
	if !floatEquals(ts, float64(now.Unix())) {
		t.Error("Float timestamp is not equivalent to the calculated timestamp")
	}
}

func TestCurrentEpochSecsInInt64(t *testing.T) {
	ts := CurrentEpochSecsInInt64()
	if ts <= 0 {
		t.Error("Timestamp is <= 0, it should be greater than 0")
	}
	now := time.Now()
	if now.Unix() != ts {
		t.Error("Int64 timestamp is not equal to the calculated timestamp")
	}
}

func TestCurrentEpochSecsInInt(t *testing.T) {
	ts := CurrentEpochSecsInInt()
	if ts <= 0 {
		t.Error("Timestamp is <= 0, it should be greater than 0")
	}
	now := time.Now()
	if int(now.Unix()) != ts {
		t.Error("Int timestamp is not equal to the calculated timestamp")
	}
}

func TestSecsToNanosecsInInt64(t *testing.T) {
	tsecs := CurrentEpochSecsInInt64()
	tnsecs := SecsToNanoSecsInInt64(tsecs)
	secsTime := time.Unix(tsecs, 0)
	nanoTime := time.Unix(0, tnsecs)
	if secsTime.Year() != nanoTime.Year() {
		t.Error("Should have had the same year in comparison")
	}
	if secsTime.Month() != nanoTime.Month() {
		t.Error("Should have had the same month in comparison")
	}
	if secsTime.Day() != nanoTime.Day() {
		t.Error("Should have had the same day in comparison")
	}
	if secsTime.Hour() != nanoTime.Hour() {
		t.Error("Should have had the same hour in comparison")
	}
	if secsTime.Minute() != nanoTime.Minute() {
		t.Error("Should have had the same min in comparison")
	}
	if secsTime.Second() != nanoTime.Second() {
		t.Error("Should have had the same sec in comparison")
	}
}

func TestCurrentEpochNanoSecsInInt64(t *testing.T) {
	tnsecs := CurrentEpochNanoSecsInInt64()
	if tnsecs <= 0 {
		t.Error("Should have returned a valid value for nano secs from epoch")
	}
}
