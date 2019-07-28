# week19

---

# Algorithm [894. All Possible Full Binary Trees](https://leetcode.com/problems/all-possible-full-binary-trees/)
## 1. 问题描述
给出所有可能的完全二叉树。

完全二叉树是指一棵二叉树的每个节点要么没有子节点，要么有2个子节点。

对于有 N 个节点的二叉树，给出所有可能的完全二叉树。

每个节点的值都为 0

示例1：
* 输入：7
* 输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
* 解释：
![fivetrees](fivetrees.png)

其中 1 <= N <=20

## 2. 解题思路
采用递归的方式。

假设函数 FBT(N) 返回所有完全二叉树的列表
1. 如果 N == 1 那么返回只有一个节点组成的完全二叉树
2. 如果 N % 2 == 0 那么返回空的二叉树，因为完全二叉树的节点数必须是奇数
3. 如果 N >= 3, 那么可以递归的调用 FBT(N),其中左子树 FBT(x),右子树为 FBT(N-1-x) (循环调用 x 从 1 到 N)

## 3. 代码
```go
func allPossibleFBT(N int) []*TreeNode {
	res := []*TreeNode{}
	if N % 2 == 0 {
    		return res
    }
	if N == 1 {
		res = append(res, &TreeNode{0,nil,nil})
		return res
	}
	N -= 1
	for i := 1; i < N; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(N - i)
		for _,nl := range left {
			for _,nr := range right {
				cur := &TreeNode{0,nl,nr}
				res = append(res, cur)
			}
		}
	}
	return res
}
```
## 4. 复杂度分析
* 时间复杂度：2<sup>N</sup> 递归调用 N/2 次，且每次循环 N 次
* 空间复杂度：2<sup>N</sup>

---

# Review []()

---

# Tip

## 

---
    
# Share : 21 为什么我只改一行的语句，锁这么多？—— 极客时间 MySQL实战45讲
间隙锁和行锁的加锁规则。

规则前提：MySQL 5.x 系列 <=5.7.24，8.0 系列 <=8.0.13

间隙锁只有在可重复读的隔离级别下才有效。以下默认都在可重复隔离级别下。

## 加锁规则
两个"原则"，两个"优化"，一个"bug"
1. 原则1：加锁的基本单位是 next-key lock (前开后闭)
2. 原则2：查找过程中访问到的对象才会加锁
3. 优化1：索引上的等值查询，给唯一索引加锁的时候，next-key lock 退化为行锁
4. 优化2：索引上的等值查询，向右遍历时且最后一个值不满足等值条件的时候，next-key lock 退化为间隙锁
5. 一个 bug：唯一索引上的范围查询会访问到不满足条件的第一个值为止

建表语句
```sql
CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t values(0,0,0),(5,5,5),
(10,10,10),(15,15,15),(20,20,20),(25,25,25);
```
## 案例一：等值查询间隙锁
| session A | session B | session C |
| --- | --- | --- |
| begin; <br> update t set d=d+1 where id=7; | | |
| | insert into t values(8,8,8); <br> (blocked) | | 
| | | update t set d=d+1 where id=10; <br> (Query OK) |

1. 根据原则1，A 加锁 (5,10]
2. 根据优化2，这是一个等值查询，且 id=10 不满足条件，所以 netx-key lock 退化为间隙锁 (5,10)

因此 B 往间隙 (5,10) 中插入会被锁住，而 C 修改 id=10 这一行是可以的

## 案例二：非唯一索引等值锁
| session A | session B | session C |
| --- | --- | --- |
| begin; <br> select id from t where c=5 lock in share mode; | | |
| | update t set d=d+1 where id=5; <br> (Query OK) | |
| | | insert into t values(7,7,7); <br> (blocked) |

A 要给索引 c 上的 c=5 这一行加上读锁。
1. 根据原则1，给 (0,5] 加上 next-key lock
2. 注意 c 是普通索引，因此仅访问到 c=5 这一条记录是不能马上停下来的（因为可能存在重复值），需要向右遍历，查到 c=10 才放弃。
根据原则2，访问到的都要加锁，因此要给 (5,10] 加 next-key lock
3. 根据优化2，等值判断，向右遍历，最后一个值不满足 c=5 这个等值条件，因此退化成间隙锁 (5,10)
4. 根据原则2，只有访问到的对象才会加锁，这个查询使用覆盖索引，并不需要访问主键索引，所以主键索引上没有加任何锁，这就是为什么 B 的 update 语句可以执行完成

C 插入 (7,7,7) 会被 A 的间隙锁 (5,10) 锁住

* lock in share mode 只锁覆盖索引。但是 for update 时，系统会认为你接下来要更新数据，因此会顺便给主键索引上满足条件的行加上行锁。
* 锁是加在索引上的
* 如果要用 lock in share mode 给行加读锁避免数据被更新的话，就必须绕过覆盖索引优化，在查询字段中加入索引中不存在的字段。
比如，将 A 中的查询改为 select d from t where c=5 lock in share mode

## 案例三：主键索引范围锁
下面这两条语句加锁范围相同吗？
```sql
mysql> select * from t where id=10 for update;
mysql> select * from t where id>=10 and id<11 for update;
```
这两条语句的加锁规则不同

| session A | session B | session C |
| --- | --- | --- |
| begin; <br> select * from t where id>=10 and id<11 for update; | | |
| | insert into t values(8,8,8); <br> (Query OK) <br> insert into t values(13,13,13); <br> (blocked) | | 
| | | update t set d=d+1 where id=15; <br> (blocked) |

1. 开始执行时，要找到 id=10 的行，因此加 next-key lock (5,10]。根据优化1，主键 id 上的等值条件，退化成行锁，只加了 id=10 这一行的行锁。
2. 范围查找就往后继续查找，找到 id=15 这一行停下来，因此需要加上 next-key lock (10,15]

所以 A 此时的锁在主键索引上，行锁 id=10 和 next-key lock (10,15]
## 案例四：非唯一索引范围锁
