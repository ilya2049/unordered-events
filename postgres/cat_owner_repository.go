package postgres

import (
	"unordered-events/domain/catowner"
)

type CatOwnerRepository struct {
	abstractRepository[*catowner.CatOwner]
}

func (r CatOwnerRepository) ListByCatID(catID, offset, limit int) []*catowner.CatOwner {
	catOwners := r.filterByCatID(catID)

	start := offset
	end := offset + limit

	if end > len(catOwners) {
		end = len(catOwners)
	}

	return catOwners[start:end]
}

func (r CatOwnerRepository) CountCatOwnersByCatID(catID int) int {
	return len(r.filterByCatID(catID))
}

func (r *CatOwnerRepository) filterByCatID(catID int) []*catowner.CatOwner {
	var catOwners []*catowner.CatOwner

	for _, aCatOwner := range r.entities {
		if (*catowner.CatOwner)(aCatOwner).CatID == catID {
			catOwners = append(catOwners, aCatOwner)
		}
	}

	return catOwners
}
