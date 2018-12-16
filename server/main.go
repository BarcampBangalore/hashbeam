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
	log.Fatal(http.ListenAndServe(":8080", logRequestsMiddleware(a.Router)))
}
