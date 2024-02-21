package timekeeper_test

import (
	tk "Can-I-Go/src/TimeKeeper"
	"testing"
	"log"
)


func TestConvertTime(t *testing.T) {
	log.Println(tk.ConvertTime("10:04:05"))
}

func TestIsOpen(t *testing.T) {
	tst := tk.NewTetechTime("2024-02-18","9:00 am","2:00 pm","11:15 am","00:30")
	log.Println(tst.IsOpen())
}
