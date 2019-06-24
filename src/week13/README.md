#week13

---

## Algorithm [1079. Letter Tile Possibilities](https://leetcode.com/problems/letter-tile-possibilities/)
### 1. 问题描述
字母瓷砖的可能性问题

你有一组瓷砖，每个瓷砖上印有一个字母。返回可以制作的非空字母序列的数量。

例子1：
* 输入："AAB"
* 输出：8
* 解释：可能的序列有 "A" "B" "AA" "AB" "BA" "AAB" "ABA" "BAA"

例子2：
* 输入："AAABBC"
* 输出：188

注意：
1. 1 <= tiles.length <= 7
2. tiles 由大写英文字母组成


### 2. 解题思路
dfs（深度优先搜素），参考[https://leetcode.com/problems/letter-tile-possibilities/discuss/308284/Concise-java-solution](https://leetcode.com/problems/letter-tile-possibilities/discuss/308284/Concise-java-solution)

由于 tiles 中只包含大写英文字母，所以可以设置一个长度为 26 的 int 型数组，其中每个元素存储的是对应字母出现的次数。

递归地遍历数组，记录次数。

1. 我们只需要计数，不需要记住每个字符串
2. 如果我们计算每个长度的字符串的个数，那我们必须要记住之前已经出现过的字符串
3. 所以我们使用递归。这样每次递归返回后，我们只需要再将减去的字母加回来就好了。

### 3. 代码
```go
func dfs(arr []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if arr[i] == 0 {
			continue
		}
		arr[i]--
		sum = sum + 1 + dfs(arr)
		arr[i]++
	}
	return sum
}

func numTilePossibilities(tiles string) int {
	var arr []int = make([]int,26)
	for i := 0; i < len(tiles); i++ {
		arr[tiles[i] - 'A']++
	}
	return dfs(arr)
}
```

### 4. 复杂度分析
* 时间复杂度：O(n<sup>n</sup>), 其中n是字符串的长度。

最坏情况，假设字符串中每个字母各不相同，且长度为n。
* 计算长度为n的字符串的排列情况是 n*(n-1)*(n-2)...1 = n! , 需要递归的次数是 n!
* 计算长度为 n-1 的长度的字符串的排列情况 n*(n-1)...2
* 计算长度为 n-2 的长度的字符串的排列情况 n*(n-1)...3
* ...
* 计算长度为 1 的长度的字符串的排列情况 n

综上，计算所有可能需要遍历的次数是 n<sup>n</sup> 数量级的

* 空间复杂度： O(1) 只需要一个长度为 26 的 int 型数组

---

## Review []()

---

## Tip

### 

---
    
## Share
### 15 答疑文章（一）：日志和索引相关问题 —— 极客时间 MySQL实战45讲
### 日志相关问题
#### 问题1: 两阶段提交的不同瞬间，MySQL 如果发生异常重启，是怎么保证数据完整性的？
![two-phase_commit](two_phase_commit.png)

```sql
update T set c=c+1 where ID=2;
```
上图是该语句的执行步骤

这里的 commit 并非 MySQL 语法中的 commit 语句，而是 commit 步骤，是指事务提交过程中的一个小步骤，也是最后一步，
当这个步骤执行完成后，这个事务就提交完成了。

commit 语句执行的时候，会包含 commit 步骤。

在两阶段提交的不同时刻，MySQL异常重启会出现什么现象。

* 如果在时刻 A 崩溃

也就是在 redo log 处于 prepare 阶段之后，写入 binlog 之前的时间点。

此时 binlog 还没有写，redo log 还没有提交，所以崩溃恢复的时候这个事务会回滚。这时候 binlog 还没写，所以不会传到备库。

* 在时刻 B 崩溃

也就是 binlog 写完了， redo log 还没有 commit 前的时间点。

崩溃恢复时的判断规则：
1. 如果 redo log 里的事务是完整的，也就是已经有了 commit 标识，则直接提交
2. 如果 redo log 里面的事务只有完整的 prepare，则判断对应的事务 binlog 是否存在并完整：
    * a. 如果是，则提交事务
    * b. 否则，回滚事务

时刻 B 对应的是 2.a 的情况，所以崩溃恢复过程中事务会被提交。

#### 追问1： MySQL 怎么知道 binlog 是完整的
回答：一个事务的 binlog 是有完整格式的
* statement 格式的 binlog，最后会有 COMMIT
* row 格式的 binlog，最后会有一个 XID event

在 MySQL 5.6.2 版本之后，引入了 binlog-checksum 参数，用来验证 binlog 的内容的正确性。
对于 binlog 日志由于磁盘的原因，可能会在日志中间出错，MySQL 可以通过校验 checksum 的结果来发现。

#### 追问2：redo log 和 binlog 是怎么关联起来的
回答：它们有一个共同的数据字段，叫 XID。崩溃恢复的时候，会按顺序扫描 redo log
* 如果碰到既有  prepare，又有 commit 的 redo log，就直接提交
* 如果碰到只有 prepare，而没有 commit 的 redo log， 就拿着 XID 去 binlog 找对应的事务
#### 追问3： 处于 prepare 阶段的 redo log 加上完整的 binlog， 重启就能恢复， MySQL 为什么要这么设计


