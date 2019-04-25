package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Id   string
	Age  int
}

type Book struct {
	Name  string
	Count int
}
type BookArray []Book
type StudentArray []Student

func (p BookArray) Len() int {
	return len(p)
}
func (p BookArray) Less(i, j int) bool {
	return p[i].Count > p[j].Count
}
func (p BookArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p StudentArray) Len() int {
	return len(p)
}

func (p StudentArray) Less(i, j int) bool {
	return p[i].Name > p[j].Name
}

func (p StudentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var books BookArray
	for i := 0; i < 10; i++ {
		book := Book{
			Name:  fmt.Sprintf("book%d", i),
			Count: i + 1,
		}
		books = append(books, book)
	}

	sort.Sort(books)
	for _, v := range books {
		fmt.Println(v)
	}
}
