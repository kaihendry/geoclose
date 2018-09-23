package main

import (
	"log"
	"reflect"
	"testing"
)

var busStops Places

func init() {
	var err error
	busStops, err = loadBusJSON("all.json")
	if err != nil {
		log.Fatal(err)
	}
}

func TestPlaces_closest(t *testing.T) {
	tests := []struct {
		name   string
		places Places
		w      place
		wantC  place
	}{
		{
			name:   "All bus stops",
			places: busStops,
			w: place{
				Title: "Middle earth",
				Point: Point{
					lat: 0.0,
					lng: 0.0,
				},
			},
			wantC: place{
				Title: "BEF TUAS STH AVE 14",
				Point: Point{lat: 1.27637, lng: 103.621508},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := tt.places.closest(tt.w); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Places.closest() = %+v, want %+v", gotC, tt.wantC)
			}
		})
	}
}
