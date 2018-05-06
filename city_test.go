package main

import (
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestRemoveConnection(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
		"c4": "east",
		"c5": "west",
	})

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	c2After := NewCity("c2", map[string]string{})

	// Test cases
	cases := []struct {
		testName string
		city     string
		input    *City
		want     *City
	}{
		{
			testName: "no connection",
			city:     "dne",
			input:    c1,
			want:     c1,
		},
		{
			testName: "destroy connection",
			city:     "c1",
			input:    c2,
			want:     c2After,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			tc.input.RemoveConnection(tc.city)

			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("got: %+v, want %+v", tc.input, tc.want)
			}
		})
	}
}

func TestHasAlien(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{
		"c2": "north",
		"c3": "south",
		"c4": "east",
		"c5": "west",
	})

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	c2.SetAlien(1)

	// Test cases
	cases := []struct {
		testName string
		city     *City
		want     bool
	}{
		{
			testName: "alien dne",
			city:     c1,
			want:     false,
		},
		{
			testName: "alien exits",
			city:     c2,
			want:     true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			got := tc.city.HasAlien()

			if got != tc.want {
				t.Errorf("got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

func TestSetAlien(t *testing.T) {
	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	if c2.Alien != -1 {
		t.Errorf("got: %+v, want: %+v", c2.Alien, -1)
	}

	c2.SetAlien(1)
	if c2.Alien != 1 {
		t.Errorf("got: %+v, want: %+v", c2.Alien, 1)
	}

	c2.RemoveAlien()
	if c2.Alien != -1 {
		t.Errorf("got: %+v, want: %+v", c2.Alien, -1)
	}

}

func TestGetMove(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	// Initial setup
	c1 := NewCity("c1", map[string]string{})

	c2 := NewCity("c2", map[string]string{
		"c1": "south",
	})

	// Test cases
	cases := []struct {
		testName string
		city     *City
		want     string
	}{
		{
			testName: "no connections",
			city:     c1,
			want:     "",
		},
		{
			testName: "has one connection",
			city:     c2,
			want:     "c1",
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			got := tc.city.GetMove()

			if got != tc.want {
				t.Errorf("got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}
