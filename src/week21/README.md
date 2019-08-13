# week21

---

# Algorithm [763. Partition Labels](https://leetcode.com/problems/partition-labels/)
## 1. 问题描述
分区标签

给一个全由小写字母组成的字符串 S。将 S 分成若干区间，要求同一个字母只能出现在其中一区间中，且要求尽可能多地分区。返回每个段的长度。

* 示例1：
* 输入：S="ababcbacadefegdehijhklij"
* 输出：[9,7,8]
* 解释：分区是 "ababcbaca", "defegde", "hijhklij"。这个分区中一个字母只出现在其中一个分区中。
"ababcbacadefegde", "hijhklij" 这样的分区是不对的，因为不是最大分区。

## 2. 解题思路
思路：
1. 找到第一字字母，假设是 a，那么第一个区间包含所有的 a，所以要找到最后出现的 a 的位置
2. 第一个 a 和最后一个 a 之间所有出现的字母，如果在最后一个 a 后边的子串中出现，则要将该区间右移到最后一个字母出现的位置
3. 直到第一个区间中任何字母，在后边的子串中没有出现过为止
4. 重复1、2、3步，找到所有区间

算法：
1. 设置一个长度为 26 的整型数组 last[S[i] - 'a'] ，将每个字母在 S 中出现的最后的位置 index 存储在这个数组中
2. 设置两个变量 start，end 代表要查找区间的起始位置
3. 遍历 S
4. 查看当前字母在 S 中出现的最后位置 last[S[i] - 'a']
5. 如果 end < last[S[i] - 'a'], 则 end = last[S[i] - 'a']
6. 如果 end == i，则找到了一个区间，将第一个区间的长度 end - start + 1 加入到结果中，同时重新计算 start = i + 1

## 3. 代码
```go
func partitionLabels(S string) []int {
	res := []int{}
	l := len(S)
	last := make([]int,26)
	for i := 0; i < l; i++ {
		last[S[i] - 'a'] = i
	}
	start,end := 0,0
	for i := 0; i < l; i++ {
		if last[S[i] - 'a'] > end {
			end = last[S[i] - 'a']
		}
		if end == i {
			r := end - start + 1
			start = end + 1
			res = append(res, r)
		}
	}
	return res
}
```
## 4. 复杂度分析
* 时间复杂度： O(N) N 为 S 的长度
* 空间复杂度： O(1)

---

# Review [10 Steps to Become a Millionaire in 5 Years (or Less)](https://medium.com/better-marketing/10-skills-to-becoming-a-millionaire-in-5-years-or-less-e16b8b20500c)
5年内成为百万富翁的10步法
## 1. 设定一个愿景
* 想想比知识更重要
* 创造力比经验更重要
## 2. 开发一个90天的系统来衡量进步和未来的节奏
每隔90天进行一次复盘：
* 跟踪你的钱
* 跟踪你的健康
* 跟踪你的时间
* 跟踪你想要成功的领域的进展
## 3. 养成一种生活在心流或巅峰状态的日常习惯
你想成为谁？
* 想象他就是自己
* 感受自己已经成为了他
* 假设你的愿望已经实现
* 知道你想要什么，你就能拥有什么
* 承诺
* 投身于现实
* 从现在开始，按照现实行事
* 享受来自现在和和谐的心流
## 4. 为清晰、恢复和创造性设计你的环境
你需要重塑你的环境。

当你改变一个部分时，你就改变了整个系统。不要让一个烂苹果坏了一筐。

## 5. 关注结果，而非习惯或过程
目标是手段，而不是目的。它们是增长和进步的手段。一旦你达到了一个目标，你就会吸取你所学到的，并继续扩展。
## 6. 确定理想的导师或合作伙伴
通过做有用的人，你可以发展导师关系和伙伴关系。你奉献你的思想和努力去帮助他们。通过帮助他们，你把自己定位在一个独特的位置。
在这个独特的新职位上，赚很多钱变得很容易。
## 7. 成为一名出色的倾听者和观察者
作为一个倾听者，认证倾听，其他人知道你真正倾听他们并且你真正想要帮助他们。 他们会爱和尊重你，因为与大多数人不同，你是真诚的。 你是一个倾听者。
## 8. 关注是谁，而不是如何
不要关注怎么做这件事，而是关注谁能帮你做这件事。

