# week3 ARTS
## Algorithm 
### Leetcode [1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/)
### 问题描述
有效的括号对的定义：（）或（A)或A+B, 其中A和B也是有效的括号对。

简单的括号对的定义：给定一个括号对有效的括号对S，若S非空，且不可被拆分成任意个有效括号对，则S是简单括号对。

给定一个有效括号对S， S = P_1 + P_2 + ... + P_k ，其中P_i是简单括号对。要求去掉S的最外层的括号并返回。

### 举例
1. 
    * 输入： "(()())(())"
    * 输出： "()()()"
2. 
    * 输入："(()())(())(()(()))"
    * 输出："()()()()(())"
3.  
    * 输入："()()"
    * 输出：""
### 分析：
典型的考察栈的问题。
1. 初始化一个空栈stack
2. 遇到左括号，如果此时stack非空，则结果字符串result += "(", 入栈 stack.push("(")
3. 遇到右括号，出栈 stack.pop() ，如果此时stack非空，则结果字符串 result += ")"

注意：
   1. 遇到左括号要先判断stack是否为空，然后再入栈
   2. 遇到右括号要先出栈，然后再判断stack是否为空
### golang 代码
```go
func removeOuterParentheses(S string) string {
	strArr := []rune(S)
	var sta,ret []rune
	for i := 0; i < len(S); i++ {
		switch strArr[i] {
		case '(':
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
			sta = append(sta, strArr[i])
		case ')':
			sta = sta[:len(sta) - 1]
			if len(sta) > 0 {
				ret = append(ret, strArr[i])
			}
		}
	}
	return string(ret)
}
```

这里用rune数组代替栈

##Review
### 阅读文章 [Understanding The Memory Model Of Golang : Part 1](https://medium.com/@edwardpie/understanding-the-memory-model-of-golang-part-1-9814f95621b4)
这篇文章的其实就是为了让你看视频，Golang的内存模型主要都在视频中讲解的。

视频内容如下：
#### 主要讲解内容
* 程序内存模型
* 分配内存和垃圾回收
* 怎样监控内存的使用

#### 经典的内存模型
* 静态内存区：存储静态数据和代码指令，通常比较小
* 栈：后进先出，存储程序当前位置，存储局部变量等
    * 存储当前程序的执行状态
    * 由栈帧组成
    * 每个栈帧存储一个函数调用和本地变量
    * 一个栈相当于一个执行状态
    * 栈的大小通常是预先分配的
    * 对于并发程序需要复制整个实体模型
* 堆：
    * 具有可伸缩性
    * 大小受限于系统内存
    * 存储指针、数组、大的数据结构
#### Go的内存模型
* 每个协程都包含一个堆和一个栈
* 栈大小是可伸缩的
* 栈是从堆中借用内存的

```go
type Artist struct {
	Name                string
    Performances        []Event
	InstrumentsPlayed   []Instrument
}
```
* 在堆上分配内存的写法
```go
    slash := new(Artist)
```
* 在栈上分配内存的写法
```go
    axel := &Arist{}
    axel := &Arist{
    	Name: "Axel Rose",
    	...
    }
```

  如果是在c语言中，栈比较小的情况下可能导致栈内存溢出，但是在go中，go会扩展栈内存
  
#### 分配内存和垃圾回收
* 在一个函数中new一个对象时会分配一个堆内存
* 当函数返回之后，被分配的内存会被标记为可复用，仅仅是标记
* go的垃圾回收算法：标记清除法
    1. 第一步，垃圾收集器标记所有可达的对象
    2. 第二步，清除所有不可达的
    3. go现在已经从stop-the-world模式转到了并发清除模式
#### 监控内存的使用
* Benchmark 监控内存
* struct中成员变量的顺序会影响struct占用内存的大小
```go
type A1 struct {
	b bool
	a float64
	c int32
}
type A2 struct {
	a float64
	c int32
	b bool
}

func main() {
	a1 := A1{}
	a2 := A2{}
	fmt.Println(unsafe.Sizeof(a1)) //24
	fmt.Println(unsafe.Sizeof(a2)) //16
}
```
* 使用 profiler 监控

#### go在什么情况下会出现内存泄露
* 僵尸协程或者未关闭reader
* go申请了系统资源但是使用之后没有还回给系统
* 使用 runtime.GC() 强制执行垃圾回收

## Tip
### gdb调试代码时如何打印一块内存
* 使用 p *(a.b.data)@1000 ,但是打印太多最后会出现...
* 使用 x /1000uh buf
    * 其中 1000 表示1000个单位
    * u 表示按十六进制显示
    * h 表示双字节一个单位
* 更多用法可以参考 [http://visualgdb.com/gdbreference/commands/x](http://visualgdb.com/gdbreference/commands/x)

## Share
### 04 深入浅出数据库索引(上)
* 索引的出现就是为了提高数据的查询效率，就像书的目录
#### 索引常见的模型
* 哈希
    * key-value 存储
    * 采用链表解决hash冲突
    * 插入和查找单值的速度快O(1)
    * 不适合范围查找
    * 适合于等值查询的场景，比如memcached等nosql
* 有序数组
    * 在等值查询和范围查询的场景中性能都非常优秀，可以采用二分法O(lg(N)))
    * 插入成本高
    * 只适用于静态存储引擎，保存不需要改变的静态数据
* 搜索树
    * 二叉搜索树(每个节点的左儿子小于父节点，父节点小于右儿子), 查询和更新的复杂度都是O(log(N))
    * N叉树更适配磁盘的访问模式，在读写性能上优于二叉树
        * 例如，对于100万的数据，二叉树需要用20层来存储，需要读取20次磁盘，而InnoDB一个整数字段索引大概有1200叉，树高为4的时候就能大约
        存储17亿的数据量
#### InnoDB 的索引模型
* 每个索引对应一棵B+树
* InnoDB中分为主键索引和非主键索引2种索引
    * 主键索引，也称聚簇索引(clustered index)，叶子节点存储整行的数据
    * 非主键索引，也称二级索引(secondary index)，叶子节点内容存储的是主键的值
* 基于主键索引和普通索引(非主键索引)的查询的区别 
    * select * from t where primary_key=1 基于主键索引的只需要搜索主键索引这棵B+树就能查询到该行数据
    * select * from t where other_key=1 基于非主键索引需要先查询非主键索引这棵B+树，找到对应的主键，然后在用该主键在主键索引的
    B+树中找到该行数据。需要进行2次B+树的查询操作。所以如果要查询整行数据要优先使用主键索引。
#### 索引的维护
* 当按照索引顺序插入数据时，则只需在B+树的最后申请一块内存，顺序插入即可
* 当按照索引的乱序插入数据时，或者索引是无序的情况下插入数据时，则需要将插入点后面的数据进行逻辑右移。如果此时B+树的数据页已满，则还会导致
页分裂。
    * 页分裂会导致性能下降
    * 页分裂还会导致利用率下降。原本放在一页的数据，现在分裂成2页存放，空间利用率降低了大约50%
* 主键长度越小，普通索引的叶子节点就越小（因为普通索引的叶子节点存储的是主键的内容），普通索引占用的空间就越小。

* 建表时使用一个自增主键比使用业务字段作为主键的好处：
    * 可以减少页分裂的发生
    * 使得普通索引占用更少的存储空间
* 什么情况适合业务字段直接作为主键？
    1. 只有一个索引
    2. 该索引必须是唯一索引
    * 典型的场景就是KV存储 