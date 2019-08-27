# week

---

# Algorithm [890. Find and Replace Pattern](https://leetcode.com/problems/find-and-replace-pattern/)
## 1. 问题描述
给一个单词列表，给一个模式，找出符合模式的单词。

返回符合模式的所有单词列表。

示例：
* 输入: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
* 输出: ["mee","aqq"]
* 解释: "mee" {a -> m, b -> e}, "aqq" {a -> a, b -> q}

## 2. 解题思路
用两个 map : m,n 分别记录 word -> pattern 和 pattern -> word 的映射。

* m 可以校验 "aa" -> "xy" 类型的错误，m[a] = x, m[a] = y
* n 可以校验 "ab" -> "xx" 类型的错误, n[x] = a, n[x] = b

## 3. 代码
```go
func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	lp := len(pattern)
	for _,word := range words {
		var m = make(map[uint8]uint8)
		var n = make(map[uint8]uint8)
		l := len(word)
		if  l != lp {
			continue
		}
		flg := true
		for i := 0; i < l; i++ {
			if m[pattern[i]] == 0 && n[word[i]] == 0 {
				m[pattern[i]] = word[i]
				n[word[i]] = pattern[i]
			} else if m[pattern[i]] == word[i] && n[word[i]] == pattern[i] {
				continue
			} else {
				flg = false
				break
			}
		}
		if flg {
			res = append(res, word)
		}
	}
	return res
}
```
## 4. 复杂度分析
* 时间复杂度: O(N*K) N 是单词个数，K 是单词长度
* 空间复杂度: O(N*K) res 使用的空间

---

# Review []()

---

# Tip

## 

---
    
# Share 25 MySQL是怎么保证高可用的？—— 极客时间 MySQL实战45讲
最终一致性：正常情况下，主库执行的更新操作生成的所有 binlog，都可以正确地传到备库执行，备库就能达到跟主库一致的状态。
## 主备延迟

