<?php
require "TreeNode.php";

/**
 * Class Solution
 * @package php
 * 问题：输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
 * @link https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
 */
class MaxDepth
{

    /**
     * @param TreeNode $root
     * @return int
     */
    public function _maxDepth($root)
    {
        return $this->recur($root);
    }

    /**
     * recur方法先递归到叶子节点，然后逐层回溯+1，最终得到的最大值就是树的深度
     * @param $root
     * @return int|mixed
     */
    public function recur($root)
    {
        if (is_null($root)) return 0;
        return max($this->recur($root->left), $this->recur($root->right)) + 1;
    }
}

$tree = new TreeNode(3, new TreeNode(1, null, new TreeNode(2)), new TreeNode(4));

var_dump((new MaxDepth())->_maxDepth($tree));


