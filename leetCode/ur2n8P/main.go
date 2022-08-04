package main

import "fmt"

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	m := map[int]bool{}
	for _, sequence := range sequences {
		for i := 0; i < len(sequence)-1; i++ {
			m[(sequence[i]<<10)|sequence[i+1]] = true
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if !m[(nums[i]<<10)|nums[i+1]] {
			return false
		}
	}
	return true
}

func main() {
	//nums := []int{4, 1, 5, 2, 6, 3}
	nums := []int{1, 2, 3}
	sequences := [][]int{{1, 2}, {1, 3}, {2, 3}}
	//sequences := [][]int{{5, 2, 6, 3}, {4, 1, 5, 2}}
	fmt.Println(sequenceReconstruction(nums, sequences))
}

//[4,1,5,2,6,3]
//[[5,2,6,3],[4,1,5,2]]
