package main

import (
	"log"
	"net/http"
	"server/api"
	"server/conf"
	"server/db"
)

func logRequestsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	config, err := conf.Init("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	sessionProvider, err := db.Init(config.Mongo.URL, config.Mongo.User, config.Mongo.Password)
	if err != nil {
		log.Fatal(err)
	}

	a := api.NewAPI(*sessionProvider, *config)
	log.Fatal(http.ListenAndServe(":8080", logRequestsMiddleware(a.Router)))
}
