package main

import "fmt"

func ReverseString(str string) string {
	rStr := []rune(str)
	for i, j := 0, len(rStr)-1; i != len(rStr)/2; i, j = i+1, j-1 {
		rStr[i], rStr[j] = rStr[j], rStr[i]
	}
	return string(rStr)
}

func main() {
	var test1 string = "abcdfg" // четное кол-во символов
	fmt.Printf("String before reversing: %s\n", test1)
	test1 = ReverseString(test1)
	fmt.Printf("String after reversing: %s\n", test1)
	test2 := "aaaaaaaaaaaa"
	fmt.Printf("String before reversing: %s\n", test2)
	test2 = ReverseString(test2)
	fmt.Printf("String after reversing: %s\n", test2)
	test3 := "qwertyu" // нечетное кол-во символов
	fmt.Printf("String before reversing: %s\n", test3)
	test3 = ReverseString(test3)
	fmt.Printf("String after reversing: %s\n", test3)

}
