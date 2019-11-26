package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type event struct {
	id   int
	kind eventKind
	time time.Time
}

func (e event) String() string {
	date := e.time.Format("01/02 15:04")
	switch e.kind {
	case eventStart:
		return fmt.Sprintf("[%s] Guard #%d starts", date, e.id)
	case eventAsleep:
		return fmt.Sprintf("[%s] Guard #%d sleeps", date, e.id)
	case eventAwake:
		return fmt.Sprintf("[%s] Guard #%d wakes", date, e.id)
	}

	return fmt.Sprintf("unknown event type: %#v", e)
}

type eventKind byte

const (
	eventStart eventKind = iota
	eventAsleep
	eventAwake
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)

	var events []event

	for _, line := range lines {
		dateStart := strings.Index(line, "]")
		dateText := line[1:dateStart]
		date, err := time.Parse("2016-01-02 15:04", dateText)
		if err != nil {
			log.Fatalf("could not parse date %q: %v", dateText, err)
		}
		e := event{time: date}
		pieces := strings.Fields(line[dateStart+2:])
		switch pieces[0] {
		case "Guard":
			id, err := strconv.Atoi(pieces[1][1:])
			if err != nil {
				log.Fatalf("could not parse id %q: %v", pieces[1][1:], err)
			}
			e.id = id
			e.kind = eventStart
		case "falls":
			e.id = events[len(events)-1].id
			e.kind = eventAsleep
		case "wakes":
			e.id = events[len(events)-1].id
			e.kind = eventAwake

		}
		events = append(events, e)
	}

	id, minute := findGuard(events)
	fmt.Println(id * minute)

}

func findGuard(events []event) (id, minute int) {
	for _, e := range events {
		fmt.Println(e)
	}
	return 0, 0

}
