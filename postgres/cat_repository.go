package postgres

import "unordered-events/domain/cat"

type CatRepository struct {
	abstractRepository[*cat.Cat]
}
