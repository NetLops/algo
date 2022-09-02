package main

/*
机器人
  W：表示下一步机器人将向上移动

  A：表示下一步机器人将向左移动

  S：表示下一步机器人将向下移动

  D：表示下一步机器人将向右移动

  越过了是否还要执行？？？？？？？？？
*/
import "fmt"

// 标记一波
func main() {
	n, m, k := 0, 0, 0
	_, _ = fmt.Scanln(&n, &m, &k)
	codes := []byte{}
	var code byte
	for i := 0; i < k; i++ {
		_, _ = fmt.Scanf("%c", &code)
		codes = append(codes, code)
	}
	//fmt.Println(m, n, k, string(codes))
	count := 1           // 起始位置0,0
	x, y := 0, 0         // 位于 m * n 的位置
	pos := map[int]int{} // 记录位置
	pos[0] = 1
	num := 0 // 命令计数器
	for _, c := range codes {
		if count != m*n {
			num++
		}
		if c == 'W' { // 上
			if x-1 >= 0 {
				x -= 1
			}
		} else if c == 'S' { // 下
			if x+1 < n {
				x += 1
			}
		} else if c == 'A' { // 左
			if y-1 >= 0 {
				y -= 1
			}
		} else if c == 'D' { // 右
			if y+1 < m {
				y += 1
			}
		}
		if pos[x*n+y] == 0 {
			count++
			pos[x*n+y] = 1
		}
	}
	//fmt.Println(pos)
	if count == m*n {
		fmt.Println("Yes")
		fmt.Println(num)
	} else {
		fmt.Println("No")
		fmt.Println(m*n - count)
	}
}
