package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	for k := 0; k < len(t); k++ {
		if s[i] == t[k] {
			i++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}
func main() {

}
