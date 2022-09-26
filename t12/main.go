package main

import "fmt"

// SliceSet реализация через слайс
type SliceSet[T comparable] struct {
	buf []T
}

func (set *SliceSet[T]) Contains(value T) bool {
	for _, val := range set.buf {
		if val == value {
			return true
		}
	}
	return false
}

func (set *SliceSet[T]) Add(value T) {
	if set.buf == nil {
		set.buf = make([]T, 0, 10)
	}
	if ok := set.Contains(value); !ok {
		set.buf = append(set.buf, value)
	}
}

func main() {
	test := []string{"cat", "cat", "dog", "cat", "tree"}
	set := SliceSet[string]{}
	for _, val := range test {
		set.Add(val)
	}
	fmt.Println(set)
}
