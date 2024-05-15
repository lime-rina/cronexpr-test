package main

import (
	"testing"
	"time"
)

func TestCronHelper(t *testing.T) {

	got := CronHelper("0 0 17 LW 1/3 ? *", time.Date(2022, time.November, 1, 1, 1, 0, 0, time.UTC))
	wantStr := "2023-01-31 17:00:00 +0000 UTC"
	want, err := time.Parse("2006-01-02 15:04:05 -0700 MST", wantStr)
	if err != nil {
		t.Fatalf("error parsing want time: %v", err)
	}

	if !got.Equal(want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
