package main

import "fmt"

func typeOf(any interface{}) string {
	switch any.(type) {
	case int:
		return "int"
	case int64:
		return "int64"
	case string:
		return "string"
	case bool:
		return "bool"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case chan int:
		return "chan int"
	case chan float32:
		return "chan float32"
	case chan float64:
		return "chan float64"
	default:
		return "unknown type"
		// to be continued
	}
}

func main() {
	fmt.Println(typeOf("string"))
	fmt.Println(typeOf(3.14))
	fmt.Println(typeOf(true))
	fmt.Println(typeOf(make(chan int)))
}
