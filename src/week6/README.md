#week

---

## Algorithm [654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/)
### 1. 问题描述
使用给定的数组构建最大二叉树。

最大二叉树定义：数组中的最大值为该树的根，数组中最大值左侧的子数组为二叉树的左子树，最大值的右侧为二叉树的右子树，
其中左子树和右子树也都是最大二叉树

### 2. 解题思路
采用递归的算法，找到最大值，递归构造最大值左侧和右侧的子数组

1. 构造一个树节点的结构体
2. 找到数组中的最大值的索引 maxIndex
3. 用 arr[0:maxIndex] 递归构造左子树 left
4. 用 arr[maxIndex + 1: len(arr)] 构造右子树 right
5. 用最大值, left, right 构造树的根节点 root 并返回

### 3. 代码
```go
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findMaxIndex(nums []int) int {
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 1 {
		return &TreeNode{nums[0],nil, nil}
	}
	maxIndex := findMaxIndex(nums)
	val := nums[maxIndex]
	var left, right *TreeNode
	if maxIndex > 0 {
		left = constructMaximumBinaryTree(nums[0:maxIndex])
	}
	if maxIndex + 1 < len(nums) {
		right = constructMaximumBinaryTree(nums[maxIndex + 1: len(nums)])
	}
	return &TreeNode{val, left, right}
}

```
### 4. 复杂度分析
* 时间复杂度：平均复杂度 O(nlogn), 最坏情况 O(n<sup>2</sup>) 。
    * 为了找到最大值，每次需要遍历一个数组 O(n), 递归深度取决于树的高度, 平均树高 logn，所以平均时间复杂度为 O(nlogn)
    * 当数组是一个有序数组的时候, 树高是 n, 所以时间复杂度退化为 O(n<sup>2</sup>)
* 空间复杂度： 
---

## Review [How to Become a Better Software Developer](https://medium.com/devtrailsio/how-to-become-a-better-software-developer-dd16072c974e)
本篇文章给出了一些提高工作技能的方法。
### 理解端到端的过程
软件开发生命周期
1. 需求分析
2. 设计
3. 开发
4. 测试
5. 运维

以上阶段都是由不同的部分执行运作的，他们之间缺少沟通和反馈。
所以你需要了解整个生命周期，尝试去了解每个部分为什么要这么做事情。
### 理解你的用户需求
你的用户不是专家，他们不了解软件开发，也不了解软件开发的专有名词。尽量使用他们能听得懂的话与他们沟通。

跟你的客户沟通他们想要解决什么问题，而不是讨论解决方案。

跟客户之间建立信任关系。
### 用合适的工具工作
* 如果你手里有个锤子，你看什么都像钉子

不要试图用单一的技术去解决所有问题。

你可以准备一个问题列表，来尝试找出解决问题的合适的工具：
* 需要支持哪些平台和设备
* 都有哪些非功能性需求：性能？内存使用？
* 买商业版还是用开源版
* 这个解决方案是否提供了开箱即用的一切所需，还是你需要自己写一些东西
* 是否还有其他约束条件：公司政策、法律问题、团队中是否缺少专家

安全试验
* 你是否有足够的时间准备？
* 找到你需要先测试的东西
### 站在巨人的肩膀上
导致我们重复发明轮子的一些常见误解：
* 自己实现比学习第3方库更简单。
* 这个解决方案功能太多了，我只需要使用其中一小部分。
* 我们能做的更好。
* 代码所有权和长期维护将成为问题。

重复造轮子的好处：
* 我们可以通过重复造轮子来学习
### 研究你的工作方式
改进方法比学习技术更重要。如何建立一个高效的工作流：
* 团队管理和项目管理方法：可以考虑敏捷流程，使用scrum和看板
* 保持个性：可以保持自己的个性，单不要影响团队
* 工程实践：TDD、BDD、Code Review、持续集成、持续部署

选择合适团队、合适自己的工作方法。将使用的工作方法调整到合适团队合适自己的状态。
#### 消除障碍
保持警惕，将问题扼杀在摇篮里。
### 专注于基础知识
IT行业发展迅速。学习基础知识，把基础知识学牢，比无限的追求新技术更有用。很多新技术都是基于基础知识构建的。

花时间阅读一些经典书籍例如：Gregor Hohpe 和 Bobby Woolf 的《企业集成模式》，
四人组（Gang of Four）著名的《设计模式：可重用面向对象软件的元素》
### 额外提示
* 分享知识
* 不要责备自己或他人
* 不要成为一个混蛋
### 总结
这些建议可以帮助你找到一条适合你自己的路，成为一个更好的开发者。

---

## Tip
### SQL 语句分类 DDL, DML, DCL 和 TCL
| |DDL|DML|DCL|TCL|
|---|---|---|---|---|
|全称|Data Define Language|Data Manipulation Language|Data Control Language|Transaction Control Language|
|中文释义|数据定义语言|数据操作语言|数据控制语言|事务控制语言|
|SQL命令|create - 创建 database 和它的对象(table, index, views, store procedure, function, triggers)|select - 从表中检索数据|grant - 给用户赋权|commit - 提交一个事务|
| |alter - 修改已存在的数据库的结构|insert - 向表中插入数据|revoke - 收回通过grant给用户赋予的权限|rollback - 回滚一个事务|
| |drop - 从数据库中删除对象|update - 更新表中已经存在的数据| |savepoint - 回滚事务到组内创建的生成点|
| |truncate - 删除表中的所有记录，同时删除给表分配的存储空间|delete - 从表中删除数据| |set transaction - 指定事务特征|
| |comment - 给数据加注释|merge - upsert(insert or update) 更新并插入| | |
| |rename - 重命名|call - 调用 PL/SQL 或 Java 子程序| | |
| | |explain plan - 解释数据访问路径| | |
| | |lock table - 并发控制，锁表| | |

---
    

## Share
### 07  行锁功过：怎么减少行锁对性能的影响 —— 极客时间-MySQL实战45讲
行锁是针对数据库表中行记录的锁。
#### 两阶段锁
* 行锁是在需要的时候才加上，但并不是不需要了就立刻释放，而要等到事务结束时才释放。这就是两阶段锁协议。

如果你的事务中需要锁多行，要把最可能造成锁冲突的、最可能影响并发度的锁尽量往后放。

#### 死锁和死锁检测
| 事务A | 事务B |
| --- | --- |
| begin;| begin;|
|update t set k = k + 1 where id = 1; | |
| | update t set k = k + 1 where id = 2;|
|update t set k = k + 1 where id = 2;| |
| | update t set k = k + 1 where id = 1;|

A持有id=1的行锁，等待B释放id=2的行锁；B持有id=2的行锁，等待A释放id=1的行锁

出现死锁之后的两种策略：
* 等待超时。通过参数 innodb_lock_wait_timeout 来设置，默认是50s
* 发起死锁检测，主动回滚死锁链条中的某一个事务，让其他事务得以继续执行。 
innodb_deadlock_detect 设置为 on, 表示开启这个逻辑。默认开启。

死锁检测需要耗费大量的CPU资源。那么如何解决热点行更新导致的性能问题：
* 如果确保业务一定不会出现死锁，可以临时把死锁检测关闭
* 控制并发度。可以考虑采用中间件或者修改MySQL源码，对于相同行的更新采用排队策略。


