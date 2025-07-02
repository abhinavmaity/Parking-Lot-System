package main

import (
	"testing"
)

func TestAssignCarToParkingLot(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "John"}

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parked := attendant.AssignCarToParkingLot(parkingLot, car1)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Check if the parking lot's available spots decreased
	if parkingLot.AvailableSpots != 0 {
		t.Errorf("Expected available spots to be 0, but got %d", parkingLot.AvailableSpots)
	}

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked = attendant.AssignCarToParkingLot(parkingLot, car2)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}

func TestUnparkCarByAttendant(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "John"}
	car := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}

	// Park the car
	attendant.AssignCarToParkingLot(parkingLot, car)

	// Unpark the car and check if available spots have increased
	unparked := parkingLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car to be unparked, but it wasn't")
	}

	// Check if available spots have increased
	if parkingLot.AvailableSpots != 1 {
		t.Errorf("Expected available spots to be 1, but got %d", parkingLot.AvailableSpots)
	}
}

func TestAttendantParkingLotFull(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "John"}

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.AssignCarToParkingLot(parkingLot, car1)

	// Try to park another car in a full lot
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked := attendant.AssignCarToParkingLot(parkingLot, car2)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}

func TestAttendantNotifySecurityWhenFull(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(1) // Create a parking lot with 1 spot
	attendant := ParkingAttendant{Name: "John"}

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	attendant.AssignCarToParkingLot(parkingLot, car1)

	// Check if security is notified when the lot is full
	parkingLot.CheckIfFull() // This simulates the security notification being triggered
	// You can expand this by checking if the print statement "Airport security has been notified..." is triggered
}
