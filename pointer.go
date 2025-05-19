package main

import "fmt"

//Without pointers

type person struct {
	name string
	age int
}

// func updateAge(p person) {
// 	p.age = 30
// }

// func update() {
// 	p := person{name: "Alice", age: 25}
// 	updateAge(p)
// 	fmt.Println(p.age)
// }

//With pointers

func updateAge(p *person) {
	p.age = 30
}

func update() {
	p := person{name: "Alice", age: 25}
	updateAge(&p)
	fmt.Println(p.age)
}