package parking_lot_system

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

	pl.ParkedCars[car.LicensePlate] = car
	pl.AvailableSpots--
	fmt.Printf("Car %s parked at %v\n", car.LicensePlate, time.Now())
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

	// Park a few cars
	car1 := Car{LicensePlate: "ABC123", Make: "Toyota", Model: "Camry", Color: "Blue", ParkedAt: time.Now()}
	car2 := Car{LicensePlate: "XYZ456", Make: "Honda", Model: "Civic", Color: "Red", ParkedAt: time.Now()}

	parkingLot.ParkCar(car1)
	parkingLot.ParkCar(car2)

	// Unpark a car
	parkingLot.UnparkCar("ABC123")
}
