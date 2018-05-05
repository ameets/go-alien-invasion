package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"testing"
)

func TestDestroyCity(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := City{
		Name: "c1",
		Connections: map[string]string{
			"c2": "north",
			"c3": "south",
		},
	}

	c2 := City{
		Name: "c2",
		Connections: map[string]string{
			"c1": "south",
		},
	}

	c3 := City{
		Name: "c3",
		Connections: map[string]string{
			"c1": "north",
		},
	}

	c2After := City{
		Name:        "c2",
		Connections: map[string]string{},
	}

	c3After := City{
		Name:        "c3",
		Connections: map[string]string{},
	}

	w := NewWorld()
	w.Cities = map[string]City{
		"c1": c1,
		"c2": c2,
		"c3": c3,
	}

	wAfter := NewWorld()
	wAfter.Cities = map[string]City{
		"c2": c2After,
		"c3": c3After,
	}

	// Test cases
	cases := []struct {
		testName string
		city     string
		input    World
		want     World
	}{
		{
			testName: "city dne",
			city:     "dne",
			input:    w,
			want:     w,
		},
		{
			testName: "destroy city",
			city:     "c1",
			input:    w,
			want:     wAfter,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.input.DestroyCity(tc.city, 1)

			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("got: %+v, want %+v", tc.input, tc.want)
			}
		})
	}
}

func TestGetCities(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := City{
		Name: "c1",
		Connections: map[string]string{
			"c2": "north",
			"c3": "south",
		},
	}

	c2 := City{
		Name: "c2",
		Connections: map[string]string{
			"c1": "south",
		},
	}

	c3 := City{
		Name: "c3",
		Connections: map[string]string{
			"c1": "north",
		},
	}

	w := NewWorld()
	w.Cities = map[string]City{
		"c3": c3,
		"c2": c2,
		"c1": c1,
	}

	wEmpty := NewWorld()

	// Test cases
	cases := []struct {
		testName string
		world    World
		want     sort.StringSlice
	}{
		{
			testName: "empty world",
			world:    wEmpty,
			want:     []string{},
		},
		{
			testName: "populated world",
			world:    w,
			want:     []string{"c1", "c2", "c3"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			var got sort.StringSlice = tc.world.getCities()
			sort.Sort(got)
			sort.Sort(tc.want)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestCreateAliens(t *testing.T) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := City{
		Name: "c1",
		Connections: map[string]string{
			"c2": "north",
			"c3": "south",
		},
	}

	w := NewWorld()
	w.Cities = map[string]City{
		"c1": c1,
	}

	// Test cases
	cases := []struct {
		testName string
		n        int
		world    World
		want     int
	}{
		{
			testName: "no aliens",
			n:        0,
			world:    w,
			want:     0,
		},
		{
			testName: "invalid input",
			n:        -1,
			world:    w,
			want:     0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.world.CreateAliens(tc.n)

			got := len(tc.world.Aliens)
			if got != tc.want {
				t.Errorf("got: %+v, want %+v", got, tc.want)
			}
		})
	}
}
