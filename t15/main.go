package main

var justString string

/*
основная проблема в том, что после такого присваивания
в памяти останется весь кусок от v
*/
func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]           // instead of this

	newJustString := make([]byte, 0, 100)
	newJustString = append(newJustString, v[:100]...) // we can use this (or copy)
}

func main() {
	someFunc() // incorrect
}

// for compiler
func createHugeString(size int) string {
	return ""
}
