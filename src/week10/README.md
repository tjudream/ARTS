#week10

---

## Algorithm [1008. Construct Binary Search Tree from Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)
### 1. 问题描述
用先序遍历数组构造二叉搜索树。

给定一个二叉搜索树的先序遍历数组，构造二叉搜索树。
### 2. 解题思路
先序遍历，即先遍历根，然后遍历左子树，最后遍历右子树。所以数组的第一个元素即是根。

构造出树的根之后，剩下的元素可以采用递归插入，或者非递归插入。
### 3. 代码
* 递归
```go
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root *TreeNode
	for _,num := range preorder {
		root = insertNode(num, root)
	}
	return root
}
func insertNode(num int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{num, nil, nil}
	}
	if (num > root.Val) {
		root.Right = insertNode(num, root.Right)
	} else {
		root.Left = insertNode(num, root.Left)
	}
	return root
}
```
* 非递归
```go
func bstFromPreorderNoRecursion(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	for i,num := range preorder {
		if i == 0 {
			continue
		}
		node := &TreeNode{num, nil, nil}
		tmp := root
		for tmp != nil {
			if num > tmp.Val {
				if tmp.Right == nil {
					tmp.Right = node
					tmp = nil
				} else {
					tmp = tmp.Right
				}
			} else {
				if tmp.Left == nil {
					tmp.Left = node
					tmp = nil
				} else {
					tmp = tmp.Left
				}
			}
		}
	}
	return root
}
```
### 4. 复杂度分析
* 时间复杂度: O(nlogn), n为数组中元素个数，遍历数组 n， 插入logn
* 空间复杂度: O(n), n为数组中元素个数

---

## Review [10 Git Commands You Should Know](https://towardsdatascience.com/10-git-commands-you-should-know-df54bea1595c)
你应该知道的10个 Git 命令

在这篇文章中我们将讨论你作为一个开发者、数据科学家或产品经理应该知道的各种各样的 Git 命令。我们将看看如何用 Git 进行检查、删除和整理。
同时我们也会涉及到退出 Vim 的方法、如何使用 bash 的别名来节省时间和 Git 的编辑器的配置。

