package catowner

import (
	"unordered-events/domain"
	"unordered-events/pkg/stringer"
)

type CatOwner struct {
	domain.Entity

	Name string

	CatID  int
	CatAge int
}

func (co *CatOwner) String() string {
	return stringer.ToString(co)
}

func (co *CatOwner) ActualizeCatAge(actualCatAge int) {
	co.CatAge = actualCatAge
}

type ActualizeCatBirthdayTaskCreatedEvent struct {
	CatID          int
	NewCatAge      int
	CatOwnerLimit  int
	CatOwnerOffset int
}

func (e *ActualizeCatBirthdayTaskCreatedEvent) String() string {
	return stringer.ToString(e)
}

type Repository interface {
	ListByCatID(catID, offset, limit int) []*CatOwner
	CountCatOwnersByCatID(catID int) int
	Update(aCatOwner *CatOwner)
}
