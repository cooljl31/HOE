package main

import "fmt"

func main() {
	is := []int{1, 34, 5}
	// print out the slice
	fmt.Println(is)

	// range over the slice printing out just the index
	for i := range is {
		fmt.Printf("%v\n", i)
	}

	// range over the slice printing out both the index and the value
	for i, v := range is {
		fmt.Printf("%v - %v\n", i, v)
	}
}
