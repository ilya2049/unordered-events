package cat

import (
	"unordered-events/domain"
	"unordered-events/pkg/stringer"
)

type Cat struct {
	domain.Entity

	Name string
	Age  int
}

func (c *Cat) String() string {
	return stringer.ToString(c)
}

func (c *Cat) CelebrateBirthday() []domain.Event {
	c.Age += 1

	return []domain.Event{
		&BirthdayCelebratedEvent{
			CatID:  c.ID,
			NewAge: c.Age,
		},
	}
}

type BirthdayCelebratedEvent struct {
	CatID  int
	NewAge int
}

func (e *BirthdayCelebratedEvent) String() string {
	return stringer.ToString(e)
}

type Repository interface {
	Get(id int) *Cat
	Update(aCat *Cat)
}
