<?php

/**
 * Class MaxProfit
 * 问题：假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 * @link https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
 */
class MaxProfit
{
    /**
     * 最小买入值记为代价cost, 最大卖出利润记为profit，状态转移方程如下
     * profit=max(profit,prices[i]−min(cost,prices[i])
     * 时间复杂度O(n),空间复杂度O(1)
     * @param $prices
     * @return int|mixed
     */
    public function _maxProfit($prices)
    {
        $profit = 0;
        $cost = $prices[0];
        foreach ($prices as $price) {
            $cost = min($price, $cost); //获得最小代价
            $profit = max($profit, $price - $cost); //获得最大利润
        }
        return $profit;
    }
}

$c = new MaxProfit();
var_dump($c->_maxProfit([7, 6, 4, 3, 1]));

