//go:build wireinject
// +build wireinject

package wire_demo

import "github.com/google/wire"

// wire.go

func InitializeEvent(phrase string) (Event, error) {
	// woops! NewEventNumber is unused.
	wire.Build(NewEvent, NewGreeter, NewMessage, CreateEventNumber)
	return Event{}, nil
}
