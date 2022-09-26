package main

import (
	"fmt"
	"sort"
)

type Number interface {
	int16 | int32 | int | int64 | float32 | float64
}

// SliceBinarySearch return -1 if value doesn't find or index otherwise
func SliceBinarySearch[T Number](slice []T, value T) (int, bool) {
	if ok := sort.SliceIsSorted(slice, func(i, j int) bool {
		return i < j
	}); !ok {
		return -1, false
	}
	start, end := 0, len(slice)-1
	for start <= end {
		middle := start + (end-start)/2
		middleValue := slice[middle]
		if middleValue == value {
			return middle, true
		} else if middleValue < value {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1, false
}

func main() {
	testSliceSorted := []int{-5, 1, 54, 89, 692, 782, 888, 902, 982}
	testSliceUnsorted := []float64{-4.56, 9.0123, 6.8992, 891.56, 12.78, 0.56, -56.89213}
	test1Idx, test1Find := SliceBinarySearch(testSliceSorted, 89)
	fmt.Println("Slice1: ", testSliceSorted)
	fmt.Println("Slice2: ", testSliceUnsorted)
	fmt.Printf("Binary search value 89 on slice1: %d %t\n", test1Idx, test1Find)
	test1Idx, test1Find = SliceBinarySearch(testSliceSorted, 1000)
	fmt.Printf("Binary search value 1000 on slice1: %d %t\n", test1Idx, test1Find)
	test2Idx, test2Find := SliceBinarySearch(testSliceUnsorted, 0.56)
	fmt.Printf("Binary search value 0.56 on slice2: %d %t\n", test2Idx, test2Find)
}
