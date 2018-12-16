package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/conf"
	"server/db"
	"server/models"
	"time"
)

type API struct {
	dbSessionProvider db.SessionProvider
	config conf.Config
	Router *httprouter.Router
}

func NewAPI(dbSessionProvider db.SessionProvider, config conf.Config) *API {
	router := httprouter.New()

	api := &API{dbSessionProvider, config, router}

	api.Router.GET("/announcements", api.getAnnouncements())
	api.Router.POST("/announcements", api.createAnnouncement())

	api.Router.PanicHandler = api.panicHandler
	return api
}

func (a *API) panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("Panic recovered: %s -- %v\n", r.URL.Path, err)
	http.Error(w, "server error", http.StatusInternalServerError)
}

func (a *API) getAnnouncements() httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		results, err := a.dbSessionProvider.GetSession().GetAnnouncementsByDate(time.Now())
		if err != nil {
			log.Printf("HTTP handler getAnnouncements failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			log.Printf("HTTP handler getAnnouncements encoding results to JSON failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	}
}

func (a *API) createAnnouncement() httprouter.Handle {
	type RequestBody struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		announcement := models.NewAnnouncement(requestBody.Message)
		err = a.dbSessionProvider.GetSession().SaveAnnouncement(announcement)
		if err != nil {
			log.Printf("HTTP handler createAnnouncement saving announcement to db failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(announcement)
		if err != nil {
			log.Printf("HTTP handler saveAnnouncement marshaling response to JSON failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}