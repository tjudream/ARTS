#week9

---

## Algorithm [804. Unique Morse Code Words](https://leetcode.com/problems/unique-morse-code-words/)
### 1. 问题描述
26个字母中每个字母都对应一个 morse 码。题目中给出了每个字母对应的 morse 码。
多个单词可能对应同一个 morse 码。

要求输入一个单词数组 words ，计算该数组中的单词一共对应几个不同的 morse 码。
### 2. 解题思路
将 words 数组中的每个单词都转换成 morse 码，然后用一个结合存储这些 morse 码，最后返回集合中元素的个数。
### 3. 代码
```go
var dict = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
func uniqueMorseRepresentations(words []string) int {
	morseMap := make(map[string]string)
	for _,word := range words {
		morse := ""
		for i := 0; i < len(word); i++ {
			morse = morse + dict[word[i] - 'a']
		}
		morseMap[morse] = word
	}
	return len(morseMap)
}
```
### 4. 复杂度分析
* 时间复杂度: O(S), 其中 S 为单词数组 words 中所有单词中字数个数的总和。
* 空间复杂度: O(S), 每个字母转成 morse 码，最多4个字符，所以集合 set 中最多占用 O(4S) 个空间
---

## Review [10 Common Software Architectural Patterns in a nutshell](https://towardsdatascience.com/10-common-software-architectural-patterns-in-a-nutshell-a0b47a1e9013)
10种常见的软件架构模式的概括
#### 架构模式的定义
架构模式是针对给定上下文的软件架构中的常见问题的通用的、可重用的解决方案。架构模式类似于软件设计模式，但是范围更广

本文将介绍如下10个架构模式的用法和优缺点:
1. Layered pattern 分层模式
2. Client-server pattern 客户端-服务器模式
3. Master-slave pattern 主从模式
4. Pipe-filter pattern 管道过滤模式
5. Broker pattern 代理模式
6. Peer-to-peer pattern 端到端模式
7. Event-bus pattern 事件总线模式
8. Model-view-controller pattern MVC模式
9. Blackboard pattern 黑板模式
10. Interpreter pattern 解释器模式

#### 1. 分层模式
此模式用于构建可以分解为子任务组的程序，每个子任务组都处于一个特定的抽象级别。每一层都对其上一层提供服务。

最常见的4层模型：
* 表现层（UI 层）
* 应用层(服务层)
* 业务逻辑层（领域层）
* 数据访问层（持久层）

使用场景：
* 一般的桌面应用
* 电子商务应用

![layered_pattern](layered_pattern.png)
#### 2. 客户端-服务器模式
此模式包含2个部分，一个服务器端和多个客户端。服务器组件为多个客户端组件提供服务。客户端向服务器端发送请求，服务器端提供
相应的服务给客户端。此外，服务器端持续监听客户端的请求。

使用场景：
* 在线服务，如：邮件，文档共享和银行等

![client-server_pattern](client-server_pattern.png)
#### 3. 主从模式
此模式包含2个部分，主服务和从服务。主服务将工作分发给各个从服务，然后将结果汇总。

使用场景：
* 数据库中的数据复制，主数据库中的数据是权威，所有的从数据库都要从主库中同步数据。
* 计算机总线连接的外围设备（主从驱动）

![master-slave_pattern](master-slave_pattern.png)
#### 4. 管道-过滤器模式
此模式可以用于构建生产和处理流数据的系统。每个处理步骤都包含在过滤器组件中。要处理的数据是通过管道传递的。这些管道可以
用于缓冲和同步。

使用场景：
* 编译器。通过连续的过滤器来执行词法分析、解析、语义分析和代码生成。
* 生物信息学的工作流。

![pipe-filter_pattern](pipe-filter_pattern.png)
#### 5. 代理模式
此模式用于构建具有解耦组件的分布式系统。这些组件之间可以通过远程服务调用进行互相通讯。
代理组件负责协调各个组件之间的通讯。

服务发布者将它们的能力（服务和特征）发布给代理。客户端向代理发送请求，然后代理将请求发送给相应的服务。

使用场景：
* 消息代理软件，如 Apache ActiveMQ, Apache Kafka, RabbitMQ 和 JBoss Messaging

