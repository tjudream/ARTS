#week11

---

## Algorithm [905. Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/)
### 1. 问题描述
按照奇偶校验排序数组。

给定一个数组A，其元素都是非负正整数。对 A 进行排序，要求所有偶数都在奇数之前。任意结果都接收。

例如：
* 输入 [3,1,2,4]
* 输出 [2,4,1,3], [4,2,3,1], [2,4,1,3] 或 [4,2,1,3] 4个中的任何一个都是对的。
### 2. 解题思路
1. 用两个指针 s, e 分别指向数组的头和尾。
2. 从头开始找到第一个奇数，从尾向前找到第一个偶数，交换这2个元素
3. 重复第2步，直到 s 大于 e

* 这里还有一个小技巧，用 A[s]&1 == 1 来判断奇偶，而不是用 A[s]%2 == 1 来判断。
因为位运算比取余运算性能更高。
### 3. 代码
```go
func sortArrayByParity(A []int) []int {
    s := 0
	e := len(A) - 1
	for s < e {
		for A[s]&1 == 0 && s < e {
			s++
		}
		for A[e]&1 == 1 && s < e {
			e--
		}
		if (s < e) {
			A[s],A[e] = A[e],A[s]
		}
	}
	return A
}
```
### 4. 复杂度分析
* 时间复杂度： O(N), N 是数组 A 的元素个数。 s, e 分别从头尾遍历数组，对数组只遍历了1遍
* 空间复杂度： O(1), 只需要2个指针的空间来存储 s, e

---

## Review [Why Defensive Programming is the Best Way for Robust Coding](https://medium.com/swlh/why-defensive-programming-is-the-best-way-for-robust-coding-cfa790fe04cd)
### 为什么防御性编程是健壮代码的最佳实践
防御性编程是指程序员预测问题并且写代码来处理他们。

这就像将来发生车祸……但是你依旧淡定，因为你已经买了保险。

也就是说，防御性编程的关键在于防范你不期望发生的错误。一个防御性程序员会在一个麻烦成为实际问题之前避免它。其主要思想是永远不要编写永不失败的代码。
这是一个乌托邦式的梦想。其主要思想是在遇到任何没有预期到的问题时，使代码优雅的失败。
优雅的失败指的是如下几种情况：
* 尽早失败：你的代码应该确保所有重要的操作都提前终止，尤其是如果这些操作的计算量很大或者可能对数据造成不可逆的影响。
* 安全的失败：在失败的地方，你的代码应该要确保释放所有的锁，不会申请新的锁，不会写文件等。
* 清晰的失败：当某些东西出问题时，你的代码应该返回一条清晰的错误信息，并且这个信息描述可以使得支持团队解决这个问题。

好了，到这里你可能会反对。

现在没有问题。我的代码运行的很好。为什么我要花时间和精力在"未来预期"的问题上呢？毕竟，我们总是被反复的教授"你不需要它(You Ain’t Gonna Need It)"(YAGNI).
而且你是一个专业的程序员，而不是一个随意添加代码的业余爱好者。

关键在于实用主义。

Andrew Hunt 在他的书《程序员修炼之道》中将防御性编程描述为"注重实效的偏执(Pragmatic Paranoia)"。

保护你的代码不受其他人和你自己错误的影响。如果有疑虑，就验证它。检查数据一致性和完整性。你不可能测试到每一个错误，因此使用断言和异常处理来处理那些"不可能发生"的事情。

健康的偏执编程是正确的编程方式。但是不能太过于偏执。关键在于适当的平衡。

以下是一些防御性编程的方式。

#### 问你自己：这里会失败吗？
每一行代码都会做一些事情，所以第一道防线就是问你自己如果代码失败了，会怎么样。

例如，考虑以下不符合规范的代码
```java
CASE SY-INDEX. // Noncompliant; missing WHEN OTHERS clause
WHEN ONE.
WRITE ‘One’.
WHEN 2.
WRITE ‘Two’.
ENDCASE.
```

我们会问如下两个问题：
* 如果 sy-index 不是 1 会发生什么。
* 如果 sy-index 不是 2 会发生什么。

为了解决这些问题，我们增加了 OTHERS 代码段：
```java
CASE SY-INDEX.
WHEN ONE.
WRITE ‘One’.
WHEN 2.
WRITE ‘Two’.
WHEN OTHERS. // Compliant
WRITE ‘Unexpected result’
ENDCASE.
```

