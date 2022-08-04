package main

import "fmt"

func replaceSpace(s string) string {
	temp := []byte{}
	for i := range s {
		if s[i] != ' ' {
			temp = append(temp, s[i])
		} else {
			temp = append(temp, []byte("%20")...)
		}

	}
	return string(temp)
}
func main() {
	fmt.Println(replaceSpace("We are happy."))
}
