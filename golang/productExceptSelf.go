package golang

/**
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
输入: [1,2,3,4]
输出: [24,12,8,6]

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
链接：https://leetcode-cn.com/problems/product-of-array-except-self
思路：第i位的结果等于nums数组中，小于i的元素的乘积（左边，或称前缀），乘以大于i的元素的乘积（右边，或称后缀）
*/
func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	// result先保存前缀。刚开始左边没有元素，所以result[0]=1
	result[0] = 1
	// 算出前缀。
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// R用于记录每次遍历的后缀。刚开始右边没有元素，所以R=1
	R := 1
	// 遍历乘上后缀，即所求数组
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * R
		R *= nums[i]
	}

	return result
}
