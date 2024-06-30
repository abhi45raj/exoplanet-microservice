package utils

import "bitbucket.mynt.myntra.com/api/exoplanets/models"

func CalculateGravity(planet models.Exoplanet) float64 {
	switch planet.Type {
	case models.GasGiant:
		return 0.5 / (planet.Radius * planet.Radius)
	case models.Terrestrial:
		return planet.Mass / (planet.Radius * planet.Radius)
	}
	return 0
}
