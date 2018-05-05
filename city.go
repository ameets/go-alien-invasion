package main

type Direction int

const (
	N Direction = iota + 1
	S
	E
	W
)

func (d Direction) String() string {
	names := [...]string{
		"north",
		"south",
		"east",
		"west",
	}

	// direction is out of range
	if d < N || d > W {
		return ""
	}

	// name of direction constant
	return names[d]
}

type City struct {
	Name        string
	Connections map[string]string // city -> direction
	Alien       Alien
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
	if c.Alien.Name == -1 {
		return false
	}
	return true
}

// NewCity initializes a city with name `n` and
// connections `c`.
func NewCity(n string, c map[string]string) City {
	return City{
		Name:        n,
		Connections: c,
		Alien:       NewAlien(-1),
	}
}
