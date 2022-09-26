package main

func canFormArray(arr []int, pieces [][]int) bool {
	visited := map[int]bool{}
	n := len(arr)
	t := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(pieces); j++ {
			if !visited[j] {
				for _, v := range pieces[j] {
					if arr[t] == v {
						visited[j] = true
						t++
					}
					if t == n {
						return true
					}
				}
			}
		}
	}
	return t == n
}
func main() {

}