是不是很简单？

正是这种"假设"的想法将优秀的程序员与那些编写代码并希望代码永远不失败的程序员分开来。"永远"总是比期望的来的要早，到那时，
代码早已被埋没在程序的一个长期被遗忘的部分了，同时错误信息并没有说明问题出在了哪以及如何解决。

这种防御性编程技术的美妙之处在于，它几乎没有花费任何时间就可以在你的代码中添加详细的类型检查。

#### 仔细的检查边界条件
第一个检查是确定是否需要边界条件，毕竟循环是昂贵的。

边界条件是所有动作发生的地方。从0循环到100和从1循环到98几乎相同（当然，除了代码中的条件）。但是0是代码进入循环的地方，并且初始条件已经被设置
（并可能设置错误）。同样，最后一个循环就是离开循环的地方，无论循环对值如何操作，都会停止。

最多迭代一次循环相当于使用 if 语句有条件地执行一段代码。任何开发者都不应该期望找到循环语句的这种用法。如果作者的最初意图确实是想有条件地执行
一段代码，那么应该使用 if 语句。

请考虑以下不合格和合格的代码。在这个示例中我们不需要使用循环。一个简单的 if 就搞定了。
```java
    Noncompliant Code Example

DATA remainder TYPE i.
DO 20 TIMES.
remainder = sy-index MOD 2.
cl_demo_output=>write_text().
EXIT. “ noncompliant, loop only executes once. We can use IF
ENDDO.


    Compliant Code Example

DATA remainder TYPE i.
DO 20 TIMES.
remainder = sy-index MOD 2.
cl_demo_output=>write_text().
ENDDO.
```

永远记住，调试循环总是需要包括开始循环和结束循环时的大量的工作，确保进入循环和退出循环都是正确的。因此，一旦你弄清楚了边界条件，你的代码就没有
其他任何问题了。

#### 使用 TDD （测试驱动开发）
TDD的基本思想是"首先写单元测试，然后写代码，然后重构，然后重复"。

单元测试是用于检查函数是否按行为是否符合预期的自动化测试。你的第一次单元测试应该是失败的，因为它是在你写代码之前编写的。

你在测试代码中增加了一些内容。你在生产代码中增加一些内容。两个代码流同时增长为互补组件。这些测试适合生产代码，就像抗体适合抗原一样。

测试代码的问题是你必须隔离该代码。如果一个函数调用另一个函数，那么通常会较难测试。要编写改测试，你必须找出一些方法将该函数与其他函数解耦。
换句话说，对于测试的需求首先迫使你寻求一个良好的设计。

这样就创造了一个更好的和解耦的设计，在这个设计里你可以很好地控制代码的开发。

虽然预先编写测试用例可能会耗费时间，但是这会带来很多好处。开发人员承认以前他们曾经编写过代码行，意识到他们的解决方案无关紧要，然后从头开始重新编码。

与过时的编写实践不同，TDD允许开发人员回到绘图板并专注于预先设计轻量级，灵活的架构。

提前编写测试用例的事实可以防止任何可能在稍后出现的bug，从而节省时间、精力和心力。

#### 始终编写优秀的代码
一些程序（或程序员）非常喜欢资源。但是尽可能少的使用它们。同时为了使用最小化的资源，你的代码应该尽可能的优化。

通常，一个确定的优化方法是打开编辑器提供的内置优化。

编译器优化通常可以将运行时间从几百个百分点提高到2倍。有时它可能会减慢产品速度，所以在进行最后一次编译之前要仔细测量。然而，现代的编译器在这
方面表现的非常好，因为它们消除了程序员对小规模更改的大部分需求。

除了标准的编译器优化之外，还可以使用其他几种调优技术。

##### 收集常见的子表达式
在一个地方进行计算并记录结果好于在多个地方进行昂贵的计算。除非必要，否则不要将这些计算放在循环中。

##### 用便宜的操作替换昂贵的操作
字符串操作可能是任何程序中最常见的操作之一。但是，如果操作不正确，字符串操作很可能是一项昂贵的操作。同样，在某些情况下，你可以用一系列的位移
操作来替换乘法，已获得更好的性能。即使这是有效的（并且并非总是如此），它也会产生非常令人迷惑的代码。因此做决定时也需要考虑代码的可读性。

