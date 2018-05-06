package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var n int
	flag.IntVar(&n, "n", 0, fmt.Sprintf("specify number of aliens, defaults to 0, max %d.", getMax()))

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./go-alien-invalsion -n <numAliens>")
		os.Exit(1)
	}

	flag.Parse()

	world := NewWorld()
	world.Cities = parseMap()

	world.CreateAliens(n)
	for !world.GameOver() {
		world.Move()
	}

	printWorld(world.Cities)
}

// parseMap assumes that each line entry of the map is well-formed and
// that the map geography makes sense.
// Duplicate entries are ignored and empty lines are skipped.
// No restrictions are enforced on city names.
// Allowed cardinal directions are north, south, east, west. Other connections
// are skipped.
// If a city is specified as a connection without a line entry in the map,
// no aliens can be placed or move there.
func parseMap() map[string]*City {
	//world.Initialize
	f, err := os.Open("testworld.txt")
	if err != nil {
		log.Printf("error opening map file, exiting")
		os.Exit(1)
	}
	defer f.Close()

	m := make(map[string]*City)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		// skip empty lines
		if len(t) == 0 {
			continue
		}
		tArr := strings.Fields(t)
		name := tArr[0]
		// skip duplicate entries
		if _, ok := m[name]; ok {
			continue
		}
		c := make(map[string]string)
		// city has 1-4 connections
		if len(tArr) > 1 || len(tArr) > 4 {
			for _, v := range tArr[1:] {
				pair := strings.Split(v, "=")
				// skip malformed connections
				if len(pair) != 2 {
					continue
				}
				dir := pair[0]
				cityName := pair[1]
				// skip dupe connections
				if _, ok := c[cityName]; ok {
					continue
				}
				if isValidDir(dir) {
					c[cityName] = dir
				}
			}
		}
		city := NewCity(name, c)
		m[city.Name] = city
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error parsing map")
		os.Exit(1)
	}
	return m
}

func isValidDir(s string) bool {
	switch s {
	case "north":
	case "south":
	case "east":
	case "west":
	default:
		return false
	}
	return true
}

func printWorld(m map[string]*City) {
	for _, c := range m {
		var b bytes.Buffer
		b.WriteString(c.Name)

		for city, dir := range c.Connections {
			b.WriteString(" ")
			b.WriteString(dir)
			b.WriteString("=")
			b.WriteString(city)
		}
		fmt.Println(b.String())
	}
}

//Returns max int value based on system architecture
func getMax() int {
	const intSize = 32 << (^uint(0) >> 63)
	if intSize == 32 {
		return math.MaxInt32
	} else {
		return math.MaxInt64
	}
}
