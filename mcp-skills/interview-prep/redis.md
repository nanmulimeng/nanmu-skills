# Redis 面试题库

> AI 注意：只用于**出题**，不要把 checkpoints 当答案直接告诉用户。

## 数据类型与底层结构

### Q1: String 类型的底层存储方式？
tags: [SDS, embstr, raw, int, 编码转换]
checkpoints:
- 三种编码：int / embstr / raw
- embstr vs raw 的分界线（44 字节，为什么？）
- SDS 和 C 字符串的区别（O(1) 长度、预分配、二进制安全）

### Q2: Hash 类型底层是怎么存的？
tags: [ZipList, Listpack, HashTable, 渐进式rehash]
checkpoints:
- 小数据量：ZipList/Listpack
- 大数据量：HashTable（两个 dictht）
- 渐进式 rehash 的过程（为什么用两个表？迁移期间怎么查？）

### Q3: ZSet 的底层实现？为什么同时用 SkipList 和 HashTable？
tags: [ZSet, SkipList, HashTable, 范围查询]
checkpoints:
- SkipList 按 score 排序（范围查询 O(log n)）
- HashTable 存 member → score（单查 O(1)）
- 两者互补，缺一不可
- ZipList/Listpack 到 SkipList 的转换阈值

---

## 单线程与高性能

### Q4: Redis 为什么单线程还这么快？
tags: [单线程, IO多路复用, 内存, 无锁]
checkpoints:
- 四个因素：纯内存 + 无锁 + IO 多路复用 + 高效数据结构
- IO 多路复用是核心（没有它单线程处理不了多连接）
- 不是"单线程进程"，有后台线程和 IO 线程

### Q5: Redis 6.0 多线程做了什么？命令执行为什么不改多线程？
tags: [多线程IO, 命令执行, 网络瓶颈]
checkpoints:
- 多线程只做网络读写，命令执行仍单线程
- 为什么不改命令执行？（加锁代价、原子性、事务语义）
- 什么场景下多线程 IO 有效（网络密集型 vs CPU 密集型）

---

## 持久化

### Q6: RDB 和 AOF 的区别？
tags: [RDB, AOF, 快照, 命令日志, 重写]
checkpoints:
- 文件内容：二进制快照 vs 文本命令
- 数据安全：可能丢数据 vs 取决于刷盘策略
- 恢复速度：快 vs 慢
- AOF 重写原理（fork 子进程 + 重写缓冲区）
- 混合持久化（RDB + AOF）

### Q7: RDB 的 fork 为什么会阻塞？
tags: [fork, 页表, COW, 大内存]
checkpoints:
- fork 需要复制页表（不是复制内存）
- 大内存实例 fork 耗时可达百毫秒
- COW（写时复制）机制
- 优化方向：控制单实例内存、用物理机

---

## 缓存问题

### Q8: 缓存穿透、击穿、雪崩的区别和解决方案？
tags: [穿透, 击穿, 雪崩, 布隆过滤器, 互斥锁]
checkpoints:
- 穿透：查不存在的数据 → 布隆过滤器/空值缓存
- 击穿：热点 key 过期 → 互斥锁/逻辑过期
- 雪崩：大量 key 同时过期 → 随机过期时间/多级缓存
- 三者的区别（从现象到原因到解决方案）

### Q9: Redisson 分布式锁的原理？
tags: [Redisson, 看门狗, Lua, 续期]
checkpoints:
- 加锁：SET NX + Lua 脚本保证原子性
- 看门狗：自动续期（1/3 过期时间）
- 可重入锁：Hash 结构存 {UUID:threadId:count}
- 主从切换导致锁丢失的问题（RedLock 的争议）

---

## 集群

### Q10: Redis Cluster 的数据分片规则？
tags: [Cluster, 16384槽, CRC16, MOVED, ASK]
checkpoints:
- 16384 槽，CRC16(key) % 16384
- MOVED vs ASK 重定向
- 为什么是 16384 不是更多（心跳包大小）
- 主从复制 + 自动故障转移
