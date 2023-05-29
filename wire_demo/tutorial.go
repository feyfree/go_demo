package wire_demo

import (
	"errors"
	"fmt"
	"time"
)

type Message string

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}
type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

//func NewGreeter(m Message) Greeter {
//	return Greeter{Message: m}
//}

func (g Greeter) Greet() Message {
	return g.Message
}

//func NewEvent(g Greeter) Event {
//	return Event{Greeter: g}
//}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {

}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

//func NewEventNumber() int {
//	return 1
//}

type EventNumber int

func CreateEventNumber() EventNumber {
	return 1
}
