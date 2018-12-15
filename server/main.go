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

	SessionProvider, err := db.Init(config.Mongo.URL, config.Mongo.User, config.Mongo.Password)
	if err != nil {
		log.Fatal(err)
	}

	Session := SessionProvider.GetSession()

	err = Session.AddAnnouncement(models.NewAnnouncement("this is a test announcement"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := Session.GetAnnouncements()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
