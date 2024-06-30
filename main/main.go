package main

import (
	"bitbucket.mynt.myntra.com/api/exoplanets/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", handlers.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", handlers.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handlers.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", handlers.EstimateFuel).Methods("GET")
	http.ListenAndServe(":8080", r)
}
