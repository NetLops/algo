package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	collect := make([]byte, len(digits))
	var dfs func(index int)
	collects := []string{}
	dfs = func(index int) {
		if index == len(digits) {
			collects = append(collects, string(collect[:]))
			return
		}
		for _, v := range m[digits[index]] {
			collect[index] = byte(v)
			dfs(index + 1)
		}
	}
	dfs(0)
	return collects
}
