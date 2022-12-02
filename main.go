package main

import (
	"dont/hexagonal/app"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//lambda.Start(app.Handler)
	app.StartServer()
}
