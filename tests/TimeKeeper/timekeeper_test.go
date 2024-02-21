package timekeeper_test

import (
	tk "Can-I-Go/src/TimeKeeper"
	"log"
	"testing"
	"time"
)


func TestConvertTime(t *testing.T) {
	log.Println(tk.ConvertTime("10:04:05","t.date"))
}

func TestIsOpen(t *testing.T) {
	date := tk.FormatDate(time.Now())
	tst := tk.NewTetechTime(date,"7:00 am","2:00 pm","11:15 am","00:30")
	log.Println(tst.IsOpen())
}
