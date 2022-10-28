package main

//
//func leastInterval(tasks []byte, n int) int {
//
//}

func leastInterval(tasks []byte, n int) (ans int) {
	letters := make([]int, 26)
	max := 0
	count := 0
	for i := 0; i < len(tasks); i++ {
		letters[tasks[i]-'A']++
		if max < letters[tasks[i]-'A'] {
			max = letters[tasks[i]-'A']
			count = 1
		} else if letters[tasks[i]-'A'] == max {
			count++
		}
	}
	if n == 0 || (max-1)*(n+1)+count < len(tasks) {
		return len(tasks)
	}
	ans = (max-1)*(n+1) + count
	return
}
