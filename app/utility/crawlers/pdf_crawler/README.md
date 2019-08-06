1. Given a website Url crawl for all the links and filters pdf urls. 
2. Download and Parse the pdf files and extract the text content.
3. Use Stanford NER library to identify Named Entities in extracted text
4. Save to GIG API

## How to Run:
    1. set category var according the source category. eg. (Tenders, Gazettes, etc.)
    2. go run pdf_crawler.go "https://site.lk"