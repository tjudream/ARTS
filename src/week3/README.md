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