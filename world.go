package main

import (
	"log"
	"math/rand"
	"sort"
)

type World struct {
	Aliens map[int]Alien
	Cities map[string]City
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

	cities := w.getCities()
	for i := 0; i < n; i++ {
		alien := Alien{
			Name:     i,
			Location: cities[rand.Intn(n)], //[0,n)
		}
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
