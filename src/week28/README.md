# week28

---

# Algorithm [44. Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)
## 1. 问题描述
通配符匹配

给定一个字符串 s 和一个模式串 p，实现 '?' 和 '*' 的匹配
* '?' 匹配任意单一字符
* '*' 匹配任何字符串，包括空串

p 需要匹配 s 的整个串，而非部分
#### 示例 1 :
* 输入 : s="aa" , p="a"
* 输出 : false
#### 示例 2 :
* 输入 : s="aa" , p="*"
* 输出 : true
#### 示例 3 :
* 输入 : s="cb" , p = "?a"
* 输出 : false
#### 示例 4 :
* 输入 : s="adceb" , p = "*a*b"
* 输出 : true
#### 示例 5 :
* 输入 : s="acdcb" , p = "a*c?b"
* 输出 : false
## 2. 解题思路
动态规划

设 dp[i][j] 表示 p[0,j) 是否匹配 s[0,i) 

如果 p[j] == s[i] || p[j] == '?' , 则 dp[i][j] = dp[i-1][j-1]

如果 p[j] == '*' , 则 dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-2][j-1] || ... || dp[0][j-1] ;

推到 :
* (1) dp[i][j] = dp[i][j-1] || dp[i-1][j]
* (2) 根据(1) dp[i-1][j] = dp[i-1][j-1] || dp[i-1-1][j] = dp[i-1][j-1] ||dp[i-2][j]
* (3) dp[i][j] = dp[i][j-1] || dp[i-1][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-2][j]
= dp[i][j-1] || dp[i-1][j-1] || dp[i-2][j-1] || dp[i-3][j] = ...
= dp[i][j-1] || dp[i-1][j-1] || dp[i-2][j-1] || ... || dp[1][j-1] || dp[0][j]
* (4) 因为 p[j] == '*' , '*' 可以匹配空串，所以 dp[0][j] = dp[0][j-1]
* (5) 所以 dp[i][j] = dp[i][j-1] || dp[i-1][j-1] || dp[i-2][j-1] || ... || dp[0][j-1]

所以当 p[j] = '*' 时，状态转移方程可以简化为 dp[i][j] = dp[i][j-1] || dp[i-1][j]

初始化 dp[0][0] = true , 空串匹配空串为真



## 3. 代码

## 4. 复杂度分析


---

# Review []()

---

# Tip
 

---
    
# Share


