package main

import (
	"fmt"
	"strings"
)

func encodeDigits(Phrase, parseDig string) string {
	parseDig = strings.ToUpper(parseDig)
	//fmt.Println(parseDig)
	m := map[byte]string{}
	phrase := getPhrase(Phrase)
	for k, v := range phrase {
		if v == 'I' {
			m['J'] = k
			m['I'] = k
		} else if v == 'J' {
			m['I'] = k
			m['J'] = k
		} else {
			m[v] = k
		}
	}
	//fmt.Println(m)
	ans := ""
	//fmt.Println(parseDig)
	for i := range parseDig {
		//fmt.Println(string([]byte{parseDig[i]}))
		ans += m[parseDig[i]]
	}
	return ans
}

func getDigits(Phrase, parseDig string) string {
	//phrase := getPhrase("FINAL FANTASY")
	phrase := getPhrase(Phrase)
	ans := []byte{}
	temp := parseDig
	for len(temp) > 0 {
		s := temp[0:2]
		temp = temp[2:]
		var letter byte
		if phrase[s] == 'J' || phrase[s] == 'I' {
			if len(temp) != 0 {
				letterNext := phrase[temp[0:2]]
				if letterNext == 'A' || letterNext == 'E' || letterNext == 'O' || letterNext == 'U' {
					letter = 'J'
				} else {
					letter = 'I'
				}
			}
		} else {
			letter = phrase[s]
		}

		ans = append(ans, letter)
	}
	return strings.ToLower(string(ans))
}

func getPhrase(str string) map[string]byte {
	str = strings.ToUpper(str)
	ans := map[string]byte{}
	temp := ""
	m := map[byte]int{} // 去重一波
	for i := range str {
		if m[str[i]] == 0 && 'A' <= str[i] && str[i] <= 'Z' {
			temp += string([]byte{str[i]})
		}
		m[str[i]]++
	}
	//fmt.Println("去重后：", temp)
	m = map[byte]int{}
	char := byte('A')
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			sprintf := fmt.Sprintf("%d%d", i, j)
			if len(temp) != 0 {
				ans[sprintf] = temp[0]
				if temp[0] == 'I' || temp[0] == 'J' {
					if m['I'] > 0 || m['J'] > 0 {
						temp = temp[1:]
						j--
						continue
					} else {
						m[temp[0]]++
					}
				}
				m[temp[0]]++
				temp = temp[1:]
			} else {
				if char == 'I' || char == 'J' {
					if m['I'] > 0 || m['J'] > 0 {
						char++
						j--
						continue
					} else {
						ans[sprintf] = char
						char++
					}
				}
				if m[char] > 0 {
					char++
					j--
					continue
				}
				ans[sprintf] = char

				char++
			}

		}
	}
	//
	//i := 1
	//l := []string{}
	//for s := range ans {
	//	l = append(l, s)
	//}
	//sort.Strings(l)
	//for _, s := range l {
	//	if ans[s] == 'I' || ans[s] == 'J' {
	//		fmt.Printf("|%2s %4s  ", s, "I/J")
	//	} else {
	//		fmt.Printf("|%2s %4s  ", s, string([]byte{ans[s]}))
	//	}
	//
	//	if i%5 == 0 {
	//		fmt.Println("|")
	//	}
	//	i++
	//}

	ans["00"] = ' '  // 空格
	ans["10"] = '\n' // 换行
	ans["01"] = ','  // 逗号
	ans["02"] = '.'  // 点
	return ans

}
func main() {
	//fmt.Println(getDigits("FINAL FANTASY", "53343213325552324500221413330041230022421333221042130021343200222114333201004213004123004253131053343213325232450022141231004123005342453122105312223412133300213432230053425115310024320034321445311012002214530023425100224112151213330014210041321053142200122100453214150042450012512221004123001114132114222310"))
	//fmt.Println(getDigits("1000010214242531321133341235154113424344452221515253542355"))
	//fmt.Println(encodeDigits("here is a hero\n,If you look inside your heart\nYou do not have to be afraid of what you are\n.There is an answer If you reach into your soul\nAnd the sorrow that you know will melt away"))
	//fmt.Println(encodeDigits("This is my image format JPEG"))
	//fmt.Println(encodeDigits("On My God", "i am a joy enjoy "))
	fmt.Println(encodeDigits("SIIIIIJJJJJJJIJIJIJIJIJSSADASDJIJDBAJBDHAS", "i am a joy enjoy "))
	fmt.Println(getDigits("SIIIIIJJJJJJJIJIJIJIJIJSSADASDJIJDBAJBDHAS", "1200133300130012355400233412355400"))
	fmt.Println(encodeDigits("SJIIIIIJJJJJJJIJIJIJIJIJSSADASDJIJDBAJBDHAS", "i am a joy enjoy "))
	fmt.Println(getDigits("SJIIIIIJJJJJJJIJIJIJIJIJSSADASDJIJDBAJBDHAS", "1200133300130012355400233412355400"))
	//fmt.Println(getDigits("On My God", "3300221300220033111400251233111400"))
	//	fmt.Println(encodeDigits("RIPPLE", `NO WAR, NO PAIN
	//LOVE AND PEACE`))
	//fmt.Println(encodeDigits("subject"))
	//fmt.Println(encodeDigits("jujube"))
	//fmt.Println(encodeDigits("So let it out and let it in, hey Jude, begin,\nYou are waiting for someone to perform with.\nAnd do not you know that it is just you, hey Jude, you will do,\nThe movement you need is on your shoulder."))
	//fmt.Println(getDigits("3432453200122200140034324542100112110023425100154242350012132212313200234251450034321445211023425100314200134221003414523200214200243200141145141231004211005334142100234251001445321002213432453200122200141300141322533245001211002342510045321425340012132142002342514500224251151014133100213432002242454542530021341421002342510035134253005312151500413215210014531423"))
	//fmt.Println(getDigits("2242001532210012210042512100141331001532210012210012130100343223001251313201002432331213011023425100144532005314122112133300114245002242413242133200214200433245114245410053122134021014133100314200134221002342510035134253002134142100122100122200125122210023425101003432230012513132010023425100531215150031420110213432004142523241321321002342510013323231001222004213002342514500223442511531324502"))
	//fmt.Println(getDigits("125112512432"))

}
