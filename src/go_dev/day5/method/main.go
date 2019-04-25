package main

import "fmt"

type Car struct {
	weight int
	with   int
}
type Bike struct {
	Car
	weight int
	height int
}

type Train struct {
	Car
	Bike
}

func (p Car) run() {
	fmt.Println("running")
}

func main() {
	var c Train
	c.run()
	c.Car.weight = 100
	fmt.Println(c)
}
