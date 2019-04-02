# week1 ARTS
## [Algorithm](algorithm/Algorithm.md)  Leetcode [711. Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/)
* 求S中所有出现在J中的字符的个数，大小写敏感
* 比如：J=aB S=aABBb ，3个

#### 算法1 -- 集合法 
* 用一个集合JSet存放J中所有字符
* 遍历S中的字符，如果字符在JSet中则结果+1
* 时间复杂度O(len(J) + len(S))
#### 算法2 —— 对算法1的优化 
* 用int数组替代算法1中集合
* int数组的创建比HashSet的创建更高效
* int数组的标记比Set的插入更高效，因为Set的出入需要算hash值
* 综上所述，算法2比算法1的效率更高
#### 算法3 -- 用正在表达式1行代码实现
* S.replaceAll("[^" + J + "]", "").length(); 
* 将S的非J中的字符替换为空，剩下的字符串长度就是结果
#### 算法4 -- 用lambda表达式实现
```java
        (int)S.chars().parallel()
                .mapToObj(i -> (char)i)
                .filter(c -> J.contains(c+""))
                .count()
```

* 第1行将S转成IntStream（整数并行流）
* 第2行将int转成char流
* 第3行过滤掉J中存在的字符
* 第4行聚合操作，计算剩余的字符数

* 提交后发现性能并没有算法1、2好，可能是由于线上运行的服务器只有单核，也可能是S的长度不够长不足以发挥并行的作用。
* 这里仅作为学习lambad表达式用
## Review
* 阅读文章 [How to write code you will love in the future](https://medium.freecodecamp.org/how-to-write-code-you-will-love-in-the-future-ee5decae5ce4)
* 怎样写出让未来的你喜欢的代码 
    * 对于代码质量，永远不要妥协
    * 一定要写文档和注释
    * 除非你确保重复造轮子是可维护的，否则永远不要重复造轮子
        * 商业上重复造轮子是不好的，可是如果为了学习，你一定要重复造轮子
    * 一定测试你的代码，尤其要写单元测试
    * 保持学习
* 我觉得要写出可维护行强的代码，首先一定要学会如何设计，其次一定要遵循代码规范。
    * 设计良好的代码，思路清晰，让人容易阅读，且可扩展性强
    * 遵循代码规范，可以减少误读。否则很有可能会引发严重的事故。笔者（我）就曾经遇到过由于if语句没有带大括号，其他人修改代码后导致生产数据库数据混乱的问题。
    
## Tip
* 如何使用gdb调试c++
    1. 编写c++源码文件 hello.cpp
    2. 编译： c++ -g hello.cpp -o hello
        * -g 参数代表加入调试代码
    3. 用gdb运行： gdb hello
    4. 设置断点进行调试，gdb常用命令：
        * help 获取帮助信息
        * break 设置断点，如：break main 在main函数入口设置断点，break 5 在第5行设置断点
        * run 运行程序，在断点处会停止
        * next 单步执行，跳过函数调用
        * step 单步执行，进入函数调用
        * continue 继续执行，直到下一个断点
        * list 列出部分源代码
        * watch 监视变量，有变化时显示
        * quit 退出
## Share
### 学习极客时间MySQL实战45讲——02 日志系统：一条SQL更新语句是如何执行的？

##### 一条update语句的执行过程
 1. 通过连接器连接到服务器
 2. 从查询缓存中删除相关表的缓存
 3. 分析器做词法分析和语法分析
 4. 优化器选择使用哪些索引，如果需要多表关联则选择关联顺序
 5. 执行器调用存储引擎接口更新数据，同时写binlog（归档日志）
 6. 存储引擎更新数据，如果是InnoDB引擎，则同时写redolog（重做日志）
##### 重做日志 redolog
* WAL(Write-Ahead Logging)技术:先写日志，再写磁盘
* redolog大小固定，比如可以配置一组4个文件，每个文件大小1G，总共可以记录4G大小的日志。
* redolog是循环写的，从头开始写，写的末尾后在从头循环写，
* write pos：当前记录位置，一边写一边后移。
* checkpoint：检查点。当前需要擦除的位置，往后推移并循环写入。擦除之前要先将记录更新到数据文件中。
* crash-safe：使用redolog，InnoDB可以保障，即使数据库异常重启也不会丢失数据
##### 归档日志 binlog 与 redolog 的区别
1. redolog是InnoDB存储引擎特有的。binlog是MySQL的Server层实现的。
2. redolog是物理日志，更底层，记录“在某个数据页上做了什么修改”，由于redolog在存储引擎层，所以其信息更贴近于磁盘的存储；
binlog是逻辑日志，记录语句的原始逻辑，是通过语句分析得到的。
3. redolog大小固定，循环写。binlog是追加写的，理论上无大小限制。
##### "update T set c=c+1 where ID=2;"语句的执行流程
1. 执行器找InnoDB要ID=2（ID是主键）这一行。InnoDB通过B+树索引找到ID=2这一行数据，如果在内存中则直接返回，否则先从磁盘读入到内存，然后返回。
2. 执行器拿到这行数据库把c值+1，然后得到一行新的数据，再调用InnoDB接口写入数据。
3. InnoDB更新此数据到内存中，同时记录redolog，此时redolog出于prepare状态。告诉执行器执行完成，随时可以提交事务。
4. 执行器写binlog，并把binlog写磁盘。
5. 执行器调用InnoDB接口提交事务，InnoDB将刚刚写入的redolog改为commit状态，更新完成。

* 以上就是两阶段提交的过程
##### 如何让数据库恢复到半个月内的任意一秒
* 首先需要定期全量备份数据库，每天一备，或每周一备
* 找到最近一次全量备份，从这个备份恢复到临时库
* 从备份点开始，将备份的binlog依次取出，重复到需要恢复到的那个时刻
* 将临时库改为生产库
##### 相关参数建议值
* innodb_flush_log_at_trx_commit = 1; 表示每次事务的redolog都直接持久化到磁盘中。可以保证MySQL异常重启后数据不丢失。
* sync_binlog = 1; 表示事务的每次binlog都持久化到磁盘，保证MySQL异常重启后binlog不丢失。

