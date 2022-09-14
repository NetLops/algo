package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func maximumSwap1(num int) int {
	temp := strconv.Itoa(num)
	numsTemp := []byte(temp)
	mPos := map[byte][]int{}
	nums := []byte(temp)
	for i := range nums {
		if _, ok := mPos[nums[i]]; !ok {
			mPos[nums[i]] = []int{}
		}
		mPos[nums[i]] = append(mPos[nums[i]], i)
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] > nums[j] {
			return true
		}
		return false
	})
	pos := -1
	var val byte
	for i := 0; i < len(nums); i++ {
		if numsTemp[i] != nums[i] {
			pos = i
			val = nums[i]
			break
		}
	}
	if pos != -1 {
		beforeUpdate := numsTemp[pos]
		afterUpdate := val
		// 找替换前最靠前的索引、找替换后最小的索引
		numsTemp[mPos[val][len(mPos[afterUpdate])-1]], numsTemp[mPos[beforeUpdate][0]] = numsTemp[mPos[beforeUpdate][0]], numsTemp[mPos[val][len(mPos[afterUpdate])-1]]
	}
	ans, _ := strconv.Atoi(string(numsTemp))
	return ans
}
func maximumSwap(num int) int {
	// stack 单调递减
	chars, stack := []byte(strconv.Itoa(num)), []int{}
	n := len(chars)
	// 左边 应该 更偏 左边，右边应该更偏向右（即使相等 也要像右）
	left, right := n, n-1
	for i, char := range chars {
		for len(stack) > 0 && chars[stack[len(stack)-1]] < char {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if idx < left {
				left = idx
				// 当出现更大的数将右边交换的坐标弹出，右边坐标成为新坐标
				right = i
			}
			//if idx == right {
			//	right = i
			//}
		}
		// 当出现与右边坐标一样大的但更靠右时，右边坐标变为新坐标
		if right < n && chars[right] <= char {
			right = i
		}
		stack = append(stack, i)
	}
	if left < n {
		chars[left], chars[right] = chars[right], chars[left]
		ans, _ := strconv.Atoi(string(chars))
		return ans
	}
	return num
}

func main() {
	time1 := time.Now()
	fmt.Println(maximumSwap(99897892638999))
	fmt.Println(time.Since(time1))
	time2 := time.Now()
	fmt.Println(maximumSwap1(99897892638999))
	fmt.Println(time.Since(time2))
}
