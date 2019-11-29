# GIG
General Information Graph

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

## Get Started
### Deployment Requirements
* Golang
* Revel
* Mongo Server for Database Hosting
* Docker for Deploying Minio Server
* Minio Server for File Hosting
* Python for NER Recognition Server
* Google Custom Search API
### Run Server:
Configure conf/app.conf. Refer [How to Configure the Server](conf/README.md)

    revel run
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