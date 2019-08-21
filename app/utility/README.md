# Documentation of Utility functions used
## Entity Handlers:
    1. AddEntitiesAsLinks -> Add list of related entities to a given entity
    2. AddEntityAsAttribute -> Add entity as an attribute to a given entity
    3. AddEntityAsLink -> Add entity as an link to a given entity
    4. CreateEntities -> Create a list of new entities and save to GIG
    5. CreateEntity -> Create a new entity and save to GIG
## Normalizers
    1. Normalize -> Normalize a given string to a Entity or Location
    1. NormalizeLocation -> Normalize a given string to a Entity
    1. NormalizeName -> Normalize a given string to a Location
## Request Handlers:
    1. GetRequest -> get the response string for a given url
    2. PostRequest -> Post to an url with data
## Commons:
    1. FileTypeCheck -> check if the file type of given source path matches given file type
    2. DownloadFile -> download a file given the source and destination
    3. ExtractDomain -> extract the main domain from a given source path
    4. ExtractFileName -> extract filename from a given source path
    5. FixUrl -> convert relative urls to absolute urls
    6. ObjectIdInSlice -> check if a given string exists in a given slice
    7. StringContainsAnyInSlice -> check if a given string is contained in any string in a given slice
    8. StringInSlice -> check if a given string exists in a given slice
    9. StringMatchPercentage -> check the similarity percentage of two given strings
    10. Maximum -> return maximum of a positive number slice
    11. Minimum -> return minimum of a positive number slice