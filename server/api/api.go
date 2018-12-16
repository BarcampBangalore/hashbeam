package api

import (
	"fmt"
	"net/http"
	"server/conf"
	"server/db"
)

type API struct {
	dbSessionProvider db.SessionProvider
	config conf.Config
	Mux *http.ServeMux
}

func NewAPI(dbSessionProvider db.SessionProvider, config conf.Config) *API {
	mux := http.NewServeMux()

	a := &API{dbSessionProvider, config, mux}

	a.Mux.HandleFunc("/", a.root())
	a.Mux.HandleFunc("/announcements", a.getAnnouncements())

	return a
}

func (a *API) root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "root")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (a *API) getAnnouncements() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "getAnnouncements")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
