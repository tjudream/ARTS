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
### 假设1: 先假设只给 id=5 这一行加锁，其他行不加锁，那么会怎么样呢？

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
### 首先，语义问题
session A 在 T1 时刻声明，"我要把所有 d=5 的行锁住，不准别的事务进行读写操作"。而实际上，这个语义被破坏了。

| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin; <br> select * from t where d=5 for update; /* Q1 */ | | |
| T2 | | update t set d=5 where id=0; <br> update t set c=5 where id=0; | |
| T3 | select * from t where d=5 for update; /* Q2 */ | | |
| T4 | | | insert into t values(1,1,5); <br> update t set c=5 where id=1; |
| T5 | select * from t where d=5 for update; /* Q3 */ | | |
| T6 | commit; | | |

session B 的第二条语句 update t set c=5 where id=0 的语义是 "我把id=0, d=5 这一行的 c 改为 5"

由于 T1 时，A 只给 id=5 这一行加锁，id=0 这一行并没有加锁。因此 B 的 T2 时刻可以执行这两条 update 。这就破坏了 A 的 Q1 中锁住 d=5 的行的语义。

### 其次，数据一致性问题
| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin; <br> select * from t where d=5 for update; /* Q1 */ <br> update t set d=100 where d=5; | | |
| T2 | | update t set d=5 where id=0; <br> update t set c=5 where id=0; | |
| T3 | select * from t where d=5 for update; /* Q2 */ | | |
| T4 | | | insert into t values(1,1,5); <br> update t set c=5 where id=1; |
| T5 | select * from t where d=5 for update; /* Q3 */ | | |
| T6 | commit; | | |

分析此表
1. T1 之后，id=5 这一行变成了 (5,5,100), 此结果在 T6 正式提交
2. T2 之后，id=0 这一行变成了 (0,5,5)
3. T4 之后，表里多了一行 (1,5,5)
4. 其他行保持不变

binlog 里的内容
```sql
/* T2 B 写入两条语句 */
update t set d=5 where id=0; /*(0,0,5)*/
update t set c=5 where id=0; /*(0,5,5)*/
/* T4 C 写入两条 */
insert into t values(1,1,5); /*(1,1,5)*/
update t set c=5 where id=1; /*(1,5,5)*/
/* T6 A 写入 */
update t set d=100 where d=5;/* 所有 d=5 的行，d 改成 100*/
```
如果拿此 binlog 去备库执行，或者恢复数据，这三行都会变成 (0,5,100),(1,5,100),(5,5,100)。

id=0 和 id=1 这两行发生了数据不一致。

主要原因是我们假设 select * from t where d=5 for update 这条语句只给 d=5 这一行加锁，也就是只给 id=5 这一行加锁导致的。

因此假设不合理

### 假设2: 把扫描过程中碰到的行都加上写锁
| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin; <br> select * from t where d=5 for update; /* Q1 */ <br> update set d=100 where d=5; | | |
| T2 | | update t set d=5 where id=0;(blocked) <br> update t set c=5 where id=0; | |
| T3 | select * from t where d=5 for update; /* Q2 */ | | |
| T4 | | | insert into t values(1,1,5); <br> update t set c=5 where id=1; |
| T5 | select * from t where d=5 for update; /* Q3 */ | | |
| T6 | commit; | | |

A 把所有行都加了写锁，所以 B 执行第一条 update 的时候就被锁住了。需要等到 T6 时 A 提交之后 B 才能继续执行。

binlog如下
```sql
/* C */
insert into t values(1,1,5); /*(1,1,5)*/
update t set c=5 where id=1; /*(1,5,5)*/
/* A T6*/
update t set d=100 where d=5;/* 所有 d=5 的行，d 改成 100*/
/* B */
update t set d=5 where id=0; /*(0,0,5)*/
update t set c=5 where id=0; /*(0,5,5)*/
```
应用 binlog 的结果为 (0,5,5),(1,5,100),(5,5,100)

