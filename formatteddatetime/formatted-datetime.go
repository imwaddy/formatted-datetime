package formatteddatetime

import (
	"strconv"
	"time"

	"github.com/araddon/dateparse"
)

const (
	dash  = "-"
	slash = "/"
	space = " "
	dot   = "."
)

// TimeStruct is a Receiver struct
// Which accepts current time through API
//
//ex. formatteddatetime.TimeStruct{CurrentTime: time.Now()}
type TimeStruct struct {
	CurrentTime time.Time
}

// ConvertUTCFormatDate : This service will return an UTC timestamp formatted date and error.
// This service has a single parameter string date.
func ConvertUTCFormatDate(date string) (time.Time, error) {
	return dateparse.ParseLocal(date)
}

// GetFormattedDatetime : This service will return an formatted datetime.
// This service has a single parameter i.e. Format string which is any of 4  dash(-), slash(/), space( ), and dot(.).
// this will returns output as  four formats like 30-Jul-2019 10:25:37 AM, 30/Jul/2019 10:25:38 AM, 30 Jul 2019 10:25:39 AM, 30.Jul.2019 10:25:40 AM
func (t *TimeStruct) GetFormattedDatetime(format string) string {
	ampm := "AM"

	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := t.CurrentTime.In(loc)

	minutes := strconv.Itoa(now.Minute())
	seconds := strconv.Itoa(now.Second())

	if now.Second() < 10 {
		seconds = "0" + seconds
	}

	if now.Minute() < 10 {
		minutes = "0" + minutes
	}

	var date string

	switch format {
	case dash:
		date = now.Format("02-Jan-2006")
	case slash:
		date = now.Format("02/Jan/2006")
	case space:
		date = now.Format("02 Jan 2006")
	case dot:
		date = now.Format("02.Jan.2006")
	}

	hourTemp := now.Hour()

	if hourTemp >= 12 {
		if hourTemp == 12 {
			ampm = "PM"
		} else {
			hourTemp -= 12
			ampm = "PM"
		}
	}
	hour := strconv.Itoa(hourTemp)

	if hourTemp < 10 {
		hour = "0" + hour
	}

	return date + " " + hour + ":" + minutes + ":" + seconds + " " + ampm
}

// GetFormattedDate : This service will return onlu date with some formatting.
func (t *TimeStruct) GetFormattedDate(format string) string {
	// Get date-time from GetFormattedDatetime method
	datetime := t.GetFormattedDatetime(format)

	return datetime[:11]
}

// GetFormattedTime : This service will return the formatted time.
func (t *TimeStruct) GetFormattedTime() string {
	// Get date-time from GetFormattedDatetime method
	// default "/" set to method as it has no use in returning time
	datetime := t.GetFormattedDatetime("/")

	return datetime[12:len(datetime)]
}
