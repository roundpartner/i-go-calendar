package ical

import (
    "testing"
    "strings"
    "ical/event"
)

func TestCreateCalendarBeginsCalendar(t *testing.T) {
    if (!strings.Contains(CreateCalendar(), "BEGIN:VCALENDAR")) {
        t.Error("Begin calendar not written")
    }
}

func TestCreateCalendarEndsCalendar(t *testing.T) {
    if (!strings.Contains(CreateCalendar(), "END:VCALENDAR")) {
        t.Error("End calendar not written")
    }
}

func TestAddEvent(t *testing.T) {
    c := MakeCalendar()
    c = c.AddEvent(event.MakeEvent("test@event", "20161216", "summary content", "London", "description content"))
    if (!strings.Contains(c.ToString(), "UID:test@event")) {
        t.Error("Calendar event was not added")
    }
}