也就是说 id=1 这一行数据不一致。因为 C 的执行结果应该是 (1,5,5) 而不是 (1,5,100)

因此，假设2 也不成立

也就是说，即使把所有的记录都加上锁，还是阻止不了新插入的记录，这也是为什么"幻读"会被单独拿出来解决的原因。

## 如何解决幻读？
为了解决幻读，InnoDB 引入了间隙锁(Gap Lock)

间隙锁，锁的是两个值之间的空隙。

表 t 初始化插入 6 条记录，产生了 7 个间隙：(−∞,0),(0,5),(5,10),(10,15),(15,20),(20,25),(25,+∞)

当你执行 select * from t where d=5 for update 时，不止给数据库中已有的 6 条记录加上了行锁，还同时给 7 个间隙加上了间隙锁。
这样就确保无法再插入新的记录了。

行锁：

| | 读锁 | 写锁 |
| --- | --- | --- |
| 读锁 | 兼容 | 冲突 |
| 写锁 | 冲突 | 冲突 |

跟行锁冲突的是另一个行锁，但是跟间隙锁冲突的，是"往这个间隙中间插入一条记录"这个操作。间隙锁之间不存在冲突。

| session A | session B |
| --- | --- |
| begin; <br> select * from t where c=7 lock in share mode; | |
| | begin; <br> select * from t where c=7 for update; | 

其中 B 不会被锁住。因为 t 中并没有 c=7 这条记录，因此 A 加的是间隙锁 (5,10),而 B 也是加的间隙锁 (5,10),它们之间并不冲突，
它们有共同的目标，保护这个间隙，不允许插入新值。

间隙锁和行锁合称 next-key lock，每个 next-key lock 是前开后闭区间。

表 t 初始化之后，执行 select * from t for update 把整个表的所有记录锁起来，就形成了 7 个 next-key lock,分别是 
(-∞,0],(0,5],(5,10],(10,15],(15,20],(20,25],(25,+supremum]

由于 +∞ 是开区间，所以，实现上，InnoDB 给每个索引加了一个不存在的最大值 supremum，这样才符合前开后闭

间隙锁和 next-key lock 的引入，帮我们解决了幻读的问题，同时也带来了一些"困扰"。

需求：任意锁住一行，如果这一行不存在就插入，如果存在就更新它的数据。
```sql
begin;
select * from t where id=N for update;

/* 如果行不存在 */
insert into t values(N,N,N);
/* 如果行存在 */
update t set d=N set id=N;

commit;
```
如果有多个唯一键的时候，insert ... on duplicate key update 不能满足此需求。

这个逻辑一旦有并发，就会死锁。

| session A | session B |
| --- | --- |
| begin; <br> select * from t where id=9 for update; | | 
| | begin; select * from t where id=9 for update; |
| | insert into t values(9,9,9); <br> (blocked) |
| insert into t values(9,9,9); <br> (ERROR 1213(400001);Deadlock found) | |

分析：
1. A 执行 select ... for update, 由于 id=9 这一行不存在，因此会加间隙锁 (5,10)
2. B 执行 select ... for update, 同样加间隙锁 (5,10), 间隙锁之间不冲突
3. B 试图插入 (9,9,9), 被 A 的间隙锁锁住，等待 A
4. A 试图插入 (9,9,9), 被 B 的间隙锁锁住，等待 B ，形成死锁

间隙锁的引入，可能导致同样的语句锁住更大的范围，这其实是影响了并发度的。

间隙锁是在可重复读隔离级别下才会生效。如果改为读提交的话，就没有间隙锁了。
但同时，你要解决可能出现的数据和日志不一致的问题，需要把 binlog 格式设置为 row，这也是现在不少公司使用的配置组合。

### 思考题

| session A | session B | session C |
| begin; <br> select * from t where c>=15 and c <=20 order by c desc for update; | | |
| | insert into t values(11,11,11); | | 
| | | insert into t values(6,6,6); |

B,C 都会进入锁等待，请分析原因。

答案详见下篇文章
