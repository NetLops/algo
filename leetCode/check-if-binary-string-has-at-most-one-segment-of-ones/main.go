package main

func checkOnesSegment(s string) bool {
	left, right := 0, len(s)-1
	leftEnd, rightEnd := false, false
	for left <= right {
		if leftEnd && rightEnd {
			break
		}
		if left-1 >= 0 && s[left] == s[left-1] {
			left++
		} else {
			leftEnd = true
		}
		if right+1 <= len(s)-1 && s[right] == s[right+1] {
			right--
		} else {
			rightEnd = true
		}
		if left == 0 {
			left++
		}
		if right == len(s)-1 {
			right--
		}
	}

	if leftEnd && rightEnd {
		for i := left; i <= right; i++ {
			if s[i] != s[right-(left-i)] {
				return false
			}
		}
	}

	if !leftEnd && !rightEnd {
		return true
	}

	if left >= right {
		return true
	}
	return false
}