如果你对 Git 的基本命令不熟悉，可以查看此篇文章 ["Learn Enough Git to be Useful"](https://towardsdatascience.com/learn-enough-git-to-be-useful-281561eef959).

以下是10个命令和其常用的参数。每个命令都链接到该命令的 Atlassian Bitbucket 指南。
#### 检查信息

让我们来查看变化。

* git diff  —— 查看所有本地文件的变化。可以附加文件名，以仅显示一个文件的变化。 如： git diff myfile
* git log —— 查看所有提交历史。也可以只看一个文件的，使用命令： git log -p myfile. 按 q 退出。
* git blame myfile —— 查看谁在什么时间改了 myfile 文件中的什么内容。
* git reflog —— 显示本地存储库 HEAD 的更改日志。对找到丢失的文件很有用。

用 git 检查信息并不是很混乱。相反 git 提供了过多的选项来删除、撤销提交和文件更改。

#### 撤销信息
git reset, git checkout, 和 git revert 用于撤销对存储库所做修改的影响。这些命令可能很难理解。

git reset 和 git checkout 可以用在提交和单个文件上。 git revert 只能用于提交级别。

如果你只是处理尚未提交到远程协作仓库中的本地提交，那你可以使用以上命令中的任何一个。

如果你要协作工作，并且需要撤销远程分钟中的一个提交，那你需要使用 git revert.

每一个命令都有大量的选项。以下是一些常见的用法：
* git reset --hard HEAD —— 丢弃最近一次提交以来的阶段性和非阶段性的修改。

用一个指定的提交来替代 HEAD 来丢弃自那个提交以来的修改。--hard 指定丢弃阶段性和非阶段性的修改。

确保你不会丢弃与你协作的人正在依赖的远程分支的提交。

* git checkout mycommit —— 丢弃自 mycommit 以来的非阶段性提交。

HEAD 通常用于 mycommit， 来丢弃自从最后一次提交之后对本地的修改。

checkout 是最好的用于撤销本地撤销的参数。它不会打乱协作者所依赖的远程分支的提交历史。

如果你用一个分支来代替提交来作为 checkout 的参数， HEAD 将切换到指定分支并更新工作目录用于匹配。这个是 checkout 的最常见的用法。

* git revert mycommit —— 撤销在 mycommit 中修改的所有影响。当撤销修改时，revert 将创建一个新的提交。

revert 对于协作项目来说是安全的，因为它不会覆盖其他用户分支可能依赖的历史。

有时你只是想删除本地目录中未被追踪的文件。例如，你运行了一段代码，这段代码创建了很多类型的文件，这些文件是你不想要在你的 repo 中出现的。那么你可以瞬间将其清理掉。

* git clean -n —— 删除本地工作目录中未被追踪的文件。

-n 用于没有删除任何东西的运行。

使用 -f 来实际删除文件。

使用 -d 删除未被追踪的目录。

默认情况下，.gitignore 未追踪的文件不会被删除，但是此行为可以被更改。

现在你已经知道了 Git 中的撤销操作，下面让我们来看看另外两个用来保持秩序的命令。

#### 整理信息
* git commit --amend —— 将你的阶段性修改添加到最近一次提交中。

如果没有阶段性的修改，这条命令将允许你编辑最近一次提交的信息。仅当本次提交还没有合并到远程 master 分支的时候，使用此命令。
* git push myremote --tags —— 将本地的所有 tags 发送到远程的 repo 中。 适用于版本更改。

如果你正在使用 Python 并且修改了一个你构建的包，[bump2version](https://pypi.org/project/bump2version/) 将会自动为你创建 tags。
一旦你推送了你的 tags，你就可以在 release 中使用它们了。

#### 救命，我被困在 Vim 中出不来了
使用 git ，你有时会发现自己陷入了 Vim 编辑会话。例如，你在没有提交消息的情况下进行提交 —— Vim 就会自动打开。如果你不知道什么是 Vim ，可以看[这里](https://stackoverflow.com/a/11828573/4590385)。

以下是退出 Vim 并保存文件的4个步骤：
1. 按 i 进入插入模式
2. 输入提交信息
3. 按 Esc 进入命令模式
4. 按 :x .别忘了冒号。保存文件并退出。

现在，你自由了。

#### 修改默认的编辑器
[这里](https://www.atlassian.com/git/tutorials/setting-up-a-repository/git-config)是一些常用编辑器的文档。以下是修改默认编辑器的命令：
* git config --global core.editor "atom --wait" 将默认编辑器改为 Atom。假设你已经安装了 Atom。

#### 为 Git 命令创建快捷方式
在你的 ~/.bash_profile （如果没有可以自己创建）文件中添加如下内容：
```bash
alias gs='git status '
alias ga='git add '
alias gaa='git add -A '
alias gb='git branch '
alias gc='git commit '
alias gcm='git commit -m '
alias go='git checkout '
```
然后使其生效： source ~/.bash_profile. 或者重新登录 shell.

#### 总结
在本文中，你已经看到了一些关键的 Git 命令，并配置了环境以节省时间。现在你已经对 Git 和 GitHub 有了基本的了解。准备好下一步了吗？
* 查看 [Bitbucket Git](https://www.atlassian.com/git/tutorials/learn-git-with-bitbucket-cloud) 教程，进行深入学习。
* 探索 Git 分支的[交互指南](https://learngitbranching.js.org/)。分支较难，但绝对值得一看。
* 去使用、去学习、去向其他人讲解。



---

## Tip

### Springboot 的 WebSocket 来匹配 Sockjs 客户端
Sockjs 客户端在建立 sockjs 连接之前，首先会发起一个 http 的 xxx/info 请求来进行探测。

sockjs 的请求
```javascript
var websocket = new WebSocket("ws://127.0.0.1:8080/server/123/test");
```

SpringBoot 中的 WebSocket 需要使用对应的 sockjs 服务器才能与之匹配。

配置 Config
```java
@Configuration
@EnableWebSocket
public class WebSocketConfig implements WebSocketConfigurer {
    @Override
    public void registerWebSocketHandlers(WebSocketHandlerRegistry registry) {
        registry.addHandler(testHandler(), "/server/{id}/test")
                .setAllowedOrigins("*")
                .withSockJS();
    }

    @Bean
    public WebSocketHandler testHandler() {
        return new TestHandler();
    }
} 
```
handler
```java
public class TestHandler implements WebSocketHandler {

    @Override
    public void afterConnectionEstablished(WebSocketSession webSocketSession) throws Exception {
        log.debug("connection established! " + webSocketSession.getRemoteAddress().toString() + " , id = " + webSocketSession.getId());
        LicenseManager.checkLicense(webSocketSession);
    }

    @Override
    public void handleMessage(WebSocketSession webSocketSession, WebSocketMessage<?> webSocketMessage) throws Exception {
        log.debug("Req : " + webSocketMessage.getPayload());
        TextMessage returnMessage = new TextMessage(webSocketMessage.getPayload() + " received at server");
        webSocketSession.sendMessage(returnMessage);

    }

    @Override
    public void handleTransportError(WebSocketSession webSocketSession, Throwable throwable) throws Exception {
        if (webSocketSession.isOpen()) {
            webSocketSession.close();
        }
        log.debug(throwable.toString());
        log.debug("WS connection error,close...");
    }

    @Override
    public void afterConnectionClosed(WebSocketSession webSocketSession, CloseStatus closeStatus) throws Exception {
        log.debug("Connection closed ... " + webSocketSession.getRemoteAddress().toString());

    }

    @Override
    public boolean supportsPartialMessages() {
        return false;
    }
}
```

---
    
## Share
### 12 为什么我的MySQL会“抖”一下？ —— 极客时间 MySQL实战45讲
#### 什么叫"脏页"
MySQL 做更新操作时，只是写了 redo log 到磁盘，只有这一个写磁盘的操作。这时内存中的数据页已经被更新，但是被更新的数据页并没有被写入磁盘中。

当内存数据页跟磁盘数据页内容不一致时，我们称这个内存页为"脏页"。内存数据写入到磁盘后，内存和磁盘上的数据页内容就一致了，称为"干净页"。

MySQL "抖"的时候，可能就是在刷脏页(flush)

以下情况会触发数据库的 flush 
1. InnoDB 的 redo log 写满了。这时候系统会停止所有更新操作，把 checkpoint 往前推进， redo log 留出的空间可以继续写。checkpoint 不是随便移动位置的。
比如 checkpoint 从 CP 移动到 CP' , 那么就需要将两者之间的所有脏页都 flush 到磁盘中。
2. 系统内存不足。当需要新的内存页，而内存不够用的时候，就需要淘汰一些数据页，空出内存给其他数据页使用。如果淘汰的是"脏页"，则需要先将其写入到磁盘。
难道不能直接把内存淘汰掉，下次需要请求的时候，从磁盘读入数据页，然后拿 redo log 出来应用不就行了？这里其实是从性能方面考虑的。
如果刷脏页一定会写磁盘，就保证了每个数据页有两种状态：
    * 一种是内存里存在，内存里就肯定是正确的结果，直接返回。
    * 另一种是内存里没有数据，就可以肯定数据文件上是正确的结果，读入内存后返回。
    
    这样的效率最高。
3. MySQL 认为系统"空闲"的时候。当然即使 MySQL 较忙时，MySQL 也会见缝插针地找时间，只要有机会就刷一点"脏页"。避免内存被快速消耗完。
4. MySQL 正常关闭的时候。会把内存中所有的"脏页"都 flush 到磁盘中。

下面可以分析一下以上4种情况对性能的影响。

3、4 可以不同考虑。3 是系统空闲的时候做的，4 是 MySQL 马上就关闭了。

重点关注1、2两种情况：
1. redo log 写满了，要 flush 脏页。这种情况是 InnoDB 要尽量避免的。因为出现这种情况时，整个系统不能再接受更新了，所有的更新都必须堵住。从监控上看，此时的更新数会跌为0.
2. 内存不够用，要先将脏页写到磁盘。这种情况是常态。InnoDB 用缓冲池（buffer pool）管理内存，缓冲池中的内存页有三种状态：
* 还没有使用
* 使用了并且是干净页
* 使用了并且是脏页

InnoDB 的策略是尽量使用内存，因此对于一个长时间运行的库来说，未被使用的页面很少。

而当要读入的数据页没有在内存中的时候，就需要在 buffer pool 中申请内存页。这时需要将最久不使用的数据页从内存中淘汰掉。如果是干净页，则直接淘汰；如果是脏页，则需要刷磁盘。

刷脏页是常态，但是出现以下两种情况都会明显影响性能：
1. 一个查询需要淘汰的脏页个数太多，会导致查询的响应时间明显变长。
2. 日志写满，更新全部堵住，写性能跌到0，这种情况对敏感业务来说，是不能接受的。

InnoDB 通过控制脏页比例的机制，来尽量避免上面两种情况。
#### InnoDB 刷脏页的控制策略
首先要正确的告诉 InnoDB 所在主机的 IO 能力， 这样 InnoDB 才能知道全力刷脏页的时候能刷多快。

建议将 innodb_io_capacity 这个参数设置为磁盘的 IOPS 

IOPS 可以通过以下命令来进行测试：
```jshelllanguage
 fio -filename=$filename -direct=1 -iodepth 1 -thread -rw=randrw -ioengine=psync -bs=16k -size=500M -numjobs=10 -runtime=10 -group_reporting -name=mytest 
```

InnoDB 刷盘速度要参考以下2个指标：
* 脏页比例
* redo log 写盘速度

InnoDB首先会单独算出这2个数字。

innodb_max_dirty_pages_pct 是脏页比例上限，默认是 75% 。 InnoDB会根据当前的脏页比例（M），算出一个0~100之间的数，伪代码如下：
```jshelllanguage
F1(M)
{
  if M>=innodb_max_dirty_pages_pct then
      return 100;
  return 100*M/innodb_max_dirty_pages_pct;
}
```
InnoDB 每次写入日志都有一个序号，当前写入序号跟 checkpoint 对应的序号之间的差值，设为N。InnoDB会根据 N 算出一个0~100 之间的数字。
计算公式记为 F2(N). N 越大，F2(N) 越大。

R = max(F1(M), F2(N))

* InnoDB 按照 innodb_io_capacity 定义的能力乘以 R% 来控制刷脏页的速度。

避免 MySQL "抖"，你需要：
* 合理地设置 innodb_io_capacity 的值
* 平时要多关注脏页比例，不要让它经常接近 75%

脏页比例是通过： innodb_buffer_pool_pages_dirty/innodb_buffer_pool_pages_total 得到的。具体命令如下：
```sql
mysql> select VARIABLE_VALUE into @a from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_dirty';
select VARIABLE_VALUE into @b from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_total';
select @a/@b;
```
#### 一个有趣的策略
在刷脏页的时候，MySQL 会检查这个脏页的旁边的页是否也是脏页，如果是脏页，则会带着一起刷磁盘。然后以此类推，直到不是脏页为止。

* innodb_flush_neighbors 参数用来控制这个行为，1 表示会连坐， 0 表示不找邻居，只刷自己。

在机械磁盘的情况下，连坐机制会减少很多随机 IO。 由于机械磁盘的 IOPS 一般只有几百，减少随机 IO 意味着会大幅提高系统性能。

而使用 SSD 时，IOPS 不是瓶颈。建议将 innodb_flush_neighbors 设置为 0。

MySQL 8.0 中， innodb_flush_neighbors 参数默认是0.