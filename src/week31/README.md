# week31

---

# Algorithm [63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii/)
## 1. 问题描述
唯一路径

给定一个 m * n 的棋盘

一个机器人初始在左上角，只能向右或向下移动，其目标是移动到右下角。

路径上存在障碍物的情况下，共有多少可能的唯一路径？

障碍物和空格用 1 和 0 来表示

#### 示例 1：
* 输入 :
```code
    [
        [0,0,0],
        [0,1,0],
        [0,0,0]
    ]
```
* 输出 : 2
* 解释 : 
    1. Right -> Right -> Down -> Down
    2. Down -> Down -> Right -> Right

## 2. 解题思路
dp[i,j] 表示从 (0,0) 走到 (i,j) 的路径数
```java
if dp[0,0] =1 return 0
dp[0,0]=1

if obstacleGrid[k,0]=1 then 
    dp[i,0] = 1 (i < k)
    dp[j,0] = 0 (j >= k)
if obstacleGrad[0,k]=1 then
    dp[0,i] = 1 (i < k)
    dp[0,j] = 0 (j >=k)
if obstacleGrad[i-1,j] = 1 && obstacleGrad[i,j-1] = 1 then
    dp[i,j] = 0
else if obstacleGrad[i-1,j] = 1 then
    dp[i,j] = dp[i,j-1]
else if obstacleGrad[i,j-1] = 1 then
    dp[i,j] = dp[i-1,j]
else     
    dp[i,j] = dp[i-1,j] + dp[i,j-1]
```
最终求出 dp[m-1,n-1]

## 3. 代码
```go
if obstacleGrid[0][0] == 1 {
        return 0
    }
    m,n := len(obstacleGrid), len(obstacleGrid[0])
    var dp [][]int
    dp = make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    dp[0][0] = 1
    for i := 1; i < n; i++ {
        if obstacleGrid[0][i] == 1 {
            break
        }
        dp[0][i] = 1
    }
    for i := 1; i < m; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
                dp[i][j] = 0
            } else if obstacleGrid[i-1][j] == 1 {
                dp[i][j] = dp[i][j-1]
            } else if obstacleGrid[i][j-1] == 1 {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i][j-1] + dp[i-1][j]
            }
        }
    }
    return dp[m-1][n-1]
```
## 4. 复杂度分析
* 时间复杂度 : O(m*n)
* 空间复杂度 : O(m*n)

---

# Review [How to (quickly) Build a Tensorflow Training Pipeline](https://towardsdatascience.com/how-to-quickly-build-a-tensorflow-training-pipeline-15e9ae4d78a0)
## Tensorflow 模型使用数据的 3 种方式：如何正确的使用？
1. 使用 feed_dict 命令。其缺点是性能差，只能使用 python 加载数据到内存，不能使用多线程。但是在某些情况下是很好的解决方案。
2. 使用 TfRecords。使用 TfRecords, 10 次有 9 次都是糟糕的选择。
3. 使用 tensorflow 的 tf.data.Dataset 对象。

## 基本的人脸识别
输入是两张图片，输出是 1 或者 0，其中 1 表示他们是同一个人，0 表示不是同一个人。
model.py
```python
import tensorflow as tf
from tensorflow import Tensor


class Inputs(object):
    def __init__(self, img1: Tensor, img2: Tensor, label: Tensor):
        self.img1 = img1
        self.img2 = img2
        self.label = label

class Model(object):
    def __init__(self, inputs: Inputs):
        self.inputs = inputs
        self.predictions = self.predict(inputs)
        self.loss = self.calculate_loss(inputs, self.predictions)
        self.opt_step = tf.train.AdamOptimizer(learning_rate=0.001).minimize(self.loss)

    def predict(self, inputs: Inputs):
        with tf.name_scope("image_substraction"):
            img_diff = (inputs.img1 - inputs.img2)
            x = img_diff
        with tf.name_scope('conv_relu_maxpool'):
            for conv_layer_i in range(5):
                x = tf.layers.conv2d(x,
                                     filters=20 * (conv_layer_i + 1),
                                     kernel_size=3,
                                     activation=tf.nn.relu)
                x = tf.layers.max_pooling2d(x,
                                            pool_size=3,
                                            strides=2)
        with tf.name_scope('fully_connected'):
            for conv_layer_i in range(1):
                x = tf.layers.dense(x,
                                    units=200,
                                    activation=tf.nn.relu)
        with tf.name_scope('linear_predict'):
            predicted_logits = tf.layers.dense(x, 1, activation=None)

        return tf.squeeze(predicted_logits)

    def calculate_loss(self, inputs: Inputs, prediction_logits: Tensor):
        with tf.name_scope('calculate_loss'):
            return tf.reduce_mean(tf.nn.sigmoid_cross_entropy_with_logits(labels=inputs.label,
                                                                          logits=prediction_logits))
```

