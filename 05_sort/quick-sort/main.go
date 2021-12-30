package main

import "fmt"

func partition(a []int, lo, hi int) int {
	i := lo
	j := hi
	for i <= j {
		for a[i] < a[lo] {
			i++
		}
		for a[lo] < a[j] {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	arr := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
