package main

import (
	"fmt"
)

func matrix() {
	var matrix [3][3]int
	//var nums [5]int
	//nums := []int{10, 20, 30, 40, 50}

	//nums := [4]int{10, 20, 30, 40}
	// for i, v := range nums {
		//fmt.Printf("index: %d, Value: %d \n", i, v)
	//}


	// If only the value
	// for _, v := range nums {
		//fmt.Println(v)
	//}

	//Using range on maps
	// students := map[string]int {
	// 	"Alice": 21,
	// 	"Bob": 19,
	// 	"Eve": 22,
	// }

	// for name, age := range students {
	// 	fmt.Printf("Name: %s, Age: %d", name, age)
	// }

	//Using range on Strings
	// text := "Go!"
	// for i, ch := range text {
	// 	fmt.Printf("Index: %d, Character: %c", i, ch)
	// }

	fmt.Println("Enter elements for a 3x3 matrix:")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter value for matrix[%d][%d]: ", i, j)
			_, err := fmt.Scan(&matrix[i][j])

			if err != nil {
				fmt.Println("Error reading input: ", err)
			}
		}
	}

	fmt.Println("The matrix is: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}