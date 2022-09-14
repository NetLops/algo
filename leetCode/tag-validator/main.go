package main

import "fmt"

func isUpperLetter(c byte) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}
	return false
}

func isValid(code string) bool {
	// 以下两个 针对 cdata 字符串 进行优先处理
	cdataPrefix := "<![CDATA[" // 9个字符
	cdataSuffix := "]]>"       // 3个

	n := len(code)
	stack := []string{}

	i := 0
	for i < n {
		// 找一手 cdata
		if i+8 < n && code[i:i+9] == cdataPrefix {
			// 首位就是 那直接返回错，因为此时没有 TAG_NAME
			if i == 0 {
				return false
			}
			j := i + 9
			ok := false // 代表找到了
			for j < n && !ok {
				// 找一手 cdata2
				if j+2 < n && code[j:j+3] == cdataSuffix {
					j += 3
					ok = true
				} else {
					j++
				}
			}
			if !ok {
				// 找到cdataPrefixx, 找不到cdataSuffix，直接抛出错误
				return false
			}
			i = j // 找到了 那就 直接 解析后面的
		} else if code[i] == '<' { // 老实人做法了， 看看是否是 TAG_NAME_PREFIX
			if i == n-1 { // 到头就G了，面向实例编程
				return false
			}
			isEnd := code[i+1] == '/' // 最好不要是结尾标签
			next := i + 1             // 不是结尾标签 那就移一个位置
			if isEnd {
				next = i + 2 //  是结尾标签，那就移动两个位置（相对于 i而言）
			}
			j := next
			for j < n && code[j] != '>' {
				// 标签内的字母不是大写字母，返回false
				if !isUpperLetter(code[j]) {
					return false
				}
				j++
			}
			if j == n { // 面向 用例编程
				// 到头了，返回false 没找到 '>'
				return false
			}
			length := j - next
			if length < 1 || length > 9 {
				// 标签中的字符长度小于1或者大于9，返回false
				return false
			}
			tagContent := code[next:j] // <HTML>  tagContent = "HTML" </HTML> tagContent = "HTML"
			// i 指向 '>' 后的第一个字符（直到头那就无了）
			i = j + 1
			if !isEnd {
				//  不是isEnd 就好说
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
				// 标签已经匹配完，但i未遍历到code尾部,就是后面还有东西 比如 <A> </A> <B></B> 这种 恶心情况
				if len(stack) == 0 && i < n {
					return false
				}
			}
		} else {
			if i == 0 { // 开头都不是 '<' 直接G
				return false
			}
			i++ // 其他分类 就直接 往下继续遍历
		}
	}
	return len(stack) == 0 // 如果 还存有 那就。。。。没有结尾了，面向实例编程 害
}

func main() {
	//cdataPrefix := "<![CDATA["
	//
	//temp := "345345<![CDATA[s9"
	//
	//fmt.Println(string(temp[6 : 6+9]))
	//fmt.Println(cdataPrefix == temp[6:6+9])
	//isValid("<A>  <B> </A>   </B>")
	//fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
	//fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]><DIV>"))
	fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
}
