package main

import (
	"testing"
)

func TestAttendantDirectsCarToLotWithLeastCars(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5) // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5) // Create another parking lot with 5 spots
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB}, car1)

	// Verify that the car is parked in lot A
	if _, exists := lotA.ParkedCars[car1.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot A, but it wasn't", car1.LicensePlate)
	}

	// Park the second car in lot B
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB}, car2)

	// Verify that the car is parked in lot B
	if _, exists := lotB.ParkedCars[car2.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot B, but it wasn't", car2.LicensePlate)
	}

	// Park the third car when both lots are equal
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB}, car3)

	// Verify that the third car is parked in lot A (first available lot)
	if _, exists := lotA.ParkedCars[car3.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot A, but it wasn't", car3.LicensePlate)
	}
}

func TestParkingLotWhenFull(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.DirectCarToLot([]*ParkingLot{lotA}, car1)

	// Verify the lot is full after the first car
	if !lotA.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked := attendant.DirectCarToLot([]*ParkingLot{lotA}, car2)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}
