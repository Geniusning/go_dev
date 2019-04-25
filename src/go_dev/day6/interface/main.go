package main

import "fmt"

type Carer interface {
	GetName() string
	Runing()
	Didi()
}

//BMW is type of car
type BMW struct {
	Name string
}

//GetName get the car name
func (p BMW) GetName() string {
	return p.Name
}

// Runing is action
func (p BMW) Runing() {
	fmt.Printf("%s is running", p.Name)
}

//Didi is action
func (p BMW) Didi() {
	fmt.Printf("%s is didi", p.Name)
}
func main() {
	var car Carer
	var c BMW
	c.Name = "BMW"
	car = c
	car.Runing()
}
