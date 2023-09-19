package core

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	hasTits bool
}

func Run() {
	fmt.Println("QMS Integration (Go) is good to Go")
	p := Person{}

	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.hasTits)
}
