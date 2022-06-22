package index_manager

import (
	"log"
	"sync"
)

type IndexManagerInterface interface {
	CreateEntityIndexes(wg *sync.WaitGroup)
	CreateNormalizedNameIndexes(wg *sync.WaitGroup)
	CreateUserIndexes(wg *sync.WaitGroup)
}

func CreateDBIndexes(manager IndexManagerInterface) {
	var wg sync.WaitGroup
	log.Println("creating database indexes")
	wg.Add(3)
	go manager.CreateEntityIndexes(&wg)
	go manager.CreateNormalizedNameIndexes(&wg)
	go manager.CreateUserIndexes(&wg)
	wg.Wait()
	log.Println("indexes created successfully")
}
