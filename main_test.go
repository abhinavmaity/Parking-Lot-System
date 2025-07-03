package main

import (
	"testing"
	"time"
)

func TestFindCar(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.AssignCarToParkingLot(parkingLot, car1)

	// Test finding the car by license plate
	parkingLot.FindCar("ABC123")

	// Test searching for a car that doesn't exist
	parkingLot.FindCar("NONEXISTENT")
}

func TestFindCarAndCheckParkingTime(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.AssignCarToParkingLot(parkingLot, car1)

	// Wait for a second to simulate parking time
	time.Sleep(time.Second)

	// Test finding the car by license plate and check parking time
	parkingLot.FindCar("ABC123")

	// Test if the car's parked time is greater than zero
	if car1.ParkedAt.IsZero() {
		t.Errorf("Expected parked time to be set, but it wasn't")
	}
}
