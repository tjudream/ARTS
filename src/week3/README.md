# week3 ARTS
## Algorithm 
### Leetcode [1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/)
### 问题描述
有效的括号对的定义：（）或（A)或A+B, 其中A和B也是有效的括号对。

简单的括号对的定义：给定一个括号对有效的括号对S，若S非空，且不可被拆分成任意个有效括号对，则S是简单括号对。

给定一个有效括号对S， S = P_1 + P_2 + ... + P_k ，其中P_i是简单括号对。要求去掉S的最外层的括号并返回。

### 举例
1. 
    * 输入： "(()())(())"
    * 输出： "()()()"
2. 
    * 输入："(()())(())(()(()))"
    * 输出："()()()()(())"
3.  
    * 输入："()()"
    * 输出：""
### 分析：
典型的考察栈的问题。
1. 初始化一个空栈stack
2. 遇到左括号，如果此时stack非空，则结果字符串result += "(", 入栈 stack.push("(")
3. 遇到右括号，出栈 stack.pop() ，如果此时stack非空，则结果字符串 result += ")"

注意：
   1. 遇到左括号要先判断stack是否为空，然后再入栈
   2. 遇到右括号要先出栈，然后再判断stack是否为空
### golang 代码
```go
func removeOuterParentheses(S string) string {
	strArr := []rune(S)
	var sta,ret []rune
	for i := 0; i < len(S); i++ {
		switch strArr[i] {
		case '(':
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
			sta = append(sta, strArr[i])
		case ')':
			sta = sta[:len(sta) - 1]
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
		}
	}
	return string(ret)
}
```

这里用rune数组代替栈

##Review