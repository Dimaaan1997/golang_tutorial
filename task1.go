package main

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

type Event struct {
	ID        int
	Type      string
	Data      string
	Timestamp time.Time
}

type EventStore struct {
	Events map[int]Event
	nextId int
}

func NewEventStore() *EventStore {
	newStore := EventStore{Events: make(map[int]Event, 1000), nextId: 0}
	return &newStore
}

func (es *EventStore) Add(eventType string, data string) int {
	now := time.Now().UTC()
	newEvent := Event{ID: es.nextId, Type: eventType, Data: data, Timestamp: now}
	prevID := es.nextId
	es.Events[es.nextId] = newEvent
	es.nextId += 1
	return prevID
}

func (es *EventStore) GetAll() []Event {
	// решение №1
	events := slices.Collect(maps.Values(es.Events))

	//  решение №2
	// events := []Event{}
	// for _, event := range es.Events {
	// 	events = append(events, event)
	// }
	return events
}

func (es *EventStore) GetByID(id int) (Event, bool) {
	needEvent, exists := es.Events[id]
	return needEvent, exists

}

func (es *EventStore) Count() int {
	count := len(es.Events)
	return count
}

func (es *EventStore) GetByType(eventType string) []Event {
	events := []Event{}
	for _, event := range es.Events {
		if event.Type == eventType {
			events = append(events, event)
		}
	}
	return events
}

func (es *EventStore) FindAfter(timestamp time.Time) []Event {
	needEvents := []Event{}
	for _, event := range es.Events {
		if event.Timestamp.After(timestamp) {
			needEvents = append(needEvents, event)
		}
	}
	return needEvents
}

func (es *EventStore) GetRange(startID, endID int) []Event {
	events := []Event{}
	for _, event := range es.Events {
		if event.ID >= startID && event.ID <= endID {
			events = append(events, event)
		}

	}
	return events
}

func (es *EventStore) Filter(predicate func(Event) bool) []Event {
	// TODO: вернуть события, для которых predicate вернул true
	events := []Event{}
	for _, event := range es.Events {
		predict_result := predicate(event)
		if predict_result {
			events = append(events, event)
		}

	}
	return events
}

func main_task1() {
	eventStore := NewEventStore()
	id := eventStore.Add("Like", "Hello World")
	fmt.Println(id)

	id2 := eventStore.Add("Like", "Hello World")
	fmt.Println(id2)

	id3 := eventStore.Add("Like", "Hello World")
	fmt.Println(id3)

	allValues := eventStore.GetAll()
	fmt.Println(allValues)

	ev, is := eventStore.GetByID(1)
	fmt.Println(ev, is)

	count := eventStore.Count()
	fmt.Println(count)

	likeEvents := eventStore.GetByType("Like")
	fmt.Println(likeEvents)

	now := time.Date(2026, 7, 2, 0, 0, 0, 0, time.UTC)

	afterEvents := eventStore.FindAfter(now)

	fmt.Println(afterEvents)

	rangeEvents := eventStore.GetRange(1, 1)

	fmt.Println(rangeEvents)

	predicate := func(event Event) bool { return event.ID%2 == 0 }

	predicateResult := eventStore.Filter(predicate)

	fmt.Println(predicateResult)

}
