package main

import (
	"fmt"
	"sort"
)

type Number interface {
	int16 | int32 | int | int64 | float32 | float64
}

func partition[T Number](array []T, low int, high int) ([]T, int) {
	pivot := array[high]
	idx := low
	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[idx], array[j] = array[j], array[idx]
			idx++
		}
	}
	array[idx], array[high] = array[high], array[idx]
	return array, idx
}

func QuickSort[T Number](array []T, low int, high int) []T {
	if low < high {
		array, pivotIdx := partition(array, low, high)
		fmt.Println(array, pivotIdx)
		array = QuickSort(array, low, pivotIdx-1)
		array = QuickSort(array, pivotIdx+1, high)
	}
	return array
}

func main() {
	testArr1 := []int{982, 782, 902, 54, 888, 1, -5, 89, 692}
	testArr2 := []float64{-4.56, 9.0123, 6.8992, 891.56, 12.78, 0.56, -56.89213}
	fmt.Println(testArr1)
	fmt.Println(testArr2)
	fmt.Println("After sorting:")
	testArr1 = QuickSort(testArr1, 0, len(testArr1)-1)
	fmt.Println(testArr1, sort.SliceIsSorted(testArr1, func(i, j int) bool {
		return i < j
	}))
	testArr2 = QuickSort(testArr2, 0, len(testArr2)-1)
	fmt.Println(testArr2, sort.SliceIsSorted(testArr2, func(i, j int) bool {
		return i < j
	}))
}
