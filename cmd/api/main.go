package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/config"
	"github.com/schattenbrot/auth/internal/controllers"
	"github.com/schattenbrot/auth/internal/database"
	"github.com/schattenbrot/auth/internal/middlewares"
	"github.com/schattenbrot/auth/internal/routes"
)

func main() {
	app := config.Init()

	db := database.OpenDB(app)
	dbRepo := database.NewDBRepo(&app, db)

	controllers.NewRepo(&app, dbRepo)
	middlewares.NewRepo(&app, dbRepo)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      routes.Routes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Println("Starting server on port", app.Config.Port)

	err := serve.ListenAndServe()
	if err != nil {
		app.Logger.Fatalln(err)
	}
}
