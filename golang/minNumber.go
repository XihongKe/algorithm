package golang

import (
	"strconv"
	"strings"
)

/**
问题：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

解题思路：此题求拼接起来的 “最小数字” ，本质上是一个排序问题。
排序判断规则： 设 nums 任意两数字的字符串格式 "x" 和 "y" ，则
若拼接字符串 "x" + "y" > "y" + "x" ，则 "x"排在 "y" 后面；
反之，若 "x" + "y" < "y" + "x" ，则 "x" 排在 "y" 前面；
根据以上规则，套用任何排序方法对 nums 执行排序即可。
*/
func MinNumber(nums []int) string {
	var numsString []string

	// 切片转为字符串切片
	for _, num := range nums {
		numsString = append(numsString, strconv.FormatInt(int64(num), 10))
	}

	// 排序
	sortedNums := quickSortString(numsString)

	// 切片拼接成字符串返回
	return strings.Join(sortedNums, "")
}

// 快速排序
func quickSortString(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}

	v := arr[0]
	var low, up []string
	for i := 1; i < len(arr); i++ {
		// 字符串 "x" + "y" > "y" + "x"，则"x"排在"y"后面
		if v+arr[i] < arr[i]+v {
			up = append(up, arr[i])
		} else {
			low = append(low, arr[i])
		}
	}

	low = quickSortString(low)
	up = quickSortString(up)

	// 按顺序合并切片
	return sliceMerge(low, []string{v}, up)
}

// sliceMerge按顺序合并字符串切片
func sliceMerge(arr ...[]string) []string {
	var res []string
	for _, v := range arr {
		res = append(res, v...)
	}
	return res
}
