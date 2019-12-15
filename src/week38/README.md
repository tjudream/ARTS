# week38

---

# Algorithm [97. Interleaving String](https://leetcode.com/problems/interleaving-string/)
## 1. 问题描述
交叉字符串

给定字符串 s1,s2,s3 , 确定 s3 是否是由 s1 和 s2 交叉组成
#### 示例 1 :
* 输入 : s1="aabcc" , s2="dbbca", s3="aadbbcbcac"
* 输出 : true
#### 示例 2 :
* 输入 : s1="aabcc", s2="dbbca", s3="aadbbbaccc"
* 输出 : false


## 2. 解题思路
### 算法一：递归
判断 len(s3) 是否等于 len(s1) + len(s2)

如果 s1 是空串，则判断 s2 是否与 s3 相同
如果 s2 是空串，则判断 s1 是否与 s3 相同
如果 s3[0] == s1[0] 则递归判断 s1[1:len(s1)],s2,s3[1:len(s3)]
如果 s3[0] == s2[0] 则递归判断 s1,s2[1:len(s2)],s3[1,len(s3)]

递归代码
```go
    if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}
	return (s3[0] == s1[0] && isInter(s1[1:len(s1)], s2, s3[1:len(s3)])) || (s3[0] == s2[0] && isInter(s1, s2[1:len(s2)], s3[1:len(s3)]))

```
### 算法二：动态规划
dp[i][j] bool 表示 s3[0:i+j] 是否是 s1[0:i] 和 s2[0:j] 的交叉组成

dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])

如果 i==0 则 dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
如果 j==0 则 dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]

dp[0][0] = true
最后求出 dp[m][n] ,其中 m 为 s1 的长度, n 为 s2 的长度

## 3. 代码
```go
func isInterleaveDp(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}
	var dp [][]bool
	dp = make([][]bool, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]bool, l2+1)
	}
	dp[0][0] = true
	for i := 1; i < l1+1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < l2+1; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l1][l2]
}
```
## 4. 复杂度分析
动态规划算法
* 时间复杂度 : O(m*n)
* 空间复杂度 : O(m*n)

---

# Review [Exception Handling in Java Streams](https://medium.com/swlh/exception-handling-in-java-streams-5947e48f671c)
Java Stream 中的异常处理
## 未检查的异常(Unchecked Exceptions)
```java
List<String> integers = Arrays.asList("44", "373", "xyz", "145");
integers.forEach(str -> {
    try {
        System.out.println(Integer.parseInt(str));
    }catch (NumberFormatException ex) {
        System.err.println("Can't format this string");
    }
}
);
```
这样有效，但是编写的代码不够简洁易读.
可以将异常处理移到其他地方：

```java
static Consumer<String> exceptionHandledConsumer(Consumer<String> unhandledConsumer) {
    return obj -> {
        try {
            unhandledConsumer.accept(obj);
        } catch (NumberFormatException e) {
            System.err.println(
                    "Can't format this string");
        }
    };
}
public static void main(String[] args) {
    List<String> integers = Arrays.asList("44", "xyz", "145");
    integers.forEach(exceptionHandledConsumer(str -> System.out.println(Integer.parseInt(str))));
}
```
这段代码可以改为使用泛型，从而适应更多场景：
```java
static <Target, ExObj extends Exception> Consumer<Target> handledConsumer(Consumer<Target> targetConsumer, Class<ExObj> exceptionClazz) {
    return obj -> {
        try {
            targetConsumer.accept(obj);
        } catch (Exception ex) {
            try {
                ExObj exCast = exceptionClazz.cast(ex);
                System.err.println(
                        "Exception occured : " + exCast.getMessage());
            } catch (ClassCastException ccEx) {
                throw ex;
            }
        }
    };
}
``` 
使用以上代码，我们最初的代码可以简化为
```java
List<String> integers = Arrays.asList("44", "373", "xyz", "145");
integers.forEach(
        handledConsumer(str -> System.out.println(Integer.parseInt(str)), 
        NumberFormatException.class));
```
如果我们要捕获 ArithmeticException 异常，则可以：
```java
List<Integer> ints = Arrays.asList(5, 10, 0, 15, 20, 30, 0, 9);
ints.forEach(
        handledConsumer(
                i -> System.out.println(1000 / i),
                ArithmeticException.class));
```
## 检查的异常(Checked Exceptions)
```java
List<Integer> list = Arrays.asList(5, 4, 3, 2, 1);
    list.forEach(i -> {
        try {
            Thread.sleep(i);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    });
}
```
1. 创建一个新接口处理检查异常
```java
@FunctionalInterface
public interface HandlingConsumer<Target, ExObj extends Exception> {
    void accept(Target target) throws ExObj;
}
```
2. 添加一个静态方法，并将检查异常转换为 RuntimeException。因为这样可以使得我们能够
处理调用方法中的异常并释放 lambda 来完成其实际工作
```java
@FunctionalInterface
public interface HandlingConsumer<Target, ExObj extends Exception> {
    void accept(Target target) throws ExObj;
    static <Target> Consumer<Target> handlingConsumerBuilder(
            HandlingConsumer<Target, Exception> handlingConsumer) {
        return obj -> {
            try {
                handlingConsumer.accept(obj);
            } catch (Exception ex) {
                throw new RuntimeException(ex);
            }
        };
    }
}
```
现在我们的代码可以简化为：
```java
List<Integer> list = Arrays.asList(5, 4, 3, 2, 1);
list.forEach(handlingConsumerBuilder(i->Thread.sleep(i)));
```