![broker_pattern](broker_pattern.png)
#### 6. 端到端模式
此模式中，单个组件被称为端。端可以作为一个客户端，从其它端请求服务，也可以作为服务端，为其它端提供服务。
一个端可以同时既作为客户端也作为服务端，并且可以动态地切换角色。

使用场景：
* 像 Gnutella 和 G2 这样的文件共享网络
* 多媒体协议，如 P2PTV 和 PDTP

![peer-to-peer_pattern](peer-to-peer_pattern.png)
#### 7. 事件总线模式
此模式主要处理事件，包含4个组件：事件源、事件监听者、通道和事件总线。
事件源给一条事件总线上的特定通道发送消息。事件监听者订阅特定的通道。
监听者会收到发布到它们之前订阅的频道的消息的通知。

使用场景：
* Android 开发
* 通知服务

![event-bus_pattern](event-bus_pattern.png)
#### 8. MVC 模式
MVC 模式把一个交互式应用分成3个部分：
1. model 模型 —— 包含最核心的功能和数据
2. view 视图 —— 将信息展现给用户
3. controller 控制器 —— 处理来自用户的输入

这样做是为了将信息的内部表示与信息的呈现给用户并从用户接收的方式分开。它解耦了组件，并且允许有效代码复用。

使用场景：
* 主流语言开发的互联网应用架构
* web 框架，如 Django 和 Rails

![mvc_pattern](model-view-controller_pattern.png)
#### 9. 黑板模式
此模式对于没有确定性解决方案策略的问题很有用。

黑板模式包含3个主要组件：
* 黑板 —— 包含了来自解决方案空间对象的结构化的全局内存
* 知识源 —— 具有自表述性的专业模块
* 控制组件 —— 选择、配置和执行模块

所有的模块都可以访问黑板。组件可以生产添加到黑板中的新的数据对象。
组件在黑板中寻找特定类型的数据，可以通过利用现有数据源进行正则匹配的方式找到这些数据。

使用场景：
* 语音识别
* 车辆识别和追踪
* 蛋白质结构识别
* 声呐信号的解释

[blackboard_pattern](blackboard_pattern.png)
#### 10. 解释模式
此模式用于设计用专用语言撰写的程序组件。它主要指定如何评估程序行，即用特定语言写的句子或表达式。
其基本思想是为语言的每个符号都设置一个类。

使用场景：
* 数据库查询语言，如 SQL
* 用于描述通讯协议的语言

[interpreter_pattern](interpreter_pattern.png)
#### 各架构模式对比

| 名称 | 优点 | 缺点 |
| --- | --- | --- | 
| 分层模式 |  低层模块可以被多个高层模块使用。<br>层级概念使得标准化更容易，同时使得我们可以更容易地定义级别。<br>在层级内部的变化不会影响其他层级。| 不是普遍适用的。<br>某些层在特定情况下会被跳过。|
| 客户端-服务器模式| 可以给客户端请求的服务集合很好地建模。| 请求通常由服务器中单独的线程处理。<br>因为不同的客户端有不同的表示，所以进程间通讯会导致额外的开销。|
| 主从模式| 准确性 —— 一个服务的执行会被委托给不同的从服务，每个从服务的实现都是不同的。| 从服务是独立的，且没有共享状态。<br>在实时系统中，主从延迟是个问题。<br>此模式只能用于可解耦的问题。|
| 管道-过滤器模式| 展示并发处理。<br>当输入输出由流组成时，过滤器在接收数据时开始计算。<br>系统很容易扩展。<br>过滤器可重用。<br>可以通过重新组合给定的过滤器的集合来构建不同的管道。| 性能取决于最慢的过滤器。<br>当数据从一个过滤器传输到另一个过滤器时，有性能开销。|
| 代理模式| 允许动态地改变、添加、删除和重新定位对象，并且使得分发对开发者透明。| 需要标准化的服务描述。|
| 端到端模式| 支持分散计算。<br>对于任一节点故障都有很强的健壮性。<br>对于算力和资源有很强的可扩展性。| 由于节点是自愿合作的，所以对于服务质量没有保障。<br>安全性很难保障。<br>性能取决于节点数量。|
| 事件总线模式| 可以轻松添加新的发布者、订阅者和连接器。<br>适用于高度分布式的应用。| 由于所有的消息都要通过同一个事件总线传输，所以可伸缩性较差。|
| MVC 模式| 对于同一个模型可以轻松拥有不同的视图，这些视图可以在运行时连接或者断开。| 增加复杂性。<br>可能导致许多不必要的用户行为的更新。|
| 黑板模式| 轻松添加新的应用。<br>轻松扩展数据空间的结构。| 修改数据空间结构很困难，会影响到所有应用。<br>可能需要同步和访问控制。|
| 解释器模式| 高度动态的行为是可能的。<br>有利于终端用户的可编程性。<br>因为替换一个解释程序很容易，所以提高了灵活性。| 由于解释性的语言比编译后的语言要慢，所以性能可能是个问题。|


