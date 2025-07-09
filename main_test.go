package main

import (
	"testing"
)

func TestNotifyPoliceForAllCarsInParkingLot(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a car in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a car in lot B
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park a car in handicap lot
	car3 := Car{LicensePlate: "DEF789", Make: "BMW", Model: "X5", Color: "Black", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Simulate parking a car of another color
	car4 := Car{LicensePlate: "GHI012", Make: "Ford", Model: "Focus", Color: "White", Size: "small"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car4)

	// Notify police about all cars in the lot
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}

func TestNotifyPoliceWhenNoCarsInLot(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars

	// Notify police about all cars in the lot (should return nothing as no cars are parked)
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}
