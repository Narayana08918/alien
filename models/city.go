package models

// City has the Name of the city, Ways in and out of city and the list of aliens in the city
type City struct {
	Name   string
	In     map[string]string
	Out    map[string]string
	Aliens map[string]*Alien
}
