package event

import "time"

type Event struct {
    uid string
    date time.Time
    summary string
    location string
    description string
}

func (e Event) header() string {
    return "BEGIN:VEVENT"
}

func (e Event) footer() string {
    return "END:VEVENT"
}

func (e Event) ToString() string {
    return e.header() + "\n" +
        "UID:" + e.uid + "\n" +
        "DTSTART;VALUE=DATE:" + e.date.Format("20060102") + "\n" +
        "DTEND;VALUE=DATE:" + e.date.Format("20060102") + "\n" +
        "SUMMARY:" + e.summary + "\n" +
        "LOCATION:" + e.location + "\n" +
        "DESCRIPTION:" + e.description + "\n" +
        e.footer() + "\n"
}

func MakeEvent(uid string, date time.Time, summary string, location string, description string) Event {
    e := Event{
        uid:uid,
        date:date,
        summary:summary,
        location:location,
        description:description,
    }
    return e
}
