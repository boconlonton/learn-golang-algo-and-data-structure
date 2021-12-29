package _1_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	BubbleSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("false order.")
		}
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	SelectionSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("false order.")
		}
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	InsertionSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("false order.")
		}
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	res := MergeSort(arr)
	fmt.Println(res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			t.Errorf("false order.")
		}
	}
}
