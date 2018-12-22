package api

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
	"server/conf"
	"server/db"
)

type API struct {
	database *db.DatabaseContext
	config   conf.Config
	ws       *melody.Melody
	Router   *httprouter.Router
}

func NewAPI(database *db.DatabaseContext, config conf.Config) *API {
	router := httprouter.New()
	ws := melody.New()

	api := &API{database, config, ws, router}

	api.Router.GET("/announcements", api.getAnnouncements())
	api.Router.POST("/announcements", api.createAnnouncement())
	api.Router.GET("/ws/announcements", api.subscribeToAnnouncements())

	api.ws.HandleConnect(func(*melody.Session) {
		log.Printf("ws client connected -- current number of connections: %d\n", api.ws.Len()+1)
	})

	api.ws.HandleDisconnect(func(*melody.Session) {
		log.Printf("ws client disconnected -- current number of connections: %d\n", api.ws.Len())
	})

	api.Router.PanicHandler = api.panicHandler
	return api
}

func (a *API) panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("Panic recovered: %s -- %v\n", r.URL.Path, err)
	http.Error(w, "server error", http.StatusInternalServerError)
}
