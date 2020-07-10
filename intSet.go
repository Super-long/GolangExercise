package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) { // 扩展位图长度
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 并集,结果存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else { // 当s小于t的时候，所有多出来的项都会执行append
			s.words = append(s.words, word)
		}
	}
}

// 交集
func (s *IntSet) Intersection(t *IntSet) *IntSet {
	var inter IntSet
	for i, word := range s.words {
		if i >= len(t.words) {
			break
		}
		inter.words = append(inter.words, word&t.words[i])
	}
	return &inter
}

// 差集,s与t的差集，即属于A不属于B
func (s *IntSet) DifferenceSet(t *IntSet) *IntSet {
	var diff IntSet
	for i := range s.words {
		diff.words = append(diff.words, 0)
		if i >= len(t.words) {
			diff.words[i] = s.words[i]
			continue
		}
		for j := 0; j < 64; j++ {
			var Temp uint64
			Temp = 1 << j
			if s.words[i]&Temp != 0 && t.words[i]&Temp == 0 { //s有而t没有
				diff.words[i] |= Temp
			}
		}
	}
	return &diff
}

// 使fmt包可以直接显示
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%d", i*64+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 集合长度
func (s *IntSet) Len() (res int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res++
			}
		}
	}
	return
}

// 从集合中去除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	// C语言中非是~, go中是^
	s.words[word] &= ^(1 << bit)
}

// 删除所有元素
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var res IntSet
	num := len(s.words)
	resSum := 0 //代表返回集合目前的长度
	for resSum < num {
		res.words = append(res.words, s.words[resSum])
		resSum++
	}
	return &res
}

// 多参数插入 在内部把list看做一个slice
func (s *IntSet) AddAll(list ...int) {
	for _, word := range list {
		s.Add(word)
	}
}

// 以slice的格式返回集合中的所有元素
func (s *IntSet) Elems() (res []int){
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, i*64+j)
			}
		}
	}
	return
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(144)
	fmt.Println(&x)

	y.Add(9)
	y.Add(2)
	fmt.Println(&y)

	x.UnionWith(&y)
	fmt.Println(&x)

	fmt.Println(x.Has(3), x.Has(9))

	fmt.Println(x.Len())

	x.Remove(2)
	fmt.Println(&x)

	y.Clear()
	fmt.Println(&y)

	var cp *IntSet = x.Copy()
	fmt.Println(cp)

	y.Add(9)

	fmt.Println("---------------------")
	fmt.Println("x : ", &x)
	fmt.Println("y : ", &y)

	var inter *IntSet = x.Intersection(&y)
	fmt.Println(inter)

	var diff *IntSet = x.DifferenceSet(&y)
	fmt.Println(diff)

	res := x.Elems()
	fmt.Println(res)
}