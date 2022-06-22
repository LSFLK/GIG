package interfaces

type DatabaseServiceInterface interface {
	InitConnection(path string, dbName string, maxPool int)
}
