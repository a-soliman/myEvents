package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-soliman/projects/myEvents/src/lib/persistence"
	"github.com/gorilla/mux"
)

// EventServiceHandler service handler
type EventServiceHandler struct {
	dbhandler persistence.DatabaseHandler
}

// New returns a pointer
func New(databasehandler persistence.DatabaseHandler) *EventServiceHandler {
	return &EventServiceHandler{
		dbhandler: databasehandler,
	}
}

// FindEventHandler handles find event
func (eh *EventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search bu id via /id/4 or search by name via /name/someName}`)
		return
	}
	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search bu id via /id/4 or search by name via /name/someName}`)
		return
	}
	var (
		event persistence.Event
		err   error
	)
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

// AllEventHandler handler
func (eh *EventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	var (
		events []persistence.Event
		err    error
	)
	events, err = eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occurred while trying to find all available events %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occurred while trying to encode events to json. err = %s}", err)
	}
}

// NewEventHandler creates a new event
func (eh *EventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occurred while trying to decode event data. err = %s}", err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{error: Error occurred while persisting event id = %d. err = %s}", id, err)
		return
	}
	fmt.Fprint(w, id)
}
