package golang

/**
问题：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
*/
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	l, t := 0, 0
	r, b := len(matrix[0])-1, len(matrix)-1
	var res []int
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		if t++; t > b {
			break
		}
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		if r--; l > r {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		if b--; t > b {
			break
		}
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		if l++; l > r {
			break
		}
	}
	return res
}
