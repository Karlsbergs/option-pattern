package main

import "fmt"

const (
	defaultSpec      = "lifestyle"
	defaultAwdOption = false
)

//Car store new car instance
type Car struct {
	HasLeatherInterior bool
	HasAWD             bool
	HasPowerFrontSeat  bool
	Specs              string
}

//CarOption store CarOption
type CarOption func(*Car)

//WithLeatherInterior modify *Car behaviour
func WithLeatherInterior() CarOption {
	return func(c *Car) {
		c.HasLeatherInterior = true
	}
}

//WithPowerFrontSeat modify *Car behaviour
func WithPowerFrontSeat() CarOption {
	return func(c *Car) {
		c.HasPowerFrontSeat = true
	}
}

//WithAWD modify *Car behaviour
func WithAWD() CarOption {
	return func(c *Car) {
		c.HasAWD = true
	}
}

//WithSpecs modify *Car behaviour
func WithSpecs(s string) CarOption {
	return func(c *Car) {
		c.Specs = s
	}
}

//NewCar create new Car instance
func NewCar(opts ...CarOption) *Car {

	//create Car instance with default options

	car := &Car{
		HasAWD: defaultAwdOption,
		Specs:  defaultSpec,
	}

	for _, opt := range opts {
		opt(car)
	}

	return car
}

func main() {

	car := NewCar(WithLeatherInterior(), WithSpecs("prestige"), WithPowerFrontSeat())

	fmt.Println(car)
}
