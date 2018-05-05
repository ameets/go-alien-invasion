package main

import (
	"log"
	"math/rand"
	"time"
)

type World struct {
	Aliens map[int]Alien
	Cities map[string]City
}

func NewWorld() World {
	return World{
		Aliens: make(map[int]Alien), //todo maybe use map[int]int for alien -> moveCount
		Cities: make(map[string]City),
	}
}

func (w *World) DestroyCity(city string, a int) {
	if v, ok := w.Cities[city]; ok {
		for _, c := range w.Cities {
			c.RemoveConnection(city)
		}
		log.Println("%s has been destroyed by alien %d and alien %d", city, v.Alien.Name, a)
		delete(w.Cities, city)
		delete(w.Aliens, v.Alien.Name)

	}
}

func (w *World) CreateAliens(n int) {
	if n < 1 {
		return
	}
	// Not security-sensitive, use math.rand instead
	// of crypto.rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cities := w.getCities()
	for i := 1; i <= n; i++ {
		if len(cities) == 0 { //all cities destroyed
			break
		}
		idx := r.Intn(len(cities)) //[0,len(cities))
		c, ok := w.Cities[cities[idx]]
		if !ok {
			panic("unexpected error")
		}
		if c.HasAlien() {
			w.DestroyCity(c.Name, i)
			cities = deleteAtIdx(cities, idx)
		} else {
			c.Alien.Name = i
		}

	}
}

func (w *World) getCities() []string {
	cities := make([]string, len(w.Cities))
	i := 0
	for k, _ := range w.Cities {
		cities[i] = k
		i++
	}
	return cities
}

func deleteAtIdx(s []string, idx int) []string {
	// do nothing if out of bounds
	if idx < 0 || idx >= len(s) {
		return s
	}

	// if idx == len(s)-1 { //last element
	// 	return append(s[:idx]) //delete last element
	// }
	return append(s[:idx], s[idx+1:]...) //delete element at idx

}
