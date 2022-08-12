package main

import "fmt"

func reverseWords(s string) string {
	temp := ""
	start := 0
	for i := range s {
		if s[i] == ' ' {
			if i > 0 && s[i-1] != ' ' {
				temp = s[start:i] + " " + temp
			}
			continue
		} else {
			if i > 0 && s[i-1] == ' ' {
				start = i
			}
			if i == len(s)-1 {
				temp = s[start:i+1] + " " + temp
			}
		}

	}
	return temp
}
func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
	t := 23223 / 10
	fmt.Println(t)
}
