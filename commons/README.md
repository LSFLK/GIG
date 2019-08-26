#Generic utility functions
## Request Handlers:
    1. GetRequest -> get the response string for a given url
    2. PostRequest -> Post to an url with data
## Commons:
    1. FileTypeCheck -> check if the file type of given source path matches given file type
    2. DownloadFile -> download a file given the source and destination
    3. EnsureDirectory -> make directory if not exist 
    4. ExtractDomain -> extract the main domain from a given source path
    5. ExtractFileName -> extract filename from a given source path
    6. FixUrl -> convert relative urls to absolute urls
    7. ObjectIdInSlice -> check if a given string exists in a given slice
    8. StringContainsAnyInSlice -> check if a given string is contained in any string in a given slice
    9. StringInSlice -> check if a given string exists in a given slice
    10. StringMatchPercentage -> check the similarity percentage of two given strings
    11. Maximum -> return maximum of a positive number slice
    12. Minimum -> return minimum of a positive number slice