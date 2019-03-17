# Leetcode [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
* 求字符串的最长子串，要求子串中没有重复的字符

## 思路1——暴力法
1. 从n(初始为0)字符开始遍历整个字符串的字符——O(n)
2. 从n+1个字符开始遍历——O(n<sup>2</sup>)
3. 查看前n个字符中是否有重复的，如果没有，则计算最大长度；如果有，则停止遍历
