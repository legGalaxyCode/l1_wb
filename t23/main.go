package main

import "fmt"

func SliceDeleteByIndex[T any](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx >= len(slice) {
		return slice, fmt.Errorf("index out of bounds")
	}
	// тот же самый слайс, но пропуская элемент slice[idx]
	return append(slice[:idx], slice[idx+1:]...), nil
}

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Slice len: %d and cap: %d\n", len(test), cap(test))
	fmt.Println("Slice: ", test)
	test, _ = SliceDeleteByIndex(test, 0)
	fmt.Println("Slice without 0 element: ", test)
	fmt.Printf("Slice len: %d and cap: %d\n", len(test), cap(test))
}
