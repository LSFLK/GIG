package entity_handlers

import "GIG/app/utility/config"

/**
Set the GIG server API url here for crawlers
 */
//var ApiUrl = "http://18.221.69.238:9000/api/"
var ApiUrl = config.GetConfig().ServerApiUrl
