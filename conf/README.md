# server.conf
* config mongo, minio server url and credentials.

        mongo.database = gig
        mongo.path = mongodb://$USERNAME:$PASSWORD@$SERVER_IP:27017/gig
        mongo.maxPool = 20
        
        minio.endpoint = 127.0.0.1:9001
        minio.accessKeyID = 
        minio.secretAccessKey =
        
        file.cache = app/cache/
        
* Provide interface for configuration of utilities

        normalizer.mapApiUrl                string      Google Location Search API URL
        normalizer.mapAppKey                string      Google Location Search API App Key
        normalizer.searchApiUrl             string      Google Search API URL
        normalizer.searchAppKey             string      Google Search API App Key    
        normalizer.cx                       string      Google Search API Secret Key
        normalizer.minMatchPercentage       integer     Mininum string match percentage
       
# routes
* configure API routes