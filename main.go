package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
)

const (
	// EarthRadius according to Wikipedia
	EarthRadius = 6371
)

// Point is a geo co-ordinate
type Point struct {
	lat float64
	lng float64
}

// BusStop describes a Singaporean (LTA) bus stop
type BusStop struct {
	BusStopCode string  `json:"BusStopCode"`
	RoadName    string  `json:"RoadName"`
	Description string  `json:"Description"`
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
}

// BusStops are many bus stops
type BusStops []BusStop

func loadBusJSON(jsonfile string) (bs []BusStop, err error) {
	content, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, &bs)
	if err != nil {
		return
	}

	return
}

func (BusStops BusStops) closest(location Point) (c BusStop) {
	c = BusStops[0]
	// fmt.Println(c)
	closestSoFar := location.GreatCircleDistance(Point{c.Latitude, c.Longitude})
	// log.Println(c.Description, closestSoFar)
	for _, p := range BusStops[1:] {
		distance := location.GreatCircleDistance(Point{p.Latitude, p.Longitude})
		// log.Printf("'%s' %.1f\n", p.Description, distance)
		if distance < closestSoFar {
			// Set the return
			c = p
			// Record closest distance
			closestSoFar = distance
		}
	}
	return
}

func (BusStops BusStops) nameBusStopID(busid string) (description string) {
	for _, p := range BusStops {
		if busid == p.BusStopCode {
			return p.Description
		}
	}
	return ""
}

// GreatCircleDistance calculates the distance between two points considering the curvature of planet earth
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
