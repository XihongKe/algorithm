<?php

class bucket{
    public $min = null;
    public $max = null;
    public $key = null;

    public function __construct($key)
    {
        $this->key = $key;
    }
}

class MaxSortedDistance{
    public function __invoke(array $array)
    {
        $min = $array[0];
        $max = $array[0];
        for ($i = 1; $i < count($array); $i ++){
            if($array[$i] < $min) $min = $array[$i];
            if($array[$i] > $max) $max = $array[$i];
        }

        if($min == $max) return 0;
        $buckets = [];
        $d = $max - $min;
        $bucketNum = count($array);
        for ($i = 0; $i < $bucketNum; $i ++){
            $buckets[] = new bucket($i);
        }

        for ($i = 0; $i < count($array); $i ++){
            $index = (int)(($array[$i] - $min) * (($bucketNum - 1) / $d));
            if($buckets[$index]->min===null || $buckets[$index]->min > $array[$i]){
                $buckets[$index]->min = $array[$i];
            }
            if($buckets[$index]->max===null || $buckets[$index]->max < $array[$i]){
                $buckets[$index]->max = $array[$i];
            }
        }

        echo json_encode($buckets, JSON_PRETTY_PRINT);
        $leftMax = $buckets[0]->max;
        $maxDistance = (int) ($buckets[0]->max - $buckets[0]->min);
        for ($i = 1; $i < count($buckets); $i ++){
            if($buckets[$i]->min === null){
                continue;
            }
            if($buckets[$i]->min - $leftMax > $maxDistance){
                $maxDistance = $buckets[$i]->min - $leftMax;
            }
            $leftMax = $buckets[$i]->max;
        }

        return $maxDistance;
    }
}

function run(){
    $maxSortedDistance = new MaxSortedDistance();
    echo $maxSortedDistance([2,6,3,4,5,10,9]);
}

run();
