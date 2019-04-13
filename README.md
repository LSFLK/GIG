# GIG
General Information Graph

## Project Structure

The directory structure of the GIG:

    conf/                       Configuration directory
        app.conf                Main app configuration file
        routes                  Routes definition file

    app/                        App sources
        init.go                 Interceptor registration
        entities/               Model classes
        controllers/            App controllers
            api/                Inbound/Outbound API controllers
        views/                  Templates directory
        services/
            crawlers/           Data Crawler classes
            storage_handlers/   Storage Handler Classes
            decoders/           API request decoder classes


    messages/                   Message files

    public/                     Public static assets
        css/                    CSS files
        js/                     Javascript files
        images/                 Image files

    tests/                      Test suites


## Help

* [Controller Documentation](app/controllers/README.md)
* [API Documentation](app/controllers/api/README.md)
* [Crawlers Documentation](app/utility/crawlers/README.md)

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).