package main

import (
	"testing"
)

func TestDirectLargeCarToLotWithMostAvailableSpace(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 2)                   // Create a parking lot with 2 spots
	lotB := NewParkingLot("B", 2)                   // Create another parking lot with 2 spots
	handicapLot := NewParkingLot("Handicap Lot", 1) // Create a Handicap lot
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a small car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "small"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a large car in lot B
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Verify that the large car is parked in the lot with the most available space (lot A)
	if _, exists := lotA.ParkedCars[car2.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot A, but it wasn't", car2.LicensePlate)
	}
}

func TestDirectCarWhenBothLotsHaveEqualAvailableSpace(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 2)                   // Create a parking lot with 2 spots
	lotB := NewParkingLot("B", 2)                   // Create another parking lot with 2 spots
	handicapLot := NewParkingLot("Handicap Lot", 1) // Create a Handicap lot
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park the second car in lot B
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park the third car when both lots have equal available space
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Verify that the third car is parked in lot A (first available lot)
	if _, exists := lotA.ParkedCars[car3.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot A, but it wasn't", car3.LicensePlate)
	}
}

func TestParkCarInFullLot(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "small"}
	attendant.DirectCarToLot([]*ParkingLot{lotA}, car1)

	// Verify the lot is full after the first car
	if !lotA.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "large"}
	parked := attendant.DirectCarToLot([]*ParkingLot{lotA}, car2)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}
