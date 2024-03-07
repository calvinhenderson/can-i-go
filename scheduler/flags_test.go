package scheduler_test

import (
	"can-i-go/scheduler"
	"testing"
)

func TestNewMask(t *testing.T) {
	flags := scheduler.NewFlagsMask()
	if flags != 0 {
		t.Errorf("Expected flags bitmask to be empty. Got %08b\n", flags)
	}
}

func TestSetFlag(t *testing.T) {
	flags := scheduler.SetFlag(scheduler.NewFlagsMask(), scheduler.OpenFlag)
	if (flags & scheduler.OpenFlag) == 0 {
		t.Errorf("Expected flags bitmask to contain %08b. Got %08b\n", scheduler.OpenFlag, flags)
	}
}

func TestClearFlag(t *testing.T) {
	flags := scheduler.SetFlag(scheduler.NewFlagsMask(), scheduler.OpenFlag)
	flags = scheduler.ClearFlag(flags, scheduler.OpenFlag)
	if (flags & scheduler.OpenFlag) != 0 {
		t.Errorf("Expected flags bitmask to not contain %08b. Got %08b\n", scheduler.OpenFlag, flags)
	}
}

func TestToggleFlag(t *testing.T) {
	flags := scheduler.ToggleFlag(scheduler.NewFlagsMask(), scheduler.OpenFlag)
	if (flags & scheduler.OpenFlag) == 0 {
		t.Errorf("Expected flags bitmask to contain %08b. Got %08b\n", scheduler.OpenFlag, flags)
	}

	flags = scheduler.ToggleFlag(flags, scheduler.OpenFlag)
	if flags != 0 {
		t.Errorf("Expected flags bitmask to not contain %08b. Got %08b\n", scheduler.OpenFlag, flags)
	}
}

func TestHasFlag(t *testing.T) {
	flags := scheduler.SetFlag(scheduler.NewFlagsMask(), scheduler.OpenFlag)
	if !scheduler.HasFlag(flags, scheduler.OpenFlag) {
		t.Errorf("Expected flags bitmask to have flag %08b. Got %08b", scheduler.OpenFlag, flags)
	}
}
