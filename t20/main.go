package main

import (
	"fmt"
	"strings"
)

func ReverseSentence(sent string) string {
	words := strings.Split(sent, " ")
	wordsLen := len(words)
	result := ""
	for i, j := 0, wordsLen-1; i != j && i != wordsLen/2; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	for idx, word := range words {
		result += word
		if idx != wordsLen {
			result += " "
		}
	}

	return result
}

func main() {
	test := "snow dog son"
	fmt.Printf("Sentence before reversing: %s\n", test)
	test = ReverseSentence(test)
	fmt.Printf("Sentence before reversing: %s\n", test)
	test1 := "alt comma separate mouse"
	fmt.Printf("Sentence before reversing: %s\n", test1)
	test1 = ReverseSentence(test1)
	fmt.Printf("Sentence before reversing: %s\n", test1)
}
