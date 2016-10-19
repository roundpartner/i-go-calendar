package ical_test

import (
    "testing"
    "strings"
    "ical/event"
)

func createTestEvent() string {
    e := event.CreateEvent("unique@id")
    return e.ToString()
}

func TestEventContainsABeginVEvent(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "BEGIN:VEVENT")) {
        t.Error("event did not have a begin event")
    }
}

func TestEventContainsAnEndVEvent(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "END:VEVENT")) {
        t.Error("event did not have an end event")
    }
}

func TestEventContainsUid(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "unique@id")) {
        t.Error("event event does not have an id")
    }
}
