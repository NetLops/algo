package main

func isUpperLetter(c byte) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}
	return false
}
func isValid(code string) bool {
	// 针对 cdata 字符串，进行优先处理
	cdataPrefix := "<![CDATA["
	cdataSuffix := "]]>"

	n := len(code)
	stack := []string{}

	i := 0

	for i < n {
		if i+8 < n && code[i:i+9] == cdataPrefix {
			// 如果首位就是，因为此时还没有TAG_NAME，所以直接返回
			if i == 0 {
				return false
			}
			j := i + 9
			ok := false // 代表找到了
			for j < n && !ok {
				// 再找找cdata2
				if j+2 < n && code[j:j+3] == cdataSuffix {
					j += 3
					ok = true
				} else {
					j++
				}
			}
			if !ok {
				// 找到cdataPrefix,找不到cdataSuffix 直接抛出错误
				return false
			}

			i = j // 找到了，直接解析后面的
		} else if code[i] == '<' { // 看看是否是 TAG_NAME_PREFIX
			if i == n-1 {
				return false
			}
			isEnd := code[i+1] == '/' // 判断是不是结尾标签
			next := i + 1             // 不是结尾标签，就移一个位置
			if isEnd {
				next = i + 2 // 是结尾标签，那就移两个位置（相当于i 而言）
			}
			j := next
			for j < n && code[j] != '>' {
				// 标签内的字母不是大写字母，返回false
				if !isUpperLetter(code[j]) {
					return false
				}
				j++
			}
			if j == n {
				return false
			}

			length := j - next
			if length < 1 || length > 9 {
				// 标签中的字符长度小于1或大于9，返回false
				return false
			}
			tagContent := code[next:j] // <HTML> tagContent = "HTML" </HTML> tagContent = "HTML"
			// i 指向 '>' 后的第一个字符
			i = j + 1
			if !isEnd {
				// 不是isEnd  直接加进去
				stack = append(stack, tagContent)
			} else {
				// 如果堆栈为空，或者 标签不匹配，返回false
				if len(stack) == 0 {
					return false
				}
				x := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if x != tagContent {
					return false
				}
				// 标签已经匹配完，但i 未表里到code 尾部，就是后面还有东西 <A> </A> <B> </B> 这种
				if len(stack) == 0 && i < n {
					return false
				}
			}
		} else {
			if i == 0 { // 开头都不是 '<' 直接无
				return false
			}
			i++ // 继续往下遍历
		}
	}
	return len(stack) == 0
}
func main() {

}
