# week38

---

# Algorithm [97. Interleaving String](https://leetcode.com/problems/interleaving-string/)
## 1. 问题描述
交叉字符串

给定字符串 s1,s2,s3 , 确定 s3 是否是由 s1 和 s2 交叉组成
#### 示例 1 :
* 输入 : s1="aabcc" , s2="dbbca", s3="aadbbcbcac"
* 输出 : true
#### 示例 2 :
* 输入 : s1="aabcc", s2="dbbca", s3="aadbbbaccc"
* 输出 : false


## 2. 解题思路

## 3. 代码

## 4. 复杂度分析

---

# Review []()

---

# Tip
 

---
    
# Share 40 | insert语句的锁为什么这么多？ —— 极客时间 MySQL实战45讲
## insert ... select 语句
```roomsql

CREATE TABLE `t` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t values(null, 1,1);
insert into t values(null, 2,2);
insert into t values(null, 3,3);
insert into t values(null, 4,4);

create table t2 like t
```
可重复读隔离级别下，binlog_format=statement，执行
```roomsql
insert into t2(c,d) select c,d from t;
```
执行这个语句时，是否需要对 t 的所有行加间隙锁呢？

| session A | session B |
| --- | --- |
| insert into t values(-1,-1,-1); | insert into t2(c,d) select c,d from t; |

如果 B 先执行，则对 t 主键索引加了 (-∞,1] 这个 next-key lock ，会在语句执行完成后，才允许 A 的 insert 语句执行

如果没有锁，就可能出现 B 的 insert 语句先执行，但写入 binlog 的情况是
```roomsql
insert into t values(-1,-1,-1);
insert into t2(c,d) select c,d from t;
```
这个语句在备库执行，就会把 id=-1 这一行也写到 t2 中，出现主备不一致

## insert 循环写入
执行 insert ... select 的时候，对目标表也不是锁全表，而是只锁住需要访问的资源

往表 t2 中插入一行数据，这一行的 c 值是表 t 中 c 值的最大值加 1
```roomsql
insert into t2(c,d)  (select c+1, d from t force index(c) order by c desc limit 1);
```
这个语句的加锁范围，就是表 t 索引 c 上的 (3,4] 和 (4,supremum] 这两个 next-key lock,以及主键索引上 id=4 这一行

执行流程：从表 t 中按照索引 c 倒序，扫描第一行，拿到结果写入到表 t2 中。
因此整个语句的扫描行数是 1
![slow_log](slow_log.png)
慢查询日志，Rows_examined=1 表示扫描行数为 1

如果要把这条数据插入到 t
```roomsql
insert into t(c,d)  (select c+1, d from t force index(c) order by c desc limit 1);
```
慢查询日志
![slow_log_t](slow_log_t.png)
explain 结果
![explain](explain.png)
Using temporary 表示使用了临时表

![Innodb_rows_read](Innodb_rows_read.png)
这个语句执行前后，Innodb_rows_read 的值增加了 4。因为默认临时表是 Memory 引擎，所以这 4 行查的都是表 t，也就是说对 t 做了全表扫描

整个执行流程:
1. 创建临时表，表里有两个字段 c 和 d
2. 按照索引 c 扫描表 t，依次取 c=4,3,2,1 ， 然后回表，读到 c 和 d 的值写入临时表。这时 Rows_examined=4
3. 由于语义里面有个 limit 1，所以只取了临时表的第一行，再插入到表 t，这时，Rows_examined=5

这个语句会在 t 上做全表扫描，并且会给索引 c 上的所有间隙都加上共享的 next-key lock。所以，这个语句执行期间，其他事务不能在这个表上
插入数据。

使用临时表示因为，这类一边遍历数据，一边更新数据的情况，如果读出来的数据直接写回原表，就可能在遍历过程中，读到刚刚插入的记录，新插入的
记录如果参与计算逻辑，就跟语义不符。

由于实现上这个语句没有在子查询中就直接使用 limit 1，从而导致了这个语句的执行需要遍历整个 t 表。优化
```roomsql
create temporary table temp_t(c int,d int) engine=memory;
insert into temp_t  (select c+1, d from t force index(c) order by c desc limit 1);
insert into t select * from temp_t;
drop table temp_t;
```
## insert 唯一键冲突
| session A | session B |
| --- | --- |
| insert into t values(10,10,10);<br/><br/>begin;<br/>insert into t values(11,10,10);<br/>(Duplicate entry '10' for key 'c') | |
| | insert into t values(12,9,9);<br/>(blocked) |

可重复隔离级别，B 进入了锁等待

A 持有 c 上的 (5,10] 共享 next-key lock （读锁）

死锁场景

| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin;<br/>insert into t values(null,5,5); | | |
| T2 | | insert into t values(null,5,5);| insert into t vlues(null,5,5); |
| T3 | rollback; | | (Deadlock found) |

A 执行 rollback 回滚时，C 几乎同时发现了死锁并返回

死锁产生的逻辑：
1. 在 T1 时刻，启动 A，并执行 insert ，此时在索引 c 的 c=5 上加了记录锁。这个索引是唯一索引，因此退化为记录锁
2. 在 T2 时刻， B 执行 insert，发现唯一键冲突，加上读锁；同样，C 也在索引 c 的 c=5 这一纪录上，加了读锁
3. T3 时刻， A 回滚，B 和 C 都试图继续执行插入操作，都要加上锁。两个 session 都要等待对方的行锁，所以就出现了死锁

状态变化图
![status_change](status_change.jpg)
## insert into ... on duplicate key update
上面的例子是主键冲突后直接报错，如果改写成
```roomsql
insert into t values(11,10,10) on duplicate key update d=100; 
```
就会给索引 c 上 (5,10] 加上一个排他的 next-key lock(写锁)

insert into ... on duplicate key update 这个语义的逻辑是，插入一行数据，如果碰到唯一键约束，就执行后面的更新语句。

如果有多个列违反了唯一键约束，就会按照索引的顺序，修改跟第一个索引冲突的行。

现在 t 中有 (1,1,1) 和 (2,2,2) 下面语句的执行过程
![key_conflict](key_conflict.png)
主键索引先判断，与 id=2 这一行冲突，所以修改 id=2 的行

affected rows 返回是 2，实际，真正更新的只有一行，只是在代码上实现了， insert 和 update 都认为自己成功了，update 计数加 1，
insert 计数也加了 1.

