# week 30

---

# Algorithm [62. Unique Paths](https://leetcode.com/problems/unique-paths/)
## 1. 问题描述
唯一路径

给定一个 m * n 的棋盘

一个机器人初始在左上角，只能向右或向下移动，其目标是移动到右下角。

共有多少可能的唯一路径？

#### 示例 1:
* 输入 : m=3, n=2
* 输出 : 3
* 解释 : 总共 3 条路
    1. 右 -> 右 -> 下
    2. 右 -> 下 -> 右
    3. 下 -> 右 -> 右
#### 示例 2 : 
* 输入 : m=7,n=3
* 输出 : 28

## 2. 解题思路
dp[i,j] 表示从 (0,0) 走到 (i,j) 的路径数

dp[i,0] = 1, dp[0,j] = 1

dp[i,j] = d[i-1,j] + d[i,j-1]

最终求出 dp[m-1,n-1]

## 3. 代码
```go
func uniquePaths(m int, n int) int {
    var dp [][]int
    dp = make([][]int, m)
    for i := 0; i < m; i++ {
    	dp[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}
```
## 4. 复杂度分析
* 时间复杂度 : O(m*n)
* 空间复杂度 : O(m*n)
---

# Review [Why Coding Your Own Makes You a Better Developer](https://medium.com/better-programming/why-coding-your-own-makes-you-a-better-developer-5c53439c5e4a)
为什么自己写代码会让你成为更好的开发者。

重复造轮子才能有更深的理解。

## 开发人员认为抽象是利索当然的
## 了解你使用的源代码
深入了解源码及其原理，帮助你深入了解如何编程。帮你成为更好的程序员。
## 重新造轮子
自己编码实现一个已有的方案，是最佳的学习方式。

---

