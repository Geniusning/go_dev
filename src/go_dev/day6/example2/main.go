package main

import "fmt"

type reader interface {
	read()
}

type writer interface {
	write()
}

type readAndWrite interface {
	reader
	writer
}

type file struct{}

func (p *file) read() {
	fmt.Println("read data")
}
func (p *file) write() {
	fmt.Println("write data")
}
func test(rw readAndWrite) {
	rw.read()
	rw.write()
}
func main() {
	var f file
	test(&f)
}
