<?php

/**
 * Class Solution
 * @package php
 * 问题：给定一棵二叉搜索树，请找出其中第k大的节点。
 * @link https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
 */
class Solution
{

    /**
     * $stack用于存放树的中序遍历结果
     * @var array
     */
    protected $stack = [];

    /**
     * 二叉查找树（又称二叉排序树）：如果左子树不为空，则左子树上所有节点的值均小于根节点的值；
     * 如果右子树不为空，则右子树上所有节点的值均大于根节点的值；左右子树也都是二叉查找树。
     * 所以，只需要进行中序遍历，得到的数组就是排好序的数组，返回数组中倒数k的元素即可。
     * @param TreeNode $root
     * @param int $k
     * @return int
     */
    public function kthLargest($root, $k)
    {
        $this->recur($root);
        $i = count($this->stack) - $k;
        return $this->stack[$i] ?? -1;
    }

    /**
     * recur方法用于进行中序遍历，并插入$stack数组
     * @param $root
     */
    public function recur($root)
    {
        if (is_null($root)) return;
        $this->recur($root->left);
        array_push($this->stack, $root->val);
        $this->recur($root->right);
    }
}

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

$solution = new Solution();
$tree = new TreeNode(3, new TreeNode(1, null, new TreeNode(2)), new TreeNode(4));

var_dump($solution->kthLargest($tree, 4));