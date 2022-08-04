package main

/**
对于N个格子（从1~N标号），机器人最开始在Start（1<=Start<=N）位置，要求在走过K（K>=1）步后（一次一格），
来到aim位置（1<=aim<=N），问机器人有多少种走法？
（注：在两端的格子只能往中间走，在中间的任意一个格子，可以选择左走或右走）
*/
func way1(n, start, aim, k int) int {
	if n < 2 || start < 1 || start > n || aim < 1 || aim > n || k < 1 {
		return -1
	}
	return process1(start, k, aim, n)
}

/**
 * 计算机器人满足条件的走法有多少种
 *
 * @param current   当前位置
 * @param remaining 剩余步数
 * @param aim       目标位置
 * @param n         格子数
 */

func process1(current, remaining, aim, n int) int {
	if remaining == 0 {
		if current == aim {
			return 1
		} else {
			return 0
		}
	}
	// 还有步数要走
	if current == 1 {
		// 在最左侧，只能往右走
		return process1(2, remaining-1, aim, n)
	} else if current == n {
		// 在最右侧， 只能往左走
		return process1(n-1, remaining-1, aim, n)
	} else {
		// 在中间位置，左走或右走都可以，所以是两种情况产生的结果之和
		return process1(current-1, remaining-1, aim, n) + process1(current+1, remaining-1, aim, n)
	}
}

func way2(n, start, aim, k int) int {
	if n < 2 || start < 1 || start > n || aim < 1 || aim > n || k < 1 {
		return -1
	}
	// 机器人当前位置范围是1~n，剩余步数范围是0~k，dp表(n + 1) * (k + 1)肯定是能够将所有的情况都包含的
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = -1
		}
	}
	return process2(start, k, aim, n, dp)
}

/**
 * 计算机器人满足条件的走法有多少种
 *
 * @param current   当前位置
 * @param remaining 剩余步数
 * @param aim       目标位置
 * @param n         格子数
 */

func process2(current, remaining, aim, n int, dp [][]int) int {
	if dp[current][remaining] != -1 {
		return dp[current][remaining]
	}
	var answer int
	if remaining == 0 {
		if current == aim {
			answer = 1
		} else {
			answer = 0
		}
	}
	// 还有步数要走
	if current == 1 {
		// 在最左侧，只能往右走
		answer = process2(2, remaining-1, aim, n, dp)
	} else if current == n {
		// 在最右侧， 只能往左走
		answer = process2(n-1, remaining-1, aim, n, dp)
	} else {
		// 在中间位置，左走或右走都可以，所以是两种情况产生的结果之和
		answer = process2(current-1, remaining-1, aim, n, dp) + process2(current+1, remaining-1, aim, n, dp)
	}
	dp[current][remaining] = answer
	return dp[current][remaining]
}

func way3(n, start, aim, k int) int {
	if n < 2 || start < 1 || start > n || aim < 1 || aim > n || k < 1 {
		return -1
	}
	// 机器人当前位置范围是1~n，剩余步数范围是0~k，dp表(n + 1)*(k + 1)肯定是能够将所有的情况都包含o的
	dp := make([][]int, n+1)
	// 剩余步数为0时， 当前位置为aim 时为1
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[aim][0] = 1
	for remaining := 1; remaining <= k; remaining++ {
		// 第一行，依赖左下方的值
		dp[1][remaining] = dp[2][remaining-1]
		// 第一行和第n行单独酸后，此处就不用判断越界问题了
		for current := 2; current < n; current++ {
			// 非边上的行，依赖在下方和左上方的值
			dp[current][remaining] = dp[current-1][remaining-1] + dp[current+1][remaining-1]
		}
		// 第n韩行，依赖左上方的值
		dp[n][remaining] = dp[n-1][remaining-1]
	}
	return dp[start][k]
}

func main() {
}