---

# Tip VS Code 配置 Golang 开发环境
1. 官网下载 [Visual Studio Code](https://code.visualstudio.com/)
2. 安装 Go 语言扩展 [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
3. 配置代理，安装 Go 语言相关包
```jshelllanguage
export http_proxy=http://127.0.0.1:1087
export https_proxy=http://127.0.0.1:1087
```
安装包
```gotemplate
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/cweill/gotests/...
go get -u -v github.com/josharian/impl
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
```
各个包的说明：
* gocode 代码补全
* godef 跳转到定义
* guru 获得代码引用
* gorename 重命名源码文件
* goimports 自动格式化 import
* gopkgs 列出包
* golangci-lint 静态代码检查
* go-outline 文件大纲
* go-symbols 工作区符号搜索
* gogetdoc 显示方法签名
* gomodifytags tags管理
* gotests 测试
* impl 自动生成接口实现
* fillstruct 自动填充 struct 的默认值


---
    
# Share 40 | insert语句的锁为什么这么多？ —— 极客时间 MySQL实战45讲
## insert ... select 语句
```roomsql

CREATE TABLE `t` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `c` int(11) DEFAULT NULL,
  `d` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `c` (`c`)
) ENGINE=InnoDB;

insert into t values(null, 1,1);
insert into t values(null, 2,2);
insert into t values(null, 3,3);
insert into t values(null, 4,4);

create table t2 like t
```
可重复读隔离级别下，binlog_format=statement，执行
```roomsql
insert into t2(c,d) select c,d from t;
```
执行这个语句时，是否需要对 t 的所有行加间隙锁呢？

| session A | session B |
| --- | --- |
| insert into t values(-1,-1,-1); | insert into t2(c,d) select c,d from t; |

如果 B 先执行，则对 t 主键索引加了 (-∞,1] 这个 next-key lock ，会在语句执行完成后，才允许 A 的 insert 语句执行

如果没有锁，就可能出现 B 的 insert 语句先执行，但写入 binlog 的情况是
```roomsql
insert into t values(-1,-1,-1);
insert into t2(c,d) select c,d from t;
```
这个语句在备库执行，就会把 id=-1 这一行也写到 t2 中，出现主备不一致

## insert 循环写入
执行 insert ... select 的时候，对目标表也不是锁全表，而是只锁住需要访问的资源

往表 t2 中插入一行数据，这一行的 c 值是表 t 中 c 值的最大值加 1
```roomsql
insert into t2(c,d)  (select c+1, d from t force index(c) order by c desc limit 1);
```
这个语句的加锁范围，就是表 t 索引 c 上的 (3,4] 和 (4,supremum] 这两个 next-key lock,以及主键索引上 id=4 这一行

执行流程：从表 t 中按照索引 c 倒序，扫描第一行，拿到结果写入到表 t2 中。
因此整个语句的扫描行数是 1
![slow_log](slow_log.png)
慢查询日志，Rows_examined=1 表示扫描行数为 1

如果要把这条数据插入到 t
```roomsql
insert into t(c,d)  (select c+1, d from t force index(c) order by c desc limit 1);
```
慢查询日志
![slow_log_t](slow_log_t.png)
explain 结果
![explain](explain.png)
Using temporary 表示使用了临时表

![Innodb_rows_read](Innodb_rows_read.png)
这个语句执行前后，Innodb_rows_read 的值增加了 4。因为默认临时表是 Memory 引擎，所以这 4 行查的都是表 t，也就是说对 t 做了全表扫描

整个执行流程:
1. 创建临时表，表里有两个字段 c 和 d
2. 按照索引 c 扫描表 t，依次取 c=4,3,2,1 ， 然后回表，读到 c 和 d 的值写入临时表。这时 Rows_examined=4
3. 由于语义里面有个 limit 1，所以只取了临时表的第一行，再插入到表 t，这时，Rows_examined=5

这个语句会在 t 上做全表扫描，并且会给索引 c 上的所有间隙都加上共享的 next-key lock。所以，这个语句执行期间，其他事务不能在这个表上
插入数据。

使用临时表示因为，这类一边遍历数据，一边更新数据的情况，如果读出来的数据直接写回原表，就可能在遍历过程中，读到刚刚插入的记录，新插入的
记录如果参与计算逻辑，就跟语义不符。

由于实现上这个语句没有在子查询中就直接使用 limit 1，从而导致了这个语句的执行需要遍历整个 t 表。优化
```roomsql
create temporary table temp_t(c int,d int) engine=memory;
insert into temp_t  (select c+1, d from t force index(c) order by c desc limit 1);
insert into t select * from temp_t;
drop table temp_t;
```
## insert 唯一键冲突
| session A | session B |
| --- | --- |
| insert into t values(10,10,10);<br/><br/>begin;<br/>insert into t values(11,10,10);<br/>(Duplicate entry '10' for key 'c') | |
| | insert into t values(12,9,9);<br/>(blocked) |

可重复隔离级别，B 进入了锁等待

A 持有 c 上的 (5,10] 共享 next-key lock （读锁）

死锁场景

| | session A | session B | session C |
| --- | --- | --- | --- |
| T1 | begin;<br/>insert into t values(null,5,5); | | |
| T2 | | insert into t values(null,5,5);| insert into t vlues(null,5,5); |
| T3 | rollback; | | (Deadlock found) |

A 执行 rollback 回滚时，C 几乎同时发现了死锁并返回

死锁产生的逻辑：
1. 在 T1 时刻，启动 A，并执行 insert ，此时在索引 c 的 c=5 上加了记录锁。这个索引是唯一索引，因此退化为记录锁
2. 在 T2 时刻， B 执行 insert，发现唯一键冲突，加上读锁；同样，C 也在索引 c 的 c=5 这一纪录上，加了读锁
3. T3 时刻， A 回滚，B 和 C 都试图继续执行插入操作，都要加上锁。两个 session 都要等待对方的行锁，所以就出现了死锁

状态变化图
![status_change](status_change.jpg)
## insert into ... on duplicate key update
上面的例子是主键冲突后直接报错，如果改写成
```roomsql
insert into t values(11,10,10) on duplicate key update d=100; 
```
就会给索引 c 上 (5,10] 加上一个排他的 next-key lock(写锁)

insert into ... on duplicate key update 这个语义的逻辑是，插入一行数据，如果碰到唯一键约束，就执行后面的更新语句。

如果有多个列违反了唯一键约束，就会按照索引的顺序，修改跟第一个索引冲突的行。

现在 t 中有 (1,1,1) 和 (2,2,2) 下面语句的执行过程
![key_conflict](key_conflict.png)
主键索引先判断，与 id=2 这一行冲突，所以修改 id=2 的行

affected rows 返回是 2，实际，真正更新的只有一行，只是在代码上实现了， insert 和 update 都认为自己成功了，update 计数加 1，
insert 计数也加了 1.

