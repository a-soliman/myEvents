package msgqueue

// EventEmitter interface
type EventEmitter interface {
	Emit(event Event) error
}