##### 消除循环
循环是主要的开销。如果迭代次数不多，尽量避免使用循环。

##### 缓存常用的值
缓存利用了局部性、程序和人员重用最近使用的数据的趋势。仅缓存最常用的字符和数据可显著提高程序的性能。

#### 用低级语言重写
这应该是最后的手段。低级语言往往更有效率，但从程序员的角度来看更耗时。有时，我们通过在低级语言中重写关键代码来获得显著的改进，但这是以降低
可移植性和维护成本为代价的。所以仔细考虑这个决定。

记住，在优化中，选择可能占用了这个过程的90%。花点时间来决定你在做什么并把它做好是值得的。当然：这也是黑魔法的所在。

#### 最后，不要相信任何人
第二任布什政府时期的国防部长唐纳德拉姆斯菲尔德曾在新闻发布会上说过："有已知的已知；有些事情我们知道我们知道。"。"我们也知道有已知的未知；
也就是说，我们知道有些事情我们不知道。但也有未知的未知，那些我们不知道我们不知道的。"

拉姆斯谈论的是伊拉克战争，但数据也是如此。简而言之，这意味着要验证你无法完全控制的所有数据。

显然，用户数据总是令人怀疑的。用户可以很好地误解你认为清楚的东西。尝试并预测问题，同时验证或以其他方式整理所有内容。

程序设置数据也容易出错。INI 文件曾经是保存程序设置的常用方法。因为它们是一个文本文件，许多人养成了用文本编辑器手动编辑它们的习惯，并且可能弄错这些值。
注册表数据，数据库文件——某人有一天可以也将会调整它们，所以验证这些配置是值得的。

简而言之，如果你希望你的代码能够按照预期的方式运行，那么输入的数据必须是干净的。如果你曾经听过"Garbage in, Garbage out"这句话，那就是它的来源。

正如爱德华德明所说的那样。

"我们相信上帝，其他所有人都必须带来数据。"（In God we trust. All others must bring data.）


---

## Tip 
### 如何用免费的 SSL 证书实现一个 HTTPS 站点
* 编译 nginx 时需要带上 http_ssl_module
```jshelllanguage
./configure --prefix=/home/dream/nginx --with-http_ssl_module
make && make install
```
* 生成证书: 域名必须存在
```jshelllanguage
apt install python-certbot-nginx
certbot --nginx --nginx-server-root=/home/dream/nginx/conf/ -d www.mydomain.top
``` 

---
    
## Share
### 13 为什么表数据删掉一半，表文件大小不变？—— 极客时间 MySQL实战45讲
为什么删了最大表的一半数据，但是表文件的大小没变？

InnoDB包含两个部分：表结构定义和数据。

在 MySQL8.0 以前，表结构是以 .frm 为后缀的文件里的。

而 MySQL8.0 已经允许将表结构定义放到系统表中了。因为表结构定义占用空间很小，所以我们主要讨论表数据。

#### 参数 innodb_file_pre_table
1. OFF 表示表的数据放在系统的共享空间，也就是跟数据字典放在一起。
2. ON 表示每个 InnoDB 表数据存储在一个以 .ibd 为后缀的文件中。

从 MySQL5.6.6 开始默认都是 ON。推荐无论哪个版本这个参数都要设置成 ON。因为如果放到共享空间中，即使表删掉了，空间也是不会回收的。

我们在删除整表的时候，可以使用 drop table 命令回收表空间。

#### 数据删除流程
InnoDB 中的数据都是用 B+ 树的结构组织的。

删除一条记录的时候，InnoDB 只是把这个记录标记为删除。如果只有再插入数据，这个位置可能被复用。因此磁盘文件并不缩小。

如果删除整个数据页，那么这个数据页就被标记位删除，并且成为可复用的。

* 数据页的复用和记录的复用是不同的
    * 记录的复用，只限于符合范围条件的数据。
    * 而数据页可以复用到任何位置。
 
如果相邻的两个数据页的利用率都很低，系统就会把这两个数据页合并到一个页中，然后将另一个页标记为可复用。

如果我们用 delete 删除整张表的数据，那么所有的数据页都会被标记为可复用，但是磁盘文件不会变小。

#### 不止删除数据会造成空洞，插入数据也会
随机插入数据有可能造成页分裂，从而产生空洞。

