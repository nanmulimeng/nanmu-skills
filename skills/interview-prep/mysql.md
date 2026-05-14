# MySQL 面试题库

> AI 注意：这是题库，只用于**出题**。不要把考点标签内容当作答案直接告诉用户。
> 考点标签是给你（AI）用来判断用户回答是否完整的检查清单。

## 索引原理

### Q1: B+ 树索引和 Hash 索引的区别？
tags: [索引类型, B+树, Hash, 范围查询]
checkpoints:
- B+ 树支持范围查询和排序，Hash 只支持等值
- B+ 树非叶子节点不存数据，树更矮
- InnoDB 主键索引是 B+ 树，不能用 Hash 做主键

### Q2: 什么情况下索引会失效？
tags: [索引失效, 最左前缀, 隐式转换]
checkpoints:
- 函数运算、隐式类型转换、前导模糊查询
- OR 条件部分列无索引
- 联合索引不满足最左前缀
- 不等于/NOT IN

### Q3: 什么是覆盖索引？什么是回表？
tags: [覆盖索引, 回表, 聚簇索引, 二级索引]
checkpoints:
- 覆盖索引：查询字段都在索引中，不需要回表
- 回表：二级索引查到主键值，再用主键查聚簇索引
- InnoDB 二级索引叶子节点自带主键

### Q4: 聚簇索引和非聚簇索引的区别？
tags: [聚簇索引, 非聚簇, InnoDB, MyISAM]
checkpoints:
- 聚簇索引：叶子节点存完整行数据（InnoDB 主键）
- 非聚簇：叶子节点存行地址/主键值（MyISAM 所有索引）
- 聚簇索引一张表只有一个

---

## 事务与隔离

### Q5: MySQL 的四种隔离级别？各自解决什么问题？
tags: [隔离级别, 脏读, 不可重复读, 幻读, MVCC]
checkpoints:
- 四个级别及各自问题
- InnoDB RR 通过 MVCC + Gap Lock 解决了幻读
- RC 和 RR 的 Read View 创建时机区别

### Q6: MVCC 的实现原理？
tags: [MVCC, Read View, undo log, trx_id]
checkpoints:
- 隐藏列：trx_id + roll_pointer
- Read View 的可见性判断规则
- RC 每次查询新建 vs RR 事务开始时创建

### Q7: 当前读和快照读的区别？
tags: [当前读, 快照读, MVCC, RR幻读]
checkpoints:
- 快照读：普通 SELECT，读 MVCC 快照
- 当前读：SELECT FOR UPDATE / INSERT / UPDATE / DELETE，读最新数据
- RR 下快照读不会幻读，但当前读可能（需 Gap Lock 防护）

---

## 锁机制

### Q8: InnoDB 的行锁有哪些类型？
tags: [Record Lock, Gap Lock, Next-Key Lock, 意向锁]
checkpoints:
- 三种行锁的锁定范围
- Next-Key Lock = Record Lock + Gap Lock
- 唯一索引等值查询且记录存在时退化为 Record Lock

### Q9: 什么情况下行锁会退化为表锁？
tags: [行锁退化为表锁, 索引失效, 锁升级]
checkpoints:
- 查询条件不走索引
- 不同事务分别用不同索引锁定同一行

---

## 日志系统

### Q10: redo log、undo log、binlog 的区别？
tags: [redo log, undo log, binlog, 两阶段提交]
checkpoints:
- 存储引擎层 vs Server 层
- 物理日志 vs 逻辑日志
- 各自用途：崩溃恢复 / 回滚+MVCC / 复制+恢复
- 两阶段提交保证一致性

### Q11: 为什么需要两阶段提交？
tags: [两阶段提交, 主从一致性, prepare, commit]
checkpoints:
- 保证 redo log 和 binlog 一致
- 异常场景分析（prepare 后崩溃、commit 前崩溃）

---

## 性能优化

### Q12: 一条慢 SQL 的排查思路？
tags: [EXPLAIN, 慢查询, 执行计划, type]
checkpoints:
- EXPLAIN 看 type/key/rows/Extra
- type 从优到劣的顺序
- 深分页优化方案
- 是否走索引、是否有临时表/文件排序
