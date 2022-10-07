package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for i := range cpdomains {
		s := cpdomains[i]
		index := strings.Index(s, " ")
		count, _ := strconv.Atoi(s[:index])
		s = s[index+1:]
		m[s] += count
		for {
			i2 := strings.Index(s, ".")
			if i2 == -1 {
				break
			}
			s = s[i2+1:]
			m[s] += count
		}
	}
	ans := []string{}
	for k, v := range m {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return ans
}
