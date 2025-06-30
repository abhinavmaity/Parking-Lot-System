package main

import (
	"testing"
)

func TestParkCar(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	car := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}

	// Test parking when there's space available
	parked := parkingLot.ParkCar(car)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Test if available spots have decreased
	if parkingLot.AvailableSpots != 1 {
		t.Errorf("Expected available spots to be 1, but got %d", parkingLot.AvailableSpots)
	}

	// Test parking when the parking lot is full
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parkingLot.ParkCar(car2) // Park second car
	parked = parkingLot.ParkCar(Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green"})

	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}

func TestUnparkCar(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	car := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car)

	// Test unparking when the car exists
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

func TestCarParkingFullLot(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	car := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car)

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked := parkingLot.ParkCar(car2)

	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}
