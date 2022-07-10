package models

import "fmt"

// City has the Name of the city, Ways in and out of city and the list of aliens in the city
type City struct {
	Name   string
	In     map[string]string
	Out    map[string]string
	Aliens map[string]*Alien
}

func (c *City) String() string {
	if len(c.Out) == 0 {
		return ""
	}

	links := ""
	for direction, name := range c.Out {
		links += fmt.Sprintf(" %s=%s", direction, name)
	}

	return fmt.Sprintf("%s%s", c.Name, links)
}
