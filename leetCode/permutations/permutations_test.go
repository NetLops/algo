package permutations

import "testing"

//
func permute(nums []int) [][]int {
	list := new([][]int)
	flags := []bool{}
	collect := []int{}
	index := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		flags = append(flags, false)
	}
	testDfs(list, flags, &collect, &nums, index, length)
	return *list
}

func testDfs(list *[][]int, flags []bool, collect *[]int, nums *[]int, index int, length int) {
	if index == length {
		temp := make([]int, length, length)
		copy(temp, *collect)
		*list = append(*list, temp)
	} else {
		for i, num := range *nums {
			// 判断是否已经访问过了
			if flags[i] {
				continue
			}
			flags[i] = true
			*collect = append(*collect, num)
			testDfs(list, flags, collect, nums, index+1, length)
			flags[i] = false
			//去掉那个元素
			*collect = append((*collect)[:index], (*collect)[index+1:]...)
		}
	}
}

func TestTest2(t *testing.T) {
	t.Log(permute([]int{1, 2, 3, 4}))
}
