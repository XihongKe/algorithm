<?php


class MaxSlidingWindow
{
    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    public function _maxSlidingWindow($nums, $k)
    {
        if (count($nums) == 0) return [];
        if ($k == 1) return $nums;
        $max = [];
        $i = 0;
        $j = $k;
        while ($j <= count($nums)) {
            if ($i == 0 || $max[$i - 1] == $nums[$i - 1]) {
                $tmp = array_slice($nums, $i, $k);
                rsort($tmp);
                array_push($max, $tmp[0]);
            } else { // 这轮刚加入的元素和上轮最大值比较，谁比较大，谁就是这轮的最大值
                if ($max[$i - 1] > $nums[$j - 1]) {
                    array_push($max, $max[$i - 1]);
                } else {
                    array_push($max, $nums[$j - 1]);
                }
            }
            $i++;
            $j++;
        }
        return $max;
    }

}

$c = new MaxSlidingWindow();
echo json_encode($c->_maxSlidingWindow([1,3,-1,-3,5,3,6,7], 3));
