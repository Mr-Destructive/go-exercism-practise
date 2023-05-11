package booking

import (
	"fmt"
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	a := Schedule(date)
	description := "You have an appointment on " + a.Weekday().String() + a.Format(", January 2, 2006, at 15:04") + "."
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	date := strconv.Itoa(year) + "-09-15 00:00:00 +0000 UTC"
	fmt.Println(date)
	anniversary_date, err := time.Parse("2006-1-2 15:04:05 +0000 UTC", date)
	if err != nil {
		panic(err)
	}
	return anniversary_date
}
