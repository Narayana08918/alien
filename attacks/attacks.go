package attacks

import (
	"github.com/Narayana08918/alien/constants"
	"github.com/Narayana08918/alien/planet"
)

type Attacks struct {
	planet *planet.Planet
	moves  map[string]int
}

func New(planet *planet.Planet) *Attacks {
	s := &Attacks{
		planet: planet,
		moves:  make(map[string]int),
	}

	for _, alien := range planet.Get().Aliens {
		s.moves[alien.Name] = 0
	}

	return s
}

func (a *Attacks) Run() error {
	for !a.done() {
		alienName, err := a.planet.MoveAliens()
		if err != nil {
			return err
		}

		_, ok := a.moves[alienName]
		if ok {
			a.moves[alienName]++

			// Once an alien has moved at least 'minAlienMoves' times, we can
			// avoid having to track/count his moves.
			if a.moves[alienName] >= constants.MaxAlienMoves {
				delete(a.moves, alienName)
			}
		}

		a.planet.ExecuteFights()
	}

	return nil
}

func (a *Attacks) done() bool {
	if len(a.planet.Get().Aliens) == 0 {
		return true
	}

	for _, totalMoves := range a.moves {
		if totalMoves < constants.MaxAlienMoves {
			return false
		}
	}

	return true
}
