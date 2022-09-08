package main

import "fmt"

type person interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	salary float64
}
type student struct {
	human
	marks float64
}
type printer struct {
}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	h := human{"John"}
	s := student{human: human{name: "Bob"}, marks: 90.0}
	t := teacher{human: human{name: "Prof Smith"}, salary: 50000.0}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
