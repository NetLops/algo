package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	f := make([][]int, len(obstacleGrid))
	for i := range f {
		f[i] = make([]int, len(obstacleGrid[i]))
	}
	f[0][0] = 1
	if obstacleGrid[0][0] == 1 {
		f[0][0] = 0
	}
	i := 0
	j := 0
	for i = 0; i < len(obstacleGrid); i++ {
		for j = 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] != 1 { // 碰到障碍物就不鸟
				if i > 0 && j > 0 { // 往下/往右
					f[i][j] = f[i-1][j] + f[i][j-1]
				} else if i > 0 { // 往下
					f[i][j] = f[i-1][j]
				} else if j > 0 { // 往右
					f[i][j] = f[i][j-1]
				}
			}
		}
	}
	return f[i-1][j-1]
}

func main() {
}
