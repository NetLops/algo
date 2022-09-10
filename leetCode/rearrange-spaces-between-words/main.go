package main

import "fmt"

func reorderSpaces(text string) string {
	words := [][]byte{}
	spaceCount := 0
	bytes := []byte(text)
	left := 0
	for left < len(bytes) {
		if bytes[left] == ' ' {
			spaceCount++
		} else {
			word := []byte{}
			for left < len(bytes) && bytes[left] != ' ' {
				word = append(word, bytes[left])
				left++
			}
			words = append(words, word)
			continue
		}
		left++
	}
	temp := spaceCount
	if len(words) > 1 {
		spaceCount = temp / (len(words) - 1)

	}
	manyCount := temp - spaceCount*(len(words)-1)
	spaceWord := make([]byte, spaceCount)
	for i := range spaceWord {
		spaceWord[i] = ' '
	}
	bytes = []byte{}
	for i := 0; i < len(words)-1; i++ {
		bytes = append(bytes, words[i]...)
		bytes = append(bytes, spaceWord...)
	}
	bytes = append(bytes, words[len(words)-1]...)
	for i := 0; i < manyCount; i++ {
		bytes = append(bytes, ' ')
	}
	return string(bytes)
}
func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
}
