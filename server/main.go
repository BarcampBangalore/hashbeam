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
	config, err := conf.Init("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	connString := config.MySQL.User + ":" + config.MySQL.Password + "@(" +
		config.MySQL.Host + ":" + config.MySQL.Port + ")/" +
		config.MySQL.Database + "?parseTime=true"

	dbCtx, err := db.Init(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCtx.DB.Close()

	err = dbCtx.SetupTables()
	if err != nil {
		log.Fatalf("couldn't ensure database schema (create table if not exists failed): %v", err)
	}

	a := api.NewAPI(dbCtx, *config)

	httpHandler := applyMiddleware(
		a.Router,
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