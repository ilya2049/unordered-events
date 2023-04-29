package kafka

import (
	"time"
	"unordered-events/domain"
	"unordered-events/pkg/slice"
)

type EventPublisher struct {
	input  chan domain.Event
	buffer []domain.Event
	output chan domain.Event
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		input:  make(chan domain.Event, 50),
		buffer: make([]domain.Event, 0, 3),
		output: make(chan domain.Event, 50),
	}
}

func (p *EventPublisher) ActivateEventBusAsync() <-chan domain.Event {
	go func() {
		ticker := time.NewTicker(time.Millisecond)

		for {
			select {
			case e := <-p.input:
				p.buffer = append(p.buffer, e)
				if len(p.buffer) == cap(p.buffer) {
					p.flushBuffer()
				}

			case <-ticker.C:
				p.flushBuffer()
			}
		}
	}()

	return p.output
}
func (p *EventPublisher) flushBuffer() {
	slice.Shuffle(p.buffer)

	for _, e := range p.buffer {
		p.output <- e
	}

	p.buffer = p.buffer[:0]
}

func (p *EventPublisher) Publish(events ...domain.Event) {
	for _, e := range events {
		p.input <- e
	}
}
