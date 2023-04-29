package kafka

import (
	"fmt"

	"unordered-events/commands"
	"unordered-events/domain"
	"unordered-events/domain/cat"
	"unordered-events/domain/catowner"
	"unordered-events/pkg/stringer"
)

type EventHandler struct {
	eventBus                 <-chan domain.Event
	catOwnerCommandExecutors *commands.CatOwnerCommandExecutors
}

func NewEventHandler(
	eventBus <-chan domain.Event,
	catOwnerCommandExecutors *commands.CatOwnerCommandExecutors,
) *EventHandler {
	return &EventHandler{
		eventBus:                 eventBus,
		catOwnerCommandExecutors: catOwnerCommandExecutors,
	}
}

func (h *EventHandler) ListenAndHandle() {
	for e := range h.eventBus {
		h.handleEvent(e)
	}
}

func (h *EventHandler) handleEvent(e domain.Event) {
	fmt.Println("Event received:", stringer.TypeOf(e), stringer.ToString(e))

	switch anEvent := e.(type) {
	case *cat.BirthdayCelebratedEvent:
		h.catOwnerCommandExecutors.CreateTasksToActualizeCatBirthday.Execute(
			anEvent.CatID,
			anEvent.NewAge,
		)

	case *catowner.ActualizeCatBirthdayTaskCreatedEvent:
		h.catOwnerCommandExecutors.ActualizeCatBirthday.Execute(
			anEvent.CatID,
			anEvent.NewCatAge,
			anEvent.CatOwnerOffset,
			anEvent.CatOwnerLimit,
		)
	}
}
