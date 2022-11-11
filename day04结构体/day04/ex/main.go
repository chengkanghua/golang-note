package main

import "fmt"

var (
	cent   = 50
	nutzer = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano",
		"Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(nutzer))
)

//uebrigCent分金币，返回剩余金币数
func uebrigCent() (int, map[string]int) {
	//遍历用户，获取每个名字
	for _, v := range nutzer {
		//fmt.Println(v)
		//遍历名字，获取名字中的字母
		for _, t := range v {
			//判断金币剩余，不足时结束
			if cent <= 0 {
				break
			}
			//e分1个，i分2个，o分3个，u分4个，用后减少相应的金币数
			switch t {
			case 'e', 'E':
				distribution[v] += 1
				cent -= 1
			case 'i', 'I':
				distribution[v] += 2
				cent -= 2
			case 'o', 'O':
				distribution[v] += 3
				cent -= 3
			case 'u', 'U':
				distribution[v] += 4
				cent -= 4
			}
		}
		fmt.Println(cent, distribution)
	}
	return cent, distribution

}

//获取剩余金币数
func main() {
	left, _ := uebrigCent()
	fmt.Println("剩下", left)
}
