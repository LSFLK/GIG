# GIG
General Information Graph (GIG) is a large scale information storage, querying and processing system for public information. With GIG, we aim to improve the quality of information which is publicly available on a multitude of social, political, and economic areas, and make it easier and more efficient to access so it can be put to gainful use.

## Project Structure

The directory structure of GIG:

    conf/                       Configuration directory
        app.conf                GIG server configuration file
        routes                  Routes definition file

    app/                        App sources
        init.go                 Interceptor registration
        cache/                  Cache Directory
        controllers/            App controllers
            api/                Inbound/Outbound API controllers
        data/                   Data Files for importing             
        models/                 Model classes
        repositories            Model Repositories
        storages/               Storage Handlers
        utilities/
            config/             Configuration Handler Class
            crawlers/           Data Crawler Classes
            entity_handlers/    Entity Management Classes
            importers/          Data Importer Classes
            normalizers/        Normalizer Classes
            parsers/            Source Parser Classes
            request_handlers/   Request Handler Classes
        views/                  Templates directory            

    messages/                   Message files

    public/                     Public static assets
        css/                    CSS files
        js/                     Javascript files
        images/                 Image files

    tests/                      Test suites
    
GIG Eco-System:

![GIG High Level Architecture](docs/images/gig_dataflow_diagram.png)

## Get Started

### Deployment Requirements
* Golang
* Revel
* Mongo Server for Database Hosting
* Docker for Deploying Minio Server
* Minio Server for File Hosting
* Python for NER Recognition Server
* Google Custom Search API

### Server Setup using Kubernetes (Optional)

To setup the GIG runtime environment and dependency servers using Kubernetes, refer to [Server Setup using Kubernetes](deployment/README.md).

### First time run:

Create cache directory:

    mkdir app/cache
    
    
Configure mongo.path at conf/app.conf using the mongodb and minio IPs. Refer [How to Configure the Server](conf/README.md)

    [dev]
    ...
    mongo.path = mongodb://developer:password@localhost:27017/gig
    ...
    minio.endpoint = localhost:9001
    
### Run Server:

### `revel run`
    
### Build Command:

    revel build -m prod -t build
    ./build/run.sh

## Help
* [API Documentation](https://app.swaggerhub.com/apis-docs/LSFGIG/GIG_API/1.0.0)
* [Utility Documentation](commons/README.md)
* [Crawlers Documentation](scripts/crawlers/README.md)

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [Revel API documentation](https://godoc.org/github.com/revel/revel).