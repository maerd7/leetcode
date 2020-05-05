# 数据结构

## 1.数组

## 2. 字符串

#### [6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/)

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。



#### [12. 整数转罗马数字](https://leetcode-cn.com/problems/integer-to-roman/)

#### [13. 罗马数字转整数](https://leetcode-cn.com/problems/roman-to-integer/)



## 3. 树

## 4. 哈希表

## 5. 栈

## 6. 链表

#### [2. 两数相加](https://leetcode-cn.com/problems/add-two-numbers/)

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

## 7. 图

## 8. 堆

## 9. 字典树

## 10. 线段树

## 11. ordered map 有序字典

## 12. 线段树

## 13. 队列

## 14. 树状数组



# 算法

## 1. 动态规划

#### [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

```
输入: "babad"
输出: "bab"
```



## 2. 贪心算法

## 3. 深度优先搜索

dfs

对于一颗二叉树，深度优先搜索(Depth First Search)是沿着树的深度遍历树的节点，尽可能深的搜索树的分支。

那么深度遍历有重要的三种方法。这三种方式常被用于访问树的节点，根节点遍历段次序决定了名称叫法，先序根是第一个，中序根是第二个，后序根是最后一个

- 先序遍历：根节点->左子树->右子树
- 中序遍历： 左中右
- 后序遍历： 左右根

#### [面试题 04.04. 检查平衡性](https://leetcode-cn.com/problems/check-balance-lcci/)

实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。 递归函数返回值

```
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
```

#### [529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/) 

题目太长还没做。，。。

#### [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/)

给出一个二维数组 A，每个单元格为 0（代表海）或 1（代表陆地）。

移动是指在陆地上从一个地方走到另一个地方（朝四个方向之一）或离开网格的边界。

返回网格中无法在任意次数的移动中离开网格边界的陆地单元格的数量。 

示例 1：

输入：[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释： 
有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。

#### 思路

理解是一个迷宫， 从四条边界开始行走，每次都走到底，（深度优先），每次行走可以走四个方向（除了来的方向），行走过后就把该点标记为已走过。 最后全部有效的点走完，就可以找到



#### [面试题 08.10. 颜色填充](https://leetcode-cn.com/problems/color-fill-lcci/)

![image-20200504000329366](/Users/xuechaowei/Library/Mobile Documents/com~apple~CloudDocs/markdown/assets/image-20200504000329366-8521809.png)

同上，直接从一个点开始一直搜索到尽头，不过需要注意一些边界条件的检查，赋值的检查

## 4. 广度优先搜索

#### 层次遍历

 bfs  使用队列。 dfs使用递归或栈（推荐递归）



## 5. 二分查找

按照模板使用，否则老是要检查边界



## 6. 双指针

#### [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

![image-20200503133611901](/Users/xuechaowei/Library/Mobile Documents/com~apple~CloudDocs/markdown/assets/image-20200503133611901-8484171.png)

## 7. 回溯算法

#### [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

## 8. 排序

## 9. 并查集

## 10. 分治算法

## 11. 递归

## 12. 极小化极大

## 13. 拓扑排序

## 14. 二叉搜索树

## 15. sliding window 滑动窗口

#### [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

给定一个字符串，请你找出其中不含有重复字符的 **最长子串** 的长度。

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
即滑动队列，主要原理是队列，满足条件队列+1，不满足条件队列出队一个
```

# 计算机常规

## 1. 位运算

#### [7. 整数反转](https://leetcode-cn.com/problems/reverse-integer/)

123 --> 321

## 2. 几何

## 3. line sweep

## 4. random

## 5. 脑筋急转弯

## 6. 记忆化

## 7. 设计