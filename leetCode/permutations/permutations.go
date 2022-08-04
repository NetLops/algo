package permutations

var list [][]int // 为最终结果

func permuteUnique(nums []int) [][]int {
	len := len(nums)
	collect := []int{} // 回朔过程收集元素
	flags := []bool{}  // 用来标记
	for i := 0; i < len; i++ {
		flags = append(flags, false)
	}
	dfs(flags, nums, collect, 0)
	return list
}

func dfs(flags []bool, nums []int, collect []int, index int) {
	len := len(nums)
	if index == len { // 停止
		temp := make([]int, len, len)
		copy(temp, collect)
		list = append(list, temp)
	} else {
		for i, num := range nums {
			if flags[i] {
				continue
			}
			collect = append(collect, num)
			flags[i] = true // 标记该元素被使用
			dfs(flags, nums, collect, index+1)
			flags[i] = false                                        // 还原
			collect = append(collect[:index], collect[index+1:]...) // 这是移除部分
		}
	}
}

//
//func main() {
//	fmt.Println(permuteUnique([]int{1, 2, 3, 4}))
//}
