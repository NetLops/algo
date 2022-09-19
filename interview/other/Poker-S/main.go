package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var gradeMap = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  10,
	"Q":  10,
	"K":  10,
}

var typeMap = map[string]int{
	"S": 4,
	"H": 3,
	"C": 2,
	"D": 1,
}

type Card struct {
	Type  string
	Grade string
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isLetter(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func getCards(strs string) []Card {
	cards := []Card{}
	i := 0
	for i < len(strs) {
		if isLetter(strs[i]) {
			card := Card{}
			card.Type = strs[i : i+1]
			i++
			if isDigit(strs[i]) {
				numBytes := []byte{}
				for i < len(strs) && isDigit(strs[i]) {
					numBytes = append(numBytes, strs[i])
					i++
				}
				card.Grade = string(numBytes)
			} else {
				card.Grade = strs[i : i+1]
				i++
			}
			cards = append(cards, card)
		} else {
			fmt.Println(strs, strs[i:i+1])
			panic("format error")
		}
	}
	return cards
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxChar(a, b string) string {
	if a < b {
		return b
	}
	return a
}
func getScore(cards []Card) (int, string, int) {
	maxGrade := "A"
	scores := []int{}
	//maxType := "D"
	maxTypeScore := 1
	for i := range cards {
		//maxGrade = maxChar(maxGrade, cards[i].Grade)
		if gradeMap[maxGrade] < gradeMap[cards[i].Grade] || (gradeMap[maxGrade] == gradeMap[cards[i].Grade] && maxGrade < cards[i].Grade) {
			maxGrade = cards[i].Grade
			maxTypeScore = typeMap[cards[i].Type]
		}
		scores = append(scores, gradeMap[cards[i].Grade])
	}
	n := len(scores)
	//fmt.Println(scores)
	ans := 0
	// 类似全排列
	for i := 0; i < 3; i++ {
		temp := scores
		temp = temp[i+1:]
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (scores[i]+scores[j]+scores[k])%10 == 0 {
					now := 0
					for l := 0; l < len(scores); l++ {
						if l != i && l != j && l != k {
							now += scores[l]
						}
					}
					if now > 10 {
						now -= 10
					}
					ans = max(ans, now)
				}

			}
		}
	}
	return ans, maxGrade, maxTypeScore
}

func whoWin(person [][]Card) int {
	//fmt.Println(person)
	LeonScore, LeonMaxChar, LenonTypeScore := getScore(person[0])
	JudyScore, JudyMaxChar, JudyTypeScore := getScore(person[1])

	if LeonScore > JudyScore {
		//fmt.Println("LeonWin")
		return 0
	} else if LeonScore < JudyScore {
		//fmt.Println("JudyWin")
		return 1
	} else {
		if LeonMaxChar > JudyMaxChar {
			//fmt.Println("LeonWin")
			return 0
		} else if LeonMaxChar < JudyMaxChar {
			//fmt.Println("JudyWin")
			return 1
		} else {
			if LenonTypeScore > JudyTypeScore {
				//fmt.Println("LeonWin")
				return 0
			} else if LenonTypeScore < JudyTypeScore {
				//fmt.Println("JudyWin")
				return 1
			} else {
				//fmt.Println("Error.......")
				return 2
			}
		}
	}

}

func main() {

	score := [][]string{}
	// 读文件
	file, err := os.Open("./poker-rec.txt")
	leon, err := os.Create("./leon.txt")
	judy, err := os.Create("./judy.txt")
	if err != nil {
		panic(err)
	}
	//readAll, err := ioutil.ReadAll(file)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		//fmt.Println(strings.TrimSpace(line))
		score = append(score, strings.Split(strings.TrimSpace(line), ";"))
		if err != nil {
			break
		}
	}

	if err != nil {
		panic(err)
	}
	//fmt.Println(string(readAll))
	defer func() {
		file.Close()
		leon.Close()
		judy.Close()
	}()

	//score := [][]string{
	//	{"H4C2H5C4S3", "HJD9H6HQD10"},
	//	{"D3C7CKS10C5", "S5DASQHKD8"},
	//	{"D5CKD10H7S10", "SKS9C10D2D7"},
	//	{"H9S7CAC2D7", "D9D5C6S5DA"},
	//	{"D9D5C6S5DA", "H9S7CAC2D7"},
	//	{"D9D5C6S5DA", "D9D5C6S5DA"},
	//}

	// 获得 两人对
	personsCards := [][][]Card{}
	for i := range score {
		for j := 0; j < len(score[i]); {
			line := [][]Card{}

			line = append(line, getCards(score[i][j]))
			line = append(line, getCards(score[i][j+1]))
			j += 2
			personsCards = append(personsCards, line)
		}
	}
	for i := range personsCards {
		who := whoWin(personsCards[i])
		if who == 0 {
			leon.WriteString(strings.Join(score[i], ";") + "\n")
		} else if who == 1 {
			judy.WriteString(strings.Join(score[i], ";") + "\n")
		} else {
			fmt.Println("木有")
		}
	}

}
