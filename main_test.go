package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	indexHandler(rec, req)

	if rec.Body.String() != "Hello, I'm functioning." ||
		rec.Code != http.StatusOK {
		t.Fatal("Server is not functioning. :(")
	}
}

func TestNotificationsHandlerGET(t *testing.T) {
	req, err := http.NewRequest("GET", "/notifications", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	notificationsHandler(rec, req)

	if rec.Body.String() != "GET" ||
		rec.Code != http.StatusOK {
		t.Fatal("Notifications handler is broken.")
	}
}

func TestNotificationsHandlerPOST(t *testing.T) {
	req, err := http.NewRequest("POST", "/notifications", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	notificationsHandler(rec, req)

	if rec.Body.String() != "POST" ||
		rec.Code != http.StatusOK {
		t.Fatal("Notifications handler is broken.")
	}
}
