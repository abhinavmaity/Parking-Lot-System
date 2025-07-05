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
	IsHandicap   bool
	Size         string // Added a Size attribute to categorize car sizes (small, medium, large)
}

// ParkingLot represents a parking lot with a certain capacity.
type ParkingLot struct {
	Name           string
	Capacity       int
	AvailableSpots int
	ParkedCars     map[string]Car // Map of license plates to cars
}

// ParkingAttendant represents the parking attendant who parks cars.
type ParkingAttendant struct {
	Name string
}

// NewParkingLot creates a new parking lot with the given capacity.
func NewParkingLot(name string, capacity int) *ParkingLot {
	return &ParkingLot{
		Name:           name,
		Capacity:       capacity,
		AvailableSpots: capacity,
		ParkedCars:     make(map[string]Car),
	}
}

// ParkCar parks a car in the parking lot.
func (pl *ParkingLot) ParkCar(car Car) bool {
	if pl.AvailableSpots == 0 {
		fmt.Printf("Parking lot %s is full!\n", pl.Name)
		return false
	}

	// Add car to parking lot and reduce available spots
	car.ParkedAt = time.Now() // Record the time the car is parked
	pl.ParkedCars[car.LicensePlate] = car
	pl.AvailableSpots--

	fmt.Printf("Car %s parked at %v in lot %s\n", car.LicensePlate, car.ParkedAt, pl.Name)
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

	// Calculate the charge based on parking time (Rs 50 per hour)
	charge := float64(timeParked.Hours()) * 50.0
	fmt.Printf("Car %s unparked. It was parked for: %v\n", car.LicensePlate, timeParked)
	fmt.Printf("Charge for parking: Rs%.2f\n", charge)
	return true
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

// DirectCarToLot directs a car to the lot with the least number of cars.
func (attendant *ParkingAttendant) DirectCarToLot(plots []*ParkingLot, car Car) bool {
	var selectedLot *ParkingLot

	// If it's a large car, prioritize parking in the lot with the most available space
	if car.Size == "large" {
		for _, lot := range plots {
			if selectedLot == nil || lot.AvailableSpots > selectedLot.AvailableSpots {
				selectedLot = lot
			}
		}
	} else {
		// For non-large cars, direct to the lot with the least number of cars
		// If both lots have equal available spots, prioritize the first lot
		for _, lot := range plots {
			if selectedLot == nil || lot.AvailableSpots < selectedLot.AvailableSpots {
				selectedLot = lot
			} else if lot.AvailableSpots == selectedLot.AvailableSpots && lot.Name < selectedLot.Name {
				// Prioritize the first lot when both have equal available spots
				selectedLot = lot
			}
		}
	}

	if selectedLot != nil {
		selectedLot.ParkCar(car)
		fmt.Printf("Parking attendant %s directed car %s to lot %s.\n", attendant.Name, car.LicensePlate, selectedLot.Name)
		return true
	}

	fmt.Println("No available lot found for parking.")
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
	// Create multiple parking lots
	lotA := NewParkingLot("A", 5)
	lotB := NewParkingLot("B", 5)
	handicapLot := NewParkingLot("Handicap Lot", 3) // Special lot for handicap cars

	// Create a parking attendant
	attendant := ParkingAttendant{Name: "Rahul"}

	// Simulate a large driver parking a car
	var licensePlate, make, model, color, size string
	fmt.Println("Enter car details to park (for Large Car Parking):")

	// Simulate user input for car details
	fmt.Print("License Plate: ")
	fmt.Scan(&licensePlate)
	fmt.Print("Make: ")
	fmt.Scan(&make)
	fmt.Print("Model: ")
	fmt.Scan(&model)
	fmt.Print("Color: ")
	fmt.Scan(&color)
	fmt.Print("Size (small, medium, large): ")
	fmt.Scan(&size)

	// Create a Car object and mark it as a large car
	car := Car{LicensePlate: licensePlate, Make: make, Model: model, Color: color, Size: size}

	// Create a slice of parking lots and direct the large car to the lot with the most available space
	attendant.DirectCarToLot([]*ParkingLot{lotA, lotB, handicapLot}, car)

	// Check if the parking lot is full and notify security
	lotA.CheckIfFull()
	lotB.CheckIfFull()
	handicapLot.CheckIfFull()

	// Driver wants to find their car by license plate
	var findPlate string
	fmt.Print("\nEnter License Plate to Find: ")
	fmt.Scan(&findPlate)

	// Simulate driver unparking a car
	var unparkPlate string
	fmt.Print("\nEnter License Plate to Unpark: ")
	fmt.Scan(&unparkPlate)

	lotA.UnparkCar(unparkPlate)
	lotB.UnparkCar(unparkPlate)
	handicapLot.UnparkCar(unparkPlate)

	// Check parking lot status after unparking
	lotA.CheckIfFull()
	lotB.CheckIfFull()
	handicapLot.CheckIfFull()

	// Notify the owner that there's space available again
	lotA.NotifyOwner()
	lotB.NotifyOwner()
	handicapLot.NotifyOwner()
}
