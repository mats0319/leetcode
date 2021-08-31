# leetcode

[![GoDoc](https://godoc.org/github.com/mats9693/leetcode?status.svg)](https://godoc.org/github.com/mats9693/leetcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/mats9693/leetcode)](https://goreportcard.com/report/github.com/mats9693/leetcode)
[![Build Status](https://travis-ci.org/mats9693/leetcode.svg?branch=master)](https://travis-ci.org/mats9693/leetcode)
[![codecov](https://codecov.io/gh/mats9693/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/mats9693/leetcode)

leetcode题库学习。相关知识点总结，会更新在我的[document repo](https://github.com/mats9693/document)。

## 目的

学习leetcode题库是为了熟悉go语言、积累使用经验，以及编程技巧和简单算法的学习。

## 特殊题目

记录一些比较特殊的题目，对自己有启发的。

### LRU & LFU

#### 概念对比

> used：使用，读或写缓存均为对缓存的“使用”

LRU：least recently used，最近最少使用；缓存淘汰策略只从```上一次使用时间```一个维度进行考量

LFU：least frequently used，最不经常使用；缓存淘汰策略需要从```使用频率```和```上一次使用时间```两个维度考量

#### solution 146：LRU缓存机制

##### 题目规则

1. 缓存容量```capacity```为正整数，缓存的```key```、```value```均为int类型
1. 读缓存```func get(key int) int```：
    1. ```key```已存在，返回对应```value```
    1. ```key```不存在，返回```-1```
1. 写缓存```func put(key int, value int)```：
    1. ```key```已存在，修改对应```value```
    1. ```key```不存在，写入该组缓存，若写入前缓存容量已达上限，则应淘汰最久未使用的缓存（读、写缓存均算作使用）

##### 分析过程

> 目标：```O(1)```复杂度

1. 使用双向链表维护缓存的**上一次使用时间**：
    1. 预期效果：链表正方向（从头部到尾部）节点按照使用时间排序——越早使用（即久未使用）的节点，越靠近链表尾部
    1. 实现方法：每使用一次缓存，就将该缓存对应的链表节点移动到链表头部
    1. 解决问题：缓存淘汰时，只需要删除尾部节点即可
1. 如果只使用双向链表，每次判断```key```是否存在时，都要遍历链表，考虑增加一个map，记录key到链表节点的映射关系
    1. 解决问题：```O(1)```时间判断缓存是否存在

##### 数据结构

1. cache：```map[int]*listNode```，```key```到```节点```的映射
    1. listNode data: ```key```, ```value```
1. list：```*list```，双向链表，维护缓存的**上一次使用时间**
1. capacity：```int```，链表容量

##### 功能逻辑与伪代码

> 功能逻辑将从整体到局部，分析可能遇到的每种情况及其处理方式，不同分支情况可能包含相同的处理逻辑；
> 伪代码则从代码编写角度，提炼可能存在的相同处理逻辑，重新组合代码顺序

###### 读缓存```func get(key int) int```

1. 查询```key```是否存在：
    1. 存在：
        1. 在原链表中删除该缓存节点，重新插入到链表头部
        1. 返回对应的```value```
    1. 不存在：返回```-1```

```go 
func get(key int) int {
   node, ok := cache.isExist(key)
   if !ok { // key不存在
      return -1
   }
   
   // key存在
   list.remove(node)
   list.addToHead(node)
   
   return node.value
}
```

###### 写缓存```func put(key int, value int)```

1. 查询```key```是否存在：
    1. 存在：
        1. 在原链表中删除该缓存节点，重新插入到链表头部
        1. 更新缓存节点的```value```
    1. 不存在：判断当前链表容量是否已达上限：
        1. 容量已达上限：
            1. 在链表中删除尾部节点（记录该节点的```key```）
            1. 根据上一步中记录的```key```，删除对应的映射关系
            1. 根据输入参数构造新的节点：
                1. 将新的节点插入链表头部
                1. 新增```key```到```新的节点```的映射关系
        1. 容量未达上限：
            1. 根据输入参数构造新的节点：
                1. 将新的节点插入链表头部
                1. 新增```key```到```新的节点```的映射关系

```go 
func put(key int, value int) {
   node, ok := cache.isExist(key)
   if ok { // key已存在
      list.remove(node)
      node.value = value
      list.addToHead(node)
   
      return 
   }
   
   // key不存在
   if cache.isFull() { // 缓存容量已达上限
      list.removeTail()
      cache.delete(tailNode.Key)
   }
   
   newNode := &listNode {
      key: key
      value: value
   }
   
   cache[key] = newNode
   list.addToHead(newNode)
}
```

#### solution 460：LFU缓存机制

##### 题目规则

1. 缓存容量```capacity```、缓存的```key```和```value```均为自然数
1. 读缓存```func get(key int) int```：（与lru相同）
    1. ```key```已存在，返回对应```value```
    1. ```key```不存在，返回```-1```
1. 写缓存```func put(key int, value int)```：
    1. ```key```已存在，修改对应```value```
    1. ```key```不存在，写入该组缓存，若写入前缓存容量已达上限，则应淘汰使用次数最少的缓存（记其使用次数为```n```）；  
       若使用次数为```n```的缓存数大于一个，则淘汰最久未使用的缓存（即，此时遵守lru规则）

##### 分析过程

> 目标：```O(1)```复杂度

1. 从lru的解决方案入手，使用双向数组主要维护缓存的使用次数，再使用map维护```key```到节点的映射关系
    1. 存在问题：缓存节点移动时（主要指插入环节），如何快速定位节点应该在的位置？
        1. 场景分析：
            1. 删除节点：使用```key```，通过映射关系直接定位到节点
            1. 插入新节点：新缓存使用次数为```1```，应插入到所有使用次数为```1```的缓存节点的前面；  
               若不存在使用次数为```1```的缓存，则应插入到链表末尾
            1. 插入移除的节点：缓存使用次数```n+1```，应插入到所有使用次数为```n+1```的缓存节点的前面；  
               若不存在使用次数为```n+1```的缓存，则应插入到所有使用次数为```n```的缓存节点的前面；
        1. 分析结论：通过新增一个map，维护**缓存使用次数(记为```n```)** 到**所有使用次数为n的缓存中，最近使用的缓存节点**的映射关系,  
           可以快速定位节点应该在的位置。

##### 数据结构

1. recent：```map[int]*listNode```，```frequency```到**使用次数为```frequency```的节点中，最近使用的一个**的映射
2. count：```map[int]int```，```frequency```到**对应频率的节点数量**的映射
3. cache：```map[int]*listNode```，```key```到**节点**的映射
    1. listNode data: ```key```, ```value```, ```frequency```
4. list：```*list```，双向链表，维护缓存的**使用次数（优先）**和**上一次使用时间**
5. capacity：```int```，链表容量

##### 功能逻辑与伪代码

###### 读缓存```func get(key int) int```

1. 查询```key```是否存在：
    1. 存在：（记节点```frequency```为```n```）
        1. 若存在其他```frequency = n+1```的节点，则将节点移动到所有```frequency = n+1```的节点的前面；  
           否则，若存在其他```frequency = n```的节点，且当前节点不是最近节点，则将节点移动到所有```frequency = n```的节点的前面；    
           否则，不移动节点（该情况下，节点就应该呆在它现在的位置）
        2. 更新```recent```
        3. 更新```count```
        4. 将节点```frequency``` ```+1```
        5. 返回节点的```value```
    1. 不存在：返回```-1```

```go 
func get(key int) int {
    node, ok := cache.isExist(key)
    if !ok { // key不存在
        return -1
    }
    
    // key已存在
    if count[node.frequency+1] > 0 {
        // 存在其他使用次数为n+1的缓存，将指定缓存移动到所有使用次数为n+1的节点之前
        list.removeNode(node)
        list.addBefore(recent[node.frequency+1], node)
    } else if count[node.frequency] > 1 && recent[node.frequency] != node {
        // 不存在其他使用次数为n+1的缓存，但存在其他使用次数为n的缓存，且当前节点不是最近的节点
        // 将指定缓存移动到所有使用次数为n的节点之前
        list.removeNode(node)
        list.addBefore(recent[node.frequency], node)
    }     
    
    // 更新recent
    recent[node.frequency+1] = node
    if count[node.frequency] <= 1 { // 不存在其他freq = n的节点，recent置空
        recent[node.frequency] = nil
    } else if recent[node.frequency] == node { // 存在其他freq = n的节点，且recent = node，将recent向后移动一位
        recent[node.frequency] = node.next
    }
    
    // 更新使用次数对应的节点数
    count[node.frequency+1]++
    count[node.frequency]--
    
    // 更新缓存使用次数
    node.frequency++
    
    return node.value
}
```

###### 写缓存```func put(key int, value int)```

1. 查询key是否存在：
    1. 存在：参考读缓存-key存在，额外修改对应的value即可
    1. 不存在：
        1. 若当前缓存容量已达上限：
            1. 淘汰尾部的缓存节点
            2. 若不存在其他```freq = n```的节点，则将```recent```置空
            3. 更新```cache```
            4. 更新```count```
        2. 构造新节点：key，value，frequency = 1
            1. 是否存在其他```frequency = 1```的节点：
               1. 存在：插入到它们的前面
               2. 不存在：插入链表尾部
            2. 更新```recent```、```cache```、```count```

```go 
func put(key int, value int) {
    node, ok := cache.isExist(key)
    if ok { // key已存在
        get(key)
        node.value = value
        
        return 
    }
    
    // key不存在
    if cache.isFull() { // 缓存已满，删除最后一个节点，相应更新cache、count、recent（条件）
        list.removeTail()
        
        if count[tailNode.frequency] <= 1 {
            recent[tailNode.frequency] = nil
        }
        count[tailNode.frequency]--
        cache.delete(tailNode.key)
    }
    
    newNode := &listNode{
        key: key,
        value: value,
        frequency: 1,
    }
    
    // 插入新的缓存节点
    if count[1] > 0 {
        list.addBefore(recent[1], newNode)
    } else {
        list.addToTail(newNode)
    }
    
    // 更新recent、count、cache
    recent[1] = newNode
    count[1]++
    cache[key] = newNode
}
```

---

###### Mario

I've been pretending to work hard, but you're really growing up.
