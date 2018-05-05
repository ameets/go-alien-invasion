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
}

// RemoveConnection takes in a city name and removes it
// from the connections of a City struct.
func (c *City) RemoveConnection(city string) {
	if _, ok := c.Connections[city]; ok {
		delete(c.Connections, city)
	}
}
