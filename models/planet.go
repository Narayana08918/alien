package models

import (
	"strings"

	"github.com/Narayana08918/alien/constants"
)

// Planet is the main model with both the map and the aliens
type Planet struct {
	Map
	Aliens map[string]*Alien
}

func NewPlanet() *Planet {
	return &Planet{
		Map: Map{
			Cities: make(map[string]*City),
		},
		Aliens: make(map[string]*Alien),
	}
}

func (p *Planet) AddNodeAndEdge(cityName, direction, edgeCityName string) {
	// Add the origin city to the map of cities
	if _, ok := p.Cities[cityName]; !ok {
		p.Cities[cityName] = &City{
			Name:   cityName,
			In:     make(map[string]string),
			Out:    make(map[string]string),
			Aliens: make(map[string]*Alien, constants.MaxOccupancy),
		}
	}

	// Add the edge city to the map of cities
	if _, ok := p.Cities[edgeCityName]; !ok {
		p.Cities[edgeCityName] = &City{
			Name:   edgeCityName,
			In:     make(map[string]string),
			Out:    make(map[string]string),
			Aliens: make(map[string]*Alien, constants.MaxOccupancy),
		}
	}

	// Add outbound and inbound edges to the cities
	p.Cities[cityName].Out[strings.ToLower(direction)] = edgeCityName
	p.Cities[edgeCityName].In[constants.DirectionOpposites[strings.ToLower(direction)]] = cityName
}
