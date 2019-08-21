# GIG
General Information Graph

## Project Structure

The directory structure of the GIG:

    conf/                       Configuration directory
        app.conf                Main app configuration file
        config.json             Configuration for utilities
        routes                  Routes definition file

    app/                        App sources
        init.go                 Interceptor registration
        cache/                  Cache Directory
        controllers/            App controllers
            api/                Inbound/Outbound API controllers
        data/                   Data Files for importing             
        models/                 Model classes
        repository/             Model Repositories
        routes/                 Generated Routes
        storage/                Storage Handlers
        tmp/                    Main App Directory
        utility/
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

## Deployment Requirements
* Mongo Server for Database Hosting
* Docker for Deploying Minio Server
* Minio Server for File Hosting
* Python for NER Recognition Server

## Help

* [API Documentation](https://app.swaggerhub.com/apis-docs/LSFGIG/GIG_API/1.0.0)
* [Utility Documentation](app/utility/README.md)
* [Crawlers Documentation](app/scripts/crawlers/README.md)

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [Revel API documentation](https://godoc.org/github.com/revel/revel).