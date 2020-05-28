package golang

/**
问题：在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
 */
func SingleNumber(nums []int) int {
	numsMap := map[int]bool{}
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num] = false
		} else {
			numsMap[num] = true
		}
	}
	for num, _ := range numsMap {
		if numsMap[num] {
			return num
		}
	}
	return 0
}
