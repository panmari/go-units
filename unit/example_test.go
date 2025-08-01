package unit_test

import (
	"fmt"

	"github.com/golang/geo/earth"
	"github.com/golang/geo/s2"
	"github.com/google/go-units/unit"
)

func ExampleLength() {
	// Define a length using a constant.
	distance := 500 * unit.Mile

	// Define a length from a variable.
	var altitudeInFeet = 29031
	altitude := unit.Length(altitudeInFeet) * unit.Foot

	// Convert to different units.
	fmt.Printf("I would walk %.0f miles, but also %.1f kilometers.\n", distance.Miles(), distance.Kilometers())
	fmt.Printf("Altitude: %.2f m\n", altitude.Meters())

	// Perform calculations.
	totalDistance := 2 * distance
	fmt.Printf("And after I walk %.0f miles more, that's a total of %.1f nautical miles.\n", distance.Miles(), totalDistance.NauticalMiles())

	// Output:
	// I would walk 500 miles, but also 804.7 kilometers.
	// Altitude: 8848.65 m
	// And after I walk 500 miles more, that's a total of 869.0 nautical miles.
}

func ExampleLength_fromS2() {
	sanFrancisco := s2.LatLngFromDegrees(37.7749, -122.4194)
	zurich := s2.LatLngFromDegrees(47.3769, 8.5417)

	d := earth.LengthFromLatLngs(sanFrancisco, zurich)

	fmt.Printf("The distance from San Francisco to Zurich is %.2f km or %.2f miles.\n", d.Kilometers(), d.Miles())
	// Output:
	// The distance from San Francisco to Zurich is 9370.33 km or 5822.45 miles.
}
