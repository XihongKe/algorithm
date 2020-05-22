package golang

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
link: https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
 */
func ReverseLeftWords(s string, n int) string {
	sb := []byte(s)
	return string(sb[n:]) + string(sb[0:n])
}
