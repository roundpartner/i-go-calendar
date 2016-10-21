package event

import (
    "testing"
    "strings"
)

func createTestEvent() Event {
    return MakeEvent("test@event", "20161216", "summary content", "London", "description content")
}

func createTestEventString() string {
    e := createTestEvent()
    return e.ToString()
}

func TestEventContainsABeginVEvent(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "BEGIN:VEVENT")) {
        t.Error("event did not have a begin event")
    }
}

func TestEventContainsAnEndVEvent(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "END:VEVENT")) {
        t.Error("event did not have an end event")
    }
}

func TestEventContainsUid(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "UID:test@event")) {
        t.Error("event event does not have an id")
    }
}

func TestEventContainsStartDate(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "DTSTART;VALUE=DATE:20161216")) {
        t.Error("event event does not have a start date")
    }
}

func TestEventContainsEndDate(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "DTEND;VALUE=DATE:20161216")) {
        t.Error("event event does not have an end date")
    }
}

func TestEventContainsSummary(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "SUMMARY:summary content")) {
        t.Error("event event does not have a summary")
    }
}

func TestEventContainsLocation(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "LOCATION:London")) {
        t.Error("event event does not have a location")
    }
}

func TestEventContainsDescription(t *testing.T) {
    if (!strings.Contains(createTestEventString(), "DESCRIPTION:description content")) {
        t.Error("event event does not have a description")
    }
}
