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
	database *db.DBContext
	config   conf.Config
	ws       *melody.Melody
	router   *httprouter.Router
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func NewAPI(database *db.DBContext, config conf.Config) *API {
	router := httprouter.New()
	ws := melody.New()

	router.PanicHandler = panicHandler

	ws.HandleConnect(func(*melody.Session) {
		log.Printf("ws client connected -- current number of connections: %d\n", api.ws.Len()+1)
	})

	ws.HandleDisconnect(func(*melody.Session) {
		log.Printf("ws client disconnected -- current number of connections: %d\n", api.ws.Len())
	})

	api := &API{database, config, ws, router}

	api.router.GET("/announcements", api.getAnnouncements())
	api.router.POST("/announcements", api.createAnnouncement())
	api.router.GET("/ws/announcements", api.subscribeToAnnouncements())

	return api
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("Panic recovered: %s -- %v\n", r.URL.Path, err)
	http.Error(w, "server error", http.StatusInternalServerError)
}
