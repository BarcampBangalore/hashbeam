package main

import (
	"fmt"
	"log"
	"server/conf"
	"server/db"
	"server/models"
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

	dbContext := dbContextProvider.GetSession()

	err = dbContext.AddAnnouncement(models.NewAnnouncement("this is a test announcement"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := dbContext.GetAnnouncements()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
