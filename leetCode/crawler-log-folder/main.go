package main

func minOperations(logs []string) int {
	ans := 0

	for _, log := range logs {
		if len(log) == 3 && log == "../" {
			if ans != 0 {
				ans--
			}
		} else if len(log) == 2 && log == "./" {
			continue
		} else {
			ans++
		}

	}
	return ans
}
func main() {

}