你需要建立一个团队。像其他所有事情一样，你想在准备好之前做到这一点。
事实上，在你开始之前，你从来都没有准备好。
你从来都没有资格做任何事情。
它总是飞跃的本身，然后通过工作的过程，是你合格。
## 9. 不断更新你对成功的价值观或定义
变革性的经历可以改变你的生活。 同样，变革关系可以改变你的生活。

如果你在过去的12个月中对“成功”的定义没有改变，那么你还没有学到东西还不够多。
如果你对成功的定义没有改变，那么你就没有足够的经验。
## 10. 当你知道是时候改变的时候，不要等太久
不断地成长、转变、改变、奋斗。这才是健康的生活方式。

---

# Tip

## 

---
    
# Share 23 MySQL是怎么保证数据不丢的？ —— 极客时间 MySQL实战45讲
## binlog 的写入机制
事务执行过程中，先把日志写到 binlog cache，事务提交的时候，再把 binlog cache 写到 binlog 文件中

一个事务的 binlog 是不能被拆开的，因此不论这个事务多大，也要确保一次写入。这涉及到 binlog cache 的保存问题

系统给 binlog cache 分配了一片内存，每个线程一个，参数 binlog_cache_size 用于控制单个线程内 binlog cache 所占内存的大小，
如果超过这个大小，就要暂存到磁盘上。

事务提交时，执行器把 binlog cache 里完整的事务写入到 binlog 中，并清空 binlog cache

![binlog_disk](binglog_disk.png)

每个线程有自己的 binlog cache，但是共用同一份 binlog 文件。
* 图中 write，指的就是把日志写入到文件系统的 page cache，并没有把数据持久化到磁盘，所以速度比较快。
* 图中 fsync，才是将数据持久化到磁盘，fsync 才占用磁盘的 IOPS

write 和 fsync 的时机，是由参数 sync_binlog 控制的：
1. sync_binlog = 0 的时候，表示每次提交事务都只 write，不 fsync
2. sync_binlog = 1 的时候，表示每次提交事务都会执行 fsync
3. sync_binlog = N (N > 1) 的时候，表示每次提交事务都 write，但是累计 N 个事务后才 fsync

因此，在出现 IO 瓶颈的场景里，将 sync_binlog 设置成一个比较大的值，可以提升性能。

在实际的业务场景中，考虑到丢失日志量的可控性，一般不建议将这个参数设置成 0，比较常见的是将其设置成 100~1000 中的某个数值

如果将 sync_binlog 设置为 N，对应的风险是：如果主机发生异常重启，会丢失最近 N 个事务的 binlog 日志
## redo log 的写入机制
事务在执行过程中，生成的 redo log 要先写到 redo log buffer

redo log buffer 里面的内容，并不是每次生成后都直接写入到磁盘

如果事务执行期间 MySQL 发生异常重启，这部分日志就丢失了，但是由于事务并没有提交，所以这时日志丢了也不会有损失。

但是在事务没提交之前，redo log buffer 中有可能有部分日志已经被持久化到磁盘。

![redolog_status](redolog_status.png)
redo log 的三种状态：
1. 存在 redo log buffer 中，物理中存在于 MySQL 进程的内存中，图中红色部分
2. 写到磁盘(write),但是没有持久化(fsync),物理上是在文件系统的 page cache 中，图中黄色部分
3. 持久化到磁盘，对应的 hard disk， 图中绿色部分

InnoDB 有一个后台线程，每隔 1 秒，就会把 redo log buffer 中的日志，调用 write 写到文件系统的 page cache 中，然后调用 fsync 持久化到磁盘。

事务执行期间，redo log 也是直接写到 redo log buffer 中的，这些 redo log 也会被后台线程一起持久化到磁盘中。
一个没有提交的事务的 redo log，也可能已经持久化到磁盘

