package ical

import (
    "testing"
    "strings"
    "github.com/roundpartner/i-go-calendar/src/ical/event"
    "time"
)

func makeCalendar() Calendar {
    return MakeEmptyCalendar("Personal Events")
}

func createCalendarString() string {
    return CreateCalendar("Personal Events")
}

func TestCreateCalendarBeginsCalendar(t *testing.T) {
    if (!strings.Contains(createCalendarString(), "BEGIN:VCALENDAR")) {
        t.Error("Begin calendar not written")
    }
}

func TestCreateCalendarEndsCalendar(t *testing.T) {
    if (!strings.Contains(createCalendarString(), "END:VCALENDAR")) {
        t.Error("End calendar not written")
    }
}

func TestCalendarNameIsSet(t *testing.T) {
    if (!strings.Contains(createCalendarString(), "X-WR-CALNAME:Personal Events")) {
        t.Error("Calendar does not have a name")
    }
}

func TestAddEvent(t *testing.T) {
    c := makeCalendar()
    date := time.Date(2016, time.December, 16, 0, 0, 0, 0, time.UTC)
    c = c.AddEvent(event.MakeEvent("test@event", date, "summary content", "London", "description content"))
    if (!strings.Contains(c.ToString(), "UID:test@event")) {
        t.Error("Calendar event was not added")
    }
}
