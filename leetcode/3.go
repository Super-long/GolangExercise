package main

import (
	"fmt"
)

/*func lengthOfLongestSubstring(s string) int {
	lhs, rhs := 0, 0
	length := len(s)
	mp := make(map[byte]int)
	res := 0.0
	for rhs < length {
		fmt.Println(lhs,rhs)
		if mp[s[rhs]] == 0{
			res = math.Max(float64(rhs-lhs+1),float64(res))
			mp[s[rhs]]++
			rhs++
		}else{
			for mp[s[rhs]]!=0 {
				mp[s[lhs]]--
				lhs++
			}
		}
	}
	return int(res)
}*/

func lengthOfLongestSubstring(s string) int {
	has := make(map[rune]int)
	start := 0 //左区间
	res := 0
	for i, val := range s {
		if index, ok := has[val]; ok {
			if index + 1 > start {
				start = index + 1 //上一个相同位置
			}
		}
		if i - start + 1 > res {
			res = i - start + 1
		}
		has[val] = i //相当于记录下次要跳的位置 这里保证记录的值都是目前有效区间内的值
	}
	return res
}

func main() {
	str := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str))
}
