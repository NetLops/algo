package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	props := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	ret := 0
	for _, item := range items {
		_type, color, name := item[0], item[1], item[2]
		if (_type == ruleValue && props[ruleKey] == 0) ||
			(color == ruleValue && props[ruleKey] == 1) ||
			(name == ruleValue && props[ruleKey] == 2) {
			ret++
		}
	}
	return ret
}
