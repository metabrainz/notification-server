package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	log.Println("Starting server (127.0.0.1:8080)...")
	http.ListenAndServe("127.0.0.1:8080", makeRouter())
}

func makeRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", indexHandler).
		Methods(http.MethodGet)
	r.HandleFunc("/notifications", notificationsHandler).
		Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/notifications/marker", notificationsMarkerHandler).
		Methods(http.MethodPut)
	return r
}

/*
 * Handlers
 */

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// This is for testing purposes only.
	fmt.Fprint(w, "Hello, I'm functioning.")
}

// notificationsHandler handles requests for creation of new notifications
// (POST) and retrieval of existing ones (GET).
func notificationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		fmt.Fprintf(w, "POST")
	}
}

// notificationsMarkerHandler handles requests that modify "read" marker,
// which is a timestamp that splits read and unread notifications.
func notificationsMarkerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT (Marker)")
}

func main() {
	StartServer()
}
