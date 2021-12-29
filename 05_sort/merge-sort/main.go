package main

import "fmt"

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	// Copy the rest to the final array
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	// Copy the rest to the final array
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func main() {
	data := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	sorted := mergeSort(data)
	fmt.Println(sorted)
}
