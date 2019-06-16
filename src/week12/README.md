#week12

---

## Algorithm [832. Flipping an Image](https://leetcode.com/problems/flipping-an-image/)
### 1. 问题描述
水平翻转图像

给定一个二进制矩阵 A，A 表示图像，水平翻转图像，然后反转它，并返回得到的图像。

水平翻转图像，意思是图像的每一行都是反向的。如 水平翻转 [1,1,0] 的结果是 [0,1,1]

反转图像意思是把其中的 0 都替换成 1， 1 都替换成 0. 例如，反转 [0,1,1] 的结果是 [1,0,0]

示例1：
输入： [[1,1,0],[1,0,1],[0,0,0]]
输出： [[1,0,0],[0,1,0],[1,1,1]]
解释：翻转每一行之后：[[0,1,1],[1,0,1],[0,0,0]]
；然后反转图像：[[1,0,0],[0,1,0],[1,1,1]]

示例2：
输入： [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
输出： [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
解释：翻转每一行之后：[[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
然后反转图像：[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

### 2. 解题思路
对于二维数组A，对于每一个行 A[i] 兑换 A[i][j] 和 A[i][len(A[i]) - j - 1] 的值，并取反

### 3. 代码
```go
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		leni := len(A[i]);
		for j := 0; j < (leni + 1)/2; j++ {
			A[i][j],A[i][leni - j - 1] = A[i][leni - j - 1]^1,A[i][j]^1
		}
	}
	return A
}
```
### 4. 复杂度分析
* 时间复杂度：O(N*M) 其中 N 是行数，M 是列数, 确切的说应该是 N*(M/2) 因为每一行只需要遍历一半
* 空间复杂度：O(1)

---

## Review [1. Spring WebFlux](https://docs.spring.io/spring/docs/current/spring-framework-reference/web-reactive.html#webflux)
## Web on Reactive Stack
Version 5.1.7.RELEASE

这部分文档涵盖了对构建在 Reactive Streams API 上的 Reactive 技术栈的 web 应用的支持，这些应用运行在非阻塞服务器上，如 Netty, Undertow
和 Servlet 3.1+ 容器上。共分为以下几个章节 Spring WebFlux 框架，the reactive WebClient， 对测试的支持和 reactive 库。
Servlet 技术栈的 web 应用，请参见  [Web on Servlet Stack](https://docs.spring.io/spring/docs/current/spring-framework-reference/web.html#spring-web)

## 1. Spring WebFlux
Spring Framwork 中包含的原生的 web 框架 Spring Web MVC 是专门为 Servlet API 和 Servlet 容器构建的。
reactive-stack web 框架 Spring WebFlux 是在 5.0 版本中添加进来的。它是完全非阻塞的，支持响应式流背压
([Reactive Streams](https://www.reactive-streams.org/) back pressure, 在数据流从上游生产者向下游消费者传输的过程中，
上游生产速度大于下游消费速度，导致下游的 Buffer 溢出，这种现象就叫做 Backpressure 出现), 并且在 Netty, Undertow 和 Servlet 3.1+ 容器上运行。

两个 web 框架都反映了其源模块的名称(spring-mvc 和 spring-webflux) 并在 Spring 框架中并存。每个模块都是可选的。
应用可以用其中的一个或者另一个，或者同时用两个。例如，使用 Spring MVC 的控制层，同时使用 reactive WebClient。

### 1.1 概览
Spring WebFlux 为什么会出现？

部分原因是需要一个为阻塞的 web 技术栈，来使用少量的线程处理并发性，并且可以使用更少的硬件资源进行扩展。Servlet 3.1 确实提供了非阻塞 I/O 的 API.
但是，使用它会远离 Servlet API 的其他部分，其中契约是同步的(Filter, Servlet)或阻塞(getParameter, getPart)。这是新的通用 API 作为跨
任何非阻塞运行时的基础的动机。这是很重要的，因为这是建立在异步和非阻塞空间中的服务（例如 Netty）。

另一部分原因是函数式编程。就像在 Java5 中添加了注解(例如 REST controller 或单元测试的注解)创造的机会一样，在 Java8 中添加 lambda 
表达式为 Java 中的函数式 API 创造了机会。这对于非阻塞程序和 continuation-style API (由 CompletableFuture 和 ReactiveX 推广)
来说是一个福音，它允许异步逻辑的声明性组合。在编程模型级别，Java 8 使 Spring WebFlux 能够提供函数式 Web endpoint 以及带注解的 controller。

---

## Tip

### RabbitMQ Exchange 类型
| exchange 类型 | 默认的预先声明的名称 |
| --- | --- |
| direct exchange | 空串和 amq.direct |
| topic exchange | amp.topic |
| fanout exchange | amq.fanout |
| header exchange | amq.match(和 amq.header) |

---
    
## Share
### 14 count(*)这么慢，我该怎么办？ —— 极客时间 MySQL实战45讲
#### count(*) 的实现方式
* MyISAM 引擎把一个表的总行数存储在了磁盘上，因此执行 count(*) 的时候直接返回，效率很高
* InnoDB 引擎，执行 count(*) 时，需要把数据一行一行地把数据从引擎中读出来，然后累计计数

注意：这里是无条件的 count(*) ,如果是有 where 条件的 MyISAM 也不能直接返回

* 在保证逻辑正确的前提下，尽量减少扫描的数据量，是数据库系统设计的通用法则之一。

#### 为什么 InnoDB 不能把总数存起来
因为即使是同一时刻的多个查询，由于多版本并发控制 (MVCC) 的原因，InnoDB 表应该返回多少行也是不确定的。

例如：表 t 有 10000 行数据

|会话A|会话B|会话C|
|---|---|---|
|begin;| | |
|select count(*) from t;| | |
| | |insert into t(插入一行);|
| |begin;| |
| |insert into t(插入一行);| |
|返回10000|select count(*) from t;(返回10002)|select count(*) from t;(返回10001)|

InnoDB 对 count(*) 做了优化：MySQL 优化器会找到最小的那棵树去遍历。

由于 InnoDB 是索引组织表，主键索引树的叶子节点是数据，普通索引的叶子节点是主键值。所以普通索引树比主键索引树小很多。
而对于 count(*) 而言，扫描哪个索引树得到的结果逻辑上都是一样的。

show table status; 命令返回的结果中有一行 TABLE_ROWS 显示当前表有多少行。但是这个是从采样估算得来的，因此不准确。官网文档给的误差可能达到 40% 到 50%

小结：
* MyISAM 表 count(*) 很快，但不支持事务
* show table status 命令返回很快，但不准
* InnoDB 表直接 count(*) 会遍历全表，虽然结果准确，但是会导致性能问题

#### 用缓存系统保存计数（有可能丢数据，而且可能不精确）
用 Redis 记录总行数
##### 丢数据问题
缓存可能会丢失更新。比如，插入一条数据，然后更新 Redis 中的值加 1，然后 Redis 异常重启，重启后你要从存储 Redis 数据的地方把这个值读回来，而刚刚加的 1 丢失了。
##### 逻辑上不精确的问题
假设一个页面要显示总数，而且要显示最近操作的100条记录。不精确的定义如下：
1. 一种是，查询到 100 行结果里面有最新的插入记录，而 Redis 中的计数还没加 1
2. 一种是，查询到 100 行结果里面没有最新插入的记录，而 Redis 中的计数已经加 1 了

#### 在数据库中保存计数
1. InnoDB 支持崩溃恢复不丢数据。所以丢数据的问题可以解决。
2. 可以利用 InnoDB 的事务来解决逻辑不精确问题。将计数加 1 和插入一条数据放到同一个事务中。

#### 不同的 count 的用法
count的不同用法：
* count(*)
* count(主键id)
* count(1)
* count(字段)

count() 的语义： count() 是一个聚合函数，对于返回的结果集，一行行地判断，如果 count() 函数的参数不是 NULL，累计值加 1，否则不加。最后返回累计值。

所以 count(*)、count(主键id)和 count(1) 都表示返回满足条件的结果集的总行数。
而 count(字段) 则表示返回满足条件的数据行里面，参数"字段"不为 NULL 的总个数。

记住以下原则：
1. server 层要什么就给什么
2. InnoDB 只给必要字段
3. 现在的优化器只优化了 count(*) 的语义为"取行数",其他"显而易见"的优化并没有做

* count(主键id) : InnoDB 会遍历整张表，把每一行的 id 值都取出来，返回给 server 层。 server 层拿到 id 后，判断是不可能为空的，就按行累加。
* count(1): InnoDB 遍历整张表，但不取值。 server 层对于返回的每一行，放一个数字"1"进去，判断是不可能为空的，按行累加。
单看这两个用法的差别，count(1) 的性能优于 count(主键id)。因为从引擎返回 id 会涉及到解析数据行，以及拷贝字段值的操作
* count(字段):
    1. 如果字段定义为 not null，一行行地从记录里读出这个字段，判断不能为 null，按行累加
    2. 如果字段允许为 null，在执行的时候，判断有可能是 null，还要把值取出来再判断一下，不是 null 才累加
* count(*): 并不会把全部字段的值取出来，而是专门做了优化，不取值。count(*) 肯定不是 null，按行累加

结论：按照效率排序： count(字段) < count(主键id) < count(1) ≈ count(*)

所以建议使用 count(*)

#### 问题
* 将计数保存在数据库中，用事务来保证精确性。在事务中先插入行和先插入计数对逻辑结果都没有影响。那么从并发系统性能的角度考虑应该先插入行还是计数呢？

答：应该先插入行，再插入计数。

因为更新计数表涉及到行锁竞争，先插入行再更新计数能最大程度减少事务之间的锁等待，提升并发度。
