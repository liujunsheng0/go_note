package algorithm

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 字母串计算
// https://www.lintcode.com/problem/1579/
// 是有一串英文字母，每次可以删掉一个英文字母，求最少删多少次可以让每个字母出现的次数不同
// 只有小写字母，英文串的长度不会超过100000
// 输入："aaabbb"  输出：1  解释：因为'a'的个数有三个，'b'的个数也有三个，所以删掉一个'a'或者一个'b'都可以满足题意
// 输入："abcd"    输出：3  解释：因为'a','b','c','d'各有一个，所以需要删掉任意三个来满足题意

func LetterStringCalculate(str string) int {
	countMap := map[int]int{}
	for _, c := range str {
		countMap[int(c)]++
	}
	countSlice := make([]int, 0, len(countMap))
	for _, count := range countMap {
		countSlice = append(countSlice, count)
	}

	// 由大至小排序
	sort.Slice(countSlice, func(i, j int) bool {
		return countSlice[i] > countSlice[j]
	})
	appeared := map[int]bool{}
	ans := 0
	// 最近未出现过的数
	for _, count := range countSlice {
		for ; count > 0; count-- {
			if !appeared[count] {
				appeared[count] = true
				break
			}
			count--
			ans++
		}
	}
	return ans
}

func TestLetterStringCalculate(t *testing.T) {
	ar := assert.New(t)
	ar.Equal(1, LetterStringCalculate("aaabbb"))
	ar.Equal(3, LetterStringCalculate("abcd"))
}
