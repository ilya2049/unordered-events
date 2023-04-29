package domain

type Entity struct {
	ID int
}

func (e *Entity) GetID() int {
	return e.ID
}

func (e *Entity) SetID(id int) {
	e.ID = id
}
