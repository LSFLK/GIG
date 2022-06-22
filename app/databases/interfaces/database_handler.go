package interfaces

type DatabaseHandlerInterface interface {
	GetServiceInstance() DatabaseServiceInterface
	DisconnectService()
}
