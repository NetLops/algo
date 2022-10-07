package main

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
func minAddToMakeValid(s string) int {
	q := []byte{}
	for k := 0; k < len(s); k++ {
		if s[k] == '(' {
			q = append(q, '(')
		} else if s[k] == ')' {
			if len(q) > 0 && q[len(q)-1] == '(' {
				q = q[:len(q)-1]
			} else {
				q = append(q, ')')
			}
		}
	}
	return len(q)
}
