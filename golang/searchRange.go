package golang

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
思路：二分查找找到一个target的索引，然后使用双指针发散找到边界。
*/
func SearchRange(nums []int, target int) []int {
	// 低位指针
	l := 0
	// 高位指针
	r := len(nums) - 1
	m := (r + l) / 2
	for l <= r {
		// 定位到一个等于target的元素，查找边界并返回
		if nums[m] == target {
			return getRange(nums, target, m)
		}
		// 二分查找的通用套路，不赘述
		if nums[m] > target {
			r = m - 1
		}
		if nums[m] < target {
			l = m + 1
		}
		m = (r + l) / 2
	}
	// 数组中不存在目标值
	return []int{-1, -1}
}

// getRange利用二分查找获得的index，使用双指针发散找到target的边界
func getRange(nums []int, target int, m int) []int {
	startIndex := m
	endIndex := m
	for {
		if (startIndex == 0 || nums[startIndex-1] != target) && (endIndex == len(nums)-1 || nums[endIndex+1] != target) {
			return []int{startIndex, endIndex}
		}
		if !(startIndex == 0 || nums[startIndex-1] != target) {
			startIndex--
		}
		if !(endIndex == len(nums)-1 || nums[endIndex+1] != target) {
			endIndex++
		}
	}
}
