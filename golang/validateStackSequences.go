package golang

/**
问题：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能
是该压栈序列的弹出序列。

链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

思路：使用一个辅助栈stack 遍历压入pushed,每次压入判断一次popped的第J个元素是否等于压入的元素，如果是，就弹出，j++；
当循环结束，i和j都遍历了pushed和popped的时候，说明popped是pushed的弹出序列
*/
func ValidateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	var stack []int
	i, j := 0, 0
	length := len(pushed)
	for j < len(popped) {
		if len(stack) == 0 {
			if i == length-1 { // 没有数据可压入了，跳出
				break
			}
			stack = append(stack, pushed[i]) // 压入，继续循环
			i++
			continue
		}

		if stack[len(stack)-1] == popped[j] { // 符合条件 弹出
			stack = stack[0 : len(stack)-1]
			j++
		} else { // 不符合条件 判断是否已经遍历完毕
			if i == length {
				break
			}
			stack = append(stack, pushed[i]) // 压入，继续循环
			i++
		}
	}
	return i >= length-1 && j >= length-1
}
