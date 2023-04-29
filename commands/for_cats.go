package commands

import (
	"unordered-events/domain"
	"unordered-events/domain/cat"
)

type CatCommandExecutors struct {
	CelebrateCatBirthday *CelebrateCatBirthdayCommandExecutor
}

func NewCatCommandExecutors(
	eventPublisher domain.EventPublisher,
	catRepository cat.Repository,
) *CatCommandExecutors {
	return &CatCommandExecutors{
		CelebrateCatBirthday: &CelebrateCatBirthdayCommandExecutor{
			eventPublisher: eventPublisher,
			catRepository:  catRepository,
		},
	}
}

type CelebrateCatBirthdayCommandExecutor struct {
	eventPublisher domain.EventPublisher
	catRepository  cat.Repository
}

func (ex *CelebrateCatBirthdayCommandExecutor) Execute(catID int) {
	aCat := ex.catRepository.Get(catID)

	events := aCat.CelebrateBirthday()

	ex.catRepository.Update(aCat)

	ex.eventPublisher.Publish(events...)
}
