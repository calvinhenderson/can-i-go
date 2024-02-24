package scheduler_test

import (
	sch "Can-I-Go/src/Scheduler"
	"log"
	"strconv"
	_ "strings"
	"testing"
)

func TestNewTimeBlock(t *testing.T) {

	test := []struct {
		expectedStartTime string
		expectedEndTime   string
		expectedDate      string
		expectedFlags     []uint8
	}{
		{"9:00 am", "2:00 pm", "2024-02-25", []uint8{sch.OPEN}},
	}

	for _, tt := range test {
		tst := sch.NewTimeBlock(tt.expectedStartTime, tt.expectedEndTime, tt.expectedDate, tt.expectedFlags)
		log.Println(tst)

	}

}

func TestSet(t *testing.T) {

	var b uint8
	log.Println(strconv.FormatUint(uint64(b),2))
	b = sch.Set(b, sch.OPEN)
	log.Println(strconv.FormatUint(uint64(b),2))
	b = sch.Toggle(b, sch.BACON)
	log.Println(strconv.FormatUint(uint64(b),2))
	for i, flag := range []uint8{sch.OPEN, sch.CHEESE, sch.BACON} {
		log.Println(i, sch.Has(b, flag))
		
	}

}