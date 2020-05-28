package golang

import "sort"

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
示例：
输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
*/
func SingleNumbers(nums []int) []int {
	sort.Ints(nums)
	var res []int
	for i := 0; i < len(nums); i++ {
		if len(nums)-1 == i || nums[i] != nums[i+1] {
			res = append(res, nums[i])
			if len(res) == 2 {
				break
			}
			continue
		} else {
			i++
		}
	}
	return res
}
