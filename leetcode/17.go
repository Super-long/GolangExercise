package main

import "fmt"

func letterCombinations(digits string) []string {
	ans := []string{}
	dMap := map[byte]string{}
	dMap['2'] = "abc"
	dMap['3'] = "def"
	dMap['4'] = "ghi"
	dMap['5'] = "jkl"
	dMap['6'] = "mno"
	dMap['7'] = "pqrs"
	dMap['8'] = "tuv"
	dMap['9'] = "wxyz"
	PrintLetters(digits, dMap, 0, 0, []byte{}, &ans)
	return ans
}

func PrintLetters(digits string, dMap map[byte]string, cur int, dcur int, tmp []byte, ans *[]string) {
	length := len(digits)
	if cur >= length || dcur >= len(dMap[digits[cur]]) {
		return
	}
	if dcur != 0 {
		tmp[cur] = dMap[digits[cur]][dcur]
	} else {
		tmp = append(tmp, dMap[digits[cur]][dcur])
	}

	if len(tmp) == len(digits) {
		*ans = append(*ans, string(tmp))
	}
	PrintLetters(digits, dMap, cur+1, 0, tmp, ans)
	PrintLetters(digits, dMap, cur, dcur+1, tmp, ans) //当第一次执行到这里的时候tmp第一次满 后面就是在里面替换值了
}

func main() {
	para := "234"
	res := letterCombinations(para)
	for _, val := range res {
		fmt.Println(val)
	}
	fmt.Println(len(res))
}
