// Package classification General Information Graph - API
//
// General Information Graph API Documentation
//
// Terms Of Service:
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: https, http
//     Host: api.gigdemo.opensource.lk:9000/api/
//     BasePath:
//     Version: 1.0.0
//     Contact: umayangag@opensource.lk
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//	   SecurityDefinitions:
//       Bearer:
//         type: apiKey
//         name: Authorization
//         in: header
//       ApiKey:
//         type: apiKey
//         name: ApiKey
//         in: header
//
// swagger:meta
package app

import (
	"GIG/app/databases"
	"GIG/app/repositories"
	"GIG/app/storages"
	"GIG/app/utilities/normalizers"
	"github.com/revel/config"
	"github.com/revel/revel"
	"log"
	"net/http"
)

var (
	// AppVersion revel app version (ldflags)
	Version string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

)

func init() {

	//var ValidateOrigin = func(c *revel.Controller, fc []revel.Filter) {
	//	if c.Request.Method == "OPTIONS" {
	//		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
	//		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
	//		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	//		c.Response.Out.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
	//		c.Response.SetStatus(http.StatusNoContent)
	//	} else {
	//		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	//		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
	//		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//		c.Response.Out.Header().Add("Content-Type", "application/json; charset=UTF-8")
	//		c.Response.Out.Header().Add("X-Frame-Options", "SAMORIGIN")
	//		c.Response.Out.Header().Add("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
	//
	//		fc[0](c, fc[1:]) // Execute the next filter stage.
	//	}
	//}

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		ValidateOrigin,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	Config, err := config.LoadContext("app.conf",revel.CodePaths)
	if err != nil || Config == nil {
		log.Fatalf("%+v",err)
	}

	revel.OnAppStart(databases.LoadDatabaseHandler)
	revel.OnAppStart(storages.LoadStorageHandler)
	revel.OnAppStart(normalizers.LoadNormalizers)
	revel.OnAppStart(repositories.LoadRepositoryHandler)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var ValidateOrigin = func(c *revel.Controller, fc []revel.Filter) {
	if c.Request.Method == "OPTIONS" {
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Response.Out.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Response.SetStatus(http.StatusNoContent)
		// 截取复杂请求下post变成options请求后台处理方法(针对跨域请求检测)
	} else {
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Response.Out.Header().Add("Content-Type", "application/json; charset=UTF-8")
		c.Response.Out.Header().Add("X-Frame-Options", "SAMORIGIN")
		c.Response.Out.Header().Add("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")

		fc[0](c, fc[1:]) // Execute the next filter stage.
	}
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
