package main

type WordFilter struct {
	m map[string]int
}

func Constructor(words []string) WordFilter {
	m := map[string]int{}
	for i, word := range words {
		for j, n := 1, len(word); j <= n; j++ {
			for k := 1; k <= n; k++ {
				m[word[:j]+"_"+word[n-k:]] = i
			}
		}
	}
	return WordFilter{
		m: m,
	}
}

func (this *WordFilter) F(pref string, suff string) int {
	if v, ok := this.m[pref+"_"+suff]; ok {
		return v
	} else {
		return -1
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
func main() {

}
