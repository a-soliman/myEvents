package rest

import (
	"net/http"

	"github.com/a-soliman/projects/myEvents/src/lib/persistence"
	"github.com/gorilla/mux"
)

// ServeAPI runs and serves the API
func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := New(dbHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(":8181", r)
}
