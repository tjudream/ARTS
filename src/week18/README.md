# week18

---

# Algorithm : [657. Robot Return to Origin](https://leetcode.com/problems/robot-return-to-origin/)
## 1. 问题描述
机器人返回原点问题

一个机器人在一个二维坐标系中，初始点在(0,0)

给定一个字符串，表示机器人的一系列的动作

* L 向左走一步
* R 向右走一步
* U 向上走一步
* D 向下走一步

判断机器人执行这一些列的动作之后，是否又回到了原点(0,0)

示例1：
* 输入："UD"
* 输出：true

示例2：
* 输入："LL"
* 输出：false

## 2. 解题思路
设置两个变量x、y分别表示横纵坐标,初始值都为0

遍历字符串，对于每一个字符，根据其值执行如下操作：
* L x--
* R x++
* U y++
* D y--

最终判断 x,y 是否都为0,如果都是0 则表示回到了原点，否则没有回到原点

## 3. 代码
```go
x,y := 0,0
	for _, s := range moves {
		if s == 'U' {
			y++
		} else if s == 'D' {
			y--
		} else if s == 'L' {
			x--
		} else if s == 'R' {
			x++
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
```
## 4. 复杂度分析
* 时间复杂度：O(N) N 为字符串长度，只需要遍历一次字符串
* 空间复杂度: O(1) 只需要3个变量

---

# Review : []()

---

# Tip : 

## 

---
    
# Share : 20 幻读是什么，幻读有什么问题？—— 极客时间 MySQL实战45讲
建表语句如下
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
如下的sql语句序列，是如何加锁的，加的锁又是如何释放的？
```sql
begin;
select * from t where d=5 for update;
commit;
```
这个语句会命中 d=5 的这一行，即 id=5，因此 select 执行完成后 id=5 这一行会加一个写锁，而且由于两阶段锁协议，这个写锁会在执行 commit 的时候释放。

d 这个字段没有索引，因此此 select 语句会做全表扫描，那么被扫描到的行是否会加锁呢？

在可重复读的隔离级别下

## 幻读是什么？
先假设只给 id=5 这一行加锁，其他行不加锁，那么会怎么样呢？

| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin; <br> select * from t where d=5 for update;/* Q1 */ <br> <font color=#DC143C>(5,5,5) </font>| | |
| T2 | | update set d=5 where id=0; | |
| T3 | select * from t where d=5 for update;/* Q2 */ <br> (0,0,5),(5,5,5) | | |
| T4 | | | insert into t values(1,1,5); |
| T5 | select * from t where d=5 for update;/* Q3 */ <br> (0,0,5),(1,1,5),(5,5,5)
| T6 | commit; | | |

其中 Q3 读到 id=1 这一行，为"幻读"

幻读说明：
1. 在可重复读的隔离级别下，普通的查询是快照读，是不会看到别的事务插入的数据的。因此，幻读在"当前读"下才会出现。
2. session B 的修改结果，被 session A 之后的 select 语句用"当前读"看到，不能称为幻读。幻读仅专指"新插入的行"。

## 幻读有什么问题？
### 语义问题

