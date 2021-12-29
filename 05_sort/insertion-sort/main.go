package main

import "fmt"

func InsertionSort(p *[]int) {
	// Ascending order
	xi := *p
	for i := 0; i < len(xi); i++ {
		for j := i; j > 0; j-- {
			if xi[j] < xi[j-1] {
				xi[j], xi[j-1] = xi[j-1], xi[j]
			}
		}
	}
}

func main() {
	xi := []int{1, 2, 3, 4, 6, 7}
	InsertionSort(&xi)
	fmt.Println(xi)
}
