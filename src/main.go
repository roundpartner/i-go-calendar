package main

import (
    "fmt"
    "ical"
    "github.com/roundpartner/i-go-calendar/src/ical/event"
)
func main() {
    c := ical.MakeEmptyCalendar("Personal Events")
    c = c.AddEvent(event.MakeEvent("1@event", "20161116", "We got engaged 1 year ago today", "London", "We got engaged 1 year ago today"))
    c = c.AddEvent(event.MakeEvent("2@event", "20161216", "Engagement Anniversary - 1 Year 1 Month", "London", "We got engaged 1 year and 1 month ago today"))
    fmt.Print(c.ToString())
}
