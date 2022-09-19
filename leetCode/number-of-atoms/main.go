package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isNum(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func isUpperLetter(c byte) bool {
	if 'A' <= c && c <= 'Z' {
		return true
	}
	return false
}

func isLowerLetter(c byte) bool {
	if 'a' <= c && c <= 'z' {
		return true
	}
	return false
}
func countOfAtoms(formula string) string {
	elemMap := map[string]int{} // 计数小子
	stackElemMap := []map[string]int{}
	//stack := [][]string{}
	index := -1 // 栈深度
	i := 0
	n := len(formula)
	//  数字就在 if - else if - else 里面处理，嘻嘻
	for i < n {
		if formula[i] == '(' {
			index++
			if len(stackElemMap)-1 <= index {
				//stack = append(stack, []string{})
				stackElemMap = append(stackElemMap, map[string]int{})
			}
			i++
		} else if formula[i] == ')' {
			// 看一手下面还有没有数字/字母
			j := i + 1
			if j >= n {
				break
			}
			numBytes := []byte{}
			for j < n && isNum(formula[j]) {
				numBytes = append(numBytes, formula[j])
				j++
			}
			count := 1
			if len(numBytes) != 0 {
				count, _ = strconv.Atoi(string(numBytes))
			}
			// 分层套娃，每层递进
			//fmt.Println("before:", stack, elemMap, stackElemMap, count)
			for elem, _ := range stackElemMap[index] {
				if index > 0 {
					if _, ok := stackElemMap[index-1][elem]; ok {
						stackElemMap[index-1][elem] += stackElemMap[index][elem] * count
					} else {
						stackElemMap[index-1][elem] = stackElemMap[index][elem] * count
					}

				} else {
					elemMap[elem] += stackElemMap[index][elem] * count
				}
			}

			//for _, elem := range stack[index] {
			//	if index > 0 {
			//		if _, ok := stackElemMap[index-1][elem]; ok {
			//			stackElemMap[index-1][elem] += stackElemMap[index][elem] * count
			//		} else {
			//			stackElemMap[index-1][elem] = stackElemMap[index][elem] * count
			//		}
			//
			//	} else {
			//		elemMap[elem] += stackElemMap[index][elem] * count
			//	}
			//
			//}
			//fmt.Println("after:", stack, elemMap, stackElemMap)
			i = j // 另类i++
			index--
			//if index >= 0 {
			//	stack[index] = append(stack[index], stack[index+1]...)
			//}
			// 让gc 处理， 有一种情况是Mg(Mg(OH)2)30Mg(OH)2  保证重置好
			//stack[index+1] = nil
			//stack = stack[:index+1]
			stackElemMap = stackElemMap[:index+1]
		} else {
			// 一定是 大写字母， 要不然就不符合题意了
			if i >= n {
				break
			}
			// 处理那些字母 数字了（害）
			j := i + 1
			if j >= n {
				elemMap[formula[i:j]]++
				break
			}
			if isLowerLetter(formula[j]) {
				j++
			}
			elem := formula[i:j]
			if j >= n {
				elemMap[elem]++
				break
			}
			numBytes := []byte{}
			for j < n && isNum(formula[j]) {
				numBytes = append(numBytes, formula[j])
				j++
			}
			count := 1
			if len(numBytes) != 0 {
				count, _ = strconv.Atoi(string(numBytes))
			}
			//elemMap[elem] += count

			if index >= 0 {
				//stack[index] = append(stack[index], elem)
				stackElemMap[index][elem] += count
			} else {
				elemMap[elem] += count
			}
			i = j //另类i++
		}
	}
	if len(stackElemMap) != 0 {
		elemMap = stackElemMap[0]
	}

	//fmt.Println(elemMap)
	allElem := []string{}
	for k, _ := range elemMap {
		allElem = append(allElem, k)
	}
	sort.Strings(allElem)
	ans := ""
	for _, elem := range allElem {
		ans += elem
		if elemMap[elem] > 1 {
			ans += strconv.Itoa(elemMap[elem])
		}
	}
	return ans

}
func main() {
	//fmt.Println(countOfAtoms("Mg(Mg(OH)2)30Mg(OH)2"))
	//fmt.Println(countOfAtoms("Mg(OH)2"))
	//fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
	fmt.Println(countOfAtoms("(B2O39He17BeBe49)39"))
	fmt.Println(countOfAtoms("(((H)))"))
}
