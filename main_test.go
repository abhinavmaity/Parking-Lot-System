package main

import (
	"testing"
)

func TestAvailableSpaceAfterUnparking(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parked := parkingLot.ParkCar(car1)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Check if the parking lot is full after parking one car
	if !parkingLot.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Unpark the car and check available spots
	unparked := parkingLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car to be unparked, but it wasn't")
	}

	// Check if available spots have increased
	if parkingLot.AvailableSpots != 1 {
		t.Errorf("Expected available spots to be 1, but got %d", parkingLot.AvailableSpots)
	}

	// Check if the parking lot has space available after unparking
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}

func TestCheckIfFullAfterUnparking(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car1)

	// Check if the parking lot is full after parking one car
	if !parkingLot.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Unpark the car and check the status again
	parkingLot.UnparkCar("ABC123")

	// Check parking lot status after unparking
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}

func TestSpaceAvailableAfterMultipleCars(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parked := parkingLot.ParkCar(car1)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Park the second car
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked = parkingLot.ParkCar(car2)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Check if the parking lot is full after parking two cars
	if !parkingLot.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Unpark the first car and check available spots
	parkingLot.UnparkCar("ABC123")

	// Check if available spots have increased
	if parkingLot.AvailableSpots != 1 {
		t.Errorf("Expected available spots to be 1, but got %d", parkingLot.AvailableSpots)
	}

	// Check parking lot status after unparking
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}
