package golang

/**
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
输入: [1,2,3,4]
输出: [24,12,8,6]
 
提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
链接：https://leetcode-cn.com/problems/product-of-array-except-self
思路：生成lProduct、rProduct两个数组，分别是左侧所有数字的乘积和右侧所有数字的乘积，结果就是这两个数组乘积相乘就可以了
 */
func ProductExceptSelf(nums []int) []int {
	lProduct := make([]int, len(nums))
	rProduct := make([]int, len(nums))
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			lProduct[i] = nums[i]
		} else {
			lProduct[i] = nums[i] * lProduct[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rProduct[i] = nums[i]
		} else {

			rProduct[i] = nums[i] * rProduct[i+1]
		}
	}

	for key, _ := range res {
		l := 1
		r := 1
		if key != 0 {
			l = lProduct[key-1]
		}
		if key != len(res)-1 {
			r = rProduct[key+1]
		}
		res[key] = l * r
	}
	return res

}
