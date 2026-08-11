package spaceage

import "strings"

type Planet string

func Age(seconds float64, planet Planet) float64 {
	const EarthDays = 365.25
	const SecondsInADay = 60.0 * 60.0 * 24.0
	rate := map[string]float64{
		"mercury": 0.2408467,
		"venus":   0.61519726,
		"earth":   1.0,
		"mars":    1.8808158,
		"jupiter": 11.862615,
		"saturn":  29.447498,
		"uranus":  84.016846,
		"neptune": 164.79132,
	}

	if rate[strings.ToLower(string(planet))] == 0.0 {
		return -1.0
	}

	return seconds / SecondsInADay / (rate[strings.ToLower(string(planet))] * EarthDays)
}
