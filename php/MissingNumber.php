<?php

/**
 * 问题：一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
 * 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
 * @link https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
 * Class MissingNumber
 */
class MissingNumber
{
    /**
     * 二分法
     * @param Integer[] $nums
     * @return Integer
     */
    public function _missingNumber($nums)
    {
        $i = 0;
        $j = count($nums) - 1;
        $m = round(($j + $i) / 2);
        if ($nums[$j] == $j) return $j + 1; // 缺失的数字是最后一个数字的情况
        while (true) {
            if ($nums[$m] != $m) {
                // 不存在的数落在左半边的情况
                $j = $m;
            } else {
                // 不存在的数落在右半边的情况
                $i = $m;
            }
            $m = round(($i + $j) / 2);
            if ($nums[$i] != $i) return $i;
            if ($m == $j) return $j;
        }
    }
}

$c = new MissingNumber();
$number = $c->_missingNumber([0, 1, 2, 3]);
echo $number;