---

## Tip C/C++ 编译的6个步骤
* 扫描
* 语法分析
* 语义分析
* 源代码优化
* 代码生成
* 目标代码优化

Source Code ---Scanner--> Tokens ---Parser--> SyntaxTree ---Semantic Analyzer--> Commented Syntax Tree ---Source Code Optimizer--> Intermediate Representation ---Code Generator--> Target Code ---Code Optimizer--> Final Target Code

源代码 ---扫描器--> 记号 ---语法分析器--> 语法树 ---语义分析器--> 带标记的语法树 ---源码级优化器--> 中间代码 ---代码生成器--> 目标代码 ---目标代码优化器--> 最终目标代码

### 

---
    
## Share
### 10 MySQL 为什么有时会选错索引
建表
```sql
CREATE TABLE `t` (
  `id` int(11) NOT NULL,
  `a` int(11) DEFAULT NULL,
  `b` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `a` (`a`),
  KEY `b` (`b`)
) ENGINE=InnoDB；
```
#### 优化器的逻辑
优化器选择索引的原则是语句的执行代价最小。包括：
* 扫描行数
* 是否使用临时表
* 是否排序
* 等等
##### 如何判断扫描行数
MySQL 是根据统计信息来估算需要扫描的行数的。

统计信息就是索引的“区分度”。一个索引上不同的值越多，其区分度越好。一个索引上不同值的个数，我们称为“基数”。
* 基数越大，区分度越好。 show index 方法可以查看索引的基数（cardinality）。
#### MySQL 通过采样统计的方法获得基数
采样统计时， InnoDB 默认会选择 N 个数据页，统计页面上的不同值，得到一个平均值，然后乘以这个索引的页面数，就得到了索引的基数。

当变更的数据行数超过 1/M 的时候，会自动触发重新做一次索引统计。

MySQL中有两种存储索引统计的方式，可以通过设置参数 innodb_stats_persistent 的值来选择：
* 设置为 on 时， 表示统计信息会持久化存储。这时，默认的 N 是20， M 是10。
* 设置为 off 时， 表示统计信息只存储在内存中。这时，默认的 N 是8，M 是16。

analyze table t; 命令，可以用来重新统计索引信息。

#### 索引选择异常和处理
* 方法一： 采用 froce index 强行选择一个索引。
    * 写法不优美
    * 如果索引改了名字，则需要重写 SQL
    * 如果迁移到别的数据库，可能导致语句不兼容
```sql
select * from t force index(a) where a between 10000 and 20000;
```
* 方法二： 修改语句引导 MySQL 使用我们期望的索引
```sql
mysql> explain select * from t where (a between 1 and 1000) and (b between 50000 and 100000) order by b limit 1;
改为
mysql> explain select * from t where (a between 1 and 1000) and (b between 50000 and 100000) order by b,a limit 1;
```
这样改完语句后保证业务不变，且查询结果也不受影响。

之前没有选索引 a 是因为考虑到 b 需要排序，而索引是有序的，选 b 可以避免对 b 排序。

修改之后 a 也要排序，所以选 a 、 b 都一样，但是 a 扫描行数更少，所以选 a 

* 方法三：在某些场景下，我们可以新建一个更合适的索引，来提供给优化器做选择，或者删掉误用的索引
