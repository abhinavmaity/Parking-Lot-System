package main

import (
	"testing"
)

func TestNotifyPoliceForSpecificColor(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a white car in lot A
	car1 := Car{LicensePlate: "WH123", Make: "Toyota", Model: "Corolla", Color: "White", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-white car in lot B
	car2 := Car{LicensePlate: "BL456", Make: "Honda", Model: "Civic", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park another white car in handicap lot
	car3 := Car{LicensePlate: "WH789", Make: "Ford", Model: "Focus", Color: "White", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Notify police about the location of all white cars
	lotA.NotifyPolice("White")
	lotB.NotifyPolice("White")
	handicapLot.NotifyPolice("White")
}

func TestNotifyPoliceWhenNoCarsOfColor(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a non-white car in lot A
	car1 := Car{LicensePlate: "BL123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-white car in lot B
	car2 := Car{LicensePlate: "RD456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Notify police about the location of all white cars (expecting no cars)
	lotA.NotifyPolice("White")
	lotB.NotifyPolice("White")
	handicapLot.NotifyPolice("White")
}
