package main

import "fmt"

type person struct {
	fName string
	lName string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking.")
}

func main() {
	p1 := person{
		"Sirop",
		"Lesperance",
	}

	sp := p1.walk()
	// print out struct person
	fmt.Println(sp)

}
