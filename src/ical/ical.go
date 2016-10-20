package ical

import (
    "ical/event"
    "bytes"
)

type Calendar struct {
    events []event.Event
}

func (c Calendar) ToString() string {
    var buffer bytes.Buffer
    for _, event := range c.events {
        buffer.WriteString(event.ToString())
    }
    return "BEGIN:VCALENDAR\n" +
        buffer.String() +
        "END:VCALENDAR\n"
}

func (c Calendar) AddEvent(event event.Event) Calendar {
    events := append(c.events, event)
    return Calendar{events:events}
}

func MakeCalendar() Calendar {
    events := make([]event.Event, 0)
    c := Calendar{events:events}
    return c
}

func CreateCalendar() string {
    c := MakeCalendar()
    return c.ToString()
}
