package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/models"
	"time"
)

func (a *API) getAnnouncements() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var results []models.Announcement
		var err error

		if ps.ByName("all") == "1" {
			results, err = a.dbSessionProvider.GetSession().GetAnnouncements()
		} else {
			results, err = a.dbSessionProvider.GetSession().GetAnnouncementsByDate(time.Now())
		}

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

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
