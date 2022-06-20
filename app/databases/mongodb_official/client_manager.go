package mongodb_official

import "log"

func DisconnectClient() {
	err := service.client.Disconnect(Context)
	if err != nil {
		log.Println("error shutting down Mongo Client")
	}
}
