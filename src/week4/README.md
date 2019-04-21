#week4
## Algorithm [938. Range Sum of BST](https://leetcode.com/problems/range-sum-of-bst/)
### 1. 问题描述
给一个二叉查找树(BST)的根节点root，树的节点为int型，求此树中节点介于[L,R]区间之间的节点的和。
### 2. 解题思路
* 二叉查找树(BST)的特点: 一个节点的左子树都小于该节点的值，一个节点的右子树都大于该节点的值
* 由BST的特点可以知道：
    * （1）若根节点的值小于L，则在左子树中寻找解
    * （2）若根节点的值大于R，则在右子树中寻找解
    * （3）若介于L和R之间，则结果为 根节点的值 + 左子树中符合条件节点的和 + 右子树中符合条件节点的和
* 由上可知可以使用递归的方法求解

### 3. 代码
```go
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	} else {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	}
}
```
### 4. 复杂度分析
* 时间复杂度，一般情况O(N)，最坏情况需要遍历整棵树
* 空间复杂度，递归会占用大量的栈内存，尤其是递归比较深的情况下，需要注意栈内存溢出的问题。go目前还没有实现尾调用优化。
* 如果担心内存溢出问题，可以考虑使用channel来模拟递归调用，可以参考：[Go 语言中的递归和尾调用操作](https://studygolang.com/articles/16138)

---

## Review [The Go Memory Model](https://golang.google.cn/ref/mem#tmp_0)
### 简介
go内存模型规定了一个goroutine(协程)可以看到另一个goroutine修改同一个变量的值的条件。
### 建议
如果程序要访问由多个goroutine同时修改的数据，则需要顺序访问。

如果要保证顺序访问，则需要使用channel或者使用sync和sync/atomic 包中提供的同步原语。

如果你需要阅读文档的剩下部分来理解你的程序的行为，那你就太聪明了。

不要太聪明。

### Happens Before（之前发生）
* 在同一个goroutine中，读写顺序是由你代码的顺序决定的。编译器或者处理器有可能在不改变程序行为的前提下修改语句的执行顺序。比如a = 1; b =2;
b被赋值为2的行为有可能发生在a被赋值为1之前。

* 如果事件e1发生在e2之前，那我们说e2发生在e1之后。
* 如果e1没有发生在e2之前，e1也没有发生在e2之后，那我们说e1和e2是并发发生的。

* 在单个goroutine中，happens-before顺序取决于程序的表达顺序。

* 在下列2个条件都满足时，可以保证对变量v的读r可以观察到对变量v的写w的值：
    1. r没有发生在w之前
    2. 没有另外的对v的写w'发生在w之后和r之前
* 要保证对变量v的读r恰好可以读到对变量v的写w写入的值，需要保证以下2个条件：
    1. w发生在r之前
    2. 任何其它的对共享变量r的写w'，要么发生在w之前，要么发生在r之后
    
* 变量v会被初始化为v的类型的0值，这是个写内存的操作
* 读写一个大于单个机器字位数的变量v，是不保证读写顺序的。
    * 比如，机器字是32位的，变量v是64位的，那么可能先读高32位，再读低32位；也可能先读低32位，再读高32位。
### 同步
#### 初始化
go初始运行一个goroutine，但是这个goroutine有可能创建另一个并发的goroutine

如果一个包p引用了包q，那么q的init函数会在p的init函数之前执行。

main.main函数会在所有的init函数之后执行。
#### goroutine的创建
* 创建一个新的goroutine发生在goroutine执行之前
#### goroutine的销毁
* goroutine的退出不保证发生在任何事件之前
#### channel通讯
通过channel通讯是goroutine之间的主要同步方法。
* 向一个channel发送数据发生在从这个channel接收数据之前
* 一个channel的关闭发生在由于关闭而从这个channel接收到0值之前
* 从一个无缓冲的channel接收数据发生在发送数据到这个channel的完成之前

例如：
```go
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
```
1. 保证打印出"hello, world"。对a的赋值发生在从c接收之前，从c接收发生在将0发送给c之前，将0发送给c发生在print(a)之前。
2. 如果c是一个带缓冲的channel(例如：c = make(chan int, 1)),则不保证能打印出"hello, world",可能打印出空字符串。

* 对于一个容量(capacity)为C的channel来说，在此channel上的第k次接收发生在第k+C次发送之前。
    * 这一条规则说明允许通过一个带缓冲的channel对计数信号量建模。
    * channel中的元素的个数对应当前协程的个数
    * channel的容量对应最大并发数
    * 向channel发送一个元素以获取一个信号量(semaphore), 从channel接收一个元素以释放信号量

例如：以下代码确保最多有3个协程同时运行
```go
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
```
#### 锁
sync包实现了2种锁的数据类型：sync.Mutex 和 sync.RWMutex

* 对于任意的 sync.Mutex 或 sync.RWMutex 的变量 l，如果 n < m , 那么第n次调用l.Unlock()发生在第m次l.Lock()返回之前。

例如：
```go
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```
以上代码保证打印出"hello, world"。函数f中的l.Unlock()发生在函数main中的第二次l.Lock()返回之前。

* 对于sync.RWMutex变量l的任一l.RLock的调用，这里都有n个同样的调用l.RLock发生在第n次调用l.Unlock之后。
并且对应的l.RUnlock发生在第n+1次l.Lock之前。

#### 一次（Once）
sync包中提供了一个多goroutine的情况下安全初始化的机制，Once类型。多个goroutine可以同时调用once.Do(f)，
但是函数f仅会被精确的调用一次，其他的调用会被阻塞，直到f的调用返回。

* once.Do(f)仅一次对f的调用发生在所有其他once.Do(f)返回之前

示例如下：
```go
package week4

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var a string
var c int
var once sync.Once

func setup() {
	c++
	t := time.Now()
	s := string(t.Format(time.RFC3339Nano))
	fmt.Printf("c=%d\n", c)
	a = "Hello World " + s
}

func doprint() {
	once.Do(setup)
	fmt.Printf("%s: %s\n", string(time.Now().Format(time.RFC3339Nano)), a)
}

func twoprint() {
	go doprint()
	go doprint()
	time.Sleep(time.Second * 1)
}

func TestTwoprint(t *testing.T) {
	twoprint()
}
```

运行结果：
```text
=== RUN   TestTwoprint
c=1
2019-04-17T21:36:27.700258+08:00: Hello World 2019-04-17T21:36:27.700106+08:00
2019-04-17T21:36:27.700279+08:00: Hello World 2019-04-17T21:36:27.700106+08:00
--- PASS: TestTwoprint (1.00s)
PASS
```

从示例可以看出有2个gorotine都调用了once.Do(setup),但是setup函数只执行了一次且发生在2个goroutine打印语句之前。
因为c只打印了一次，2次打印中Hello World之前的时间戳不同，但其后的时间戳是相同的且都小于Hello World之前的时间戳。
#### 错误的同步
对变量v的读r也许可以观察到跟r并发的一个对v的写w对v的赋值，但是即便它发生了，也不意味着，r之后的读可以观察到w之前的写
对v的赋值。

比如：
```go
var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
```
g函数可能打印出2和0。即b被赋值为2，但a还没有被赋值为1.

这一事实，使得一些常用做法失效。

双重检查常用来避免加锁的开销。错误的用法，如下：
```go
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```
这里有可能打印出空字符串。

另一种错误的用法是循环等待。例如：
```go
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
```
main函数有可能打印出空字符串。更糟糕的是main函数可能陷入死循环，永远无法结束。

更加微妙的变体，例如：
```go
type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
```
即使在main函数中已经观察到了g != nil 并退出循环，也不能保证g.msg已经被初始化了。

####评论
* 在go语言中一定要小心，不要想当然的认为代码是顺序执行的。
* go的编译器有可能对代码进行重排序。
* 并发的goroutine之间是互不相认的，在需要同步的场景下需要注意

---

## Tip
### qmake 通过模板指定生成 makefile 的类型

TEMPLATE = subdirs

* app - 建立一个应用程序的 Makefile。这个是默认值。
* lib - 建立一个库的 Makefile
* subdirs - 在子目录中创建一个用于构建目标的 Makefile。子目录由SUBDIRS变量指定
* aux - 创建一个不构建任何东西的Makefile。当你不需要调用编译器去构建目标时，使用此选项。例如，你用解释语言开发的项目。
* vcapp - 仅用于Windows。为Visual Studio创建一个应用程序项目。
* vclib - 仅用于Windows。为Visual Studio创建一个库项目。

注意： -t 选项会覆盖 .pro 文件中指定的 TEMPLATE

---

## Share
### 05 深入浅出索引（下）- 极客时间 MySQL实战45讲

1. 回到主键索引树搜索的过程叫回表

    ```sql
    create table test(
    id int primary key,
    k int not null default 0,
    s varchar(16) not null default '',
    index k(k)
    ) engine=Innodb;
    ```

   这张表会建立2个B+树索引，一个是基于id的主键索引，另一个是基于k的索引。

   select * from test where k between 3 and 5;
   
   这条sql语句会先根据k索引找到符合条件的k值对应的主键id，然后再根据id在主键索引上查找整行数据。这个过程就叫做回表。

2. 覆盖索引
	* 如果一个索引覆盖了所要查询的所有数据，则叫做覆盖索引
	* 覆盖索引是常用的优化手段，因为覆盖索引可以减少树的搜索
	
	select id from test where k between 3 and 5;
	
	这条sql在k的索引中已经覆盖了查询需求，所以不需要回表查询数据，减少了在主键索引上的搜索过程。
	
3. 最左前缀原则
	* B+树支持利用索引的最左前缀定位数据
	
	```sql
    create table geek_user (
    id int(10) not null,
    id_card varchar(32) default null,
    name varchar(32) default null,
    age int(11) default null,
    ismale tinyint(1) default null,
    primary key (`id`),
    key `id_card` (`id_card`),
    key `name_age` (`name`, `age`)
   ) engine=innodb
    ```
	其中name_age索引的key是（"李四"，20），("王五"，10)，（"张三"，10），（"张三"，20），先安装name排序，然后按照age排序。
	
	当查询 select * from geek_user where name like '张%'; 时，可以用上name_age索引。
	
	但是当查询 select * from geek_user where age=10；或者 select * from geek_user where name like '%三'; 时
	是无法使用name_age索引的
	
4. 建立联合索引的原则：
* （1）顺序最优原则，通过调整索引中字段的顺序来使得所需建立的索引最少	

    当需要通过name、name和age，查询表中数据时，需要考虑建立索引 key name_age (name,age),而不是建立索引 key name (name)
     和 key age (age)索引。

* （2）空间最优原则，通过调整不同的索引组合使得空间最优
	
	当需要通过name、age、name和age，查询表中数据时，需要考虑建立索引 key name_age (name,age) 和 key age (age),
    而不是建立索引 key age_name (age,name) 和 key name (name)。
    因为name比age占用更多的空间。
	
5. 索引下推：
	* MySQL5.6之前版本不支持索引下推
	* MySQL5.6及其之后版本支持下推，可以减少回表次数
	
	select * from geek_user where name like '张%' and age=10 and ismale=1;
	
	在5.6之前执行以上sql：
	  * （1）通过name_age索引中找到符合name条件的数据，拿到id。
	  * （2）回表，在主键索引中找到id的数据并判断是否age和ismale字段是否符合条件

	在5.6及其之后版本执行以上sql：
	  * （1）通过name_age索引找到符合name条件的数据
	  * （2）在name_age索引上判断age是否符合条件，如果不符合则直接略过，如果符合则获取id
	  * （3）回表，在主键索引中找到id的数据，并判断ismale字段是否符合条件

