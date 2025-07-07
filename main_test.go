package main

import (
	"testing"
)

func TestNotifyPoliceForBMWCars(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a black BMW car in lot A
	car1 := Car{LicensePlate: "BMW123", Make: "BMW", Model: "X5", Color: "Black", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-BMW car in lot B
	car2 := Car{LicensePlate: "GT456", Make: "Toyota", Model: "Camry", Color: "Green", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park another black BMW car in handicap lot
	car3 := Car{LicensePlate: "BMW789", Make: "BMW", Model: "3 Series", Color: "Black", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Notify police about the location of all BMW cars
	lotA.NotifyPolice("BMW", "Black")
	lotB.NotifyPolice("BMW", "Black")
	handicapLot.NotifyPolice("BMW", "Black")
}

func TestNotifyPoliceWhenNoBMWCars(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a non-BMW car in lot A
	car1 := Car{LicensePlate: "GT123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-BMW car in lot B
	car2 := Car{LicensePlate: "RD456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Notify police about the location of all BMW cars (expecting no cars)
	lotA.NotifyPolice("BMW", "Black")
	lotB.NotifyPolice("BMW", "Black")
	handicapLot.NotifyPolice("BMW", "Black")
}
