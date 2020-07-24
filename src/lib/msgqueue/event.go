package msgqueue

// Event interface
type Event interface {
	EventName() string
}
