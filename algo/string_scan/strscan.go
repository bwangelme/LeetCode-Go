package string_scan

import (
	"strconv"
	"strings"
	"text/scanner"
)

/*
ScanInput
本程序是为了解析 leetcode 第 210 题的输入

[[0, 1], [1, 2]]

牛克网的 OJ 直接将上述字符串原样输出了，我们需要解析该字符串，并拿到期望的输入，一个二维 int 数组
*/
func ScanInput(src string) [][]int {
	var s scanner.Scanner
	var res = make([][]int, 0)

	s.Init(strings.NewReader(src))

	var nums = make([]int, 0)
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		if token == scanner.Int {
			val, _ := strconv.Atoi(s.TokenText())
			nums = append(nums, val)
		}
		if len(nums) == 2 {
			res = append(res, nums)
			nums = make([]int, 0)
		}
	}

	return res
}
