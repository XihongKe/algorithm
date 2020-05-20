<?php

/**
 * 问题：输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
 * 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
 * @link https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
 * Class FindContinuousSequence
 */
class FindContinuousSequence
{
    /**
     * _findContinuousSequence使用双指针法求解
     * i=1, j=2，如果i，j组成的等差数列和小于target，说明区间小了，j++扩大区间；
     * 如果i，j组成的等差数列和大于target，说明区间大了，i++缩小区间；
     * 如果等于target,加入解集中，i++，继续循环。
     * @param Integer $target
     * @return Integer[][]
     */
    function _findContinuousSequence($target)
    {
        $max = (int)($target / 2) + 1;
        $result = [];
        $i = 1;
        $j = 2;
        while ($i <= $max) {
            $sum = ($i + $j) * ($j - $i + 1) / 2; // 求i,j的等差数列和
            if ($sum > $target) {
                $i++;
            } elseif ($sum < $target) {
                $j++;
            } else {
                array_push($result, range($i, $j));
                $i++;
            }
        }
        return $result;
    }

}

$c = new FindContinuousSequence();
$result = $c->_findContinuousSequence(20);
echo json_encode($result);
