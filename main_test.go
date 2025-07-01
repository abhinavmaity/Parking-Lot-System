package main

import (
	"testing"
)

func TestUnparkCar(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	car := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car)

	// Test if the car can be unparked successfully
	unparked := parkingLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car to be unparked, but it wasn't")
	}

	// Test if available spots have increased
	if parkingLot.AvailableSpots != 2 {
		t.Errorf("Expected available spots to be 2, but got %d", parkingLot.AvailableSpots)
	}

	// Test unparking a car that doesn't exist in the lot
	unparked = parkingLot.UnparkCar("NONEXISTENT")
	if unparked {
		t.Errorf("Expected error when unparking non-existent car, but it succeeded")
	}
}
