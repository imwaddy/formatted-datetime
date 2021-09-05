package formatteddatetime

import (
	"time"
)

var (
	dateTimeFormatter = map[string]string{
		"-": "02-Jan-2006 03:04:05 PM",
		"/": "02/Jan/2006 03:04:05 PM",
		" ": "02 Jan 2006 03:04:05 PM",
		".": "02.Jan.2006 03:04:05 PM",
	}
	dateFormatter = map[string]string{
		"-": "02-Jan-2006",
		"/": "02/Jan/2006",
		" ": "02 Jan 2006",
		".": "02.Jan.2006",
	}
	blankString = ""
	loc, _      = time.LoadLocation("Asia/Kolkata")
)

// TimeStruct is a Receiver struct
// Which accepts current time through API
//
//ex. formatteddatetime.TimeStruct{CurrentTime: time.Now()}
type TimeStruct struct {
	CurrentTime time.Time
}

// GetFormattedDatetime : This service will return a formatted datetime.
// This service has a single parameter i.e. Format strings are any of 4 dash(-), slash(/), space( ), and dot(.).
// this will returns output as  four formats like 30-Jul-2019 10:25:37 AM, 30/Jul/2019 10:25:38 AM, 30 Jul 2019 10:25:39 AM, 30.Jul.2019 10:25:40 AM
func (t *TimeStruct) GetFormattedDatetime(format string) string {

	value, ok := dateTimeFormatter[format]
	if ok {
		return t.CurrentTime.In(loc).Format(value)
	}

	return blankString
}

// GetFormattedDate : This service will return only date with some formatting.
// valid arguments are any of 4 dash(-), slash(/), space( ), and dot(.)
func (t *TimeStruct) GetFormattedDate(format string) string {

	value, ok := dateFormatter[format]
	if ok {
		return t.CurrentTime.In(loc).Format(value)
	}

	return blankString

}

// GetFormattedTime : This service will return the formatted time.
func (t *TimeStruct) GetFormattedTime() string {

	return t.CurrentTime.In(loc).Format("03:04:05 PM")

}
