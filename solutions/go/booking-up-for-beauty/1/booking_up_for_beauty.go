package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	time, _ := time.Parse(layout, date)
    return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	timeParsed, _ := time.Parse(layout, date)
    // diff := time.Since(timeParsed)
    diff := time.Now().Sub(timeParsed)
    fmt.Println(diff)
    if diff > 0 {
        return true
    }
    return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    timeParsed, _ := time.Parse(layout, date)
    if timeParsed.Hour() >= 12 && timeParsed.Hour() < 18 {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    time, _ := time.Parse(layout, date)
    formattedMes := time.Format("Monday, January 2, 2006, at 15:04")
    mess := fmt.Sprintf("You have an appointment on %s.", formattedMes)
	return mess
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),time.September,15,0,0,0,0,time.UTC)
}
