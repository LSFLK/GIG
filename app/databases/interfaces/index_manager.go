package interfaces

import "sync"

type IndexManagerInterface interface {
	CreateEntityIndexes(wg *sync.WaitGroup)
	CreateNormalizedNameIndexes(wg *sync.WaitGroup)
	CreateUserIndexes(wg *sync.WaitGroup)
}
