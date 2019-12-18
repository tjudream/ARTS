# week39

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
 

---
    
# Share 41 | 怎么最快地复制一张表？ —— 极客时间 MySQL实战45讲
在两张表中拷贝数据。如果可以控制对源表的扫描行数和加锁范围很小的话，可以简单地使用 insert ... select 即可。

为了避免对源表加读锁，更稳妥的方案是先将数据写到外部文件，然后再写回目标表。

创建表，并插入数据
```roomsql
create database db1;
use db1;

create table t(id int primary key, a int, b int, index(a))engine=innodb;
delimiter ;;
  create procedure idata()
  begin
    declare i int;
    set i=1;
    while(i<=1000)do
      insert into t values(i,i,i);
      set i=i+1;
    end while;
  end;;
delimiter ;
call idata();

create database db2;
create table db2.t like db1.t
```

我们要把 db1.t 中 a > 900 的数据导出，并插入到 db2.t 中

## mysqldump 方法
用 mysqldump 导出一组 insert 语句,放到临时文件中
```roomsql
mysqldump -h$host -P$port -u$user --add-locks=0 --no-create-info --single-transaction  --set-gtid-purged=OFF db1 t --where="a>900" --result-file=/client_tmp/t.sql
```
参数：
1. --single-transaction 在导出数据时不对 db1.t 加锁，而是使用 start transaction with consistent snapshot 方法
2. --add-locks=0 表示输出的文件结果中，不增加 "LOCK TABLES t WRITE;"
3. --no-create-info 不需要导出表结构
4. --set-gtid-purged=off 不输出跟 GTID 相关的信息
5. --result-file 指定了输出文件的路径，其中 client 表示生成的文件是在客户端机器上的

![sql_file](sql_file.png)

一个 insert 语句中包含多个 value，这在后续执行时速度会更快

如果希望一个 insert 一行，可以使用参数 --skip-extended-insert

导入数据
```roomsql
mysql -h127.0.0.1 -P13000  -uroot db2 -e "source /client_tmp/t.sql"
```
source 是一个客户端命令

MySQL 客户端的执行流程：
1. 打开文件，默认以分号为结尾读取一条条 SQL 语句
2. 将 SQL 语句发送到服务器端

服务端执行的并不是这个 SQL 文件，而是 insert 语句，所以 slow log 和 binlog 记录的都是被真正执行的 insert 语句

## 导出 CSV 文件
```roomsql
select * from db1.t where a>900 into outfile '/server_tmp/t.csv';
```
需要注意：
1. 这条语句会将结果保存在服务器端。如果执行的客户端和 MySQL 
