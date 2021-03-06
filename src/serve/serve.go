package main

import (
	"log"
	"net/http"
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

func Methods() map[string]*webapp.Methods {
	return map[string]*webapp.Methods {
		"/api/account/": &webapp.Methods{
			Get: views.Account,
			Post: controllers.CreateAccount},
		"/api/login/": &webapp.Methods{
			Post: controllers.Login},
		"/api/logout": &webapp.Methods{
			Post: controllers.Logout}}//-- end return
}//-- end func Methods

/**
 * The actual implementation of the main function should not change.
 * However, you will need to change the implementation of the
 * functions it depends on:
 * - package "models" should implement func Models () []*model.Definition
 * - package "middleware" should implement
 * 		func Middleware () []webapp.Middleware
 * - the main package should implement func Methods () []*webapp.Methods
 */
func main() {
	config, err := webapp.LoadConfig(configFileName)
	if err != nil { log.Fatal(err.Error()) }
	//-- feel free to inject different dependencies below
	//-- check out the interfaces Server, Database, and Handler in webapp
	app, err := webapp.Init(config, &webapp.DefaultServer{},
		http.NewServeMux(), &database.Database{})
	if err != nil { log.Fatal(err) }
	err = app.RegisterModels(models.Models())
	if err != nil { log.Fatal(err) }
	app.AddMiddleware(middleware.Middleware()...)
	app.RegisterMethods(Methods())
	log.Fatal(app.ListenAndServe())
}//-- end func main


