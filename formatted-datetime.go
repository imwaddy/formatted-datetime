package formatteddatetime

import (
	"time"
)

var loc = mustLoadIST()

// mustLoadIST loads Asia/Kolkata, falling back to a fixed +5:30 offset
// if the system has no tzdata (LoadLocation would otherwise return a
// nil *Location, which panics on Time.In).
func mustLoadIST() *time.Location {
	l, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		return time.FixedZone("IST", 5*60*60+30*60)
	}
	return l
}

var months = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

func isValidSep(format string) bool {
	switch format {
	case "-", "/", " ", ".":
		return true
	default:
		return false
	}
}

// put2 writes a zero-padded 2-digit number.
func put2(b []byte, v int) {
	b[0] = byte('0' + v/10)
	b[1] = byte('0' + v%10)
}

// appendDate writes "02<sep>Jan<sep>2006" into b, returning the new length.
func appendDate(b []byte, sep byte, day int, mon time.Month, year int) []byte {
	put2(b[0:2], day)
	b[2] = sep
	m := months[mon-1]
	b[3], b[4], b[5] = m[0], m[1], m[2]
	b[6] = sep
	y := year
	b[7] = byte('0' + y/1000)
	b[8] = byte('0' + (y/100)%10)
	b[9] = byte('0' + (y/10)%10)
	b[10] = byte('0' + y%10)
	return b[:11]
}

// appendClock writes "03:04:05 PM" (12-hour clock) into b.
func appendClock(b []byte, hour, min, sec int) []byte {
	h := hour % 12
	if h == 0 {
		h = 12
	}
	ampm := "AM"
	if hour >= 12 {
		ampm = "PM"
	}
	put2(b[0:2], h)
	b[2] = ':'
	put2(b[3:5], min)
	b[5] = ':'
	put2(b[6:8], sec)
	b[8] = ' '
	b[9], b[10] = ampm[0], ampm[1]
	return b[:11]
}

// TimeStruct is a Receiver struct
// Which accepts current time through API
//
// ex. formatteddatetime.TimeStruct{CurrentTime: time.Now()}
type TimeStruct struct {
	CurrentTime time.Time
}

// GetFormattedDatetime : This service will return a formatted datetime.
// This service has a single parameter i.e. Format strings are any of 4 dash(-), slash(/), space( ), and dot(.).
// this will returns output as  four formats like 30-Jul-2019 10:25:37 AM, 30/Jul/2019 10:25:38 AM, 30 Jul 2019 10:25:39 AM, 30.Jul.2019 10:25:40 AM
func (t *TimeStruct) GetFormattedDatetime(format string) string {
	if len(format) != 1 || !isValidSep(format) {
		return ""
	}
	tm := t.CurrentTime.In(loc)
	y, mon, d := tm.Date()
	var buf [23]byte
	appendDate(buf[:11], format[0], d, mon, y)
	buf[11] = ' '
	appendClock(buf[12:23], tm.Hour(), tm.Minute(), tm.Second())
	return string(buf[:])
}

// GetFormattedDate : This service will return only date with some formatting.
// valid arguments are any of 4 dash(-), slash(/), space( ), and dot(.)
func (t *TimeStruct) GetFormattedDate(format string) string {
	if len(format) != 1 || !isValidSep(format) {
		return ""
	}
	tm := t.CurrentTime.In(loc)
	y, mon, d := tm.Date()
	var buf [11]byte
	appendDate(buf[:], format[0], d, mon, y)
	return string(buf[:])
}

// GetFormattedTime : This service will return the formatted time.
func (t *TimeStruct) GetFormattedTime() string {
	tm := t.CurrentTime.In(loc)
	var buf [11]byte
	appendClock(buf[:], tm.Hour(), tm.Minute(), tm.Second())
	return string(buf[:])
}
