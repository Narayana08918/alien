package planet

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/Narayana08918/alien/models"
)

// Read the map file to build cities and with edges and directions
// Return the planet object once the map is drawn
func BuildMap(mapFile string) (*models.Planet, error) {
	file, err := os.Open(mapFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	planet := models.NewPlanet()

	// Create a scanner to read each map entries line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")

		if len(strs) == 0 {
			return nil, errors.New("invalid line in map definition")
		}

		cityName := strs[0]

		if len(strs) > 1 {
			for _, link := range strs[1:] {
				nodes := strings.Split(link, "=")

				if len(nodes) != 2 {
					return nil, errors.New("invalid line in map definition")
				}

				planet.AddNodeAndEdge(cityName, nodes[0], nodes[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return planet, nil
}
