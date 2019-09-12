# week24

---

# Algorithm []()
## 1. 问题描述

## 2. 解题思路

## 3. 代码

## 4. 复杂度分析

---

# Review []()

---

# Tip

## 

---
    
# Share 27 主库出问题了，从库怎么办？—— 极客时间 MySQL实战45讲
![master_slave](master_slave.png)
一主多从基本结构

A和A'互为主备，从库 B、C、D 指向 A

![active_standby_switch](active_standby_switch.png)
一主多从基本结构 —— 主备切换

## 基于位点的主备切换
把 B 设置成节点 A' 的从库的时候，需要执行一条 change master 命令：
```sql
CHANGE MASTER TO 
MASTER_HOST=$host_name 
MASTER_PORT=$port 
MASTER_USER=$user_name 
MASTER_PASSWORD=$password 
MASTER_LOG_FILE=$master_log_name 
MASTER_LOG_POS=$master_log_pos  
```
6个参数：
* MASTER_HOST, MASTER_PORT, MASTER_USER, MASTER_PASSWORD 表示 A’的 IP，端口，用户名，密码
* MASTER_LOG_FILE, MASTER_LOG_POS 是同步点位，即主库对应的文件名和日志偏移量

以前 B 拥有 A 的同步点位，而没有 A' 的，所以切换时，首先需要找到切换点位。

但是这个点位只能取到一个大概的位置，很难精确取到

因为不能丢数据，所以找点位的时候要“稍微往前”，然后再跳过那些从库 B 上已经执行过的事务。

一种取同步点位的方法：
1. 等待新主库 A' 把中转日志(relay log)全部同步完成
2. 在 A' 上执行 show master status 命令，得到当前 A' 上最新的 File 和 Position
3. 取原主库 A 故障的时刻 T
4. 用 mysqlbinlog 工具解析 A' 的 File ，得到 T 时刻的点位

```sql
mysqlbinlog File --stop-datetime=T --start-datetime=T
```
![binlog_result](binlog_result.png)

end_log_pos = 123 表示 A' 在 T 时刻写入新的 binlog 的位置。
我们将 123 作为 $master_log_pos ,用在节点 B 的 change master 命令里。

为什么是不精确的？假设 T 时刻，A 已经执行完一个 insert 插入一行数据 R,并且已经将 binlog 传给了 A' 和 B，
然后传完瞬间 A 掉电。

此时的系统状态：
1. 在从库 B 上，由于同步了 binlog，R 这一行已经存在
2. 在新的主库 A' 上，R 这一行也已经存在，日志是写在 123 之后的
3. 我们在从库 B 上执行 change master 命令，指向 A' 的 File 文件的 123 位置，
就会把插入 R 这一行数据的 binlog 又同步到从库 B 去执行

这时候， B 同步线程报 Duplicate entry 'id_of_R' for key 'PRIMARY' 错误，提示主键冲突，然后停止

通常情况下，我们再切换任务的时候，要先主动跳过这些错误，有两种方法：
1. 主动跳过一个事务 
```sql 
set global sql_slave_skip_counter=1;
start slave;
```
2. 通过设置 slave_skip_errors 直接跳过指定错误
    * 1062 插入数据时唯一键冲突
    * 1032 删除数据时找不到行

slave_skip_errors="1062,1032"

主备同步关系建立完成后，并稳定执行一段时间之后，需要把这个参数设置为空

## GTID
MySQL 5.6 引入了 GTID (Global Transaction Identifier) 全局事务 ID，是一个事务提交的时候生成的，是这个事务的唯一标识。
由两部分组成：
```sql
GTID=server_uuid:gno
GTID=source_id:transaction_id  #MySQL 官方文档里的定义
```
* server_uuid 是一个实例第一次启动时自动生成的，是一个全局唯一的值
* gno 是一个整数，初始值是 1，每次提交事务的时候分配给这个事务，并加 1

启动 GTID 需要设置参数：
* gtid_mode=on
* enforce_gtid_consistency=on





