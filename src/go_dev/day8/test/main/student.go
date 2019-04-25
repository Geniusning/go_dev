package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	name string
	age  int
}

func (p *Student) Save() (err error) {
	data, err := json.Marshal(p)
	fmt.Println(data)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("D:/test1.dat", data, 0633)
	return
}

func (p *Student) Load() (err error) {
	data, err := ioutil.ReadFile("D:/test1.dat")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, p)
	return
}
