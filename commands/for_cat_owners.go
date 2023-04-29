package commands

import (
	"unordered-events/domain"
	"unordered-events/domain/cat"
	"unordered-events/domain/catowner"
)

type CatOwnerCommandExecutors struct {
	ActualizeCatBirthday              *ActualizeCatBirthdayCommandExecutor
	CreateTasksToActualizeCatBirthday *CreateTasksToActualizeCatBirthdayCommandExecutor
}

func NewCatOwnerCommandExecutors(
	catRepository cat.Repository,
	catOwnerRepository catowner.Repository,
	eventPublisher domain.EventPublisher,
) *CatOwnerCommandExecutors {
	return &CatOwnerCommandExecutors{
		ActualizeCatBirthday: &ActualizeCatBirthdayCommandExecutor{
			catRepository:      catRepository,
			catOwnerRepository: catOwnerRepository,
		},
		CreateTasksToActualizeCatBirthday: &CreateTasksToActualizeCatBirthdayCommandExecutor{
			catRepository:      catRepository,
			catOwnerRepository: catOwnerRepository,
			eventPublisher:     eventPublisher,
		},
	}
}

type ActualizeCatBirthdayCommandExecutor struct {
	catRepository      cat.Repository
	catOwnerRepository catowner.Repository
}

func (ex *ActualizeCatBirthdayCommandExecutor) Execute(
	catID,
	newCatAge,
	catOwnerOffset,
	catOwnerLimit int,
) {
	if aCat := ex.catRepository.Get(catID); aCat.Age != newCatAge {
		return
	}

	catOwnersBatch := ex.catOwnerRepository.ListByCatID(catID, catOwnerOffset, catOwnerLimit)

	for _, catOwner := range catOwnersBatch {
		catOwner.ActualizeCatAge(newCatAge)

		ex.catOwnerRepository.Update(catOwner)
	}
}

type CreateTasksToActualizeCatBirthdayCommandExecutor struct {
	catRepository      cat.Repository
	catOwnerRepository catowner.Repository
	eventPublisher     domain.EventPublisher
}

func (ex *CreateTasksToActualizeCatBirthdayCommandExecutor) Execute(catID, newCatAge int) {
	if aCat := ex.catRepository.Get(catID); aCat.Age != newCatAge {
		return
	}

	catOwnersTotal := ex.catOwnerRepository.CountByCatID(catID)

	const limit = 1

	var events []domain.Event

	for offset := 0; offset < catOwnersTotal; offset += limit {
		events = append(events, &catowner.ActualizeCatBirthdayTaskCreatedEvent{
			CatID:          catID,
			NewCatAge:      newCatAge,
			CatOwnerLimit:  limit,
			CatOwnerOffset: offset,
		})
	}

	ex.eventPublisher.Publish(events...)
}
