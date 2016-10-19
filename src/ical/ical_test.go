package ical_test

import (
    "testing"
    "strings"
    "ical"
)

func createTestEvent() string {
    e := ical.CreateEvent("unique@id")
    return e.ToString()
}

func TestEventContainsABeginVEvent(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "BEGIN:VEVENT")) {
        t.Error("ical did not have a begin event")
    }
}

func TestEventContainsAnEndVEvent(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "END:VEVENT")) {
        t.Error("ical did not have an end event")
    }
}

func TestEventContainsUid(t *testing.T) {
    if (!strings.Contains(createTestEvent(), "unique@id")) {
        t.Error("ical event does not have an id")
    }
}
