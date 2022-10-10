package golang

import (
	"fmt"
)

// 返回x的n次方
func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1.00
	}
	var isPositiveInteger bool
	if n > 0 {
		isPositiveInteger = true
	} else {
		isPositiveInteger = false
		n = -n
	}
	left := 1.0

	for {
		// 二分到二次方，跳出
		if n == 1 {
			break
		}
		if n%2 == 1 {
			n = (n - 1) / 2
			left *= x
		} else {
			n /= 2
		}
		x *= x
		fmt.Printf("n = %d, x = %f, left = %f \r\n", n, x, left)
	}

	if isPositiveInteger {
		return x * left
	}
	return 1 / (x * left)
}
