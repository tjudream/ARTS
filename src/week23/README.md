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
主备切换：
* 主动切换：运维、软件升级、主库所在机器按计划下线等
* 被动切换：主库所在机器掉电

同步延迟

数据同步有关的时间点：
1. 主库 A 执行完成一个事务，写入 binlog， 此时刻记为 T1
2. 之后传给备库 B，B 库接收完这个 binlog，此时刻记为 T2
3. 备库 B 执行完成这个事务，此时刻记为 T3

主备延迟即 T3 - T1

在备库上执行 show slave status 命令，返回结果中的 seconds_behind_master 表示主备延迟多少秒

seconds_behind_master 的计算方法：
1. 每个事务的 binlog 里都有一个时间字段，用于记录主库上写入的时间
2. 备库取出当前正在执行的事务的时间字段的值，计算它与当前系统时间的差值，得到 seconds_behind_master, 即 T3 - T1

备库连接到主库的时候，会通过执行 select unix_timestamp() 来获取当前主库的系统时间。如果这时候发现主库的系统时间与自己不一致，
备库在执行 seconds_behind_master 计算的时候会自动扣掉这个差值。

网络正常的情况下，日志从主库传给备库所需的时间很短，即 T2 - T1 的值很小。即网络正常时，主备延迟主要来源是备库接收完 binlog 和执行完
这个事务之间的时间差。

主备延迟最直接的表现是，备库消费中转日志（relay log）的速度，比主库生产 binlog 的速度要慢

## 主备延迟的来源
### 首先，有些部署条件下，备库所在的机器的性能要比主库所在机器的性能差
如此部署的想法：
1. 备库无请求（实际和主库的请求时一样的）
2. 20个主库放在4台机器上，而备库放在一台机器上

更新过程中触发大量的读操作，所以，当备库主机上的多个备库都在争抢资源的时候，就可能会导致主备延迟

现在比较常见的部署方式是 对称部署，即主备机器一样

对称部署也会有主备延迟，原因如下

### 备库的压力大
主库提供写能力，备库提供读能力。或者一些运营后台需要的分析语句，不能影响正常业务，所以只能在备库上执行。

这样就会导致备库上的查询耗费大量的 CPU 资源，影响了同步速度，造成主备延迟。

解决方案：
1. 一主多从。除了备库外，可以多接几个从库，让这些从库来分担读的压力
2. 通过 binlog 输出到外部系统，比如 Hadoop 这类系统，让外部系统提供统计类查询的能力

### 大事务
如果一个事务在主库上执行 10 分钟，那这个事务很可能就会导致从库延迟 10 分钟。
常见的大事物：
1. 一次性地用 delete 语句删除太多数据
2. 大表的 DDL 

### 备库的并行复制能力
后续文章讨论

## 由于主备延迟的存在，所以在主备切换时，就相应的有不同的策略。
## 可靠性优先策略
![two_m](tow_m.png)
图中从 状态1 到 状态2 的详细过程：
1. 判断备库 B 现在的 seconds_behind_master，如果小于某个值（比如 5秒）继续下一步，否则持续重试这一步
2. 把主库 A 改成只读状态，即把 readonly 设置为 true
3. 判断备库 B 的 seconds_behind_master 的值，直到这个值变成 0 为止
4. 把备库 B 改成可读写状态，也就是把 readonly 设置成 false
5. 把业务请求切到备库 B

这个切换流程，一般由专门的 HA 系统来完成
![reliability_first_strategy](reliability_first_strategy.png)
图中 SBM 即 seconds_behind_master 的简写

在步骤 2 ~ 5 两个库都是 readonly 状态，此期间数据库服务不可用

步骤 3 最耗时，可能需要几秒的时间，因此要先在步骤 1 做判断

## 可用性优先策略

