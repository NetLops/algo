package main

var res = 0

func sumNums(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	res += n
	return res
}

func main() {

}
