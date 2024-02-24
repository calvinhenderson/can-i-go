package timekeeper

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"
)


const (
	OPEN uint8 = 1 << iota
	LM //Low Man -> short staffed
	HD // Holiday
	ED // Early Dismissal
)



type TimeBlock struct {
	startDate time.Time
	endDate   time.Time
	flags      uint8
  }


/* Sets a bit of a uint to a desired flag -> 
EX. b = 0010, flag = 1001. Becomes 1011
*/ 
func Set(b, flag uint8) uint8 {
	return b | flag
}

/* Clears the desired flag from a uint -> 
EX. b = 0011, flag = 0001. Becomes 0010
*/ 
func Clear(b, flag uint8) uint8 {
	return b &^ flag
}

/* Clears all the flags (bits) of a uint and sets them to 0  -> 
EX. b of 0011 Becomes 0000
*/ 
func ClearAll(b uint8) uint8 {
	return (b & 0)
}

/* Toggles the desired flag from a uint -> 
EX. b = 0011, flag = 0001. Becomes 0010
*/ 
func Toggle(b, flag uint8) uint8 {
	return b ^ flag
}

/* Checks to see if uint has the desired flag -> 
EX. b = 0001, flag = 0001 -> true
*/ 
func Has(b, flag uint8) bool {
	return b&flag != 0
}

/* Checks to see if open by checking if the OPEN flag is in TimeBlock.flags -> 
EX. b = 0011, flag = 0001 -> true
*/ 
func (t TimeBlock) IsOpen() bool {
	return Has(t.flags,OPEN)
}

/* Checks to see if the current time is between t.endDate & tt.startDate -> 
EX. t.endDate = 11:00 am - tt.startDate = 12:00 pm - current time = 11:15 am -> true
*/ 
func (t TimeBlock) IsBetween(tt TimeBlock) bool {
	return time.Now().After(t.endDate) && time.Now().Before(tt.startDate)
}

/* Returns the time between tt.startDate & t.endDate
EX. t.endDate = 11:00 am - tt.startDate = 12:00 pm -> 1h0m0s
*/ 
func (t TimeBlock) TimeBetween(tt TimeBlock) time.Duration {
	return tt.startDate.Sub(t.endDate)
}

/* Returns the time between the current time and t.startDate or t.endDate
 return depends on if start is set to true or false
EX. t.startDate = 12:00 pm - time.Now() = 11:00 am -> 1h0m0s
*/ 
func (t TimeBlock) TimeTil(start bool) time.Duration{
	if start {
		return time.Until(t.startDate)
	} else {
		return time.Until(t.endDate)
	}
}


/* Converts a string of time and date into a time.Time struct ->
EX. s = 2:30 pm - d = 2024-02-25 becomes 2024-02-25 14:30:00 -0500 EST
*/
func ConvertTime(s string,d string) time.Time {
	
	format :=  "2006-01-02 3:04 pm"
	s = d +" "+ s 
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Println(err)
	}
	t,err := time.ParseInLocation(format,s,loc)
	if err != nil {
		log.Println(err)
	}
	return t
}


/* Creates and returns a new TimeBlock Struct
startTime and endTime need to be formatted as kitchen time -> 2:40 pm

date needs to be formated in YYYY-MM-DD -> 2024-02-25

flags needs to hold at least one flag -> {OPEN, HD, LM, ED}
*/
func NewTimeBlock(startTime string, endTime string, date string, flags []uint8) TimeBlock{
	st := ConvertTime(startTime,date)
	

	et := ConvertTime(endTime,date)
	

	var f uint8
	for _,flag := range flags {
		f = Set(f,flag)
	}

	tb := TimeBlock{
		startDate: st,
		endDate: et,
		flags: f,
	}
	return tb
}

/* Creates and returns a new TimeBlock Struct
Uses time.Time for StartDate and EndTime args

flags needs to hold at least one flag -> {OPEN, HD, LM, ED}
*/
func NewTimeBlockNoStrings(StartDate time.Time, EndTime time.Time, flags []uint8) TimeBlock{
	
	var f uint8
	for _,flag := range flags {
		f = Set(f,flag)
	}

	tb := TimeBlock{
		startDate: StartDate,
		endDate: EndTime,
		flags: f,
	}
	return tb
}


func NewTimeBlocksFromCSV(fpath string) []TimeBlock {
	file, err := os.Open(fpath) 
      
    
    if err != nil { 
        log.Println(err) 
    } 
  
    
    defer file.Close() 
  
    reader := csv.NewReader(file) 
      
    records, err := reader.ReadAll() 

    if err != nil { 
		log.Println(err)
    } 

    var tb []TimeBlock

    for _, eachrecord := range records[1:]  { 
		var flags []uint8
		fgs := strings.Split(eachrecord[3], ",")
		for _,f := range fgs {
			if strings.Contains(f,"OPEN") {
				flags = append(flags, OPEN)
			}
			if strings.Contains(f,"LM") {
				flags = append(flags, LM)
			}
			if strings.Contains(f,"HD") {
				flags = append(flags, HD)
			}
			if strings.Contains(f,"ED") {
				flags = append(flags, ED)
			}
		}
		tb = append(tb, NewTimeBlock(eachrecord[0],eachrecord[1],eachrecord[2],flags))
    } 
	return tb
}



