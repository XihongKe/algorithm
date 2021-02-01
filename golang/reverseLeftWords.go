package golang

import (
	"strings"
)

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
link：http://www.gonglin91.com/2018/03/30/go-code-review-comments/
*/
func ReverseWords(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == "" {
			continue
		}
		res = append(res, list[i])
	}
	return strings.Join(res, " ")
}

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
link: https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
*/
func reverseLeftWords(s string, n int) string {
	sb := []byte(s)
	return string(sb[n:]) + string(sb[0:n])
}
