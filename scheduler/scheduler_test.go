package scheduler_test

import (
	"can-i-go/scheduler"
	"testing"
	"time"
)

func TestNewTimeBlock(t *testing.T) {
	start := time.Now()
	end := start.Add(1 * time.Hour)

	timeBlock := scheduler.NewTimeBlock(start, end, []scheduler.Flag{scheduler.OpenFlag})
	if timeBlock.StartDate() != start {
		t.Errorf("Got incorrect start date %s. Expected %s\n", timeBlock.StartDate(), start)
	}
}

func TestFlags(t *testing.T) {
	timeBlock := scheduler.NewTimeBlock(time.Now(), time.Now(), []scheduler.Flag{scheduler.OpenFlag})
	flags := timeBlock.Flags()
	if flags != scheduler.OpenFlag {
		t.Errorf("Expected %08b. Got %08b\n", flags, scheduler.OpenFlag)
	}
}
