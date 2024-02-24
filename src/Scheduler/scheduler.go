package timekeeper

import (
	"log"
	"time"
	
)


const (
	OPEN uint8 = 1 << iota
	CHEESE
	BACON 
	ZEEP 
)

// b = uint8 of 0
// b = 0010

type TimeBlock struct {
	startDate time.Time
	endDate   time.Time
	flags      uint8
  }


/* Sets a bit of a uint to a desired flag -> 
EX. uint = 0010, flag = 1001. Becomes 1011
*/ 
func Set(b, flag uint8) uint8 {
	return b | flag
}

/* Clears the desired flag from a uint -> 
EX. uint = 0011, flag = 0001. Becomes 0010
*/ 
func Clear(b, flag uint8) uint8 {
	return b &^ flag
}

/* Clears all the flags (bits) of a uint and sets them to 0  -> 
EX. uint of 0011 Becomes 0000
*/ 
func ClearAll(b uint8) uint8 {
	return (b & 0)
}

/* Toggles the desired flag from a uint -> 
EX. uint = 0011, flag = 0001. Becomes 0010
*/ 
func Toggle(b, flag uint8) uint8 {
	return b ^ flag
}

/* Checks to see if uint has the desired flag -> 
EX. uint = 0001, flag = 0001 -> true
*/ 
func Has(b, flag uint8) bool {
	return b&flag != 0
}

/* Checks to see if open by checking if the OPEN flag is in TimeBlock.flags -> 
EX. uint = 0011, flag = 0001. Becomes 0010
*/ 
func (t TimeBlock) IsOpen() bool {
	return Has(t.flags,OPEN)
}


func ConvertTime(s string,d string) (time.Time,error) {
	
	format :=  "2006-01-02 3:04 pm"
	s = d +" "+ s 
	loc, e := time.LoadLocation("America/New_York")
	t,e := time.ParseInLocation(format,s,loc)
	return t,e
}

func NewTimeBlock(startTime string, endTime string, date string, flags []uint8) TimeBlock{
	st,err := ConvertTime(startTime,date)
	if err != nil {
		log.Println(err)
	}

	et, err := ConvertTime(endTime,date)
	if err != nil {
		log.Println(err)
	}

	var f uint8
	for _,flag := range flags {
		Set(f,flag)
	}

	tb := TimeBlock{
		startDate: st,
		endDate: et,
		flags: f,
	}
	return tb
}


