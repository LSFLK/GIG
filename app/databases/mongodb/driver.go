package mongodb

var MaxPool int
var PATH    string
var DBNAME  string


func CheckAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = PATH
		err := service.New()
		if err != nil {
			panic(err)
		}
	}
}


