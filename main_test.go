package main

import (
	"testing"
)

func TestHandicapCarParking(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a handicap car in the Handicap Lot
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", isHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Verify that the car is parked in the Handicap Lot
	if _, exists := handicapLot.ParkedCars[car1.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in Handicap Lot, but it wasn't", car1.LicensePlate)
	}

	// Park a non-handicap car
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", isHandicap: false}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Verify that the car is parked in the lot with fewer cars (since Handicap Lot is full)
	if _, exists := lotA.ParkedCars[car2.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot A, but it wasn't", car2.LicensePlate)
	}
}

func TestHandicapLotFullAndDirectToOtherLot(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 1)                   // Create a parking lot with 1 spot
	lotB := NewParkingLot("B", 1)                   // Create another parking lot with 1 spot
	handicapLot := NewParkingLot("Handicap Lot", 1) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a handicap car in the Handicap Lot
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", isHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park another handicap car when Handicap Lot is full
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", isHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Verify that the second car is directed to the lot with the fewest cars (lot B)
	if _, exists := lotB.ParkedCars[car2.LicensePlate]; !exists {
		t.Errorf("Expected car %s to be parked in lot B, but it wasn't", car2.LicensePlate)
	}
}

func TestUnparkHandicapCar(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)
	lotB := NewParkingLot("B", 5)
	handicapLot := NewParkingLot("Handicap Lot", 3)
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a handicap car in the Handicap Lot
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", isHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Simulate unparking the car and check if available spots are updated
	unparked := handicapLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car %s to be unparked, but it wasn't", car1.LicensePlate)
	}

	// Check if available spots are updated
	if handicapLot.AvailableSpots != 3 {
		t.Errorf("Expected available spots to be 3, but got %d", handicapLot.AvailableSpots)
	}
}
