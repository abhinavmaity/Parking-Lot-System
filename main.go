package main

import (
	"fmt"
	"time"
)

// Car represents a car with relevant details.
type Car struct {
	LicensePlate string
	Make         string
	Model        string
	Color        string
	ParkedAt     time.Time
}

// ParkingLot represents a parking lot with a certain capacity.
type ParkingLot struct {
	Capacity       int
	AvailableSpots int
	ParkedCars     map[string]Car // Map of license plates to cars
}

// NewParkingLot creates a new parking lot with the given capacity.
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		Capacity:       capacity,
		AvailableSpots: capacity,
		ParkedCars:     make(map[string]Car),
	}
}

// ParkCar parks a car in the parking lot.
func (pl *ParkingLot) ParkCar(car Car) bool {
	if pl.AvailableSpots == 0 {
		fmt.Println("Parking lot is full!")
		return false
	}

	// Add car to parking lot and reduce available spots
	car.ParkedAt = time.Now() // Record the time the car is parked
	pl.ParkedCars[car.LicensePlate] = car
	pl.AvailableSpots--

	fmt.Printf("Car %s parked at %v\n", car.LicensePlate, car.ParkedAt)
	return true
}

// UnparkCar removes a car from the parking lot.
func (pl *ParkingLot) UnparkCar(licensePlate string) bool {
	car, exists := pl.ParkedCars[licensePlate]
	if !exists {
		fmt.Println("Car not found!")
		return false
	}

	delete(pl.ParkedCars, licensePlate)
	pl.AvailableSpots++
	fmt.Printf("Car %s unparked. Time parked: %v\n", car.LicensePlate, time.Since(car.ParkedAt))
	return true
}

// Main function to simulate the parking lot operations.
func main() {
	// Create a new parking lot with 10 spots.
	parkingLot := NewParkingLot(10)

	// Simulate a driver parking a car
	var licensePlate, make, model, color string
	fmt.Println("Enter car details to park:")

	// Simulate user input for car details
	fmt.Print("License Plate: ")
	fmt.Scan(&licensePlate)
	fmt.Print("Make: ")
	fmt.Scan(&make)
	fmt.Print("Model: ")
	fmt.Scan(&model)
	fmt.Print("Color: ")
	fmt.Scan(&color)

	// Create a Car object and attempt to park it
	car := Car{LicensePlate: licensePlate, Make: make, Model: model, Color: color}

	if parkingLot.ParkCar(car) {
		fmt.Printf("Successfully parked car %s.\n", car.LicensePlate)
	} else {
		fmt.Println("Failed to park the car, parking lot is full.")
	}
}
