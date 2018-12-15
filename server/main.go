package main

import (
	"fmt"
	"log"
	"server/conf"
	"server/db"
)

func main() {
	config, err := conf.Init("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	dbContextProvider, err := db.Init(config.Mongo.URL, config.Mongo.User, config.Mongo.Password)
	if err != nil {
		log.Fatal(err)
	}

	dbContext := dbContextProvider.GetContext()
	fmt.Println(dbContext.GetTweets())
}