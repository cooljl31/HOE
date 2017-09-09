package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	forWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln(s.color, "is the best color of Sedan model")
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln(t.color, "is elegance for a Truck")
}

func main() {
	t1 := truck{
		vehicle{
			2,
			"red",
		},
		true,
	}
	s1 := sedan{vehicle{4, "blue"}, true}
	fmt.Println(t1)
	fmt.Println(s1)
	// Truck descriptions
	fmt.Println(t1.transportationDevice())
	fmt.Println(t1.vehicle.doors)
	fmt.Println(t1.vehicle.color)
	fmt.Println(t1.forWheel)
	// Sedan descriptions
	fmt.Println(s1.vehicle.doors)
	fmt.Println(s1.vehicle.color)
	fmt.Println(s1.luxury)
	fmt.Println(s1.transportationDevice())

}
