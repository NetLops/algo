package main

func threeEqualParts(arr []int) []int {
	n := len(arr)
	find := func(x int) int {
		sum := 0
		for i := 0; i < n; i++ {
			sum += arr[i]
			if sum == x {
				return i
			}
		}
		return 0
	}
	count := 0
	for _, c := range arr {
		count += c
	}
	if count%3 != 0 {
		return []int{-1, -1}
	}
	if count == 0 {
		return []int{0, n - 1}
	}
	count /= 3
	i := find(1)
	j := find(count + 1)
	k := find(count*2 + 1)
	for k < n && arr[i] == arr[j] && arr[j] == arr[k] {
		i++
		j++
		k++
	}
	if k == n {
		return []int{i - 1, j}
	}
	return []int{-1, -1}
}
