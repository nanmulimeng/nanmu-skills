# Java 基础面试题库

> AI 注意：只用于**出题**，不要把 checkpoints 当答案直接告诉用户。

## 集合框架

### Q1: HashMap 的 put 流程？
tags: [HashMap, hash, 桶, 链表, 红黑树, 扩容]
checkpoints:
- hash 定位桶 → 桶空直接放 → 桶有值遍历链表/树
- key 存在覆盖，不存在尾插
- 链表 >= 8 且数组 >= 64 转红黑树
- 元素 > threshold 扩容 2 倍
- JDK7 头插 vs JDK8 尾插（为什么改？）

### Q2: HashMap 为什么线程不安全？ConcurrentHashMap 怎么解决？
tags: [线程安全, ConcurrentHashMap, CAS, synchronized, Segment]
checkpoints:
- HashMap：put 并发数据覆盖、扩容死循环(JDK7)、size 不准确
- JDK7 ConcurrentHashMap：Segment 分段锁
- JDK8 ConcurrentHashMap：CAS + synchronized 锁桶头节点
- 为什么 JDK8 去掉 Segment？

### Q3: HashMap 的扩容机制？
tags: [扩容, rehash, 高低位拆分,负载因子]
checkpoints:
- 负载因子为什么是 0.75（时间和空间折中）
- JDK8 高低位拆分避免重新计算 hash
- 扩容时链表和红黑树的处理

---

## JVM

### Q4: JVM 内存结构？
tags: [堆, 栈, 元空间, 程序计数器, 直接内存]
checkpoints:
- 堆（新生代 Eden/S0/S1 + 老年代）
- 栈帧内容（局部变量表、操作数栈、动态链接、返回地址）
- 元空间替代永久代的原因
- 直接内存与 NIO

### Q5: G1 垃圾收集器的特点？
tags: [G1, Region, Mixed GC, 可预测停顿, ZGC]
checkpoints:
- Region 划分 + 优先回收价值最大的 Region
- 可预测停顿时间模型
- Mixed GC vs Full GC
- 什么情况退化成 Full GC
- ZGC 的染色指针和并发整理

### Q6: 类加载机制和双亲委派？
tags: [类加载, 双亲委派, Bootstrap, Tomcat]
checkpoints:
- 加载→验证→准备→解析→初始化
- 双亲委派模型（为什么需要？保证核心类不被篡改）
- Tomcat 如何打破双亲委派（先本地后委派）

---

## 并发编程

### Q7: ThreadPoolExecutor 的核心参数和执行流程？
tags: [线程池, corePoolSize, 拒绝策略, workQueue]
checkpoints:
- 7 个参数
- 执行流程：核心线程 → 队列 → 非核心线程 → 拒绝
- 四种拒绝策略
- 为什么不用 Executors 快捷方法（无界队列 OOM）

### Q8: AQS 的核心思想？
tags: [AQS, state, CLH队列, ReentrantLock, 公平锁]
checkpoints:
- volatile state + FIFO 等待队列
- 独占模式 vs 共享模式
- 公平锁 vs 非公平锁区别
- CountDownLatch vs CyclicBarrier

### Q9: volatile 和 CAS？
tags: [volatile, 可见性, 内存屏障, CAS, ABA]
checkpoints:
- volatile 保证可见性（内存屏障）但不保证原子性
- 为什么 i++ 不是原子操作
- CAS 原理 + ABA 问题（AtomicStampedReference）
- synchronized vs volatile vs ReentrantLock 对比
