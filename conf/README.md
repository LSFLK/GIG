#app.conf
* config mongo, minio server url and credentials.

        mongo.database = gig
        mongo.path = localhost
        mongo.maxPool = 20
        
        minio.endpoint = 127.0.0.1:9001
        minio.accessKeyID = 
        minio.secretAccessKey =
        
#config.json
* Provide interface for configuration of utilities

        ServerApiUrl    string      GIG Backend Server URL      
        MapApiUrl       string      Google Location Search API URL
        MapAppKey       string      Google Location Search API App Key
        SearchApiUrl    string      Google Search API URL
        SearchAppKey    string      Google Search API App Key    
        Cx              string      Google Search API Secret Key
       
#routes
* configure API routes