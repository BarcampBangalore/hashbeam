package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"server/api"
	"server/conf"
	"server/db"
	"server/twitter"
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

	twitterContext, err := twitter.NewTwitterContext(dbContext, config.Twitter)
	if err != nil {
		log.Fatalf("couldn't create twitter context: %v", err)
	}

	app := api.NewAPI(dbContext, twitterContext, *config)

	httpHandler := applyMiddleware(
		app,
		func(h http.Handler) http.Handler { return handlers.LoggingHandler(os.Stdout, h) },
		handlers.CORS(handlers.AllowedOrigins([]string{"*"})),
	)

	log.Printf("Starting to listen to Twitter stream")
	err = twitterContext.StartStream()
	if err != nil {
		log.Fatalf("couldn't start twitter stream: %v", err)
	}

	log.Printf("Starting server on port " + config.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, httpHandler))
}

func applyMiddleware(h http.Handler, middleware ...func(h http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
