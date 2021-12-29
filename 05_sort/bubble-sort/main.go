package main

import "fmt"

func BubbleSort(px *[]int) {
	// Order: Ascending
	xi := *px
	n := len(xi)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}
}

func main() {
	xi := []int{5, 2, 1, 5, 7, 9}
	BubbleSort(&xi)
	fmt.Println(xi)
}
