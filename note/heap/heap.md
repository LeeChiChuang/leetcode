# 堆
 1. 完全二叉树(从上到下，从左到右依次被填满)
 2. 每个节点的值小于等于孩子节点的值

# 堆的特点
可以在 O(logN)O(logN) 的时间复杂度内向 堆 中插入元素；

可以在 O(logN)O(logN) 的时间复杂度内向 堆 中删除元素；

可以在 O(1)O(1) 的时间复杂度内获取 堆 中的最大值或最小值。

# 分类
堆 有两种类型：最大堆 和 最小堆。

最大堆：堆中每一个节点的值 都大于等于 其孩子节点的值。所以最大堆的特性是 堆顶元素（根节点）是堆中的最大值。

最小堆：堆中每一个节点的值 都小于等于 其孩子节点的值。所以最小堆的特性是 堆顶元素（根节点）是堆中的最小值。

![image-20210515153310020](/Users/simon/Library/Application Support/typora-user-images/image-20210515153310020.png)