除了后台线程每秒轮询操作外，还有两种场景会让一个没有提交的事务的 redo log 写入到磁盘：
1. redo log buffer 占用空间即将达到 innodb_log_buffer_size 一半的时候，后台线程会主动写磁盘。
由于这个事务并没用提交，所以这个写盘动作只调用了 write，而没有调用 fsync，也就是在 page cache 中。
2. 并行事务提交的时候，顺带将这个事务的 redo log buffer 持久化到磁盘。
假设事务 A 执行到一半，已经写了一些 redo log 到 buffer， 这时候另一个线程的事务 B 提交，如果 innodb_flush_log_at_trx_commit = 1,
那么事务 B 要把 redo log buffer 里的日志全部持久化到磁盘。这时候会带上事务 A 的 redo log buffer 中的日志。

两阶段提交时，时序上 redo log 先 prepare ，再写 binlog，最后再把 redo log commit

如果 innodb_flush_log_at_trx_commit=1 那么 redo log 在 prepare 时就持久化一次，因为有一个崩溃恢复的逻辑依赖于 prepare 
的 redo log，再加上 binlog 来恢复。

每秒后台轮询刷盘，再加上崩溃恢复，InnoDB 就认为 redo log 在 commit 的时候不需要 fsync 了，只会 write 到 page cache 就够了。

通常我们说的 MySQL 的"双1"配置，就是指 sync_binlog=1 和 innodb_flush_log_at_trx_commit=1 也就是说，一个事务完整提交前，
需要等待两次刷盘，一次是 redo log (prepare阶段),一次是 binlog

这意味着我们从 MySQL 看到的 TPS 是每秒两万的话，每秒就会写四万次磁盘，但是，我们用工具测试出来，磁盘能力也就两万左右，
怎么能实现两万的 TPS 呢？这主要用到 组提交(group commit) 的机制。

日志逻辑序号(log sequence number, LSN)是单调递增的，用来对应一个 redo log 的一个个写入点。每次写入长度 length 的 redo log，
LSN 的值就会加上 length。

LSN 也会写到 InnoDB 的数据页中，来确保数据页不会被多次执行重复的 redo log。

![redolog_group](redolog_group.png)
图中是三个并发事务(trx1,trx2,trx3)在 prepare 阶段，都写完 redo log buffer，持久化到磁盘的过程，对应的 LSN 分别是 50,120,160

1. trx1 是第一个到达的，会被选为这个组的 leader
2. 等 trx1 要开始写盘的时候，这个组里已经有三个事务，这时候 LSN 变成了 160
3. trx1 去写盘的时候，带的就是 LSN=160，因此等 trx1 返回时，所有 LSN <= 160 的 redo log，都已经被持久化到磁盘
4. 这时候 trx2，trx3 就可以直接返回了

所以，一次组提交里面，组员越多，越节约磁盘的 IOPS。但如果只是单线程压测，那就只能一个事务对应一次持久化操作。

在并发场景下，第一个事务写完 redo log buffer 以后，接下来这个 fsync 越晚调用，组员可能越多，越有可能更好地节约磁盘的 IOPS。

为了让一次 fsync 带更多的组员，MySQL 进行了优化：拖时间。两阶段提交如图所示：
![two_commit](two_commit.png)
图中写 binlog 实际是两步：
1. 先把 binlog 从 binlog cache 写到磁盘的 binlog 文件中
2. 调用 fsync

MySQL 为了让组提交的效果更好，把 redo log 做 fsync 的时间拖到步骤1 之后。如图所示：
![two_commit_detail](two_commit_detail.png)
这样，binlog 也可以组提交了。上图步骤4 中，binlog 的 fsync ，如果有多个 binlog 已经写完，也是一起持久化的，这也可以减少 IOPS。

不过通常情况下，第3步执行的会很快，所以 binlog 的 write 和 fsync 间隔时间短，导致一起持久化的 binlog 较少，因此 binlog 的
组提交效果不如 redo log 的好。

如果想提高 binlog 的组提交效果可以设置以下参数：
1. binlog_group_commit_sync_delay 表示延迟多少微妙后才调用 fsync
2. binlog_group_commit_sync_no_delay_count 表示累积多少次以后才调用 fsync

