package contracts

import "time"

// EventCreatedEvent and event that firest when an event was created
type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

// EventName returns the event flag
func (e *EventCreatedEvent) EventName() string {
	return "event.created"
}