## 数据
数据集: [人脸数据库](http://vis-www.cs.umass.edu/lfw/lfw.tgz)
```xpath
/lfw
/lfw/Dalai_Lama/Dalai_Lama_0001.jpg
/lfw/Dalai_Lama/Dalai_Lama_0002.jpg
...
/lfw/George_HW_Bush/George_HW_Bush_0001.jpg
/lfw/George_HW_Bush/George_HW_Bush_0002.jpg
...
```
生成同一个人的图片对
pair_generator.py
```python
import os
import glob
import random


class PairGenerator(object):
    def __init__(self, lfw_path='resources' + os.path.sep + 'lfw'):
        self.all_people = self.generate_all_people_dict(lfw_path)

    def generate_all_people_dict(self, lfw_path):
        # generates a dictionary between a person and all the photos of that person
        all_people = {}
        for person_folder in os.listdir(lfw_path):
            person_photos = glob.glob(lfw_path + os.path.sep + person_folder + os.path.sep + '*.jpg')
            all_people[person_folder] = person_photos
        return all_people

    def get_next_pair(self):

        while True:
            # draw a person at random
            person1 = random.choice(self.all_people)
            # flip a coin to decide whether we fetch a photo of the same person vs different person

            same_person = random.random() > 0.5
            if same_person:
                person2 = person1
            else:
                person2 = random.choice(self.all_people)

            person1_photo = random.choice(self.all_people[person1])
            person2_photo = random.choice(self.all_people[person2])
            yield ({'person1': person1_photo, 
                    'person2': person2_photo, 
                    'label': same_person})

```

## Tensorflow 数据管道
建立数据管道
tf_dataset.py 
```python
import tensorflow as tf
from .pair_generator import PairGenerator
from .model import Inputs


class Dataset(object):
    img1_resized = 'img1_resized'
    img2_resized = 'img2_resized'
    label = 'same_person'

    def __init__(self, generator=PairGenerator()):
        self.next_element = self.build_iterator(generator)

    def build_iterator(self, pair_gen: PairGenerator):
        batch_size = 10
        prefetch_batch_buffer = 5

        dataset = tf.data.Dataset.from_generator(pair_gen.get_next_pair,
                                                 output_types={PairGenerator.person1: tf.string,
                                                               PairGenerator.person2: tf.string,
                                                               PairGenerator.label: tf.bool})
        dataset = dataset.map(self._read_image_and_resize)
        dataset = dataset.batch(batch_size)
        dataset = dataset.prefetch(prefetch_batch_buffer)
        iter = dataset.make_one_shot_iterator()
        element = iter.get_next()

        return Inputs(element[self.img1_resized],
                      element[self.img2_resized],
                      element[PairGenerator.label])

    def _read_image_and_resize(self, pair_element):
        target_size = [128, 128]
        # read images from disk
        img1_file = tf.read_file(pair_element[PairGenerator.person1])
        img2_file = tf.read_file(pair_element[PairGenerator.person2])
        img1 = tf.image.decode_image(img1_file)
        img2 = tf.image.decode_image(img2_file)

        # let tensorflow know that the loaded images have unknown dimensions, and 3 color channels (rgb)
        img1.set_shape([None, None, 3])
        img2.set_shape([None, None, 3])

        # resize to model input size
        img1_resized = tf.image.resize_images(img1, target_size)
        img2_resized = tf.image.resize_images(img2, target_size)

        pair_element[self.img1_resized] = img1_resized
        pair_element[self.img2_resized] = img2_resized
        pair_element[self.label] = tf.cast(pair_element[PairGenerator.label], tf.float32)

        return pair_element
```
1. tf.Data.Dataset.from_generator() 让tensorflow知道它是由python生成器提供的
2. map 操作：在这里，我们设置了从生成器输入(文件名)到我们实际需要的模型(加载和调整大小的图像)所需的所有任务。主要用 _read_image_and_resize()
3. batch 操作：是一个方便的函数，它将图像批处理成具有一致数量元素的包。
4. prefetch 操作：让Tensorflow做与设置队列相关的记录工作，以便数据管道继续读取和入列数据，直到它有N个批量的数据都已加载并准备就绪

train.py
```python
from recognizer.pair_generator import PairGenerator
from recognizer.tf_dataset import Dataset
from recognizer.model import Model
import tensorflow as tf
import pylab as plt
import numpy as np


def main():
    generator = PairGenerator()
    # print 2 outputs from our generator just to see that it works:
    iter = generator.get_next_pair()
    for i in range(2):
        print(next(iter))
    ds = Dataset(generator)
    model_input = ds.next_element
    model = Model(model_input)

    # train for 100 steps
    with tf.Session() as sess:
        # sanity test: plot out the first resized images and their label:
        (img1, img2, label) = sess.run([model_input.img1, 
                                        model_input.img2, 
                                        model_input.label])

        # img1 and img2 and label are BATCHES of images and labels. plot out the first one
        plt.subplot(2, 1, 1)
        plt.imshow(img1[0].astype(np.uint8))
        plt.subplot(2, 1, 2)
        plt.imshow(img2[0].astype(np.uint8))
        plt.title(f'label {label[0]}')
        plt.show()

        # intialize the model
        sess.run(tf.global_variables_initializer())
        # run 100 optimization steps
        for step in range(100):
            (_, current_loss) = sess.run([model.opt_step, 
                                          model.loss])
            print(f"step {step} log loss {current_loss}")


if __name__ == '__main__':
    main()
```
[代码在 github 上](https://github.com/urimerhav/tflow-dataset)

## 示例：删除图片背景
1. 搜索只有单一物品的图片
2. 搜索一些通用的背景模板，地毯或者墙壁
3. 建立一个生成器，生成成对的背景文件路径+对象文件路径，将它们融合在一起。
4. 让模型猜测哪些像素属于背景，哪些像素属于物品

生成器
```python
  def gen_item(self):
      while True:
            bkg_image = random.sample(self.all_bkg_images, 1)[0]
            png_name = self.get_complete_png(cleaned_image=False)

          yield {'png_path': str(png_name),
                 'bkg_image': bkg_image,
```

组合背景图片和物品图片
```python
def tf_blend_images(bkg_img, obj_img, object_pixels):
    object_shape = tf.shape(obj_img)
    cropped_bkg = tf.random_crop(bkg_img, object_shape)

    composed_image = (cropped_bkg * tf.cast(tf.logical_not(object_pixels), tf.uint8) +
                      obj_img * tf.cast(object_pixels, tf.uint8))

    # random saturation etc

    return composed_image
```

[模型](http://proproductpix.org/)

---

# Tip Spring MVC 中使用 Thymeleaf
1. 建立 Spring MVC 工程
2. 引入 Thymeleaf 依赖
```xml
        <dependency>
            <groupId>org.thymeleaf</groupId>
            <artifactId>thymeleaf</artifactId>
            <version>2.1.4.RELEASE</version>
        </dependency>
        <dependency>
            <groupId>org.thymeleaf</groupId>
            <artifactId>thymeleaf-spring4</artifactId>
            <version>2.1.4.RELEASE</version>
        </dependency>
        <dependency>
            <groupId>net.sourceforge.nekohtml</groupId>
            <artifactId>nekohtml</artifactId>
            <version>1.9.22</version>
        </dependency>
```
3. 配置 applicationContent.xml
```xml
...
<!-- Thymeleaf视图模板引擎配置 -->
    <bean id="templateResolver" class="org.thymeleaf.templateresolver.ServletContextTemplateResolver">
        <property name="prefix" value="/resources/static/" />
        <property name="suffix" value=".html" />
        <property name="cacheable" value="false" />
        <property name="characterEncoding" value="UTF-8" />
        <property name="templateMode" value="LEGACYHTML5" />
    </bean>

    <bean id="templateEngine" class="org.thymeleaf.spring4.SpringTemplateEngine">
        <property name="templateResolver" ref="templateResolver" />
    </bean>

    <!-- 视图解析器配置 -->
    <bean class="org.thymeleaf.spring4.view.ThymeleafViewResolver">
        <property name="templateEngine" ref="templateEngine" />
        <property name="order" value="1" />
        <!-- <property name="viewNames" value="*.html,*.xhtml" /> -->
        <property name="characterEncoding" value="UTF-8"/>
    </bean>
...
```
4. controller 层
```java
@Controller
public class MyController {
    @RequestMapping(value = "/index")
    public String index(Model model) {
        model.addAttribute("p1","v1");
        model.addAttribute("p2","v2");
        model.addAttribute("p3", "v3");
        return "index";
    }
}
```

5. view 层
```html
<head>
<body>
<script th:inline="javascript">
    /*<![CDATA[*/
    var p1 = [[${p1}]];
    var p2 = [[${p2}]];
    var p3 = [[${p3}]];
    /*]]>*/
</script>
</body>
</head>
```

---
    
# Share 33 我查这么多数据，会不会把数据库内存打爆？ —— 极客时间 MySQL实战45讲
## 全表扫描对 server 层的影响
我们对一个 200G 数据的 InnoDB 表 db1.t 做全表扫描
```roomsql
mysql -h$host -P$port -u$user -p$pwd -e "select * from db1.t" > $target_file
```
执行流程：
1. 获取一行写到 net_buffer 中。net_buffer 的大小由参数 net_buffer_length 定义，默认是 16K
2. 重复获取行数据，直到 net_buffer 写满，然后调用网络接口发送
3. 如果发送成功，则清空 net_buffer，重复步骤1、2
4. 如果发送函数返回 EAGAIN 或 WSAEWOULDBLOCK，表示本地网络栈（socket send buffer） 写满了，进入等待。直到网络栈重新可写，再继续发送

流程图:
![query_send](query_send.jpg)
1. 一个查询在发送过程中，占用的 MySQL 内部的内存最大就是 net_buffer_length
2. socked send buffer 默认定义 /proc/sys/net/core/wmem_default), 如果 socket send buffer 被写满，就会暂停读数据的流程

MySQL 是边读边发的，如果客户端慢，就会导致服务端结果发不出去，会拖长这个事务的执行时间。

如果故意让客户端不读 socket receive buffer
![send_block](send_block.png)
State的值"Sending to client" 表示服务端的网络栈写满了

对于正常的线上业务来说，如果一个查询的返回结果不会很多的，都建议使用 mysql_store_result 这个接口，直接把查询结果保存到本地内存

还有一个 “Sending data” 的状态
* MySQL 查询语句进入执行阶段后，首先把状态设置成 "Sending data"
* 然后，发送执行结果的列相关的信息（meta data）给客户端
* 再继续执行语句的流程
* 执行完成后，把状态设置成空字符串

"Sending data" 并不一定是指"正在发送数据"，而可能处于执行器过程中的任意阶段。

| session A | session B |
|---|---|
| begin;<br/> select * from t where id=1 for update; | |
| | select * from t lock in share mode;<br/> (blocked) |

![sending_data](sending_data.png)
session B 是在等锁，但是状态显示为 "Sending data"

* 仅当一个线程处于 "等待客户端接收结果" 的状态，才会显示 "Sending to client"
* 显示成 "Sending data" 的意思是 "正在执行"

## 全表扫描对 InnoDB 的影响
内存的数据页是在 Buffer Pool（BP）中管理的，在 WAL 里 Buffer Pool 起到了加速更新的作用。
实际上，Buffer Pool 还有一个更重要的作用，就是加速查询。

Buffer Pool 对查询的加速效果，依赖于一个重要的指标：内存命中率

执行 show engine innodb status 查看系统当前的 BP 命中率。一个稳定的线上服务，要求命中率在 99% 以上。
可以看到 "Buffer pool hit rate" 行
![hit_rate](hit_rate.png)

InnoDB Buffer Pool 的大小参数 innodb_buffer_pool_size 确定，一般建议设置成物理内存的 60%~80%

InnoDB 的内存管理用的是 LRU 算法，用链表实现
![LRU](LRU.jpg)
1. 状态 1 中，链表头是 P1，表示 P1 是最近刚刚被访问过的数据页，假设内存里只能放下这么多数据页
2. 这时候如果有一个请求访问 P3，因此变成状态 2，P3 被移到最前面
3. 状态 3 表示，这次访问的数据页是不存在于链表中的，所以需要在 Buffer Pool 中新申请一个数据页 Px，加到链表头部。
但是由于内存已满了，不能申请新的内存。于是，会清空链表末尾 Pm 这个数据页的内存，存入 Px 的内容，然后放到链表头
4. 从效果上看，就是最久没被访问的数据页 Pm，被淘汰了

按照这个算法，扫描一个大表的时候，当前 BP 中的页都会被淘汰，然后存入新的数据。
BP 的命中率急剧下降，磁盘压力增加，SQL语句响应变慢。

实际上，InnoDB 对 LRU 算法做了改进
![improve_LRU](improve_LRU.png)

InnoDB 中，按照 5:3 的比例把整个 LRU 链表分成了 young 区域和 old 区域。图中 LRU_old 指向的就是 old 区域的一个位置，
是整个链表的 5/8 处。也就是说，靠近链表头部的 5/8 是 young 区域，靠近链表尾部的 3/8 是 old 区域。

改进后的 LRU 算法
1. 状态 1 ，要访问数据页 P3, 由于 P3 在 young 区域，因此和优化前的 LRU 算法一样；将其移到链表头部，变成状态 2
2. 之后要访问一个新的不存在于当前链表的数据页，这时候依然是淘汰掉数据页 Pm，但是新插入的数据页 Px，是放在 LRU_old 处
3. 处于 old 区域的数据页，每次被访问到的时候都要做下面这个判断：
    * 若这个数据页在 LRU 链表中存在的时间超过了 1 秒，就把它移动到链表头部；
    * 如果这个数据页在 LRU 链表中存在的时间短于 1 秒，位置保持不变。1 秒这个时间，是由参数
    innodb_old_blocks_time 控制的。其默认值是 1000 毫秒

改进后的 LRU 算法对大表扫描的处理：
1. 扫描过程中，需要新插入的数据页，都被放到 old 区域
2. 一个数据页里面有多条记录，这个数据页会被多次访问到，但由于是顺序扫描，这个数据页第一次被访问和最后一次被访问的时间间隔不会超过 1 秒，
因此还是会被保留在 old 区域
3. 再继续扫描后续的数据，之前的这个数据页之后也不会再被访问到，于是始终没有机会移到链表头部（也就是 young 区域），很快就会被淘汰出去

## 思考题：请给出由于客户端性能问题，导致对数据库影响很严重的例子
* 答：问题的核心是造成了“长事务”
    * 如果前面的语句有更新，意味着它们在占用着行锁，会导致别的语句更新被锁住
    * 当然读的事务也有问题，就是会导致 undo log 不能被回收，导致回滚段空间膨胀
