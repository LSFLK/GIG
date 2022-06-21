package index_manager

import (
	"log"
	"sync"
)

type IndexManager interface {
	CreateEntityIndexes(wg *sync.WaitGroup)
	CreateNormalizeNameIndexes(wg *sync.WaitGroup)
	CreateUserIndexes(wg *sync.WaitGroup)
}

func CreateDBIndexes(manager IndexManager) {
	var wg sync.WaitGroup
	log.Println("creating database indexes")
	wg.Add(3)
	go manager.CreateEntityIndexes(&wg)
	go manager.CreateNormalizeNameIndexes(&wg)
	go manager.CreateUserIndexes(&wg)
	wg.Wait()
	log.Println("indexes created successfully")
}
