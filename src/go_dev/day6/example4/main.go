package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
}
type sayHi interface {
	sayHello()
}
type BMW struct {
	name string
}

func (p *BMW) GetName() string {
	return p.name
}
func (p *BMW) Run() {
	fmt.Printf("%s is runing\n", p.name)
}

func (p *BMW) sayHello() {
	fmt.Printf("%s is say hello", p.name)
}
func main() {
	var car Carer
	var say sayHi

	bmw := &BMW{
		name: "baoMa",
	}
	car = bmw
	say = bmw
	car.Run()
	say.sayHello()
}
