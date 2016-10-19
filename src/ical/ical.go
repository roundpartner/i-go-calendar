package main

import (
    "ical/event"
    "fmt"
)

func main() {
    e := event.CreateEvent("test")
    fmt.Print(e.ToString())
}
