// Write a program to store random numbers in an array and find lower and highest values. Also, store value in an array where value will be sorted
package main

import (
	"fmt"
	"sort"
)

func main() {

	var a []int
	var size int
	fmt.Println("Enter array size")
	fmt.Scan(&size)

	a = make([]int, size)
	fmt.Println("Enter array")
	for i := 0; i < size; i++ {
		fmt.Scan(&a[i])
	}
	max := a[0]
	for i := 0; i < size; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	min := a[0]
	for i := 0; i < size; i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	fmt.Println("max is", max, "min is", min)

	// var temp int

	// Sort
	// for i := 0; i < size; i++ {
	// 	for j := 0; j < size; j++ {
	// 		if a[i] < a[j] {
	// 			temp = a[i]
	// 			a[i] = a[j]
	// 			a[j] = temp
	// 		}

	// 	}
	// }
	// fmt.Println("Sorted array is")
	// for i := 0; i < size; i++ {
	// 	fmt.Print(a[i], " ")
	// }

	sort.Ints(a)

	fmt.Println("sort is", a)

}
