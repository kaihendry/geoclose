package main

import (
	"fmt"
	"log"
	"math"
)

const (
	// EarthRadius according to Wikipedia
	EarthRadius = 6371
)

type Point struct {
	lat float64
	lng float64
}

type place struct {
	Title string
	Point Point
}

var places = []place{
	{
		Title: "Table Mountain",
		Point: Point{
			lat: -33.957314,
			lng: 18.403108,
		},
	},
	{
		Title: "Statue of liberty",
		Point: Point{
			lat: 40.689167,
			lng: -74.044444,
		},
	},
}

func main() {
	whereiam := place{
		Title: "Middle Earth",
		Point: Point{
			lat: 0.0,
			lng: 0.0,
		},
	}

	fmt.Println("Closest point to", whereiam.Title, "is", closest(whereiam))

}

func closest(w place) (c place) {
	c = places[0]
	closestSoFar := w.Point.GreatCircleDistance(c.Point)
	log.Println(c.Title, closestSoFar)
	for _, p := range places[1:] {
		distance := w.Point.GreatCircleDistance(p.Point)
		log.Println(p.Title, distance)
		if distance < closestSoFar {
			// Set the return
			c = p
			// Record closest distance
			closestSoFar = distance
		}
	}
	return
}

// From https://github.com/kellydunn/golang-geo/blob/master/point.go#L70
func (p Point) GreatCircleDistance(p2 Point) float64 {
	dLat := (p2.lat - p.lat) * (math.Pi / 180.0)
	dLon := (p2.lng - p.lng) * (math.Pi / 180.0)

	lat1 := p.lat * (math.Pi / 180.0)
	lat2 := p2.lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadius * c
}
