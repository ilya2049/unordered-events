package domain

type Event any

type EventPublisher interface {
	Publish(...Event)
}
