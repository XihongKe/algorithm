<?php

/**
 * Class IsStraight
 * 问题：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
 * 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
 * @link https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
 */
class IsStraight
{
    /**
     * 顺子的条件：1.没有非0重复数字；2.最大值减去非0最小值小于5
     * @param Integer[] $nums
     * @return Boolean
     */
    public function _isStraight($nums)
    {
        $zeroCounts = 0;
        sort($nums);
        // 求值为0的元素的数量
        while ($nums[0] == 0) {
            $zeroCounts++;
            unset($nums[0]);
            $nums = array_values($nums);
        }
        // 排除0之后的元素唯一数组
        $nums = array_values(array_unique($nums));
        // 有非0重复数字
        if (count($nums) + $zeroCounts != 5) return false;
        return ($nums[count($nums) - 1] - $nums[0]) < 5;
    }
}


$c = new IsStraight();
var_dump($c->_isStraight([11, 0, 9, 0, 0]));