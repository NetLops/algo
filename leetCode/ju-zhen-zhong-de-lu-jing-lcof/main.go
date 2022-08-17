package main

func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}
		// 代表此时已经找到了
		if k == len(word)-1 {
			return true
		}

		// 重新回复
		board[i][j] = ' '
		// 上下左右
		res := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		// 恢复
		board[i][j] = word[k]
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {

}
