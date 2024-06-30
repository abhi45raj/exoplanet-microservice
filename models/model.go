package models

import (
	"errors"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`       // light years
	Radius      float64       `json:"radius"`         // Earth-radius unit
	Mass        float64       `json:"mass,omitempty"` // Earth-mass unit (only for Terrestrial)
	Type        ExoplanetType `json:"type"`
}

var Exoplanets = make(map[string]Exoplanet)

func ValidateExoplanet(planet Exoplanet) error {
	if planet.Distance < 10 || planet.Distance > 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if planet.Radius < 0.1 || planet.Radius > 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if planet.Type == Terrestrial && (planet.Mass < 0.1 || planet.Mass > 10) {
		return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial planets")
	}
	if planet.Name == "" || planet.Description == "" {
		return errors.New("name and description cannot be empty")
	}
	return nil
}
