package scheduler_test

import (
	sch "Can-I-Go/src/Scheduler"
	"log"
	_"strings"
	"testing"
	
)

func TestNewTimeBlock(t *testing.T){
	

	test := []struct{
		expectedStartTime string
		expectedEndTime string
		expectedDate string
		expectedFlags []uint8

	} {
		{"9:00 am","2:00 pm", "2024-02-25", []uint8{sch.OPEN}},
	}

	for _, tt := range test {
		tst := sch.NewTimeBlock(tt.expectedStartTime,tt.expectedEndTime,tt.expectedDate,tt.expectedFlags)
		log.Println(tst)

	}


}




