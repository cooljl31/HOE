package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favfood []string
}

func main() {
	p1 := person{
		fName:   "Sirop",
		lName:   "Lesperance",
		favfood: []string{"Katofel", "currywurst", "ketchup"},
	}

	// print out the p1
	fmt.Println(p1)

	// print out the values in favfood
	fmt.Println(p1.favfood)

	// print out the values by using a for "favfood" range loop
	for _, v := range p1.favfood {
		fmt.Println(v)
	}
}
