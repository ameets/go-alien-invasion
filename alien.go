package main

type Alien struct {
	Name      int
	MoveCount int
	Location  string // city name
}

func NewAlien(n int, l string) Alien {
	return Alien{
		Name:     n,
		Location: l,
	}
}
