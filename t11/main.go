package main

import (
	"fmt"
)

type IntSet []int

func main() {
	// первый способ, если множество - это тип map. Но это конечно не так, поэтому см. способ ниже
	firstSet := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	secondSet := map[int]int{
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
	}

	intersectionSet := make(map[int]int)
	isFirstLessSecond := len(firstSet) < len(secondSet)
	if isFirstLessSecond {
		for key, valFirst := range firstSet {
			if _, ok := secondSet[key]; ok {
				intersectionSet[key] = valFirst // предполагается, что значения с одинаковыми ключами совпадают
			}
		}
	} else {
		for key, valFirst := range secondSet {
			if _, ok := firstSet[key]; ok {
				intersectionSet[key] = valFirst // предполагается, что значения с одинаковыми ключами совпадают
			}
		}
	}
	fmt.Println(intersectionSet)
	fmt.Println("----------------------------------------------")

	intSet1 := IntSet{1, 2, 3, 4, 5, 6, 7}   // множество уникальных значений
	intSet2 := IntSet{5, 6, 7, 8, 9, 10, 11} // множество уникальных значений
	anotherIntersectionSet := make(IntSet, 0)
	for _, v := range intSet1 {
		if ok := intSet2.Contains(v); ok {
			anotherIntersectionSet = append(anotherIntersectionSet, v)
		}
	}
	fmt.Println(anotherIntersectionSet)
}

func (set *IntSet) Contains(value int) bool {
	for _, v := range *set {
		if v == value {
			return true
		}
	}
	return false
}
