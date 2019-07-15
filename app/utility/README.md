#Documentation of Utility functions used
##Crawlers:
* [Crawler Documentation](crawlers/README.md)
* [PDF Crawler Documentation](crawlers/pdf_crawler/README.md)
* [Web Crawler Documentation](crawlers/web_crawler/README.md)
* [Wiki API Crawler Documentation](crawlers/wiki_api_crawler/README.md)
##Entity Handlers:

##Importers:
* [eTender Documentation](importers/etender/README.md)
##Parsers:
    1. ParsePdf -> return the string content of a given PDF file
##Request Handlers:
    1. GetRequest -> get the response string for a given url
    2. PostRequest -> Post to an url with data
##Commons:
    1. FileTypeCheck -> check if the file type of given source path matches given file type
    2. DownloadFile -> download a file given the source and destination
    3. ExtractDomain -> extract the main domain from a given source path
    4. ExtractFileName -> extract filename from a given source path
    5. FixUrl -> convert relative urls to absolute urls
    6. ObjectIdInSlice -> check if a given string exists in a given slice
    7. StringContainsAnyInSlice -> check if a given string is contained in any string in a given slice
    8. StringInSlice -> check if a given string exists in a given slice