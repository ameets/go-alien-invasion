package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

type World struct {
	Aliens map[int]Alien
	Cities map[string]City
}

func NewWorld() World {
	return World{
		Aliens: make(map[int]Alien),
		Cities: make(map[string]City),
	}
}

func (w *World) DestroyCity(city string) {
	if _, ok := w.Cities[city]; ok {
		for _, c := range w.Cities {
			c.RemoveConnection(city)
		}
		delete(w.Cities, city)
		log.Printf("%s has been destroyed", city)
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
	for i := 0; i < n; i++ {
		idx := r.Intn(len(cities)) //[0,len(cities))
		alien := NewAlien(i, cities[idx])
		w.Aliens[i] = alien
	}
	// TODO: aliens may be randomly placed in the same city
	// and need to fight
}

func (w *World) getCities() sort.StringSlice {
	var cities sort.StringSlice = make([]string, len(w.Cities))
	i := 0
	for k, _ := range w.Cities {
		cities[i] = k
		i++
	}
	sort.Sort(cities) //ascending
	return cities
}
