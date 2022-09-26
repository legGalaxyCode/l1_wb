package main

import (
	"fmt"
	"strings"
)

// IsStringUniqueMap первый способ используя map для любых символов
func IsStringUniqueMap(str string) bool {
	str = strings.ToLower(str)
	runeStr := []rune(str)
	mp := make(map[rune]bool)
	for _, char := range runeStr {
		if _, ok := mp[char]; ok {
			return false
		} else {
			mp[char] = true
		}
	}
	return true
}

// IsStringUnique второй способ используя битовые операции только для символов a-z
func IsStringUnique(str string) bool {
	str = strings.ToLower(str)
	runeStr := []rune(str)
	mask := 0
	for _, char := range runeStr {
		value := char - 'a' // потому что вот
		if mask&(1<<value) > 0 {
			return false
		}
		mask = mask | (1 << value)
	}
	return true
}

func main() {
	test1 := "abcd"
	test2 := "abCdefAaf"
	test3 := "aabcd"
	fmt.Printf("Is %s unique? - %t\n", test1, IsStringUnique(test1))
	fmt.Printf("Is %s unique? - %t\n", test2, IsStringUnique(test2))
	fmt.Printf("Is %s unique? - %t\n", test3, IsStringUnique(test3))
	fmt.Println("--------Using map---------------")
	fmt.Printf("Is %s unique? - %t\n", test1, IsStringUniqueMap(test1))
	fmt.Printf("Is %s unique? - %t\n", test2, IsStringUniqueMap(test2))
	fmt.Printf("Is %s unique? - %t\n", test3, IsStringUniqueMap(test3))
}
