package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"rustdesk-api-server/app/controllers"
)

// Initialize routing service
func init() {
	// Cross domain solution
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// Allow access to all sources
		AllowAllOrigins: true,
		// optional parameters "GET", "POST", "PUT", "DELETE", "OPTIONS" (*for all)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Refers to the type of header allowed
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// List of exposed HTTP headers
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// If set, allows sharing of authentication credentials, such as cookies
		AllowCredentials: true,
	}))

	// Set routing information
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/api/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/ab", &controllers.AddressBookController{}, "post:Update")
	beego.Router("/api/ab/get", &controllers.AddressBookController{}, "post:List")
	beego.Router("/api/audit", &controllers.AuditController{}, "post:Audit")
	beego.Router("/api/logout", &controllers.LogoutController{}, "post:Logout")
	beego.Router("/api/currentUser", &controllers.UserController{}, "post:CurrentUser")
	beego.Router("/api/reg", &controllers.UserController{}, "get:Reg")
	beego.Router("/api/set-pwd", &controllers.UserController{}, "get:SetPwd")
	beego.Router("/api/heartbeat", &controllers.HBController{}, "post:Aliveness")

	// set error routing
	beego.ErrorController(&controllers.ErrorController{})
}
