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

type transportation interface {
	transportationDevice() string
}

func report(trans transportation) {
	fmt.Println(trans.transportationDevice())
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
	report(t1)
	report(s1)

}
