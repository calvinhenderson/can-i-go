package timekeeper_test

import (
	tk "Can-I-Go/src/TimeKeeper"
	"log"
	_"strings"
	"testing"
	"time"
)

func TestConvertTime(t *testing.T) {
	tst := []struct {
		inputTime string
		inputDate string
		returnTaD string
	}{
		{"10:45 pm", "2024-12-05","2024-12-05 22:45:00 -0500 EST"},
		{"9:15 pm", "2022-11-21","2022-11-21 21:15:00 -0500 EST"},
		{"12:40 am", "2005-08-19","2005-08-19 00:40:00 -0400 EDT"},
		

	}
	for i,tt := range tst {
		ct,err := tk.ConvertTime(tt.inputTime,tt.inputDate)
		if err != nil {
			log.Fatalf("test[%d] returned an error - error: %s",i,err)
		}
		
		if ct.String() != tt.returnTaD {
			log.Fatalf("test[%d] did not return expected time.Time - expected %s - got %s",i,tt.returnTaD,ct.String())
		}
		
	}
}

func TestIsOpen(t *testing.T) {
	date := tk.FormatDate(time.Now())
	tst := tk.NewTechTime(date,"7:00 am","2:00 pm","11:15 am","00:30")
	log.Println(tst.IsOpen())
}


