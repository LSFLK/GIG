package index_manager

import (
	"GIG/app/databases/interfaces"
	"log"
	"sync"
)

func CreateDBIndexes(manager interfaces.IndexManagerInterface) {
	var wg sync.WaitGroup
	log.Println("creating database indexes")
	wg.Add(3)
	go manager.CreateEntityIndexes(&wg)
	go manager.CreateNormalizedNameIndexes(&wg)
	go manager.CreateUserIndexes(&wg)
	wg.Wait()
	log.Println("indexes created successfully")
}
