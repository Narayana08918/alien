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

func (p *Planet) Get() *models.Planet {
	return p.planet
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

func (p *Planet) MoveAliens() (string, error) {
	for _, alien := range p.planet.Aliens {
		occupiedCity := alien.City
		city := p.planet.Cities[occupiedCity]

		for _, name := range city.Out {
			city := p.planet.Cities[name]

			if len(city.Aliens) < constants.MaxOccupancy {
				delete(city.Aliens, alien.Name)

				alien.City = city.Name
				city.Aliens[alien.Name] = alien

				return alien.Name, nil
			}
		}
	}

	return "", constants.ErrCannotMoveAliens
}

func (p *Planet) ExecuteFights() {
	for _, alien := range p.planet.Aliens {
		occupiedCity := alien.City
		city := p.planet.Cities[occupiedCity]

		// If maximum occupancy has been reached for a city, the occupying
		// aliens will fight and destroy the city. As a result, the following
		// will happen:
		//
		// 1. Both aliens will be removed from the map's known collection of
		// aliens.
		// 2. The city will be removed from the map and so are any links that
		// lead into or out of it.
		if len(city.Aliens) == constants.MaxOccupancy {
			destroyedAliens := p.destroyCity(city)
			log.Printf("%s has been destroyed by %s!", city.Name, strings.Join(destroyedAliens, " and "))
		}
	}
}
