package postgres

import (
	"context"
	"fmt"

	"unordered-events/pkg/slice"
)

type entity interface {
	fmt.Stringer

	GetID() int
	SetID(int)
}

type abstractRepository[T entity] struct {
	entities []T
}

func (r *abstractRepository[T]) Add(newEntity T) int {
	newEntity.SetID(len(r.entities) + 1)

	r.entities = append(r.entities, newEntity)

	return newEntity.GetID()
}

func (r *abstractRepository[T]) Get(id int) T {
	for _, entity := range r.entities {
		if entity.GetID() == id {
			return entity
		}
	}

	return *new(T)
}

func (r *abstractRepository[T]) Update(T) {
}

func (r *abstractRepository[T]) Delete(_ context.Context, id int) {
	for index, entity := range r.entities {
		if entity.GetID() == id {
			r.entities = slice.Remove(r.entities, index)

			return
		}
	}
}

func (r *abstractRepository[T]) Entities() []T {
	return r.entities
}
