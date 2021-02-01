package golang

import (
	"sort"
)

/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

link：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
*/
func MaxSlidingWindow(nums []int, k int) []int {
	// 极端情况判断（空的数组、k值为1）
	if len(nums) == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	var max []int
	tmp := make([]int, k)
	i := 0
	j := k
	for j <= len(nums) {
		// 初始化，以及上轮的最大值刚好是这轮被删掉的元素，需要重新遍历找最大值
		if i == 0 || max[i-1] == nums[i-1] {
			copy(tmp, nums[i:j])
			sort.Ints(tmp)
			max = append(max, tmp[k-1])
		} else { // 这轮刚加入的元素和上轮最大值比较，谁比较大，谁就是这轮的最大值
			if max[i-1] > nums[j-1] {
				max = append(max, max[i-1])
			} else {
				max = append(max, nums[j-1])
			}
		}
		i++
		j++
	}
	return max
}
