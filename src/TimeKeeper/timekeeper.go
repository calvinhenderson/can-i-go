package timekeeper

import (
	"log"
	"strings"
	"time"
	"os"
	"encoding/csv"
)

type techTime struct {
	date  	   string     // The current date formated as YYYY-MM-DD -> 2024-02-19
	startH     string     // The start time for the office -> 8:00 am
	endH       string     // The end time for the office -> 2:30 pm
	breakTimeH string     // The time when going on break  -> 11:15 am
	breakLen   string     // The length of the break formated in HM -> 03:04

}




// Returns the time.Time passed to it as a string formatted in 12h -> 3:04:05 pm (HH:MM:SS)
func FormatTime(t time.Time) string {
	format := "3:04 pm"
	f := t.Format(format)
	return f
}

// Returns the time.Time passed to it as a string formatted in YYYY-MM-DD -> 2023-01-01
func FormatDate(t time.Time) string {
	format := "2006-01-02"
	f := t.Format(format)
	return f
}

// Returns a string formatted as 3:04:05 pm (HH:MM:SS) and a string formatted as YYYY-MM-DD as a time.Time -> 00:03:27 0000-01-01 11:15:00 +0000 UTC
func ConvertTime(s string,d string) (time.Time,error) {
	var format string
	if strings.Contains(strings.ToLower(s),"pm") || strings.Contains(strings.ToLower(s),"am"){
		format =  "2006-01-02 3:04 pm"
		s = d +" "+ s 
	} else {
		format = "03:04"
		
	}
	
	
	loc, e := time.LoadLocation("America/New_York")
	t,e := time.ParseInLocation(format,s,loc)
	return t,e
}

// Returns true or false depending if the current time is within the Tech Office hours and if it is before or after break
func (t techTime) IsOpen() (bool,error) {
	start,err := ConvertTime(t.startH,t.date)
	if err != nil {
		log.Println(err)
	}
	end, err := ConvertTime(t.endH,t.date)
	if err != nil {
		log.Println(err)
	}
	brk, err := ConvertTime(t.breakTimeH,t.date)
	if err != nil {
		log.Println(err)
	}
	brkLen,err := ConvertTime(t.breakLen,t.date)
	if err != nil {
		log.Println(err)
	}

	ct := time.Now()
	open := false

	if ct.After(start) && ct.Before(end){
		distance := brk.Add(time.Minute * time.Duration(brkLen.Minute())) // Gets the break time + the break length in minutes
		
		if ct.Before(brk) || ct.After(distance){
			open = true
		}
	}
	return open,err
}

// Creates a new techTime Struct
func NewTechTime(d string, st string, et string, bkt string, bkl string) techTime {
	tt := techTime{
		date:d,
		startH: st,
		endH: et,
		breakTimeH: bkt,
		breakLen: bkl,
	}
	return tt
}


	
// Creats a slice of techTimes from a csv file at a filepath given as a string
func NewTechTimesFromCSV(f string) ([]techTime,error) {
	file, err := os.Open(f) 
      
    
    if err != nil { 
        log.Println(err) 
    } 
  
    
    defer file.Close() 
  
    reader := csv.NewReader(file) 
      
    records, err := reader.ReadAll() 

    if err != nil { 
		log.Println(err)
    } 
    var tts []techTime

    for _, eachrecord := range records[1:]  { 
		var tt techTime
        for x,r := range eachrecord{
			switch x {
			case 0:
				tt.date = r
			case 1:
				tt.startH = r
			case 2:
				tt.endH = r
			case 3:
				tt.breakTimeH = r
			case 4:
				tt.breakLen = r
			}
		}
		tts = append(tts, tt)
    } 
	return tts,err
}


	



