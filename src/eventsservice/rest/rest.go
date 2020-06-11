package rest

import (
	"net/http"

	"github.com/a-soliman/projects/myEvents/src/lib/persistence"
	"github.com/gorilla/mux"
)

// ServeAPI runs and serves the API
func ServeAPI(endpoint string, tlsendpoint string, dbHandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := New(dbHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	httpErrChan := make(chan error)
	httpTLSErrChan := make(chan error)

	go func() { httpTLSErrChan <- http.ListenAndServeTLS(tlsendpoint, "../../cert.pem", "../../key.pem", r) }()
	go func() { httpErrChan <- http.ListenAndServe(endpoint, r) }()
	return httpErrChan, httpTLSErrChan
}
