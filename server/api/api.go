package api

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/conf"
	"server/db"
)

type API struct {
	database *db.DatabaseContext
	config   conf.Config
	Router   *httprouter.Router
}

func NewAPI(database *db.DatabaseContext, config conf.Config) *API {
	router := httprouter.New()

	api := &API{database, config, router}

	api.Router.GET("/announcements", api.getAnnouncements())
	api.Router.POST("/announcements", api.createAnnouncement())

	api.Router.PanicHandler = api.panicHandler
	return api
}

func (a *API) panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("Panic recovered: %s -- %v\n", r.URL.Path, err)
	http.Error(w, "server error", http.StatusInternalServerError)
}