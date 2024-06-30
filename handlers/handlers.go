package handlers

import (
	"bitbucket.mynt.myntra.com/api/exoplanets/models"
	"bitbucket.mynt.myntra.com/api/exoplanets/utils"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var planet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil || models.ValidateExoplanet(planet) != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	planet.ID = uuid.New().String()
	models.Exoplanets[planet.ID] = planet
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(planet)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	var planets []models.Exoplanet
	for _, planet := range models.Exoplanets {
		planets = append(planets, planet)
	}
	json.NewEncoder(w).Encode(planets)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, exists := models.Exoplanets[params["id"]]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(planet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var planet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil || models.ValidateExoplanet(planet) != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	_, exists := models.Exoplanets[params["id"]]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	planet.ID = params["id"]
	models.Exoplanets[planet.ID] = planet
	json.NewEncoder(w).Encode(planet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, exists := models.Exoplanets[params["id"]]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	delete(models.Exoplanets, params["id"])
	w.WriteHeader(http.StatusNoContent)
}

func EstimateFuel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, exists := models.Exoplanets[params["id"]]
	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crewCapacity"))
	if err != nil || crewCapacity <= 0 {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	gravity := utils.CalculateGravity(planet)
	fuel := float64(planet.Distance) / (gravity * gravity) * float64(crewCapacity)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
}
