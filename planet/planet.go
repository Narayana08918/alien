package planet

import (
	"fmt"

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
}
