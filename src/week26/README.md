# week26

---

# Algorithm [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)
## 1. 问题描述
正则表达式匹配

输入一个字符串 s 和一个模式匹配串 p，实现正则表达式中的 '.' 和 '*'
* '.' 匹配任意一个字符
* '*' 匹配 0 或多个前一个字符

需要匹配整个字符串 s

* s 可能是空串，或者由 a-z 的小写字母组成
* p 可能是空串，或者由 a-z . * 组成

示例 1 :
* 输入:
    * s = "aa"
    * p = "a"
* 输出: false
* 解释: "a" 并不匹配整个字符串 "aa"

示例 2 :
* 输入 :
    * s = "aa"
    * p = "a*"
* 输出 : true
* 解释 : * 匹配一个 a

示例 3 :
* 输入 :
    * s = "ab"
    * p = ".*"
* 输出 : true
* 解释 : .* 表示 0 到多个（*）任意字符（.）

示例 4 :
* 输入 :
    * s = "aab"
    * p = "c*a*b"
* 输出 : true
* 解释 : 第一个 * 表示 c 重复 0 次，第二个 * 表示 a 重复 1 次

示例 5 :
* 输入 :
    * s = "mississippi"
    * p = "mis*is*p"
* 输出 : false

## 2. 解题思路
动态规划

设二维 boolean 型数组 dp[i][j] 表示 p[0:j] 是否匹配 s[0:i]
1. 如果 p[j] == s[i], 则 dp[i][j] = dp[i-1][j-1]
2. 如果 p[j] == '.' , 则 dp[i][j] = dp[i-1][j-1]
3. 如果 p[j] == '*' , 则
    * (1) 如果 p[j-1] != s[i], 则 dp[i][j] = dp[i][j-2]  重复 0 次
    * (2) 如果 p[j-1] == s[i] 或者 p[j-1] == '.'，那么
        * i. dp[i][j] = dp[i-1][j] 重复多次
        * 或 ii. dp[i][j] = dp[i][j-1] 重复 1 次
        * 或 iii. dp[i][j] = dp[i][j-2] 重复 0 次
## 3. 代码

## 4. 复杂度分析

---

# Review []()

---

# Tip 

## 

---
    
# Share 28 读写分离有哪些坑？—— 极客时间 MySQL实战45讲
![split_readwrite](split_readwrite.png)

读写分离结构，客户端主动做负载均衡

![proxy_readwrite](proxy_readwrite.png)

带有 proxy 层的读写分离结构

各自特点：
1. client 直连方案。性能优于 proxy，且整体架构简单，便于排查问题。但是出现准备切换、库迁移时，client 都是有感知的，需要进行调整。
此架构下一般会采用 zookeeper 管理后端组件，便于 client 调整。
2. proxy 架构。对 client 友好，不需要关注后端细节，连接维护、后端信息维护都由 proxy 完成。
但是对后端维护团队要对较高，proxy 需要高可用架构，proxy 架构相对复杂。

我们将 “在从库上会读到系统的一个过期状态”的现象，称为“过期读”

处理过期读的方案：
* 强制走主库
* sleep 方案
* 判断主备无延迟方案
* 配合 semi-sync 方案
* 等主库位点方案
* 等 GTID 方案
## 强制走主库
将查询请求分类:
1. 对于必须要拿到最新结果的请求，强制走主库
2. 允许读到旧数据的请求，发从库

此方案对于一些金融类的业务可能无效，因为金融类的业务有可能要求，所有查询都不能是过期读

## Sleep 方案
主库更新后，读从库之前先 sleep 一下
```sql
select sleep(1)
```
例如卖家发布商品，商品发布后，用 Ajax 直接把客户端输入的内容作为“新的商品”显示在页面上，
而不是去查数据库。等卖家刷新页面的时候再去查库，这时候相当于 sleep 了一段时间。

此方案不精确:
1. 如果本来 0.5 秒就能从从库上拿到结果，也要等 1 秒
2. 如果延迟超过 1 秒，就会出现过期读

## 判断主备无延迟方案
### 1. 通过 seconds_behind_master 判断
通过 show slave status 结果里的 seconds_behind_master 是否等于 0 判断。
如果不是 0 ，需要等到 0 再查询。

seconds_behind_master 单位是秒，如果精度要求较高，可以通过对比点位和 GTID 的方式
确保主备无延迟。
### 2. 对比点位
* a. Master_Log_File 和 Read_Master_Log_Pos, 表示的是读到的主库的最新点位
* b. Relay_Master_Log_File 和 Exec_Master_Log_Pos, 表示的是备库执行的最新点位

如果 a、b 的两组值一样，则表示无延迟 
### 3. 对比 GTID 集合
* Auto_Position=1 , 表示主备关系使用了 GTID
* Retrieved_Gtid_Set, 是备库收到的所有日志的 GTID 集合
* Executed_Gtid_Set, 是备库所有已经执行完成的 GTID 集合

如果这两个集合相同，表示备库已经同步完成

在执行查询请求之前，先判断从库是否同步完成的方法，相比于 sleep 方案，准确度确实提升了不少，但是还没有达到“精确”。

一个事务 binlog 在主备之间的状态：
1. 主库执行完成，写入 binlog，并反馈给 client
2. binlog 被从主库发送给备库，备库收到
3. 在备库执行 binlog 完成

我们上边的逻辑判断的是“备库收到的日志都执行完了”。但是，还有一部分日志，处于 client 已经
收到提交确认，而备库还没有收到日志的状态。

