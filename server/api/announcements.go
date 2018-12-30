package api

import (
	"bytes"
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

		if r.URL.Query().Get("all") == "1" {
			results, err = a.database.GetAnnouncements()
		} else {
			results, err = a.database.GetAnnouncementsByDate(time.Now())
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

		announcement := models.Announcement{ID: 0, Time: time.Now(), Message: requestBody.Message}
		savedAnnouncement, err := a.database.SaveAnnouncement(announcement)
		if err != nil {
			log.Printf("HTTP handler createAnnouncement saving announcement to db failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		b := &bytes.Buffer{}
		err = json.NewEncoder(b).Encode(savedAnnouncement)
		if err != nil {
			log.Printf("HTTP handler saveAnnouncement marshaling response to JSON failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		resp := b.Bytes()

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("HTTP handler saveAnnouncement writing created announcement bytes buffer to http.ResponseWriter failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		err = a.ws.Broadcast(resp)
		if err != nil {
			log.Printf("HTTP handler saveAnnouncement broadcasting created announcement bytes buffer to all WS clients failed: %v", err)
			return
		}
	}
}

func (a *API) subscribeToAnnouncements() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := a.ws.HandleRequest(w, r)
		if err != nil {
			log.Printf("HTTP handler subscribeToAnnouncements upgrading connection to ws failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}
}
