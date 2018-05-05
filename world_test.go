package main

import (
	"reflect"
	"testing"
)

func TestDestroyCity(t *testing.T) {
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

	w := World{
		Cities: map[string]City{
			"c1": c1,
			"c2": c2,
			"c3": c3,
		},
	}

	wAfter := World{
		Cities: map[string]City{
			"c2": c2After,
			"c3": c3After,
		},
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
			tc.input.DestroyCity(tc.city)

			if !reflect.DeepEqual(tc.input, tc.want) {
				t.Errorf("got: %+v, want %+v", tc.input, tc.want)
			}
		})
	}
}
