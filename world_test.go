package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"testing"
)

func TestDestroyCity(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
	})
	c1.Alien = 1

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	c3 := NewCity("c3", map[string]string{
		"c1": "north",
	})

	c2After := NewCity("c2", map[string]string{})

	c3After := NewCity("c3", map[string]string{})

	w := NewWorld()
	w.Cities = map[string]*City{
		"c1": c1,
		"c2": c2,
		"c3": c3,
	}
	w.Aliens = map[int]int{
		c1.Alien: 0,
	}

	wAfter := NewWorld()
	wAfter.Cities = map[string]*City{
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
			tc.input.destroyCity(tc.city, 2)

			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("got: %+v, want %+v", tc.input, tc.want)
			}
		})
	}
}

func TestGetCities(t *testing.T) {
	//log.SetFlags(0)
	//log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
	})

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	c3 := NewCity("c3", map[string]string{
		"c1": "north",
	})

	w := NewWorld()
	w.Cities = map[string]*City{
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
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
	})

	w := NewWorld()
	w.Cities = map[string]*City{
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
		{
			testName: "create 1 alien",
			n:        1,
			world:    w,
			want:     1,
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

func TestDeleteAtIdx(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Test cases
	cases := []struct {
		testName string
		input    []string
		index    int
		want     []string
	}{
		{
			testName: "idx out of lower bound",
			input:    []string{"c3", "c4", "c1", "c5"},
			index:    -1,
			want:     []string{"c3", "c4", "c1", "c5"},
		},
		{
			testName: "idx out of uppder bound",
			input:    []string{"c3", "c4", "c1", "c5"},
			index:    4,
			want:     []string{"c3", "c4", "c1", "c5"},
		},
		{
			testName: "idx is last element",
			input:    []string{"c3", "c4", "c1", "c5"},
			index:    3,
			want:     []string{"c3", "c4", "c1"},
		},
		{
			testName: "idx is first element",
			input:    []string{"c3", "c4", "c1", "c5"},
			index:    0,
			want:     []string{"c4", "c1", "c5"},
		},
		{
			testName: "idx is in bounds",
			input:    []string{"c3", "c4", "c1", "c5"},
			index:    1,
			want:     []string{"c3", "c1", "c5"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			got := deleteAtIdx(tc.input, tc.index)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestHasMinMoves(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial Setup
	wF := NewWorld()
	wF.Aliens = map[int]int{
		1: minMoves,
		3: minMoves,
		5: 0,
	}

	wT := NewWorld()
	wT.Aliens = map[int]int{
		1: minMoves + 1,
		3: minMoves,
	}
	// Test cases
	cases := []struct {
		testName string
		world    World
		want     bool
	}{
		{
			testName: "expect false",
			world:    wF,
			want:     false,
		},
		{
			testName: "expect true",
			world:    wT,
			want:     true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			got := tc.world.hasMinMoves()

			if got != tc.want {
				t.Errorf("got: %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestGameOver(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial Setup
	// world without cities
	w1 := NewWorld()

	// world without aliens
	w2 := NewWorld()
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
	})
	w2.Cities = map[string]*City{
		"c1": c1,
	}

	// world with min moves
	w3 := NewWorld()
	c1.Alien = 0
	w3.Cities = map[string]*City{
		"c1": c1,
	}
	w3.Aliens = map[int]int{
		0: minMoves + 1,
	}

	// world with moves
	w4 := NewWorld()
	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})
	c2.Alien = 1
	w4.Cities = map[string]*City{
		"c1": c1,
		"c2": c2,
	}
	w4.Aliens = map[int]int{
		0: 0,
		1: 0,
	}
	// Test cases
	cases := []struct {
		testName string
		world    World
		want     bool
	}{
		{
			testName: "no cities",
			world:    w1,
			want:     true,
		},
		{
			testName: "no aliens",
			world:    w2,
			want:     true,
		},
		{
			testName: "hasMinMoves",
			world:    w3,
			want:     true,
		},
		{
			testName: "can move",
			world:    w4,
			want:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			got := tc.world.GameOver()

			if got != tc.want {
				t.Errorf("got: %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestMove(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial setup - cities
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c5": "east",
	})

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})
	c2.SetAlien(0)

	c3 := NewCity("c3", map[string]string{})
	c3.SetAlien(0)

	c4 := NewCity("c4", map[string]string{
		"c5": "north",
	})
	c4.SetAlien(0)

	c5 := NewCity("c5", map[string]string{
		"c4": "south",
	})
	c5.SetAlien(1)

	// Initial setup - world
	w1 := NewWorld()
	w1.Cities = map[string]*City{
		"c1": c1,
	}
	w1.Aliens = map[int]int{}

	w2 := NewWorld()
	w2.Cities = map[string]*City{
		"c1": c1,
		"c2": c2,
	}
	w2.Aliens = map[int]int{
		0: 0,
	}

	w3 := NewWorld()
	w3.Cities = map[string]*City{
		"c3": c3,
	}
	w3.Aliens = map[int]int{
		0: 0,
	}

	w4 := NewWorld()
	w4.Cities = map[string]*City{
		"c4": c4,
		"c5": c5,
	}
	w4.Aliens = map[int]int{
		0: 0,
		1: 0,
	}

	// Expected after move
	c1After := NewCity("c1", map[string]string{
		"c2": "north",
		"c5": "east",
	})
	c1After.SetAlien(0)

	c2After := NewCity("c2", map[string]string{
		"c1": "south",
	})

	c4After := NewCity("c4", map[string]string{})
	c5After := NewCity("c5", map[string]string{})

	w2After := NewWorld()
	w2After.Cities = map[string]*City{
		"c1": c1After,
		"c2": c2After,
	}
	w2After.Aliens = map[int]int{
		0: 1,
	}

	w3After := NewWorld()
	w3After.Cities = map[string]*City{
		"c3": c3,
	}
	w3After.Aliens = map[int]int{
		0: 1,
	}

	w4After := NewWorld()
	w4After.Cities = map[string]*City{
		"c4": c4After,
	}
	w4After.Aliens = map[int]int{}

	w4Alt := w4After
	w4Alt.Cities = map[string]*City{
		"c5": c5After,
	}

	// Test cases
	cases := []struct {
		testName      string
		world         World
		want          World
		wantAlternate World
	}{
		{
			testName:      "no aliens to move",
			world:         w1,
			want:          w1,
			wantAlternate: NewWorld(),
		},
		{
			testName:      "move alien to connecting city",
			world:         w2,
			want:          w2After,
			wantAlternate: NewWorld(),
		},
		{
			testName:      "trapped alien",
			world:         w3,
			want:          w3After,
			wantAlternate: NewWorld(),
		},
		{
			testName:      "move alien to city with alien",
			world:         w4,
			want:          w4After,
			wantAlternate: w4Alt,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.world.Move()

			if !reflect.DeepEqual(tc.world, tc.want) {
				// Move uses map iteration, which doesn't guarantee key order
				// To test a move where aliens and cities are destroyed, have
				// an alternate world to enumerate possible results.
				// An empty NewWorld() is considered the nil case.
				if !reflect.DeepEqual(tc.wantAlternate, NewWorld()) {
					if !reflect.DeepEqual(tc.world, tc.wantAlternate) {
						t.Errorf("got: %+v, want %+v", tc.world, tc.wantAlternate)
					}
				} else {
					t.Errorf("got: %+v, want %+v", tc.world, tc.want)
				}

			}
		})
	}
}