# Tip Tomcat 特殊字符报 400 错误
http 请求 tomcat 报 400 错误
```log
org.apache.coyote.http11.Http11Processor.service Error parsing HTTP request header
 Note: further occurrences of HTTP header parsing errors will be logged at DEBUG level.
 java.lang.IllegalArgumentException: Invalid character found in the request target. The valid characters are defined in RFC 7230 and RFC 3986
```
[RFC 3986](https://www.ietf.org/rfc/rfc3986) 规范定义了Url中只允许包含英文字母（a-zA-Z）、数字（0-9）、-_.~4个特殊字符以及所有保留字符 ! * ’ ( ) ; : @ & = + $ , / ? # [ ]

url 中如果传 json，包含 {,} 等特殊字符就会直接报 400

解决此问题需要修改 tomcat 的配置，在 catalina.properties 中添加
```properties
tomcat.util.http.parser.HttpParser.requestTargetAllow=|{}
```
然后重启 tomcat 即可

---
    
# Share 32 为什么还有kill不掉的语句？ —— 极客时间 MySQL实战45讲
MySQL 有两个 kill 命令：
* kill query + 线程id ： 终止正在执行的语句
* kill connection + 线程id ： 断开这个连接，但是要等正在执行的语句执行完成

有时候 kill 之后，用 show processlist 发现是 Killed 状态，而没有退出。

大多数情况下 kill 是直接生效的，比如 kill query + threadId 终止一个执行中的语句。
语句锁等待的时候直接 kill 命令也是有效的。

| session A | session B | session C |
| --- | --- | --- |
| begin;<br/> update t set c=c+1 where id=1; | | |
| | update t set c=c+1 where id=1;<br/>(blocked) | |
| | ERROR 1317 (70100):Query execution was interrupted | kill query thread_id_B; |

C 执行 kill 之后 B 立刻终止

## 收到 kill 之后，线程做了什么
执行 kill query thread_id_B 之后 B 线程做了两件事：
1. 把 B 的运行状态改为 THD::KILL_QUERY
2. 给 B 发一个信号

发信号的目的：
1. 一个语句在执行过程中有多处“埋点”，这些“埋点”的地方判断线程状态，如果发现线程状态是 THD::KILL_QUERY,才开始进入语句终止逻辑
2. 如果处于等待状态，必须是一个可以被唤醒的等待，否则根本不会执行到“埋点”处
3. 语句从开始进入终止逻辑，到终止逻辑执行完成，是有一个过程的

* kill 不掉的例子
设置 innodb_thread_conncurrency=2 

| session A | session B | session C | session D | session E |
| --- | --- | --- | --- | --- |
| select sleep(100) from t; | select sleep(100) from t; | | | |
| | | select * from t;<br/>(blocked) | | |
| | | | kill query C; | |
| | | ERROR 2013(HY000):Lost connection to MySQL server during query | | kill C; |

1. C 执行的时候被堵住2
2. D 的 kill query C 没效果
3. E 执行 kill connection，才断开了 C 的连接
4. 此时在 E 中执行 show processlist
![kill_connection](kill_connection.png)

#### 为什么 kill query 的时候，没有像 update 语句一样退出
等行锁时，使用的是 pthread_cond_timewait 函数，这个等待可以被唤醒 

C 的等待逻辑： 每 10 毫秒判断一下是否可以进入 InnoDB 执行，如果不行，就调用 nanosleep 函数进入 sleep 状态

在循环等待进入 InnoDB 的过程中，无法判断线程的状态

E 执行 kill connection 时：
1. 把 12 号线程状态设置为 KILL_CONNECTION
2. 关掉 12 号线程的网络连接，因为此操作，所以看到 C 收到了断开连接的提示

执行 show processlist 时，如果一个线程状态是 KILL_CONNECTION，就把 Command 列显示成 Killed

只有等满足进入 InnoDB 的条件后，C 继续执行，然后才有可能判断线程状态，再进入终止逻辑阶段。

* 无法kill 的第一类情况：线程没有执行到判断状态的逻辑。IO 压力过大，读写 IO 函数一直没有返回，导致不能及时判断状态。
* 另一种情况是：终止逻辑耗时较长

主要有以下场景：
1. 超大事务被 kill，回滚操作耗时很长
2. 大查询回滚。期间生成较大的临时文件，再加上此时系统 IO 压力较大，需要等待 IO 资源
3. DDL 执行到最后阶段被 kill 。需要删除中间过程中的临时文件，需要 IO 资源

直接在客户端执行 Ctrl + C 也不可以直接终止线程。因为此操作实际上是 MySQL 客户端另外启动一个连接，然后发送一个 kill query 命令

## 两个关于客户端的误解
### 误解一：如果库里的表特别多，连接就会很慢
有些线上库有很多表（有 6 万张表的情况），每次客户端连接会卡在如下界面：
![conn_wait](conn_wait.png)
客户端连接的时候，需要做的事情是 TCP 握手、用户校验、获取权限。这些都与表的数量无关。

当使用默认参数连接时，MySQL 客户端会提供一个本地库名和表名补全的功能。为了实现这个功能，客户端会做如下工作：
1. 执行 show databases;
2. 切到 db1 库，执行 show tables;
3. 把这两个命令的结果用于构建一个本地的哈希表

其中主要时间用在了第 3 步，即本质是客户端慢。

连接时加上 -A 参数，可以关闭此功能，然后客户端就可以快速返回。

如果不用表名的自动补全功能，强烈建议关闭此功能，也可以用 -quick(-q) 参数关闭

-quick 有可能降低服务端的性能

MySQL 客户端接收服务端返回结果的方式有两种：
1. 本地缓存。mysql_store_result 方法
2. 不缓存，读一个处理一个。mysql_use_result 方法

MySQL 客户端默认使用第一种，加上 -q 之后会使用第二种。

因此如果本地处理慢，就会导致服务端发送结果被阻塞，让服务器端变慢。

quick 参数可以达到如下效果：
1. 跳过表名自动补全功能
2. mysql_store_result 需要申请本地内存来缓存查询结果，如果结果太大，会耗费很多本地内存，可能会影响本地机器性能
3. 不会把执行命令记录到本地的命令历史文件

quick 是让客户端变的更快，而不是让服务端变的更快。

## 思考题：如果一个被 killed 的事务一直处于回滚状态，是直接把 MySQL 进程强制重启，还是让它自己执行完呢？
* 答：从恢复的角度，应该让它自己结束
1. 因为重启之后该做的回滚动作还是不能少
2. 如果这个语句占用别的锁，或者占用 IO 资源过多，从而影响到了其它语句的执行。
此时可以先做主备切换，切到新主库提供服务。切换之后别的线程都断开了，自动停止执行。
接下来等它自己都执行完成。此属于减少系统压力，加速终止逻辑。






