package main

import (
	"encoding/json"
	"log"
	"memories/application"
	"memories/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	addRoutes(router)
	application.InitDB()
	application.InitLogger()

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second, ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/api/users", services.ListUsers).Methods("GET")
	router.HandleFunc("/api/users", services.AddUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", services.DeleteUser).Methods("DELETE")

	router.HandleFunc("/api/photos", services.PhotosHandler)
}
