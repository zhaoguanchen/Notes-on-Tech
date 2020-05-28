# swagger使用相关

## 报错问题

### 问题1：启动报错
```
        java.lang.NoSuchMethodError: com.google.common.collect.Multimaps.asMap(Lcom/google/common/collect/ListMultimap;)Ljava/util/Map;
```
#### 解决：grava包冲突pom文件依赖首航添加guava
```
 <dependency>
            <groupId>com.google.guava</groupId>
            <artifactId>guava</artifactId>
            <version>20.0</version>
        </dependency>
```