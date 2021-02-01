package golang

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third]+nums[first] > 0 {
				third--
			}
			if second == third {
				continue
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
