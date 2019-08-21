# week22

---

# Algorithm [1161. Maximum Level Sum of a Binary Tree](https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/)
## 1. 问题描述
找到二叉树中层中节点和最大的那一层

给定一个二叉树的根 root，root 的层是 1，root 的左右儿子的层是 2，以此类推。

将节点按照层进行分类，计算每一层所有节点的和，返回和最大的那个层的层号，如果有和相同的层，则返回最小的那层的层号。

示例 1：
![capture](capture.jpeg)
* 输入：[1,7,0,7,-8,null,null]
* 输出：2
* 解释：
    * 第 1 层 sum = 1
    * 第 2 层 sum = 7 + 0 = 7
    * 第 3 层 sum = 7 + (-8) = -1
    * 所以和最大的层的层号为 2

## 2. 解题思路
广度优先搜索，计算每一层的所有节点的和。

## 3. 代码
```go
type Element struct {
	node *TreeNode
	Level int
}
func maxLevelSum(root *TreeNode) int {
    if root == nil {
		return 0
	}
	maxSum := 0
	maxLevel := 1
	curSum := 0
	curLevel := 1
	elements := []*Element{&Element{root, 1}}

	for len(elements) > 0 {
		ele := elements[0]
		elements = elements[1:]
		node := ele.node
		level := ele.Level
		if node.Left != nil {
			elements = append(elements, &Element{node.Left, level + 1})
		}
		if node.Right != nil {
			elements = append(elements, &Element{node.Right, level + 1})
		}

		if level == curLevel {
			curSum += node.Val
		} else {
			if curSum > maxSum {
				maxSum = curSum
				maxLevel = curLevel
			}
			curSum = node.Val
			curLevel = level
		}
	}

	return maxLevel
}
```
## 4. 复杂度分析
* 时间复杂度： O(N) N 为树的节点个数，只需遍历一遍树的所有节点
* 空间复杂度： O(2<sup>H</sup>) H 为树高，root 的高为 1，因为数组 elements 最多需要存储一层的所有节点。
一个满二叉树的最后一层的节点数最多，为 2<sup>H</sup> 个。

---

# Review []()

---

# Tip

## 

---
    
# Share 24 MySQL是怎么保证主备一致的？ —— 极客时间 MySQL实战45讲
binlog 可以用来归档，也可以用来做主备同步。
## MySQL 主备的基本原理
![master-slave](master-slave.png)
状态1 ，客户端读写 A，B 是 A 的备库，只是将 A 的更新都同步过来，到本地执行。

状态2，A 是 B 的备库。

状态 1 中，虽然 B 没有被直接访问，但是依然建议设置成 readonly 模式，因为：
1. 有时候一些运营类的语句会放到 B 库上去查，设置为 readonly 可以防止误操作
2. 防止切换逻辑有 bug， 比如切换时造成双写，造成主备不一致
3. 可以用 readonly 状态，来判断节点的角色

readonly 设置对超级（super）权限用户是无效的，而用于同步更新的线程是拥有超级权限的。所以 readonly 不影响同步。

同步的详细流程：
![master-slave-workflow](master-slave-workflow.png)
B 和 A 直接维持一个长连接。A 内部有一个线程，专门用于服务备库 B 的这个长连接。

一个事务日志同步的完整流程如下：
1. 在备库 B 上通过 change master 命令，设置主库 A 的 IP、端口、用户名、密码，以及要从哪个位置开始请求 binlog，这个位置包含文件名和日志偏移量。
2. 在备库 B 上执行 start slave 命令，这时候备库会启动两个线程，就是图中的 io_thread 和 sql_thread。其中 io_thread 负责与主库建立连接。
3. 主库 A 校验完用户名、密码后，开始按照备库 B 传过来的位置，从本地读取 binlog，发给 B
4. 备库 B 拿到 binlog 后，写到本地文件，称为中转日志(relay log)
5. sql_thread 读取中转日志，解析出日志里的命令，并执行

后来由于多线程方案的引入，sql_thread 演化成为了多个线程。

## binlog 的三种格式对比
binlog 有三种格式
* statement : 记录sql语句
* row : 记录数据行
* mixed : 以上两种格式的混合

建表和初始化：
```sql
mysql> CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `a` int(11) DEFAULT NULL,
  `t_modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `a` (`a`),
  KEY `t_modified`(`t_modified`)
) ENGINE=InnoDB;

insert into t values(1,1,'2018-11-13');
insert into t values(2,2,'2018-11-12');
insert into t values(3,3,'2018-11-11');
insert into t values(4,4,'2018-11-10');
insert into t values(5,5,'2018-11-09');

```
用 MySQL 客户端做实验的话需要加上 -c 参数。

当 binlog_format=statement 时，binlog 中记录的是 SQL 的原文：
```sql
mysql> delete from t /*comment*/  where a>=4 and t_modified<='2018-11-10' limit 1;
mysql> show binlog events in 'master.000001';
```
![statement_binlog](statement_binlog.png)
* 第一行 SET @@SESSION.GTID_NEXT='ANONYMOUS' 后续文章会介绍
* 第二行 BEGIN，跟第四行 COMMIT 对应，表示中间是一个事务
* 第三行 是真实的执行语句。在 delete 之前还有一个 use test ，这条命令是 MySQL 自行添加的，这条语句可以防止备库在其他库上执行。
后边的 delete 就是我们执行的 delete 语句。
* 最后一行是 COMMIT。其中还有一个 xid=61。
XID 用于关联 redo log 和 binlog，redo log 和 binlog 有一个共同的数据字段，xid。崩溃恢复的时候，会按顺序扫描 redo log，
如果碰到只有 prepare 而没有 commit 的 redo log ，就拿着 xid 去 binlog 中找对应的事务。

为了说明 statement 和 row 的区别，我们来看一下 delete 语句的执行效果：
![delete_warning](delete_warning.png)
产生了一个 warning，原因是当前 binlog 设置的是 statement 格式，并且语句中有 limit，所以这个命令是 unsafe 的。

delete 带 limit 可能会导致主备不一致：
1. 如果 delete 语句使用的是索引 a，那么会根据索引 a 找到第一个满足条件的行，也就是说删除的是 a=4 这一行
2. 但如果使用的是索引 t_modified ，那么删除的就是 t_modified='2018-11-09' 也就是 a=5 的这一行。

由于 statement 记录的是语句原文，因此可能会出现这样一种情况：在主库执行 SQL 的时候，用的是索引 a；而在备库执行这条 SQL 的时候却使用的是索引 t_modified.

设置 binlog_format='row'
![row_binlog](row_binlog.png)

