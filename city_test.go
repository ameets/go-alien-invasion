package main

import (
	"reflect"
	"testing"
)

func TestRemoveConnection(t *testing.T) {
	// Initial setup
	c1 := City{
		Name: "c1",
		Connections: map[string]string{
			"c2": "north",
			"c3": "south",
			"c4": "east",
			"c5": "west",
		},
	}

	c2 := City{
		Name: "c2",
		Connections: map[string]string{
			"c1": "south",
		},
	}
	// Test cases
	cases := []struct {
		testName string
		city     string
		input    City
		want     City
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
			want:     City{Name: "c2", Connections: map[string]string{}},
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