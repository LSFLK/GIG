#server.conf
* config mongo, minio server url and credentials.

        mongo.database = gig
        mongo.path = localhost
        mongo.maxPool = 20
        
        minio.endpoint = 127.0.0.1:9001
        minio.accessKeyID = 
        minio.secretAccessKey =
        
* Provide interface for configuration of utilities

        normalizer.mapApiUrl       string      Google Location Search API URL
        normalizer.mapAppKey       string      Google Location Search API App Key
        normalizer.searchApiUrl    string      Google Search API URL
        normalizer.searchAppKey    string      Google Search API App Key    
        normalizer.cx              string      Google Search API Secret Key
        normalizer.tolerance       integer     Google Search API Secret Key
       
#routes
* configure API routes