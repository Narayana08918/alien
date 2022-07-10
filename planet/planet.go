package planet

import (
	"fmt"
	"log"
	"strings"

	"github.com/Narayana08918/alien/constants"
	"github.com/Narayana08918/alien/models"
	"github.com/Narayana08918/alien/pq"
)

type Planet struct {
	planet *models.Planet
}

func NewPlanet() *Planet {
	return &Planet{}
}

func (p *Planet) GetNumCities() int {
	return len(p.planet.Cities)
}

func (p *Planet) PlaceAliens(n int) {
	q := pq.New()

	// Place all aliens in the map based on the cities with most edges for better maneuverability
	for _, city := range p.planet.Cities {
		q.Push(city)
	}

	placedAliens := int(0)
	for placedAliens != n {
		city := q.Pop()

		for i := 0; i < constants.MaxOccupancy && placedAliens != n; i++ {
			alien := &models.Alien{
				Name: fmt.Sprintf("alien %d", placedAliens+1),
				City: city.Name,
			}

			city.Aliens[alien.Name] = alien
			p.planet.Aliens[alien.Name] = alien

			placedAliens++
		}
	}

	// Run once to check if there are cities to be destroyed
	p.Attack()
}

func (p *Planet) Attack() {
	for _, alien := range p.planet.Aliens {
		occupiedCity := alien.City
		city := p.planet.Cities[occupiedCity]

		if len(city.Aliens) == constants.MaxOccupancy {
			aliens := p.destroyCity(city)
			log.Printf("%s has been destroyed by %s!", city.Name, strings.Join(aliens, " and "))
		}
	}
}

func (p *Planet) destroyCity(city *models.City) []string {
	aliens := []string{}

	for alien := range city.Aliens {
		aliens = append(aliens, alien)
		delete(p.planet.Aliens, alien)
	}

	for _, cityName := range city.In {
		inCity := p.planet.Cities[cityName]

		for direction, edgeCityName := range inCity.Out {
			if edgeCityName == city.Name {
				delete(inCity.Out, direction)
				break
			}
		}

		for direction, edgeCityName := range inCity.In {
			if edgeCityName == city.Name {
				delete(inCity.In, constants.DirectionOpposites[direction])
				break
			}
		}
	}

	for _, cityName := range city.Out {
		outCity := p.planet.Cities[cityName]

		for direction, edgeCityName := range outCity.In {
			if edgeCityName == city.Name {
				delete(outCity.In, direction)
				break
			}
		}
	}

	delete(p.planet.Cities, city.Name)
	return aliens
}
