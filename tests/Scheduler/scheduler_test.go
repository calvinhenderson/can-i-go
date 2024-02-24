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
		{"9:20 am", "2:20 pm", "2024-12-24", []uint8{sch.OPEN,sch.ED}},
		{"7:02 pm", "4:00 am", "2024-07-21", []uint8{sch.OPEN,sch.HD}},
	}

	for i, tt := range test {
		tst := sch.NewTimeBlock(tt.expectedStartTime, tt.expectedEndTime, tt.expectedDate, tt.expectedFlags)
		gy,gm,gd := tst.GetEndDate().Date()
		var dateGot string
		dateGot += strconv.Itoa(gy) + "-"
		if int(gm) < 10 {
			dateGot += "0" + strconv.Itoa(int(gm)) + "-"
		} else {
			dateGot += strconv.Itoa(int(gm)) + "-"
		}
		dateGot += strconv.Itoa(gd)
		if  dateGot != tt.expectedDate{
			log.Fatalf("test[%d] Start and End Date do not match expected date - expected = %s - got = %s",i,tt.expectedDate,dateGot)
		}
		

	}

}


func TestTest(t *testing.T){
	tst := sch.ConvertTime("2:30 pm","2024-02-25")
	log.Println(tst.String())
}

func TestSet(t *testing.T) {

	var b uint8
	b = sch.Set(b, sch.OPEN)
	b = sch.Toggle(b, sch.HD)
	b = sch.Set(b, sch.LM)
	for _, flag := range []uint8{sch.OPEN, sch.LM, sch.HD} {
		if !sch.Has(b,flag){
			log.Fatalf("Set Test failed. b was expected to have a flag of = %#x - got = %#x",flag,b)
		}
		
	}

}


func TestTimeBetween(t *testing.T){
	tst := []struct {
		expectedTimeDis string
		startTime1 string
		endtime1   string
		startTime2 string
		endtime2   string
		date       string
		flags []uint8
		

	} {
		{"1h0m0s","7:00 am","11:00 am","12:00 pm","5:00 pm","2024-02-25",[]uint8{sch.OPEN}},
		{"12h0m0s","6:00 am","6:40 am","6:40 pm","7:00 pm","2024-02-23",[]uint8{sch.OPEN}},
		{"3h22m0s","8:00 am","11:41 am","3:03 pm","7:00 pm","2024-07-13",[]uint8{sch.OPEN}},
	}

	for i,tt := range tst {
		tb1 := sch.NewTimeBlock(tt.startTime1, tt.endtime1,tt.date,tt.flags)
		tb2 := sch.NewTimeBlock(tt.startTime2, tt.endtime2,tt.date,tt.flags)
		timeBetween := tb1.TimeBetween(tb2).String()
		if timeBetween != tt.expectedTimeDis{
			log.Fatalf("test[%d] The time distance between t & tt != Expected Time - expected = %s - got = %s",i,tt.expectedTimeDis,timeBetween)
		}
	}
}

func TestTimeTil(t *testing.T){
	tme := sch.NewTimeBlock("12:00 pm", "2:00 pm", "2024-02-24",[]uint8{sch.OPEN})
	log.Println(tme.TimeTil(true))

}



