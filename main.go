package main

import (
	"log"

	"github.com/erdemkosk/go-config-service/api/routes"
	_ "github.com/erdemkosk/go-config-service/docs"
)

// @title Config Service
// @version 1.0.0
// @description Swagger structure prepared for config service

// @contact.name Mustafa Erdem Köşk
// @contact.email erdemkosk@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//@securityDefinitions.apikey Authorization
//@in header
//@name Authorization

// @BasePath /
func main() {

	app := routes.GenerateFiberApp()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
