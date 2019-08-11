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

# Review []()

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


