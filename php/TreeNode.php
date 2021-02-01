<?php


/**
 * 树类
 * Class TreeNode
 * @package php
 */
class TreeNode
{
    public $val, $left, $right;

    public function __construct($val, $left = null, $right = null)
    {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }

}