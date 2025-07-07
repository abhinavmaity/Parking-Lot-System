package main

import (
	"testing"
)

func TestNotifyPoliceForBlueToyotaCars(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a blue Toyota car in lot A
	car1 := Car{LicensePlate: "BT123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-blue Toyota car in lot B
	car2 := Car{LicensePlate: "GT456", Make: "Toyota", Model: "Corolla", Color: "Green", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park another blue Toyota car in handicap lot
	car3 := Car{LicensePlate: "BT789", Make: "Toyota", Model: "Highlander", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Notify police about the location of all blue Toyota cars
	lotA.NotifyPolice("Toyota", "Blue")
	lotB.NotifyPolice("Toyota", "Blue")
	handicapLot.NotifyPolice("Toyota", "Blue")
}

func TestNotifyPoliceWhenNoBlueToyotaCars(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a non-blue car in lot A
	car1 := Car{LicensePlate: "GT123", Make: "Toyota", Model: "Camry", Color: "Green", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a non-blue car in lot B
	car2 := Car{LicensePlate: "RT456", Make: "Toyota", Model: "Corolla", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Notify police about the location of all blue Toyota cars (expecting no cars)
	lotA.NotifyPolice("Toyota", "Blue")
	lotB.NotifyPolice("Toyota", "Blue")
	handicapLot.NotifyPolice("Toyota", "Blue")
}
