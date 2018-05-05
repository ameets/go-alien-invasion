package main

import "log"

type World struct {
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
