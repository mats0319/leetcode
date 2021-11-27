# leetcode

[![GoDoc](https://godoc.org/github.com/mats9693/leetcode?status.svg)](https://pkg.go.dev/github.com/mats9693/leetcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/mats9693/leetcode)](https://goreportcard.com/report/github.com/mats9693/leetcode)
[![Build Status](https://travis-ci.org/mats9693/leetcode.svg?branch=master)](https://travis-ci.org/mats9693/leetcode)
[![codecov](https://codecov.io/gh/mats9693/leetcode/branch/master/graph/badge.svg?token=YihIlqewRk)](https://codecov.io/gh/mats9693/leetcode)

## 学习leetcode题库的目的

1. 熟悉go语言，积累使用经验
2. 学习编程技巧
3. 学习算法

## 题目

### LRU & LFU

#### 概念对比

> used：使用，读或写缓存均视为对缓存的**使用**

LRU：least recently used，最近最少使用；缓存淘汰策略只与```上一次使用时间```有关

LFU：least frequently used，最不经常使用；缓存淘汰策略与```使用频率```和```上一次使用时间```有关

#### solution 146：LRU缓存机制

##### 题目规则

1. 缓存容量```capacity```为正整数，缓存的```key```、```value```均为```int```类型
1. 读缓存```func get(key int) int```：
    1. ```key```已存在，返回对应```value```
    1. ```key```不存在，返回```-1```
1. 写缓存```func put(key int, value int)```：
    1. ```key```已存在，修改对应```value```
    1. ```key```不存在，写入该组缓存，若写入前缓存容量已达上限，则应淘汰最久未使用的缓存（强调：读、写缓存均视为**使用**）

##### 分析过程

> 目标：```O(1)```复杂度

1. 使用双向链表维护缓存的**上一次使用时间**：
    1. 约定：链表正方向（从头部到尾部）节点按照使用时间排序——越早使用（即久未使用）的节点，越靠近链表尾部
    2. 维护：每使用一次缓存，就将该缓存对应的链表节点移动到链表头部；缓存淘汰时，只需要删除尾部节点即可

Q1：如果只使用双向链表，每次判断```key```是否存在时，都要遍历链表  
A1：增加一个map，记录```key```到**链表节点**的映射关系

##### 数据结构

1. cache：```map[int]*listNode```，```key```到**节点**的映射
    1. listNode data：```key```, ```value```
1. list：```*listNode```，双向链表，维护缓存的**上一次使用时间**
1. capacity：```int```，链表容量

##### 功能逻辑与伪代码

> 功能逻辑将从整体到局部，分析可能遇到的每种情况及其处理方式，不同分支情况可能包含相同的处理逻辑  
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
    1. 存在：（此场景与**读缓存——```key```存在**类似，可通过调用读缓存方法并更新```value```实现）
        1. 在原链表中删除该缓存节点，重新插入到链表头部
        1. 更新缓存节点的```value```
    1. 不存在：判断当前链表容量是否已达上限：
        1. 容量已达上限：
            1. 在链表中删除尾部节点（记录该节点的```key```）
            1. 根据上一步中记录的```key```，删除对应的映射关系
            1. 根据输入参数构造新的节点：
                1. 将新的节点插入链表头部
                1. 新增```key```到**新的节点**的映射关系
        1. 容量未达上限：
            1. 根据输入参数构造新的节点：
                1. 将新的节点插入链表头部
                1. 新增```key```到**新的节点**的映射关系

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

1. 缓存容量```capacity```、缓存的```key```和```value```均为自然数（可以为0，代码中单独处理）
1. 读缓存```func get(key int) int```：（与lru相同）
    1. ```key```已存在，返回对应```value```
    1. ```key```不存在，返回```-1```
1. 写缓存```func put(key int, value int)```：
    1. ```key```已存在，修改对应```value```
    1. ```key```不存在，写入该组缓存，若写入前缓存容量已达上限，则应淘汰使用次数最少的缓存（记其使用次数为```n```）；  
       若使用次数为```n```的缓存数大于一个，则淘汰最久未使用的缓存（即，此时遵守lru规则）

##### 分析过程——Solution 1

> 目标：```O(1)```复杂度

1. 从lru的解决方案入手，使用双向链表主要维护缓存的使用次数，再使用map维护```key```到**节点**的映射关系
    1. 存在问题：缓存节点移动时（主要指插入环节），如何快速定位节点应该插入的位置？
        1. 场景分析：
            1. 删除节点：使用```key```，通过映射关系直接定位到节点
            1. 插入新节点：新缓存使用次数为```1```，应插入到所有使用次数为```1```的缓存节点的前面；  
               若不存在使用次数为```1```的缓存，则应插入到链表末尾
            1. 插入移除的节点：缓存使用次数```n+1```，应插入到所有使用次数为```n+1```的缓存节点的前面；  
               若不存在使用次数为```n+1```的缓存，则应插入到所有使用次数为```n```的缓存节点的前面；
        2. 分析结论：通过新增一个map，维护**缓存使用次数(记为```n```)** 到**所有使用次数为n的缓存中，最近使用的缓存节点**的映射关系,  
           可以快速定位节点应该在的位置。
2. 编写伪代码时，发现使用```recent```计算如“是否存在使用次数为```n```的缓存”等问题，其表达较为繁琐、不易理解且容易出错，   
   故引入新的变量```count```保存**指定使用次数的缓存节点数量**信息

##### 数据结构——Solution 1

1. recent：```map[int]*listNode```，```frequency```到**使用次数为```frequency```的节点中，最近使用的一个**的映射
    1. listNode data：```key```, ```value```, ```frequency```
2. count：```map[int]int```，```frequency```到**对应频率的节点数量**的映射
3. cache：```map[int]*listNode```，```key```到**节点**的映射
4. list：```*listNode```，双向链表，维护缓存的**使用次数（优先）** 和**上一次使用时间**
5. capacity：```int```，链表容量

##### 功能逻辑与伪代码——Solution 1

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
    1. 存在：参考**读缓存——```key```存在**，额外修改对应的value即可
    1. 不存在：
        1. 若当前缓存容量已达上限：
            1. 淘汰尾部的缓存节点（记节点```freq```为```n```）
            2. 若不存在其他```freq = n```的节点，则将```recent```置空
            3. 更新```cache```
            4. 更新```count```
        2. 构造新节点：key，value，frequency = 1
            1. 是否存在其他```frequency = 1```的节点：
                1. 存在：插入到它们的前面
                2. 不存在：插入链表尾部
            2. 更新```recent```
            3. 更新```cache```
            4. 更新```count```

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

##### 分析过程——Solution 2

> 目标：```O(1)```复杂度

1. 使用双向链表维护缓存的使用次数（以下简称频率链表，链表节点称为频率节点）；使用map维护频率节点
    1. 约定：链表正方向，缓存使用次数逐渐升高。举个例子：```0 -> 1 -> 2 -> 0```
2. 对于使用次数相同的缓存节点，将它们挂载到对应的频率节点上，使用双向链表（以下简称缓存链表，链表节点称为缓存节点）；使用map维护缓存节点
    1. 约定：链表正方向，缓存上一次使用时间逐渐久远。举个例子：```0 -> （最近使用的节点） -> （久未使用的节点） -> 0```
    2. 约定：当一个频率节点未挂载任何缓存节点时，我们说这个频率节点是**空**的

Q1：假设现在缓存容量已满、需要淘汰一个缓存节点，如何快速定位到需要淘汰的节点呢？  
A1：维护频率节点，每一次移动都删除空的频率节点，这样需要淘汰的节点就永远是**第一个频率节点的最后一个缓存节点**

##### 数据结构——Solution 2

1. freqList：```*freqNode```，频率链表，记录当前全部缓存的使用次数
    1. freqNode data：```frequency```, ```data(cacheList)```
2. freqMap：```map[int]*freqNode```，**使用次数**到**频率节点**的映射
3. cacheList：```*cacheNode```，缓存链表，记录当前全部缓存
    1. cacheNode data：```key```, ```value```, ```frequency```
4. cacheMap：```map[int]*cacheNode```，```key```到**缓存节点**的映射
5. capacity：```int```，链表容量

##### 功能逻辑与伪代码——Solution 2

###### 读缓存```func get(key int) int```

1. 查询```key```是否存在：
    1. 存在：（记缓存节点```frequency```为```n```）
        1. 将缓存节点从原位置删除，判断```freq == n```的频率节点是否为空：
            1. 频率节点为空：删除频率节点（删除链表节点+删除映射关系）
            2. 频率节点不空：无操作
        2. 将缓存节点```frequency``` ```+1```，挂载到```freq = n+1```的频率节点的缓存链表头部
            1. 若不存在这样的频率节点，则创建（新增链表节点+新增映射关系）
        3. 更新缓存节点```frequency```
        4. 返回节点的```value```
    1. 不存在：返回```-1```

```go 
func get(key int) int {
    node, ok := cacheMap.isExist(key)
    if !ok { // key 不存在
        return -1
    }
    
    // key 存在
    node.remove()
    
    freq := node.frequency
    freqNode := freqMap[freq]
    
    newFreqNode, ok := freqMap[freq+1]
    if !ok { // 不存在 freq = n+1 的节点
        newFreqNode = &freqNode{
            frequency: freq+1,
            data: newCacheList(),
        }
        
        freqNode.addBehind(newFreqNode)
        freqMap[freq+1] = newFreqNode
    }
    
    node.frequency++
    newFreqNode.data.addToHead(node)
     
    // 频率节点(freq = n)为空。拖到此时才删除，是为了帮助freq = n+1的频率节点在新增时定位
    if freqNode.data.isEmpty() {
        freqNode.remove()
        freqMap.delete(freq)
    }
    
    return node.value
}
```

###### 写缓存```func put(key int, value int)```

1. 查询key是否存在：
    1. 存在：参考**读缓存——```key```存在**，额外修改对应的value即可
    1. 不存在：
        1. 若当前缓存容量已达上限：
            1. 淘汰第一个频率节点的最后一个缓存节点（删除链表节点+删除映射关系）
            2. 若第一个频率节点为空且频率节点的频率```>1```，删除该频率节点（删除链表节点+删除映射关系）
        2. 构造缓存节点：key，value，frequency = 1；挂载到```freq = 1```的频率节点的缓存链表头部（新增链表节点+新增映射关系）
            1. 若不存在这样的频率节点，则创建（新增链表节点+新增映射关系）

```go 
func put(key int, value int) {
    node, ok := cacheMap.isExist(key)
    if ok { // key 存在
        get(key)
        node.value = value
        
        return
    }
    
    // key 不存在
    if cacheMap.isFull() {
        delNode := freqList.firstNode.data 
        delNode.remove()
        cacheMap.delete(delNode.key)
        
        freqNode := freqList.firstNode 
        if freqNode.data.isEmpty() && freqNode.frequency > 1 { // 等于1的频率节点不需要删除
            freqNode.remove()
            freqMap.delete(freqNode.freqency)
        }
    }
    
    freqNode, ok := freqMap.isExist(1)
    if !ok {
        freqNode = &freqNode{
            freqency: 1,
            data: newCacheList(),
        }
        
        freqList.addToHead(freqNode)
        freqMap[1] = freqNode
    }
    
    newCacheNode := &cacheNode{
        key: key,
        value: value,
        frequency: 1,
    }
    
    freqNode.data.addToHead(newCacheNode)
    cacheMap[key] = newcacheNode    
}
```

---

###### Mario

I've been pretending to work hard, but you're really growing up.
