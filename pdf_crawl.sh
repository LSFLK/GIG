#!/bin/bash
# go run scripts/crawlers/pdf_crawler/pdf_web_crawler.go "http://www.buildings.gov.lk/index.php?option=com_content&view=category&layout=blog&id=47&Itemid=128&lang=en"
# go run scripts/crawlers/pdf_crawler/pdf_web_crawler.go "http://www.airforce.lk/main_tender.php"
# go run scripts/crawlers/pdf_crawler/pdf_web_crawler.go "https://www.parliament.lk/gazettes"
go run scripts/importers/pdf_importer.go "/home/umayanga/Downloads/gazette1.pdf"
