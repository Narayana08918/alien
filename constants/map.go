package constants

const (
	// MaxOccupancy of the city
	MaxOccupancy = 2

	// MaxAlienMoves is number of moves after which the attacks end
	MaxAlienMoves = 10000
)

// When constructing the directional edges
var DirectionOpposites map[string]string

func init() {
	DirectionOpposites = make(map[string]string)

	DirectionOpposites["north"] = "south"
	DirectionOpposites["south"] = "north"
	DirectionOpposites["east"] = "west"
	DirectionOpposites["west"] = "east"
}
