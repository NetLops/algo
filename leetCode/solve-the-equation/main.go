package main

import (
	"fmt"
	"time"
)

const (
	DIGIT = 1 << iota
	LETTER
	X
)

func validateDigit(letter byte) bool {
	if letter >= '0' && letter <= '9' {
		return true
	}

	return false
}
func solveEquation(equation string) string {
	ans := 0
	factor, val := 0, 0
	i, sign := 0, true // sign 代表为正
	temp := 0
	toEqual := false
	for i < len(equation) {
		if equation[i] == '=' {
			toEqual = true
			i++
			sign = false
			continue
		}
		// 左边表达式
		if !toEqual {
			//  数字拼接 判断是否是正常数字的还是x的系数
			if validateDigit(equation[i]) {
				temp = int(equation[i] - '0')
				i++
				for {
					if validateDigit(equation[i]) {
						temp = temp*10 + int(equation[i]-'0')
						i++
					} else {
						break
					}
				}
				if equation[i] == 'x' {
					// sign 为true 代表+
					if sign {
						factor += temp
					} else {
						factor -= temp
					}
					i++
				} else {
					if sign {
						val += temp
					} else {
						val -= temp
					}
				}

			} else if equation[i] == '+' {
				sign = true
				i++
			} else if equation[i] == '-' {
				sign = false
				i++
			} else {
				if sign { // 单个x
					factor += 1
				} else {
					factor -= 1
				}
				i++
			}

		} else { // 右边表达式
			//  数字拼接 判断是否是正常数字的还是x的系数
			if validateDigit(equation[i]) {
				temp = int(equation[i] - '0')
				i++
				for i < len(equation) {
					if validateDigit(equation[i]) {
						temp = temp*10 + int(equation[i]-'0')
						i++
					} else {
						break
					}
				}
				if i >= len(equation) {
					if sign {
						val += temp
					} else {
						val -= temp
					}
					break
				}
				if equation[i] == 'x' {
					// sign 为true 代表+
					if sign {
						factor += temp
					} else {
						factor -= temp
					}
					i++
				} else {
					if sign {
						val += temp
					} else {
						val -= temp
					}
				}

			} else if equation[i] == '+' {
				sign = false
				i++
			} else if equation[i] == '-' {
				sign = true
				i++
			} else {
				if sign { // 单个x
					factor += 1
				} else {
					factor -= 1
				}
				i++
			}
		}

	}
	if factor != 0 && val == 0 {
		return "Infinite solutions"
	}
	if factor == 0 && val != 0 {
		return "No solution"
	}
	//fmt.Println(factor, val)
	if factor > 0 {
		ans = -val / factor
	} else if factor < 0 {
		ans = val / factor
	}
	return fmt.Sprintf("x=%d", ans)
}
func main() {
	//fmt.Println(DIGIT)
	//fmt.Println(LETTER)

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("又过去了一秒")
		}
		fmt.Println("确实")
	}
}
