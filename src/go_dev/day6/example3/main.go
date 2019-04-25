package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["BellLabs"] = Vertex{
		40.2345, 235.235,
	}
	fmt.Println(m)
}
