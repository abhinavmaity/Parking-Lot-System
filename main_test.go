package main

import (
	"testing"
)

func TestCheckIfFull(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots

	// Park the first car
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parked := parkingLot.ParkCar(car1)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Check if the lot is full after parking one car
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}

	// Park the second car
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parked = parkingLot.ParkCar(car2)
	if !parked {
		t.Errorf("Expected car to be parked, but it wasn't")
	}

	// Check if the lot is full after parking two cars
	if parkingLot.CheckIfFull() != true {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Test parking when the lot is full
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green"}
	parked = parkingLot.ParkCar(car3)
	if parked {
		t.Errorf("Expected parking lot to be full, but car was parked")
	}
}

func TestAvailableSpaceAfterUnparking(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car1)

	// Unpark the car and check available spots
	unparked := parkingLot.UnparkCar("ABC123")
	if !unparked {
		t.Errorf("Expected car to be unparked, but it wasn't")
	}

	// Check if available spots have increased
	if parkingLot.AvailableSpots != 2 {
		t.Errorf("Expected available spots to be 2, but got %d", parkingLot.AvailableSpots)
	}

	// Check if the parking lot is not full after unparking the car
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}

func TestFullLotAfterParkingAndUnparking(t *testing.T) {
	// Setup
	parkingLot := NewParkingLot(2) // Create a parking lot with 2 spots
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue"}
	parkingLot.ParkCar(car1)

	// Simulate the parking lot becoming full
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red"}
	parkingLot.ParkCar(car2)

	// Check if the lot is full
	if !parkingLot.CheckIfFull() {
		t.Errorf("Expected parking lot to be full, but it wasn't")
	}

	// Unpark a car and check the lot status
	parkingLot.UnparkCar("ABC123")

	// Verify the lot is not full after unparking a car
	if parkingLot.CheckIfFull() != false {
		t.Errorf("Expected parking lot to have space available, but it was full")
	}
}
