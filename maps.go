package main

import (
	"fmt"
	"math/rand"
)

func maps() {
	fmt.Println("define map")
	products := make(map[string]int)
	customers := make(map[string]int)

	fmt.Println(">>>>>>insert map data")
	products["product1"] = rand.Intn(100)
	products["product2"] = rand.Intn(100)

	customers["cust1"] = rand.Intn(100)
	customers["cust2"] = rand.Intn(100)

	fmt.Println(">>>>>display map data")
	fmt.Println(products["product1"])
	fmt.Println(products["product2"])
	fmt.Println(customers["cust1"])
	fmt.Println(customers["cust2"])
}