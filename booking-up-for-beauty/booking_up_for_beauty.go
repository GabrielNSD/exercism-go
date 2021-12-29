package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()
	parsed, _ := time.Parse("January 2, 2006 15:04:05", date)
	return now.After(parsed)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	parsed, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := parsed.Hour()
	if hour >= 12 && hour <= 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parsed := Schedule(date)
	year, month, day := parsed.Date()

	resultString := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", parsed.Weekday().String(), month.String(), day, year, parsed.Hour(), parsed.Minute())

	return resultString
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
