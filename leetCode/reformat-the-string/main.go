package main

import "fmt"

func reformat(s string) string {
	str := make([]byte, len(s))
	letter := []byte{}
	digit := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit = append(digit, s[i])
		} else if s[i] >= 'a' && s[i] <= 'z' {

			letter = append(letter, s[i])
		}
	}
	if len(digit)-len(letter) >= -1 && len(digit)-len(letter) <= 1 {
		//i, j := 0, 0
		if len(digit) > len(letter) {
			for i2 := range digit {
				str[i2*2] = digit[i2]
			}
			for i := range letter {
				str[2*i+1] = letter[i]
			}
		} else {

			for i := range letter {
				str[2*i] = letter[i]
			}
			for i2 := range digit {
				str[i2*2+1] = digit[i2]
			}
		}

	} else {
		return ""
	}
	//if len(digits) != letters
	return string(str)
}
func main() {
	fmt.Println(reformat("covid2019"))
}
