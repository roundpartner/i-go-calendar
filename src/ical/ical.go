package ical

import (
    "github.com/roundpartner/i-go-calendar/src/ical/event"
    "bytes"
)

type Calendar struct {
    name string
    events []event.Event
}

func (c Calendar) ToString() string {
    var buffer bytes.Buffer
    for _, event := range c.events {
        buffer.WriteString(event.ToString())
    }
    return "BEGIN:VCALENDAR\n" +
        "VERSION:2.0\n" +
        "PRODID:-//Thomas Lorentsen//Personal Events v0.1//EN\n" +
        "X-PUBLISHED-TTL:PT1H\n" +
        "CALSCALE:GREGORIAN\n" +
        "METHOD:PUBLISH\n" +
        "X-WR-CALNAME:" + c.name + "\n" +
        buffer.String() +
        "END:VCALENDAR\n"
}

func (c *Calendar) AddEvent(event event.Event) Calendar {
    events := append(c.events, event)
    c.events = events
    return MakeCalendar(c.name, events)
}

func MakeCalendar(name string, events []event.Event) Calendar {
    c := Calendar{name:name,events:events}
    return c
}

func MakeEmptyCalendar(name string) Calendar {
    events := *new([]event.Event)
    return MakeCalendar(name, events)
}

func CreateCalendar(name string) string {
    c := MakeEmptyCalendar(name)
    return c.ToString()
}
