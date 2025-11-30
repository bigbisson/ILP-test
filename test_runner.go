package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Task1 examples
	fmt.Println("Task1:")
	fmt.Println("GO1([-1,1]) =>", GO1([]int{-1, 1}))
	fmt.Println("GO1([-1,-7,-5]) =>", GO1([]int{-1, -7, -5}))
	fmt.Println("GO1([1,2,1,6]) =>", GO1([]int{1, 2, 1, 6}))
	fmt.Println()

	// Task2 examples
	fmt.Println("Task2:")
	v, err := GO2([]int{0, 1, 0}, 1000, 1010)
	fmt.Printf("GO2([0,1,0], 1000,1010) => %d, %v\n", v, err)
	v2, err2 := GO2([]int{0, 1, 0}, 1259, 2601)
	fmt.Printf("GO2([0,1,0], 1259,2601) => %d, %v\n", v2, err2)
	fmt.Println()

	// Task3 examples
	fmt.Println("Task3:")
	fmt.Printf("GO3([1,2],[1,3]) => %v\n", GO3([]int{1, 2}, []int{1, 3}))
	fmt.Printf("GO3([1,2,2],[1,2,4]) => %v\n", GO3([]int{1, 2, 2}, []int{1, 2, 4}))
	fmt.Println()

	// Task4 examples
	fmt.Println("Task4:")
	m := GO4(1223334)
	// print map in stable order
	b, _ := json.Marshal(m)
	fmt.Println("GO4(1223334) =>", string(b))
	fmt.Println()

	// Task5 examples
	fmt.Println("Task5:")
	fmt.Println("GO5(1,2,3,4) =>", GO5(1, 2, 3, 4))
	fmt.Println("GO5(1,2,3,4,5) =>", GO5(1, 2, 3, 4, 5))
}
