# Mysql


## index索引
## 索引类型
普通索引(INDEX)：最基本的索引，没有任何限制
唯一索引(UNIQUE)：与"普通索引"类似，不同的就是：索引列的值必须唯一，但允许有空值。
主键索引(PRIMARY)：它 是一种特殊的唯一索引，不允许有空值。
全文索引(FULLTEXT )：仅可用于 MyISAM 表， 用于在一篇文章中，检索文本信息的, 针对较大的数据，生成全文索引很耗时耗空间。
组合索引：为了更多的提高mysql效率可建立组合索引，遵循”最左前缀“原则。


## Tips
MySQL常见问题解决

#### 报错
使用命令行远程连接mysql时出现Bad handshake报错：
ERROR 1043 (08S01): Bad handshake

原因：mysql客户端版本不一致。
解决方法：检查客户端版本，重新安装一致的版本。


#### 自增id
修改初始自增ID
```sql
alter table users AUTO_INCREMENT=10000;
```