package main

import (
	"log"
	"math/rand"
	"time"
)

const minMoves = 10000

type World struct {
	Aliens map[int]int
	Cities map[string]*City
}

func NewWorld() World {
	return World{
		Aliens: make(map[int]int),      // alien -> moveCount
		Cities: make(map[string]*City), // city name -> city struct
	}
}

// DestroyCity takes in a city name and invading alien name.
// The city, it's existing alien, and all connections to it
// are removed from the world.
func (w *World) destroyCity(name string, a int) {
	if city, ok := w.Cities[name]; ok {
		for _, c := range w.Cities {
			c.RemoveConnection(name)
		}
		log.Printf("%s has been destroyed by alien %d and alien %d!", city.Name, city.Alien, a)
		delete(w.Cities, city.Name)
		delete(w.Aliens, city.Alien)

		// if dueling alien is moving from another city
		// it is also destroyed
		if _, ok := w.Aliens[a]; ok {
			delete(w.Aliens, a)
		}
	}
}

// Returns true if each alien has moved at least
// minMoves times, false otherwise.
func (w *World) hasMinMoves() bool {
	for _, m := range w.Aliens {
		if m < minMoves {
			return false
		}
	}
	return true
}

// GameOver returns true if all cities are destroyed, all aliens
// are destroyed, or if the min number of moves per alien is reached.
func (w *World) GameOver() bool {
	if len(w.Cities) == 0 || len(w.Aliens) == 0 || w.hasMinMoves() {
		return true
	}
	return false
}

func (w *World) Move() {
	moved := make(map[int]bool)
	for _, city := range w.Cities {
		if !city.HasAlien() {
			continue
		}
		a := city.Alien
		if _, ok := moved[a]; ok {
			continue
		}

		// mark as moved and increment move count
		moved[a] = true
		w.Aliens[a]++

		nextMove := city.GetMove()
		if nextMove != "" {
			// city still has connections
			if nextCity, ok := w.Cities[nextMove]; ok {
				if nextCity.HasAlien() {
					w.destroyCity(nextCity.Name, a)
				} else {
					nextCity.SetAlien(a)
				}
				city.RemoveAlien()
			}
		}
	}
}

// CreateAliens takes in the number of aliens to create.
// Aliens are randomly assigned to cities. If a city already
// has an alien, the two will duel and destroy the city.
// No new aliens are created if all a world's cities are destroyed.
func (w *World) CreateAliens(n int) {
	if n < 1 {
		return
	}
	cities := w.getCities()
	for i := 0; i < n; i++ {
		// break if no cities remain in the world
		if len(cities) == 0 {
			break
		}
		// Not security-sensitive, use math.rand instead of crypto.rand
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// idx bounds are 0 (inclusive) to len (exclusive)
		// i.e. `[0,len(cities))`
		idx := r.Intn(len(cities))
		c, ok := w.Cities[cities[idx]]
		if !ok {
			panic("unexpected error")
		}
		if c.HasAlien() {
			// destroy city instead of creating alien
			w.destroyCity(c.Name, i)
			cities = deleteAtIdx(cities, idx)
		} else {
			// create alien in city and add to the world
			c.SetAlien(i)
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
