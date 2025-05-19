package main

import "fmt"

func showMapExample() {
	myMap := map[string]int {
		"apple": 5,
		"banana": 10,
	}

	fmt.Println("Fruit counts:")
	for fruit, count := range myMap {
		fmt.Printf("%s: %d\n", fruit, count)
	}
}