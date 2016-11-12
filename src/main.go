package main

import (
    "fmt"
    "ical"
    "github.com/roundpartner/i-go-calendar/src/ical/event"
    "time"
)
func main() {
    c := ical.MakeEmptyCalendar("Personal Events")
    date := time.Date(2016, time.November, 16, 0, 0, 0, 0, time.UTC)
    c = c.AddEvent(event.MakeEvent("1@event", date, "We got engaged 1 year ago today", "London", "We got engaged 1 year ago today"))
    date = date.AddDate(0, 1, 0)
    c = c.AddEvent(event.MakeEvent("2@event", date, "Engagement Anniversary - 1 Year 1 Month", "London", "We got engaged 1 year and 1 month ago today"))
    fmt.Print(c.ToString())
}
