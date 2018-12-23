package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"server/api"
	"server/conf"
	"server/db"
)

func main() {
	config, err := conf.NewConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	dbContext, err := db.NewDBContext(config.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbContext.Close()

	err = dbContext.SetupTables()
	if err != nil {
		log.Fatalf("couldn't ensure database schema (create table if not exists failed): %v", err)
	}

	app := api.NewAPI(dbContext, *config)

	httpHandler := applyMiddleware(
		app,
		func(h http.Handler) http.Handler { return handlers.LoggingHandler(os.Stdout, h) },
		handlers.CORS(handlers.AllowedOrigins([]string{"*"})),
	)

	log.Printf("Starting server on port " + config.App.Port)
	log.Fatal(http.ListenAndServe(":" + config.App.Port, httpHandler))
}

func applyMiddleware(h http.Handler, middleware ...func(h http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}