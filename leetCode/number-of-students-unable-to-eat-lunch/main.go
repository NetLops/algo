package main

func countStudents(students []int, sandwiches []int) int {
	now := []int{len(students), 0} // 0: 当前剩余的学生， 1: 判断是否走了一次循环，走了一次循环，就表示没有人可选择 心仪的三明治了

	for now[0] != 0 && now[1] <= now[0] {
		student := students[0]
		students = students[1:]

		if student != sandwiches[0] {
			students = append(students, student)
			now[1]++
		} else {
			now[0]--
			now[1] = 0
			sandwiches = sandwiches[1:]
		}
	}

	return len(students)

}
