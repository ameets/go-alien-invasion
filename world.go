package main

import (
	"log"
	"math/rand"
	"time"
)

const minMoves = 10000

type World struct {
	Aliens map[int]int
	Cities map[string]City
}

func NewWorld() World {
	return World{
		Aliens: make(map[int]int),     // alien -> moveCount
		Cities: make(map[string]City), // city name -> city struct
	}
}

// DestroyCity takes in a city name and invading alien name.
// The city, it's existing alien, and all connections to it
// are removed from the world.
func (w *World) DestroyCity(name string, a int) {
	if city, ok := w.Cities[name]; ok {
		for _, c := range w.Cities {
			c.RemoveConnection(name)
		}
		log.Println("%s has been destroyed by alien %d and alien %d!", city.Name, city.Alien, a)
		delete(w.Cities, city.Name)
		delete(w.Aliens, city.Alien)
	}
}

func (w *World) Move() bool {
	if len(w.Cities) == 0 || len(w.Aliens) == 0 {
		return false
	}
	return true
}

// CreateAliens takes in the number of aliens to create.
// Aliens are randomly assigned to cities. If a city already
// has an alien, the two will duel and destroy the city.
// No new aliens are created if all a world's cities are destroyed.
func (w *World) CreateAliens(n int) {
	if n < 1 {
		return
	}
	// Not security-sensitive, use math.rand instead
	// of crypto.rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cities := w.getCities()
	for i := 0; i < n; i++ {
		// break if no cities remain in the world
		if len(cities) == 0 {
			break
		}
		// idx bounds are 0 (inclusive) to len (exclusive)
		// i.e. `[0,len(cities))`
		idx := r.Intn(len(cities))
		c, ok := w.Cities[cities[idx]]
		if !ok {
			panic("unexpected error")
		}
		if c.HasAlien() {
			// destroy city instead of creating alien
			w.DestroyCity(c.Name, i)
			cities = deleteAtIdx(cities, idx)
		} else {
			// create alien in city and add to the world
			c.Alien = i
			w.Aliens[i] = 0
		}
	}
}

// GetCities returns all the city names in the world.
func (w *World) getCities() []string {
	cities := make([]string, len(w.Cities))
	i := 0
	for k, _ := range w.Cities {
		cities[i] = k
		i++
	}
	return cities
}

// Deletes the element of `s` at index `idx` and
// returns the updated string slice.
func deleteAtIdx(s []string, idx int) []string {
	// do nothing if out of bounds
	if idx < 0 || idx >= len(s) {
		return s
	}
	// delete element at idx
	return append(s[:idx], s[idx+1:]...)
}
