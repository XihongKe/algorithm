<?php

/**
 * Class MaxValue
 * 问题：在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，
 * 并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
 * @link https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
 */
class MaxValue
{
    /**
     * 直接在$grid数组上算出每个节点的最大价值，最终右下角节点的值即是最大价值的礼物
     * @param Integer[][] $grid
     * @return Integer
     */
    public function _maxValue($grid)
    {
        $m = count($grid) - 1;
        $n = count($grid[$m]) - 1;
        // 初始化i = 0的情况
        for ($j = 1; $j <= $n; $j++) {
            $grid[0][$j] += $grid[0][$j - 1];
        }
        // 初始化j = 0的情况
        for ($i = 1; $i <= $m; $i++) {
            $grid[$i][0] += $grid[$i - 1][0];
        }
        // 计算其他节点上的最大价值
        for ($i = 1; $i <= $m; $i++) {
            for ($j = 1; $j <= $n; $j++) {
                $grid[$i][$j] += max($grid[$i - 1][$j], $grid[$i][$j - 1]);
            }
        }
        return $grid[$m][$n];
    }

}

$grid = [[7, 1, 3, 5, 8, 9, 9, 2, 1, 9, 0, 8, 3, 1, 6, 6, 9, 5], [9, 5, 9, 4, 0, 4, 8, 8, 9, 5, 7, 3, 6, 6, 6, 9, 1, 6], [8, 2, 9, 1, 3, 1, 9, 7, 2, 5, 3, 1, 2, 4, 8, 2, 8, 8], [6, 7, 9, 8, 4, 8, 3, 0, 4, 0, 9, 6, 6, 0, 0, 5, 1, 4], [7, 1, 3, 1, 8, 8, 3, 1, 2, 1, 5, 0, 2, 1, 9, 1, 1, 4], [9, 5, 4, 3, 5, 6, 1, 3, 6, 4, 9, 7, 0, 8, 0, 3, 9, 9], [1, 4, 2, 5, 8, 7, 7, 0, 0, 7, 1, 2, 1, 2, 7, 7, 7, 4], [3, 9, 7, 9, 5, 8, 9, 5, 6, 9, 8, 8, 0, 1, 4, 2, 8, 2], [1, 5, 2, 2, 2, 5, 6, 3, 9, 3, 1, 7, 9, 6, 8, 6, 8, 3], [5, 7, 8, 3, 8, 8, 3, 9, 9, 8, 1, 9, 2, 5, 4, 7, 7, 7], [2, 3, 2, 4, 8, 5, 1, 7, 2, 9, 5, 2, 4, 2, 9, 2, 8, 7], [0, 1, 6, 1, 1, 0, 0, 6, 5, 4, 3, 4, 3, 7, 9, 6, 1, 9]];
var_dump((new MaxValue())->_maxValue($grid));