这两个条件是或的关系，满足其中之一就会触发 fsync

所以当 binlog_group_commit_sync_delay=0 的时候 binlog_group_commit_sync_no_delay_count 也是无效的。

WAL 机制是减少磁盘写，可是每次提交事务都要写 redo log 和 binlog ，那怎么减少磁盘写呢？
1. redo log 和 binlog 都是顺序写，磁盘的顺序写比随机写速度要快
2. 组提交机制，可以大幅降低磁盘的 IOPS 消耗

### 如果你的 MySQL 现在出现了性能瓶颈，而且瓶颈在 IO 上，可以通过哪些方法来提升性能呢？
1. 设置 binlog_group_commit_sync_delay 和 binlog_group_commit_sync_no_delay_count 参数，减少 binlog 的写盘次数。
这个方法是基于“额外的故意等待”来实现的，因此可能会增加语句的响应时间，但没有丢数据的风险。
2. 将 sync_binlog 设置为大于 1 的值(比较常见的是 100~1000)。这样做的风险是，主机掉电时会丢 binlog 日志。
3. 将 innodb_flush_log_at_trx_commit 设置为 2。这样做的风险是，主机掉电时会丢数据。

不建议将 innodb_flush_log_at_trx_commit 设置成 0。因为设置为 0 的话，表示 redo log 只保存在内存中，这样 MySQL 本身
异常重启也会丢数据，风险太大。而 redo log 写 page cache 速度也很快，所以设置成 2 和 0 性能差不多，但是设置成 2 的时候，
MySQL 异常重启不会丢数据，相比之下风险小很多。

## 问题

### 问题1：执行一个 update 语句以后，再去执行 hexdump 命令直接查看 ibd 文件的内容，为什么没有看到数据改变？
可能是因为 WAL 机制的原因。 update 执行完以后，InnoDB 只保证写完了 redo log、内存，可能还没来得及写磁盘。
### 问题2：为什么 binlog cache 是每个线程自己维护的，而 redo log buffer 是全局共享的？
MySQL 如此设计的原因是，binlog 不能“被打断”。一个事务的 binlog 必须连续写，因此要整个事务完成后，再一起提交写到文件里。

而 redo log 并没有这个要求，中间有生成的日志可以写到 redo log buffer 中。redo log buffer 中的内容还能“搭便车”，
其他事务提交的时候可以被一起写入磁盘。
### 问题3：事务执行期间，还没到提交阶段，如果发生 crash 的话，redo log 肯定丢了，这会不会导致主备不一致？
不会，因为这时候 binlog 也还在 binlog cache 里，没有发给备库。crash 以后 redo log 和 binlog 都没有了，从业务角度看事务
没有提交，所以数据是一致的。
### 问题4：如果 binlog 写完磁盘以后发生 crash，这时候还没给客户端答复就重启了。等客户端再重连进来，发现事务已经提交成功，这是不是 bug ？
不是。

可以想一个极端的情况，整个事务都提交成功了，redo log commit 完成了，备库也收到 binlog 并执行了。但是主库和客户端网络断开了，
导致事务成功的包返回不回去，这时候客户端也会收到“网络断开”的异常。这种情况只能算是事务成功的，不能认为是 bug。

实际上数据库的 crash-safe 保证是：
1. 如果客户端收到事务成功消息，事务就一定持久化了
2. 如果客户端收到事务失败(比如主键冲突，回滚等)的消息，事务就一定失败了
3. 如果客户端收到“执行异常”的消息，应用需要重连后通过查询当前状态来继续后续的逻辑。
此时数据库只需要保证内部(数据和日志之间，主库和备库之间)一致就可以了。

### 思考题：什么场景下需要把线上生产环境设置成“非双1”
1. 业务高峰期。一般如果有预知高峰期，DBA 会有预案，把主库设置成“非双1”
2. 备库延迟，为了让备库尽快赶上主库。
3. 用备份恢复主库的副本，应用 binlog 的过程，这个跟上一种场景类似
4. 批量导入数据时

一般把生产库改成“非双1”配置，是设置 innodb_flush_logs_at_trx_commit=2 , sync_binlog=1000

