package main

import "sort"

func rectangleArea(rectangles [][]int) int {
	MOD := int64(1e9 + 7)
	// 将所有举行的x坐标保存带list中
	allX := []int{}
	for _, rectangle := range rectangles {
		allX = append(allX, rectangle[0])
		allX = append(allX, rectangle[2])
	}
	// 对list中x从小到大排序
	sort.Ints(allX)
	ans := int64(0)
	// 每次取出相邻的x坐标
	for i := 1; i < len(allX); i++ {
		// 令相邻x坐标距离为len，如果len=0 跳出循环
		a, b, length := allX[i-1], allX[i], allX[i]-allX[i-1]
		if length == 0 {
			continue
		}
		// 定义lines 存储能够覆盖(x1,x2)的y坐标对(y1,y2)
		lines := [][]int{}
		for _, rectangle := range rectangles {
			if rectangle[0] <= a && rectangle[2] >= b { // 当矩形覆盖当前x区间，则将y坐标记录下来
				lines = append(lines, []int{
					rectangle[1], rectangle[3],
				})
			}
		}
		// 对所有的y坐标对，按照y1,y2 从小到大排序
		sort.Slice(lines, func(i, j int) bool {
			if lines[i][0] != lines[j][0] {
				return lines[i][0]-lines[j][0] < 0
			}
			return lines[i][1]-lines[j][1] < 0
		})
		// 定义tot 存储当前x区间下，y区间的并集，l,r 为上一个y区间断点
		tot, l, r := 0, -1, -1
		for _, line := range lines { // 每次取出一个区间
			if line[0] > r { // 如果和上次的区间不相交，则将上次区间计入总和，同时更新l,r
				tot += r - l
				l = line[0]
				r = line[1]
			} else if line[1] > r { // 如果上次相交，则只更新r
				r = line[1]
			}
		}

		tot += r - l // 将最后一个区间求和
		ans += int64(tot * length)
		ans %= MOD
	}
	return int(ans)
}
func main() {

}
