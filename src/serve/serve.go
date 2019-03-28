package main

import (
	"log"
	"net.http"
	// framework dependencies
	"github.com/ollykel/webapp"
	database "github.com/ollykel/webapp/databases/mysql"
	// local imports
	"models"
	"middleware"
	"views"
	"controllers"
)

const (
	configFileName = "config/config.json"
)

func main() {
	config, err := webapp.LoadConfig(configFileName)
	if err != nil { log.Fatal(err.Error()) }
	app, err := webapp.Init(config, &webapp.DefaultServer{},
		&database.Database{}, http.NewServeMux())
	if err != nil { log.Fatal(err) }
	app.RegisterModels(models.Models())
	app.AddMiddleware(middleware.Middleware())
	//-- TODO: register views and controllers
	log.Fatal(app.ListenAndServe())
}//-- end func main


