package main

// {x:{y:(x,y)数量}}
type DetectSquares struct {
	points map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points: map[int]map[int]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := this.points[x]; !ok {
		this.points[x] = map[int]int{}
		this.points[x][y] = 1
	} else {
		this.points[x][y]++
	}
}

func (this *DetectSquares) Count(point []int) int {
	xme, yme := point[0], point[1]
	// 找一手同y轴的大兄弟
	yn := this.points[xme]
	ans := 0
	for y, ps := range yn {
		// 全等那可不行
		if y == yme {
			continue
		}
		// 看手长度
		length := y - yme // 正负无所谓了，毕竟坐标轴
		// 得出 同 length 长度的 x 坐标
		xs := []int{xme + length, xme - length}
		for _, x := range xs {
			if m, ok := this.points[x]; !ok { //  不存在就算了
				continue
			} else {
				ans += ps * m[y] * m[yme]
			}
		}

	}
	return ans
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
func main() {

}
