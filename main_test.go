package main

import (
	"testing"
)

func TestCheckIfFullAndNotifySecurity(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parked := parkingLot.ParkCar(car1)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Test if the parking lot is full after parking one car
	if !parkingLot.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Test if security is notified when the lot is full
	// Simulating the NotifySecurity method being called
	parkingLot.NotifySecurity() // You can expand this by checking if the NotifySecurity logic works in the print statement

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked = parkingLot.ParkCar(car2)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}

	// Unpark the first car and check if available spots have increased
	unparked := parkingLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car to be unparked, but it wasn't")
	}

	// Check if available spots have increased
	if parkingLot.AvailableSpots != 1 {
		t.Errorf("Expected available spots to be 1, but got %d", parkingLot.AvailableSpots)
	}

	// Check if the parking lot has space after unparking
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}
