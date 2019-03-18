# Leetcode [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
* 求字符串的最长子串，要求子串中没有重复的字符（不考虑中文）

## 算法1——暴力法
1. 从start(初始为0)字符开始遍历整个字符串的字符——O(n)
2. 从end（初始值end=star+1）个字符开始遍历——O(n<sup>2</sup>)
3. 获取substr = str.substring(start,end) 
4. 遍历substr 查看第j个字符是否在substr中，如果以存在则跳出2的循环，
否则计算最大值len=max(len, substr.length + 1) ——O(n<sup>3</sup>)

* 算法时间复杂度： O(n<sup>3</sup>)
* 算法空间复杂度： O(n) 存储substr需要最大可能需要n个字符

## 算法2 —— 窗口切片法
* 算法1时间复杂度太高，性能太低
* 用2个指针（start，end）指向字符所在的位置，可以考虑记录每个字符的出现的最后位置，
如果发现之前已经出现的字符，则可以将start指针调整到此位置+1即可。
不需要每次将start+1的方式遍历

1. 用一个map（HashMap）存储 字符-字符位置信息，如果是多线程环境则需要使用ConcurrentHashMap
2. start和end都指向str的第一个字符,即start,end=0
3. 从头到尾遍历str
4. 如果map中不存在，则将此字符put到map中，同时end++，并计算len=max(len,end - start + 1)
5. 如果map中存在，则start=max(start,map.get(此字符)) ，然后将此字符put到map中

* 时间复杂度 O(n) ,只遍历了一次str
* 空间复杂度 O(n)
## 算法3 ——优化算法2
* 此题目不需要考虑中文，str的每个字符都是ascii码，ascii码总共包含256个字符，
其中可见字符128个，所以HashMap可以用一个int数组代替
* 时间复杂度 O(n),但是int数组比HashMap更轻量级，HashMap如果有Hash冲突，则查找时间会比int数组更长
* 空间复杂度 O(1),因为是一个常数128 （int[128]) 
* 算法2的空间复杂度实际也可以认为是常数，因为字符是有范围的

