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

// ParkingAttendant represents the parking attendant who parks cars.
type ParkingAttendant struct {
	Name string
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

	// Remove the car and increment available spots
	delete(pl.ParkedCars, licensePlate)
	pl.AvailableSpots++

	// Calculate how long the car was parked
	timeParked := time.Since(car.ParkedAt)
	fmt.Printf("Car %s unparked. It was parked for: %v\n", car.LicensePlate, timeParked)
	return true
}

// AssignCarToAttendant assigns a car to a parking attendant for parking.
func (attendant *ParkingAttendant) AssignCarToParkingLot(pl *ParkingLot, car Car) bool {
	if pl.ParkCar(car) {
		fmt.Printf("Parking attendant %s successfully parked the car %s.\n", attendant.Name, car.LicensePlate)
		return true
	}
	fmt.Println("Failed to park the car.")
	return false
}

// CheckIfFull checks if the parking lot is full.
func (pl *ParkingLot) CheckIfFull() bool {
	if pl.AvailableSpots == 0 {
		fmt.Println("The parking lot is full. Please put out the 'Full' sign.")
		pl.NotifySecurity()
		return true
	}
	fmt.Println("Parking lot has space available.")
	return false
}

// NotifySecurity notifies airport security when the parking lot is full.
func (pl *ParkingLot) NotifySecurity() {
	fmt.Println("Airport security has been notified that the parking lot is full.")
}

// NotifyOwner notifies the parking lot owner to take down the 'Full' sign.
func (pl *ParkingLot) NotifyOwner() {
	fmt.Println("For Owner : Parking lot has available space again.")
}

// Main function to simulate the parking lot operations.
func main() {
	// Create a new parking lot with 10 spots.
	parkingLot := NewParkingLot(1)

	// Create a parking attendant
	attendant := ParkingAttendant{Name: "Attendant"}

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

	// Create a Car object and attempt to park it via the parking attendant
	car := Car{LicensePlate: licensePlate, Make: make, Model: model, Color: color}

	// Parking attendant parks the car
	attendant.AssignCarToParkingLot(parkingLot, car)

	// Check if the parking lot is full and notify security
	parkingLot.CheckIfFull()

	// Simulate driver unparking a car
	var unparkPlate string
	fmt.Print("\nEnter License Plate to Unpark: ")
	fmt.Scan(&unparkPlate)

	if parkingLot.UnparkCar(unparkPlate) {
		fmt.Printf("Successfully unparked car %s.\n", unparkPlate)
	} else {
		fmt.Println("Failed to unpark the car.")
	}

	// Check parking lot status after unparking
	parkingLot.CheckIfFull()

	// Notify the owner that there's space available again
	parkingLot.NotifyOwner()
}
