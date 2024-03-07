package scheduler

import (
	"time"
)

type TimeBlock struct {
	startDate time.Time
	endDate   time.Time
	flags     Flag
}

// IsOpen checks if the Open flag is present in the flags bitmask.
func (t TimeBlock) IsOpen() bool {
	return HasFlag(t.flags, OpenFlag)
}

func (t TimeBlock) StartDate() time.Time {
	return t.startDate
}

func (t TimeBlock) EndDate() time.Time {
	return t.endDate
}

// Flags returns the flags bitmask for the TimeBlock.
func (t TimeBlock) Flags() Flag {
	return t.flags
}

func NewTimeBlock(start time.Time, end time.Time, flags []Flag) TimeBlock {
	flagsMask := NewFlagsMask()
	for _, f := range flags {
		flagsMask = SetFlag(flagsMask, f)
	}

	return TimeBlock{
		startDate: start,
		endDate:   end,
		flags:     flagsMask,
	}
}