更新索引上的值，可以理解为删除一个旧值，再插入一个新值。这也会造成空洞。

也就是说，经过大量增删改的表，都是可能存在空洞的。重建表可以去掉这些空洞，从而达到收缩表空间的目的。

#### 重建表
步骤一：
1. 新建与表 A 结构相同的表 B
2. 按照主键 ID 递增的顺序，把数据一行一行地从表 A 中读出来，再插入到表 B
3. 用表 B 替换表 A。 由于表 B 是新建表，且主键索引是按照顺序插入的，所以表 A 中的主键索引上的空洞，在表 B 中就不存在了。

可以使用 alter table A engine=InnoDB 命令来重建表，在 MySQL5.5 之前，执行流程如上，只不过表 B 是自动创建的。

花费时间最多的步骤是往临时表中插入数据，这个过程如果有数据写入表 A 就会造成数据丢失。因此整个 DDL 的过程中，表 A 中不能有更新。
也就是说 DDL 不是 Online 的。

#### MySQL5.6 开始引入了 Online DDL ，对这个流程做了优化
步骤二：
1. 建立一个临时表，扫描表 A 主键的所有数据页
2. 用数据页中表 A 的记录生成 B+ 树，存储到临时文件中
3. 生成临时文件的过程中，将所有对 A 的操作记录在一个日志文件（row log）中
4. 临时文件生成后，将日志文件中的操作应用到临时文件，得到一个逻辑数据上与表 A 相同的数据文件。
5. 用临时文件替换表 A 的数据文件

注意：对于很大的表来说，重建表的操作会消耗很多 IO 和 CPU 资源。

因此如果是线上服务，需要很小心地控制操作时间。如果想比较安全的操作的话，推荐使用 Github 上开源的 [gh-ost](https://github.com/github/gh-ost)

#### Online 和 inplace
在 MySQL5.5 之前临时表是 tmp_table ，是在 server 层创建的。

在MySQL5.6 之后临时表是 tmp_file, 是在 InnoDB 内部创建出来的。整个 DDL 的过程都是在 InnoDB 内部进行的。对于 server 层
来说，没有把数据挪动到临时表，是一个"原地"操作，这就是"inplace"名称的来源。

所以如果你有一个 1TB 的表，磁盘空间是 1.2TB，是不能做 inplace 的 DDL 的。因为 tmp_file 也要占用临时空间。

重建表的语句 alter table t engine=InnoDB 其隐含的意思是
```sql
alter table t engine=InnoDB, ALGORITHM=inplace;
```
跟 inplace 对应的就是拷贝表的方式：
```sql
alter table t engine=InnoDB, ALGORITHM=copy;
```
当使用 copy 时，表示强制拷贝表，对应的流程是步骤一。

只有在重建表的这个逻辑中 inplace 是 Online 的。

Online 和 inplace 之间的关系：
1. DDL 过程如果是 Online 的，就一定是 inplace 的。
2. 反过来未必，也就是说 inplcae 的 DDL，有可能不是 Online 的。截止到 MySQL8.0， 添加全文索引（FULLTEXT index)和空间索引（SPATIAL index）就属于这种情况。

optimize table ， analyze table 和 alter table 这三种方式重建表的区别：
* MySQL5.6 开始，alter table t engine=InnoDB （也就是 recreate）默认就是步骤二的流程。
* analyze table t 其实不是重建表，只是对表的索引信息做了重新统计，没有修改数据，这个过程中加了 MDL 读锁。
* optimize table t 等于 recreate + analyze

#### 问题
想要收缩表空间，结果适得其反：
1. 一个表 t 文件大小为 1TB
2. 对这个表执行 alter table t engine=InnoDB
3. 发现执行后，空间不仅没变小，还稍微大了一点儿，比如变成了 1.01TB

原因：
1. 这个表本身就没有空洞，DDL 期间，刚好有外部的 DML 在执行，这期间会引入一些新的空洞
2. 在重建表时，InnoDB 不会把整张表占满，每个页留了 1/16 给后续的更新用。也就是说，其实重建表之后不是"最"紧凑的。
 过程如下：
    * 1） 将表 t 重建一次
    * 2） 插入一部分数据，但是插入的这些数据，用掉了一部分的预留空间
    * 3） 这种情况下，再重建一次表 t， 就可能会出现多出来一部分空间的现象。


