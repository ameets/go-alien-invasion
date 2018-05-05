package main

type Alien struct {
	Name      int
	MoveCount int
}

func NewAlien(n int) Alien {
	return Alien{
		Name:      n,
		MoveCount: 0,
	}
}