![trx_no_receive](trx_no_receive.png)

主库上执行完成了 trx1、2、3，
1. trx1 和 trx2 已经传到从库，并执行完成
2. trx3 在主库完成，并已经回复给 client，但还没有传到从库

如果此时在从库 B 上执行查询，从库认为已经没有主从延迟，但是查不到 trx3。即出现了过期读。

此问题需要通过 semi-sync 解决。
## 配合 semi-sync
半同步复制 : semi-sync replication

semi-sync 设计：
1. 事务提交的时候，主库把 binlog 发给从库
2. 从库收到 binlog 以后，发回给主库一个 ack，表示收到
3. 主库收到这个 ack 以后，才能给客户端返回“事务完成”的确认

如果主库掉电的时候，有些 binlog 还来不及发给从库，会不会导致系统数据丢失？

答案是，如果使用的是普通的异步复制，就可能会丢，但 semi-sync 就可以解决此问题。

但是 semi-sync + 位点判断的方案，只对一主一备场景成立。在一主多从的场景中，主库只要等到
一个从库的 ack ，就开始给 client 返回确认。此时，在从库上执行查询请求，有两种情况：
1. 查询落在了这个响应了 ack 的从库上，能够确保读到最新数据
2. 如果查询落到了其他从库上，可能会产生过期读

另外还有一个问题，在业务高峰期，主库位点或 GTID 集合更新很快，那么上面的两个位点等值判断
会一直不成立，导致从库无法响应查询请求，或请求超时。

实际上，当发起一个查询请求后，我们要得到准确的结果，其实并不需要等到“主备完全同步”

![master_slave_delay](master_slave_delay.png)
上图就是一个等待点位方案的一个 bad case。从状态1到状态4，一直处于延迟一个事务的状态。

如果必须等到无延迟才能查询，select 语句一直到状态4都不能被执行。

但是实际上我们只需要等到 trx1 更新完成后就可以执行 select 了。即状态3时，就可以得到正确的结果了。

semi-sync 存在的两个问题：
1. 一主多从时，在某些从库执行查询会存在过期读的现象
2. 在持续延迟的情况下，可能出现过度等待的现象

## 等待主库位点方案
```sql
select master_pos_wait(file, pos[, timeout]);
```
1. 此命令是在从库上执行的
2. 参数 file 和 pos 指的是主库上的文件名和位置
3. timeout 可选，设置为正整数 N 表示这个函数最多等 N 秒

返回结果是 M 表示从命令执行，到应用完 file 和 pos 表示的 binlog 位置，执行了多少个事务

异常返回结果：
1. 如果执行期间，备库同步线程发生异常，返回 NULL
2. 如果等待超过 N 秒，返回 -1
3. 如果刚开始执行的时候，发现已经执行过这个位置，返回 0

![master_pos_wait](master_pos_wait.png)
先执行 trx1 再执行一个查询的逻辑：

1. trx1 事务更新完成后，马上执行 show master status 得到当前主库执行到的 File 和 Position
2. 选定一个从库执行查询语句
3. 在从库上执行 select master_pos_wait(File, Position, 1)
4. 如果返回值是 >=0 的正整数，则在这个从库执行查询语句
5. 否则，到主库执行查询语句

这条 select 最多等待 1 秒，如果 1 秒内返回值大于 0，则在从库执行 select，否则退化为在主库执行 select
## GTID
```sql
 select wait_for_executed_gtid_set(gtid_set, 1);
```
1. 等待，直到这个库执行的事务中包含传入的 gtid_set, 返回 0
2. 超时返回 1

MySQL 5.7.6 版本开始，允许在执行完更新类事务后，把这个事务的 GTID 返回给客户端

GTID 的流程：
1. trx1 事务更新完成后，从返回包直接获取这个事务的 GTID，记为 gtid1
2. 选定一个从库执行查询语句
3. 在从库上执行 select wait_for_executed_gtid_set(gtid1, 1)
4. 如果返回 0，则在这个从库上执行查询语句
5. 否则，到主库执行查询语句

如果超时后直接查主库，需要业务开发同学来做限流考虑。

流程图如下
![wait_for_executed_gtid_set](wait_for_executed_gtid_set.png)

要在返回包中带上 GTID，需要设置参数 session_track_gtids=OWN_GTID，然后通过 API
接口 mysql_session_track_get_first 从返回包中解析出 GTID

MySQL 提供的 [API](https://dev.mysql.com/doc/refman/5.7/en/c-api-functions.html)

修改 MySQL 客户端
```c++
const char *data;
size_t length;
if (mysql_session_track_get_first(&mysql, SESSION_TRACK_GTIDS, &data, &length) == 0)
{
    sprintf(&buff[strlen(buff)], ", GTID: %s", data);
}
```

![gtid_result](gtid_result.png)

## 思考题
采用 GTID 方案，现在需要对主库的一张大表做 DDL，可能会出现什么情况呢？为了避免这种情况要怎么做？
* 回答：
假设这条语句在主库需要执行 10 分钟，提交后传到备库也要执行 10 分钟。
那么，在主库 DDL 之后再提交的事务的 GTID，去备库查的时候，就会等 10 分钟才出现。

这个读写分离机制在 10 分钟内都会超时，然后都走主库。

这种预期的操作，应该在业务低峰期进行，确保主库能够支撑所有业务查询，然后把读请求都切换到主库，
再在主库上做 DDL。等备库延迟追上后，再把读请求切回备库。

也可以备库先做，再切主库