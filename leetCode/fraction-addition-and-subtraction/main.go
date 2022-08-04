package main

import (
	"strconv"
)

func maxGCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minHCF(a, b, maxGcd int) int {
	if maxGcd == 0 {
		return a * b / maxGCD(a, b)
	}
	return a * b / maxGcd
}

func fractionAddition(expression string) string {
	length := len(expression)
	numerators := make([]int, length)
	denominators := make([]int, length)
	operators := make([]byte, length)
	isFlag := false
	//num := 0
	//pre := make([]byte, 2)
	i, j, k := 0, 0, 0
	if length > 1 && expression[0] != '-' {
		operators[k] = '+'
		k++
	}
	for l := 0; l < length; l++ {
		if expression[l] == '+' || expression[l] == '-' {
			operators[k] = expression[l]
			k++
		} else if expression[l] == '/' {
			isFlag = true
		} else {

			if isFlag {
				denominators[j], _ = strconv.Atoi(string(expression[l]))
				if l < length-1 {
					if expression[l+1] != '+' && expression[l+1] != '-' && expression[l+1] != '/' {
						now, _ := strconv.Atoi(string(expression[l+1]))
						denominators[j] = denominators[j]*10 + now
						l++
					}
				}
				j++

				isFlag = false
			} else {
				numerators[j], _ = strconv.Atoi(string(expression[l]))
				if l < length-1 {
					if expression[l+1] != '+' && expression[l+1] != '-' && expression[l+1] != '/' {
						now, _ := strconv.Atoi(string(expression[l+1]))
						numerators[i] = numerators[i]*10 + now
						l++
					}
				}
				i++

			}

		}

	}
	isFlag = false
	var minHcf, pre int = 0, 0
	for i2 := range denominators {
		if denominators[i2] == 0 {
			break
		}
		if !isFlag {
			pre = denominators[i2]
			isFlag = true
		} else {
			//fmt.Println(pre)
			temp := denominators[i2]
			//maxGcd =
			minHcf = minHCF(pre, temp, maxGCD(pre, temp))
			pre = minHcf

		}
	}
	sum := 0
	for i2 := range operators {
		if operators[i2] == 0 {
			break
		}
		temp := 0
		if operators[i2] == '-' {
			temp = -numerators[i2]
		} else {
			temp = numerators[i2]
		}

		//atoi, _ := temp
		de := denominators[i2]
		sum += temp * (minHcf / de)
	}
	gcd := maxGCD(sum, minHcf)
	if i <= 1 {
		return expression
	}
	//fmt.Println(numerators)
	//fmt.Println(denominators)
	//fmt.Println(operators)
	//fmt.Println(minHcf)
	a, b := sum/gcd, minHcf/gcd
	if b < 0 {
		a = -a
		b = -b
	}
	left := strconv.Itoa(a)
	right := strconv.Itoa(b)
	return left + "/" + right
}
func main() {
	//fractionAddition("-1/2+1/2")
	//fmt.Println()
	//fractionAddition("1/3-1/2")
	//fmt.Println()
	//fmt.Println(fractionAddition("-5/2+10/3+7/9"))
	//fmt.Println(fractionAddition("7/3+5/2-3/10"))
	//fmt.Println(fractionAddition("-7/3"))

	//fmt.Println(maxGCD(2, 2))
	//fmt.Println(maxGCD(6, 3))
	//fmt.Println(minHCF(24, 60, 12))

}
