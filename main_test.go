package main

import (
	"testing"
	"time"
)

func TestNotifyPoliceForCarsParkedInLast30Minutes(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a car that was just parked (in the last 30 minutes)
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Wait for a second to simulate parking time
	time.Sleep(time.Second)

	// Park another car that was just parked (in the last 30 minutes)
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park a car that was parked more than 30 minutes ago
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Wait for more than 30 minutes to simulate the 30-minute threshold
	time.Sleep(30 * time.Minute)

	// Notify police about the cars parked in the last 30 minutes
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}

func TestNotifyPoliceWhenNoCarsParkedInLast30Minutes(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a car more than 30 minutes ago
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Wait for more than 30 minutes to simulate the 30-minute threshold
	time.Sleep(30 * time.Minute)

	// Park another car more than 30 minutes ago
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "medium"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Simulate parking a car more than 30 minutes ago
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green", Size: "large"}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Notify police about the cars parked in the last 30 minutes (expecting no cars)
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}
