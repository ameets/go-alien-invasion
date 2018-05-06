package main

import "math/rand"

type City struct {
	Name        string
	Connections map[string]string // city -> direction
	Alien       int
}

// RemoveConnection takes in a city name and removes it
// from the connections of a City struct.
func (c *City) RemoveConnection(city string) {
	if _, ok := c.Connections[city]; ok {
		delete(c.Connections, city)
	}
}

// HasAlien returns false if Alien.Name is -1,
// true otherwise.
func (c *City) HasAlien() bool {
	if c.Alien == -1 {
		return false
	}
	return true
}

// Returns a random move based on a city's connections,
// or empty string if the city has no connections.
// Expects that rand has been seeded at program level.
func (c *City) GetMove() string {
	if len(c.Connections) == 0 {
		return ""
	}

	moves := make([]string, len(c.Connections))
	i := 0
	for k, _ := range c.Connections {
		moves[i] = k
		i++
	}

	idx := rand.Intn(len(c.Connections))
	return moves[idx]
}

// Sets the city's alien to n.
func (c *City) SetAlien(n int) {
	c.Alien = n
}

// Sets the city's alien to -1.
func (c *City) RemoveAlien() {
	c.Alien = -1
}

// NewCity initializes a city with name `n` and
// connections `c`.
func NewCity(n string, c map[string]string) *City {
	return &City{
		Name:        n,
		Connections: c,
		Alien:       -1,
	}
}
