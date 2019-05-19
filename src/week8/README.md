#week8

---

## Algorithm [701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)
### 1. 问题描述
将一个元素插入到二叉搜索树中。
给定二叉搜索树的根 root，且新元素在树中不存在。
### 2. 解题思路
二叉索索树定义：
* 左子树中所有节点的值都小于此节点
* 右子树中所有节点的值都大于此节点
* 左右子树分别都是二叉搜索树

将新元素与当前节点的值对比，如果大于当前节点则插入到其右子树中，如果小于当前节点则插入其左子树中。
递归调用直到叶子节点为止。

### 3. 代码
#### 3.1 递归
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
```
#### 3.2 非递归
```go
func insertIntoBSTWithoutRecursion(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	node := root
	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{val, nil, nil}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{val, nil, nil}
				break
			} else {
				node = node.Right
			}
		}
	}
	return root
}
```
### 4. 复杂度分析
* 时间复杂度 O(H), 其中 H = logN 为树的高度。 最坏情况为 O(N) 二叉树退化为链表的情况。
* 空间复杂度:
    * 递归： O(H) 栈的深度
    * 非递归： O(1) 只需要一个节点的空间存储临时节点。

---

## Review [The ultimate guide to preparing for the coding interview](https://medium.freecodecamp.org/the-ultimate-guide-to-preparing-for-the-coding-interview-183251ee36c9)
### 准备编码面试的终极指南
—— 为技术面试，行为问题和工资谈判做准备
#### 任务1 阅读有 Facebook 员工 Yangshun Tay 写的[面试形式](https://github.com/yangshun/tech-interview-handbook/blob/master/non-technical/interview-formats.md),以便于理解不同公司的不同面试风格。
你需要准备如下问题：
1. 技术问题
2. 设计问题
3. 行为问题

#### 技术问题
你可以找一个平台，如 [leetcode](https://leetcode.com/), [Hackerrank](https://www.hackerrank.com/dashboard), 和 [Learneroo](https://www.learneroo.com/)

阅读由 Gayle Laakmann McDowell 写的 [“Cracking the Coding Interview”](https://www.amazon.com/Cracking-Coding-Interview-Programming-Questions/dp/0984782850) 这本书

#### 任务2 阅读 [“Cracking the Coding Interview”](https://www.amazon.com/Cracking-Coding-Interview-Programming-Questions/dp/0984782850) 的第7章 - 技术问题
#### 任务3 作为对数据结构和算法的复习，完成[此页面](https://github.com/yangshun/tech-interview-handbook/tree/master/algorithms) 中每一节的第一个编码挑战
#### 任务4 找到资源测试你自己对于最常见的概念性问题的掌握程度。如果你不知道问题的答案，在看答案之前用 Google 搜索来学习。
以下是一些资源
* 算法和数据结构（最多45分钟）
    * [https://www.geeksforgeeks.org/commonly-asked-algorithm-interview-questions-set-1/](https://www.geeksforgeeks.org/commonly-asked-algorithm-interview-questions-set-1/)
    * [https://www.geeksforgeeks.org/commonly-asked-data-...](https://www.geeksforgeeks.org/commonly-asked-data-structure-interview-questions-set-1/)
* Ruby & Rails (最多45分钟)
    * [https://www.toptal.com/ruby/interview-questions/](https://www.toptal.com/ruby/interview-questions/)
    * [https://www.upwork.com/i/interview-questions/ruby/](https://www.upwork.com/i/interview-questions/ruby/)
    * [https://rubygarage.org/blog/how-to-interview-your-ruby-on-rails-developer](https://rubygarage.org/blog/how-to-interview-your-ruby-on-rails-developer)
* HTML & CSS (最多45分钟)
    * [https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/css-questions.md](https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/css-questions.md)
    * [https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/html-questions.md](https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/html-questions.md)
* JavaScript (最多45分钟)
    * [https://www.toptal.com/javascript/interview-questi...](https://www.toptal.com/javascript/interview-questions)
    * [https://www.codementor.io/nihantanu/21-essential-j...](https://www.codementor.io/nihantanu/21-essential-javascript-tech-interview-practice-questions-answers-du107p62z)
    * [https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/javascript-questions.md](https://github.com/yangshun/front-end-interview-handbook/blob/master/questions/javascript-questions.md)
* React (最多45分钟)
    * [https://www.edureka.co/blog/interview-questions/react-interview-questions/](https://www.edureka.co/blog/interview-questions/react-interview-questions/)
    * [https://www.toptal.com/react/interview-questions](https://www.toptal.com/react/interview-questions)
#### 系统设计问题
你应该熟悉以下概念：DNS 服务，缓存，负载均衡，数据库等。
#### 任务5 阅读  “Cracking the Coding Interview”  第9章 - 系统设计和可扩展性
#### 任务6 Facebook 的一个技术经理 Donne Martin 创建了一个很好的开源资源 [“How to approach a system design interview question”.](https://github.com/donnemartin/system-design-primer#how-to-approach-a-system-design-interview-question)
#### 任务7 了解可伸缩系统的4种不同的方式：
* [Scalability for Dummies - Part 1: Clones](http://www.lecloud.net/post/7295452622/scalability-for-dummies-part-1-clones)
* [Scalability for Dummies - Part 2: Database](http://www.lecloud.net/post/7994751381/scalability-for-dummies-part-2-database)
* [Scalability for Dummies - Part 3: Cache](http://www.lecloud.net/post/9246290032/scalability-for-dummies-part-3-cache)
* [Scalability for Dummies - Part 4: Asynchronism](http://www.lecloud.net/post/9699762917/scalability-for-dummies-part-4-asynchronism)
#### 任务8 了解每个软件系统的不同组件
* [Domain Name System (DNS)](https://github.com/donnemartin/system-design-primer#domain-name-system)
* [Content Delivery Network (CDN)](https://github.com/donnemartin/system-design-primer#content-delivery-network)
* [Load Balancer](https://github.com/donnemartin/system-design-primer#load-balancer)
* [Reverse Proxy](https://github.com/donnemartin/system-design-primer#reverse-proxy-web-server)
* [Application layer](https://github.com/donnemartin/system-design-primer#application-layer)
* [Database](https://github.com/donnemartin/system-design-primer#database)
* [Cache](https://github.com/donnemartin/system-design-primer#cache)
* [Asynchronism](https://github.com/donnemartin/system-design-primer#asynchronism)
* [Communication](https://github.com/donnemartin/system-design-primer#communication)
#### 任务9 在此[列表](https://github.com/donnemartin/system-design-primer#company-architectures)中找出 2 ~ 3 个公司，了解他们的架构
#### 任务10 看下面两个设计问题的答案
* [URL Shortener](http://blog.gainlo.co/index.php/2016/03/08/system-design-interview-question-create-tinyurl-system/)
* [Collaborative Editor](http://blog.gainlo.co/index.php/2016/03/22/system-design-interview-question-how-to-design-google-docs/)
#### 任务11 在此[列表](https://github.com/yangshun/tech-interview-handbook/tree/master/design#specific-topics)中找出 2 个设计问题，然后用笔和纸写下答案。推荐你找一个朋友或同事一起完成
#### 行为问题
公司不仅想找技术高手，同时也希望找到适合企业文化的人。

一些常见的行为问题：
* 为什么你想来这里工作？
* 说一说你最成功最有成就感的一个项目
#### 任务12 阅读 “Cracking the Coding Interview” 的第5章 - 行为问题
#### 任务13 用 Google 文档完成书中的"面试准备网格"。最好使用软件项目填写
#### 任务14 根据 S.A.R （Situation 情景, Action 行动, Result 结果) 和书中的一般提示，来写下书中问题的答案（例如："弱点" 和 "告诉自己")。同时写下此[列表](https://github.com/yangshun/tech-interview-handbook/blob/master/non-technical/behavioral.md) General 小节的问题的答案。
#### 其他问题
##### 提问
在面试中提出好的问题，不仅让你看起来很聪明，而且你也能更了解这个公司以及公司中的人。
#### 任务15 查看此[列表](https://medium.freecodecamp.org/the-ultimate-guide-to-preparing-for-the-coding-interview-183251ee36c9)中的问题，以后的面试中你也许可以用到。
#### 心理学的一些技巧 查看[此网站](https://github.com/yangshun/tech-interview-handbook/blob/master/non-technical/psychological-tricks.md)
#### 谈判技巧 
#### 任务16 阅读此[列表](https://github.com/yangshun/tech-interview-handbook/blob/master/non-technical/negotiation.md)中的10条谈判规则
#### 模拟面试
模拟面试的[例子](https://docs.google.com/spreadsheets/d/1t_228bDllazltWrq7WrLaKCs3dHXlGHBybWz9nnRvSc/edit?usp=sharing)

这里有2个模拟面试的平台：
* [Pramp](https://www.pramp.com/#/)
* [Interviewing.io](https://interviewing.io/)

---

## Tip GCC 编译过程的4个步骤
```jshelllanguage
gcc hello.c 
```
整个编译过程可以分成4个步骤：
1. 预处理 (Preppressing) : 
源代码文件 .c 和相关的头文件 .h ，如 stido.h 等被预编译器 cpp 预编译成了一个 .i 文件
```jshelllanguage
gcc -E hello.c -o hello.i
#or
cpp hello.c > hello.i
```
2. 编译 (Compilation) :
编译过程就是对 .i 文件进行一系列的词法分析、语法分析、语义分析及其优化后生成汇编代码文件
```jshelllanguage
gcc -S hello.i -o hello.s
```
3. 汇编 (Assembly) :
汇编器是将汇编代码转成机器可以执行的指令，每一条汇编语句几乎都对应一条机器指令。
```jshelllanguage
as hello.s -o hello.o
#or
gcc -c hello.s -o hello.o
```
4. 链接 (Linking) :
把目标文件和静态库文件链接成一个可执行的二进制文件
```jshelllanguage
ld -static /usr/lib/crt1.o /usr/lib/crti.o /usr/lib/gcc/i486-linux-gnu/4.1.3/crtbeginT.o -L/usr/lib/gcc/i486-linux-gnu/4.1.3 -L/usr/lib -L/lib hello.o --start-group -lgcc -lgcc_eh -lc --eng-group /usr/lib/gcc/i486-linux-gnu/4.1.3/crtend.o /usr/lib/crtn.o
```

### 

---
    
## Share
### 09 普通索引和唯一索引，应该怎么选择 —— 极客时间-MySQL实战45讲
普通索引和唯一索引的区别
````sql
create table T(
    id int primary,
    key int not null,
    name varchar(16),
    index (k)
) engine=InnoDB;
````
#### 查询过程
sql 语句
```sql
select id from T where k=5;
```
查询过程：先通过 B+ 树从树根开始，按层搜索到叶子节点，拿到数据页，然后在数据页内部进行二分查找
* 对于普通索引，查找到满足条件的第一个记录后，需要查找下一个记录，知道碰到不满足条件的第一个记录为止
* 对于唯一索引，由于索引定义了唯一性，所以查找到第一满足条件的记录后，就停止继续检索

两种索引的性能差距微乎其微。
#### 更新过程
* 什么是 change buffer
    * 需要更新一个数据页时，如果该页在内存中则直接更新。如果不在内存中，在不影响数据一致性的前提下，InooDB会将更新缓存在 change buffer中，
    这样就不需要从磁盘中读取此页了。下次需要查询需要访问这个数据页时，将数据页加载到内存，然后执行 change buffer 中的相关操作。
    * change buffer 在内存中有拷贝，同时也会写入磁盘
    * change buffer 使用的是 buffer pool 中的内存。
    * 可以通过 innodb_change_buffer_max_size 来动态设置其大小。当设置为50时，表示最多能占用 50% 的 buffer pool.
* 什么是 merge
    * 将 change buffer 应用到原数据页，得到最新结果的过程称为 merge
    * 访问这个数据页会触发 merge
    * 系统后台线程会定期 merge
    * 数据库正常关闭时也会执行 merge
* 什么条件下可以使用 change buffer
    * 唯一索引的更新操作需要校验唯一性，必须要将数据页读入内存才能判断，所以唯一索引的更新不能使用 change buffer
    * 实际上只有普通索引才能使用 change buffer

更新语句的执行过程：
* 记录要更新的目标页在内存中的情况：
    * 唯一索引，找到更新值的位置，判断是否有冲突，然后插入这个值
    * 普通索引，找到位置，插入值

这种情况下普通索引和唯一索引之间的性能差距可以忽略不计

* 记录要更新的目标页不在内存中的情况：
    * 唯一索引，将目标页读入到内存（I/O 操作），判断是否有冲突，插入值
    * 普通索引，将更新记录在 change buffer 中，结束

这种情况下，由于唯一索引需要进行 I/O 操作，所以性能比普通索引差很多

* 案例：某个业务系统的库内存命中率突然从99%降低到了75%，整个系统出于阻塞状态，更新语句全部堵塞。原因：这个业务有大量的插入操作，而前一天某个普通索引被改成了唯一索引。

#### change buffer 的使用场景
change buffer 对更新过程有加速作用，而且只限于普通索引。
* 在写多读少的业务中，页面写完以后马上被访问的概率较小，此时 change buffer 使用效果最好。如账单类、日志类系统。
* 如果业务的更新模式是写入后马上做查询，那么更新时写入 change buffer，但是之后马上访问这个数据页，会立即出发 merge 过程。
这样随机访问磁盘 IO 的次数不会减少，反而还增加了 change buffer 的维护成本。所以这种业务模式不适合使用 change buffer
#### 索引选择和实战
* 普通索引和唯一索引在查询能力上没有差别，主要需要考虑更新性能。所以，建议尽量选择普通索引。
* 如果所有的更新后，都马上跟着对这个记录的查询，应该关闭 change buffer
* 普通索引和 change buffer 的配合使用对于数据量大的表的更新优化比较明显。
#### change buffer 和 redo log
现在我们要执行插入语句
```sql
insert into t(id,k) values(id1,k1),(id2,k2);
```
假设 k1 所在的数据页在内存中， k2 所在数据页不在内存中。

这条更新语句会涉及4个部分：内存，redo log (ib_log_fileX), 数据表空间(t.ibd)，系统表空间(ibdata1)

更新语句操作如下：
1. Page 1 在内存中，直接更新内存
2. Page 2 没在内存中，就在内存的 change buffer 区域，记录下"我要往 Page 2 插入一行"这个信息
3. 将上述两个动作记入 redo log 中

执行这条更新语句的成本很低，写了2处内存，然后写了1处磁盘，而且还是顺序写

读请求的处理
```sql
select * from t where k in (k1,k2)
```
如果读语句发生在更新语句后不久，内存中的数据都还在，那么此时的这两个读操作就与系统表空间(ibdata1) 和 redo log (ib_log_fileX) 无关了。

1. 读 Page1 的时候，直接从内存返回。
2. 要读 Page2 的时候，需要把 Page2 从磁盘读入到内存中，然后应用 change buffer 里面的操作日志，生成一个正确的版本并返回结果。

redo log 和 change buffer 对比：
* redo log 主要节省的是随机写磁盘 IO 消耗（转成顺序写）
* change buffer 主要节省的是随机读磁盘 IO 的消耗。



