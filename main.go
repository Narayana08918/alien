package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Narayana08918/alien/attacks"
	"github.com/Narayana08918/alien/planet"
)

func main() {
	var (
		mapFile      string
		finalMapFile string
		numAliens    int
	)

	flag.StringVar(&mapFile, "map", "", "file containing the map data")
	flag.StringVar(&finalMapFile, "out", "", "output file to write final map to")
	flag.IntVar(&numAliens, "n", 0, "number of aliens to use in the attacks")

	flag.Parse()

	if mapFile == "" {
		cmdErrMsg("Invalid map input file: map cannot be empty")
	}

	if finalMapFile == "" {
		cmdErrMsg("Invalid map output file: out cannot be empty")
	}

	if numAliens == 0 {
		cmdErrMsg("invalid number of aliens: n must be greater than zero")
	}

	planet, err := planet.BuildMap(mapFile)
	if err != nil {
		log.Fatalf("Failed to build planet from file: %v", err)
	}

	if numAliens > (planet.GetNumCities())*2 {
		log.Fatalf("Invalid number of aliens: cannot have more than 2x of cities. All cities will be destroyed")
		os.Exit(1)
	}

	planet.PlaceAliens(numAliens)
	attack := attacks.New(planet)
	err = attack.Run()
	if err != nil {
		log.Fatal("Failed to execute alien invasion attacks. Err:", err)
		os.Exit(1)
	}

	err = planet.WriteFinalMap(finalMapFile)
	if err != nil {
		log.Fatal("Failed to write the map to file. Err:", err)
		os.Exit(1)
	}
}

func cmdErrMsg(msg string) {
	fmt.Println(msg)
	fmt.Println("usage:")
	flag.PrintDefaults()
	os.Exit(1)
}
