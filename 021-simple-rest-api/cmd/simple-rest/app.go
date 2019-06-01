package main

import (
	"log"
	"net/http"

	ndb "github.com/haidlir/x-golang-course/021-simple-rest-api/db"
	nhandler "github.com/haidlir/x-golang-course/021-simple-rest-api/handler"

	"github.com/gorilla/mux"
)

const (
	// ListeningPort is the API listerner port
	ListeningPort = ":8080"
)

// Handler hold the function handler for API's endpoint.
type Handler interface {
	AddSiswa(w http.ResponseWriter, r *http.Request)
	GetAllSiswa(w http.ResponseWriter, r *http.Request)
	GetDetailSiswa(w http.ResponseWriter, r *http.Request)
	DeleteSiswa(w http.ResponseWriter, r *http.Request)
	UpdateSiswa(w http.ResponseWriter, r *http.Request)
}

// NewRouter returns router.
func NewRouter(handler Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/siswa", handler.GetAllSiswa).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", handler.GetDetailSiswa).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa", handler.AddSiswa).Methods(http.MethodPost)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", handler.UpdateSiswa).Methods(http.MethodPut)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", handler.DeleteSiswa).Methods(http.MethodDelete)
	return r
}

func main() {
	log.Println("Start service...")

	db := ndb.NewDummyDB()
	log.Println("Successfully Conneceted to DB")

	handler := nhandler.NewHandler(db)
	r := NewRouter(handler)

	// Run Web Server
	log.Printf("Starting http server at %v", ListeningPort)
	err := http.ListenAndServe(ListeningPort, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}
	log.Println("Stopping API Service...")
}
