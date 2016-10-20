package main

import (
    "fmt"
    "ical"
    "ical/event"
)
func main() {
    c := ical.MakeCalendar()
    e := event.MakeEvent("test@event")
    c = c.AddEvent(e)
    fmt.Print(c.ToString())
}
