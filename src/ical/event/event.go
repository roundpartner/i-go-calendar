package event

type Event struct {
    uid string
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
        e.footer() + "\n"
}

func CreateEvent(uid string) Event {
    e := Event{uid:uid}
    return e
}
