package main

import (
	"testing"
)

func TestNotifyPoliceForSmallHandicapCarsOnRowsBandD(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a small handicap car on row B in lot A
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "small", Row: "B", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a small handicap car on row D in lot B
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "small", Row: "D", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park a small non-handicap car (shouldn't be notified)
	car3 := Car{LicensePlate: "DEF789", Make: "Ford", Model: "Focus", Color: "Green", Size: "small", Row: "A", IsHandicap: false}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Park a non-small handicap car (shouldn't be notified)
	car4 := Car{LicensePlate: "LMN123", Make: "BMW", Model: "X5", Color: "Black", Size: "large", Row: "B", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car4)

	// Park a small handicap car on a row that is not B or D (shouldn't be notified)
	car5 := Car{LicensePlate: "PQR456", Make: "Mercedes", Model: "C-Class", Color: "White", Size: "small", Row: "C", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car5)

	// Notify police about the small handicap cars on rows B or D
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}

func TestNotifyPoliceWhenNoSmallHandicapCarsOnRowsBandD(t *testing.T) {
	// Setup
	lotA := NewParkingLot("A", 5)                   // Create a parking lot with 5 spots
	lotB := NewParkingLot("B", 5)                   // Create another parking lot with 5 spots
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars
	attendant := ParkingAttendant{Name: "Rahul"}

	// Park a small non-handicap car on row A (shouldn't be notified)
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", Size: "small", Row: "A", IsHandicap: false}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car1)

	// Park a large handicap car on row B (shouldn't be notified)
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", Size: "large", Row: "B", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car2)

	// Park a non-small handicap car (shouldn't be notified)
	car3 := Car{LicensePlate: "DEF789", Make: "BMW", Model: "X5", Color: "Black", Size: "large", Row: "B", IsHandicap: true}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car3)

	// Park a small non-handicap car on row C (shouldn't be notified)
	car4 := Car{LicensePlate: "PQR456", Make: "Ford", Model: "Focus", Color: "Green", Size: "small", Row: "C", IsHandicap: false}
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car4)

	// Notify police about the small handicap cars on rows B or D (expecting no cars)
	lotA.NotifyPolice()
	lotB.NotifyPolice()
	handicapLot.NotifyPolice()
}
