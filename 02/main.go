package main

import "fmt"

func main() {
	ms := map[string]int{
		"Sirop": 1,
		"Kaffe": 2,
		"Kachu": 3,
	}
	// print out the map
	fmt.Println(ms)
	// range over the map printing out just the key
	for k := range ms {
		fmt.Println(k)
	}

	// range over the map printing out both the key and the value
	for k, v := range ms {
		fmt.Printf("%v: %v\n", k, v)
	}
}
