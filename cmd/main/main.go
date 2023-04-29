package main

import (
	"fmt"
	"time"

	"unordered-events/commands"
	"unordered-events/domain/cat"
	"unordered-events/domain/catowner"
	"unordered-events/input/kafka"
	"unordered-events/postgres"
)

func main() {
	catRepository := postgres.CatRepository{}

	felixID := catRepository.Add(&cat.Cat{Name: "Felix", Age: 2})
	rockyID := catRepository.Add(&cat.Cat{Name: "Rocky", Age: 3})

	catOwnerRepository := postgres.CatOwnerRepository{}

	_ = catOwnerRepository.Add(&catowner.CatOwner{Name: "Bob", CatID: felixID, CatAge: 2})
	_ = catOwnerRepository.Add(&catowner.CatOwner{Name: "Mike", CatID: felixID, CatAge: 2})
	_ = catOwnerRepository.Add(&catowner.CatOwner{Name: "Bill", CatID: felixID, CatAge: 2})
	_ = catOwnerRepository.Add(&catowner.CatOwner{Name: "Jack", CatID: rockyID, CatAge: 3})
	_ = catOwnerRepository.Add(&catowner.CatOwner{Name: "Matt", CatID: rockyID, CatAge: 3})

	eventPublisher := kafka.NewEventPublisher()

	catCommandExecutors := commands.NewCatCommandExecutors(eventPublisher, &catRepository)
	catOwnerCommandExecutors := commands.NewCatOwnerCommandExecutors(
		&catRepository, &catOwnerRepository, eventPublisher,
	)

	eventHandler := kafka.NewEventHandler(eventPublisher.ActivateEventBusAsync(), catOwnerCommandExecutors)

	go eventHandler.ListenAndHandle()

	catCommandExecutors.CelebrateCatBirthday.Execute(felixID)
	catCommandExecutors.CelebrateCatBirthday.Execute(rockyID)
	catCommandExecutors.CelebrateCatBirthday.Execute(felixID)
	catCommandExecutors.CelebrateCatBirthday.Execute(rockyID)
	catCommandExecutors.CelebrateCatBirthday.Execute(rockyID)
	catCommandExecutors.CelebrateCatBirthday.Execute(felixID)

	time.Sleep(1 * time.Second)

	fmt.Println("*** Cats after updates")
	for _, aCat := range catRepository.Entities() {
		fmt.Println(aCat)
	}

	fmt.Println("*** Cat owners after updates")
	for _, aCatOwner := range catOwnerRepository.Entities() {
		fmt.Println(aCatOwner)
	}
}
