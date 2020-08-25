# Mybatis Plus

##QueryWrapper
```java
        QueryWrapper<CccSeatUser> wrapper = new QueryWrapper<>();
        wrapper.eq("ag_id", freeAgID);
        wrapper.isNotNull("priority");
        wrapper.isNotNull("user_id");
        wrapper.eq("is_del", 0);
        wrapper.orderByDesc("priority");
        List<CccSeatUser> cccSeatUserResultList = cccSeatUserMapper.selectList(wrapper);

```
```java
 /**
     * 根据 entity 条件，查询全部记录
     *
     * @param queryWrapper 实体对象封装操作类（可以为 null）
     */
    List<T> selectList(@Param(Constants.WRAPPER) Wrapper<T> queryWrapper);
```

# HuTool java工具集


# RESTful架构


# snowFlake 雪花算法
第一位为未使用，接下来的41位为毫秒级时间(41位的长度可以使用69年)
然后是5位datacenterId和5位workerId(10位的长度最多支持部署1024个节点）
最后12位是毫秒内的计数（12位的计数顺序号支持每个节点每毫秒产生4096个ID序号）
#### 应用
HuTool工具库
```java
Snowflake snowflake = IdUtil.createSnowflake(1, 1);
long id = snowflake.nextId();
```
#### 源码
```java
package cn.hutool.core.lang;

import java.io.Serializable;

import cn.hutool.core.date.SystemClock;
import cn.hutool.core.util.StrUtil;

/**
 * Twitter的Snowflake 算法<br>
 * 分布式系统中，有一些需要使用全局唯一ID的场景，有些时候我们希望能使用一种简单一些的ID，并且希望ID能够按照时间有序生成。
 * 
 * <p>
 * snowflake的结构如下(每部分用-分开):<br>
 * 
 * <pre>
 * 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000
 * </pre>
 * 
 * 第一位为未使用，接下来的41位为毫秒级时间(41位的长度可以使用69年)<br>
 * 然后是5位datacenterId和5位workerId(10位的长度最多支持部署1024个节点）<br>
 * 最后12位是毫秒内的计数（12位的计数顺序号支持每个节点每毫秒产生4096个ID序号）
 * 
 * 并且可以通过生成的id反推出生成时间,datacenterId和workerId
 * <p>
 * 参考：http://www.cnblogs.com/relucent/p/4955340.html
 * 
 * @author Looly
 * @since 3.0.1
 */
public class Snowflake implements Serializable{
    private static final long serialVersionUID = 1L;

    // Thu, 04 Nov 2010 01:42:54 GMT
    private final long twepoch = 1288834974657L;
    private final long workerIdBits = 5L;
    private final long datacenterIdBits = 5L;
    //// 最大支持机器节点数0~31，一共32个
    private final long maxWorkerId = -1L ^ (-1L << workerIdBits);
    // 最大支持数据中心节点数0~31，一共32个
    private final long maxDatacenterId = -1L ^ (-1L << datacenterIdBits);
    // 序列号12位
    private final long sequenceBits = 12L;
    // 机器节点左移12位
    private final long workerIdShift = sequenceBits;
    // 数据中心节点左移17位
    private final long datacenterIdShift = sequenceBits + workerIdBits;
    // 时间毫秒数左移22位
    private final long timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits;
    private final long sequenceMask = -1L ^ (-1L << sequenceBits);// 4095

    private long workerId;
    private long datacenterId;
    private long sequence = 0L;
    private long lastTimestamp = -1L;
    private boolean useSystemClock;

    /**
     * 构造
     * 
     * @param workerId 终端ID
     * @param datacenterId 数据中心ID
     */
    public Snowflake(long workerId, long datacenterId) {
        this(workerId, datacenterId, false);
    }

    /**
     * 构造
     * 
     * @param workerId 终端ID
     * @param datacenterId 数据中心ID
     * @param isUseSystemClock 是否使用{@link SystemClock} 获取当前时间戳
     */
    public Snowflake(long workerId, long datacenterId, boolean isUseSystemClock) {
        if (workerId > maxWorkerId || workerId < 0) {
            throw new IllegalArgumentException(StrUtil.format("worker Id can't be greater than {} or less than 0", maxWorkerId));
        }
        if (datacenterId > maxDatacenterId || datacenterId < 0) {
            throw new IllegalArgumentException(StrUtil.format("datacenter Id can't be greater than {} or less than 0", maxDatacenterId));
        }
        this.workerId = workerId;
        this.datacenterId = datacenterId;
        this.useSystemClock = isUseSystemClock;
    }
    
    /**
     * 根据Snowflake的ID，获取机器id
     *
     * @param id snowflake算法生成的id
     * @return 所属机器的id
     */
    public long getWorkerId(long id) {
        return id >> workerIdShift & ~(-1L << workerIdBits);
    }

    /**
     * 根据Snowflake的ID，获取数据中心id
     *
     * @param id snowflake算法生成的id
     * @return 所属数据中心
     */
    public long getDataCenterId(long id) {
        return id >> datacenterIdShift & ~(-1L << datacenterIdBits);
    }

    /**
     *根据Snowflake的ID，获取生成时间
     *
     * @param id snowflake算法生成的id
     * @return 生成的时间
     */
    public long getGenerateDateTime(long id) {
        return (id >> timestampLeftShift & ~(-1L << 41L)) + twepoch;
    }

    /**
     * 下一个ID
     * 
     * @return ID
     */
    public synchronized long nextId() {
        long timestamp = genTime();
        if (timestamp < lastTimestamp) {
            // 如果服务器时间有问题(时钟后退) 报错。
            throw new IllegalStateException(StrUtil.format("Clock moved backwards. Refusing to generate id for {}ms", lastTimestamp - timestamp));
        }
        if (lastTimestamp == timestamp) {
            sequence = (sequence + 1) & sequenceMask;
            if (sequence == 0) {
                timestamp = tilNextMillis(lastTimestamp);
            }
        } else {
            sequence = 0L;
        }

        lastTimestamp = timestamp;

        return ((timestamp - twepoch) << timestampLeftShift) | (datacenterId << datacenterIdShift) | (workerId << workerIdShift) | sequence;
    }
    
    /**
     * 下一个ID（字符串形式）
     *
     * @return ID 字符串形式
     */
    public String nextIdStr() {
        return Long.toString(nextId());
    }

    // ------------------------------------------------------------------------------------------------------------------------------------ Private method start
    /**
     * 循环等待下一个时间
     * 
     * @param lastTimestamp 上次记录的时间
     * @return 下一个时间
     */
    private long tilNextMillis(long lastTimestamp) {
        long timestamp = genTime();
        while (timestamp <= lastTimestamp) {
            timestamp = genTime();
        }
        return timestamp;
    }

    /**
     * 生成时间戳
     * 
     * @return 时间戳
     */
    private long genTime() {
        return this.useSystemClock ? SystemClock.now() : System.currentTimeMillis();
    }
    // ------------------------------------------------------------------------------------------------------------------------------------ Private method end
}

```
# HttpClient

##### post中文参数乱码
增加header设置编码为UTF-8
```java
headerMap.put("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8");
```



## HttpClient下载文件

```java
public void run() {
        CloseableHttpClient client = HttpClientBuilder.create().build();
        String url = cccTellCall.getRecord_url();
        if (url == null || url.isEmpty()) {
            log.error("getRecord_url录音地址为空");
            return;
        }

        String fileName = url.substring(url.lastIndexOf("/") + 1);
        String parentPath = "call/" + fileName.substring(0, 8) + "/";
        String targetName = parentPath + fileName;
        HttpGet request = new HttpGet(url);
        try {
            HttpResponse response = client.execute(request);
            HttpEntity entity = response.getEntity();
            InputStream is = entity.getContent();
            String ossUrl = OssServiceUtil.uploadSftpFileFolder(is, targetName);
            if (StringUtils.isEmpty(ossUrl)) {
                log.error("上传录音文件失败, oss返回结果为空: {}", ossUrl);
            }
            cccTellCall.setCrm_record_url(ossUrl);
            log.info("上传录音成功。oss地址为：" + ossUrl);
            cccTellCall.setRecord_down(true);
            cccTellCallMapper.updateById(cccTellCall);
        } catch (Exception e) {
            log.error("上传录音文件失败", e);
        } finally {
            request.releaseConnection();
        }
}
```
# HuTool


## 字符串
字符串分割 .split() 
特殊字符如“.”作为分隔符需要添加转义

## spring Boot 配置文件不生效
重新install一下。


##Hashmap  treemap 原理和区别


## List

Arrays.asList()之后使用remove()
为啥使用了Arrays.asList()之后使用remove是错误用法，我们看一下asList()的源码就能知道了。Arrays.asList()返回的是一个指定数组长度的列表，所以不能做Add、Remove等操作。至于为啥是返回的是固定长度的，看下面源码，asList()函数中调用的new ArrayList<>()并不是我们常用的ArrayList类，而是一个Arrays的内部类，也叫ArrayList，而且这个内部类也是基于数组实现的，但它有一个明显的关键字修饰，那就是final。都用final修饰了，那是肯定不能再对它进行add/remove操作的。


Arrays.asList()之后需要进行add/remove操作，可以使用下面这种方式：
```java
    List list = new ArrayList(Arrays.asList(arr));
```




## text

MySQL 3种text类型：

    TEXT
    MEDIUMTEXT
    LONGTEXT

text的长度

    TEXT：65,535 bytes ~64kb
    MEDIUMTEXT： 16,777,215 bytes ~16Mb
    LONGTEXT： 4,294,967,295 bytes ~4Gb
 

## @ResponseBody    @RequestBody

@ResponseBody的作用其实是将java对象转为json格式的数据。
@RequestBody 注解则是将 HTTP 请求正文插入方法中，使用适合的 HttpMessageConverter 将请求体写入某个对象。



## 函数式编程
```java
   seatResultList = seatResultList.stream().filter(seat -> existAgId.contains(seat.getAg_id())).collect(Collectors.toList());
```   
## 字符串分割转Long
```java
List<Long> checkIdList = Arrays.stream(checkIdStr.split(",")).map(checkId -> Long.parseLong(checkId.trim())).collect(Collectors.toList());
```

## 分组 计数
```java
    Map<String, Long> map =
                    qa.stream().collect(Collectors.groupingBy(String::intern, Collectors.counting()));
```

## map排序
转为TreeMap，使其有序
```java
    TreeMap<String, Long> map = list.stream().collect(Collectors.groupingBy(item -> sf.format(new Date(item)), TreeMap::new, Collectors.counting()));
```


## 权重
```java
    private String getFreeAgID() {

        List<SeatResult> freeSeatList = getFreeSeatList();
        List<String> freeAgIdList = freeSeatList.stream().map(SeatResult::getAg_id).collect(Collectors.toList());

        Random random = new Random();
        int weightSum = 0;
        HashMap agIdWeightMap = JSON.parseObject(seatWeightValue, HashMap.class);
        Map<String, Integer> map = new HashMap<>();

        if (freeAgIdList.isEmpty()) {
            map = agIdWeightMap;
        } else {
            for (String item : freeAgIdList) {
                if (agIdWeightMap.containsKey(item)) {
                    map.put(item, (Integer) agIdWeightMap.get(item));
                }
            }
        }

        for (Map.Entry<String, Integer> item : map.entrySet()) {
            weightSum += item.getValue();
        }

        String resAgID = null;
        int n = random.nextInt(weightSum);
        int m = 0;
        for (Map.Entry<String, Integer> item : map.entrySet()) {
            if (m <= n && n < m + item.getValue()) {
                resAgID = item.getKey();
                break;
            }
            m += item.getValue();
        }
        return resAgID;
    }
    ```

# 注解

#### @Deprecated
标志此方法过时，不推荐调用

## 事务回滚
####  @Transactional(rollbackFor = Exception.class)
事务一致性

测试代码：
```java
 @Override
    @Transactional(rollbackFor = Exception.class)
    public ResponseData test() {
        try {
            save1();
            save2();
            return ResponseData.success();
        } catch (Exception e) {
            throw new RuntimeException(e.getMessage());
        }
    }


    private void save1() {

        QualitycheckChangeRecord qualitycheckChangeRecord = new QualitycheckChangeRecord();
        qualitycheckChangeRecord.setOperatorId(11111L);
        qualitycheckChangeRecord.setCheckId(11111L);
        qualitycheckChangeRecord.setCreateTime(new Date());
        qualitycheckChangeRecord.setIsDel(1);
        qualitycheckChangeRecordMapper.insert(qualitycheckChangeRecord);
        log.info("测试新增质检单变更记录1:{}", qualitycheckChangeRecord.toString());

    }


    private void save2() {
 
        int a = 5 / 0;

        QualitycheckChangeRecord qualitycheckChangeRecord = new QualitycheckChangeRecord();
        qualitycheckChangeRecord.setOperatorId(222222L);
        qualitycheckChangeRecord.setCheckId(222222L);
        qualitycheckChangeRecord.setCreateTime(new Date());
        qualitycheckChangeRecord.setIsDel(1);
        qualitycheckChangeRecordMapper.insert(qualitycheckChangeRecord);
        log.info("测试新增质检单变更记录2:{}", qualitycheckChangeRecord.toString());

    }
```

##  数据返回前端
#### @JSONField(serialize = false)
添加该注解过滤返回字段
```java
@JSONField(serialize = false)
private Date passTime;
```


## Spring MVC 的过程
###  Spring MVC 的过程
  

## 集合排序
```java
  List<PhoneRecordResult> phoneRecordResultList = new ArrayList<>();
   phoneRecordResultList.sort((o1, o2) -> {
            int flag;
            if (null == o1.getStartTime() || null == o2.getStartTime()) {
                flag = o1.getId().compareTo(o2.getId());
            } else {
                flag = o1.getStartTime().compareTo(o2.getStartTime());
            }
            if (flag < 0) {
                flag = 1;
            } else if (flag > 0) {
                flag = -1;
            }
            return flag;
        });
```



# 多线程
#### execute
```Java
        handleInfo.execute(() -> {
            try {
                log.info("方法体");
                handleCustomInfo(systemTaskRecord);
                distribute(systemTaskRecord);
                this.baseMapper.updateById(systemTaskRecord);
            } catch (Exception e) {
                log.error("处理补充线索与分配异常：{}", e.getMessage());
            }
        });
```

 

#### 实现runnable线程类
```java
@Slf4j
class RecordDownloadThread implements Runnable {

    CccTellCall cccTellCall;
    CccTellCallMapper cccTellCallMapper;

    RecordDownloadThread(CccTellCallMapper cccTellCallMapper, CccTellCall cccTellCall) {
        this.cccTellCall = cccTellCall;
        this.cccTellCallMapper = cccTellCallMapper;
    }

    /**
     * @Author zhaoguanchen
     * @Description 下载录音文件到oss
     * @Date 2019/12/16 14:39
     **/
    @Override
    public void run() {
        CloseableHttpClient client = HttpClientBuilder.create().build();
        String url = cccTellCall.getRecord_url();
        if (url == null || url.isEmpty()) {
            log.error("getRecord_url录音地址为空");
            return;
        }

        String fileName = url.substring(url.lastIndexOf("/") + 1);
        String parentPath = "call/" + fileName.substring(0, 8) + "/";
        String targetName = parentPath + fileName;
        HttpGet request = new HttpGet(url);
        try {
            HttpResponse response = client.execute(request);
            HttpEntity entity = response.getEntity();
            InputStream is = entity.getContent();
            String ossUrl = OssServiceUtil.uploadSftpFileFolder(is, targetName);
            if (StringUtils.isEmpty(ossUrl)) {
                log.error("上传录音文件失败, oss返回结果为空: {}", ossUrl);
            }
            cccTellCall.setCrm_record_url(ossUrl);
            log.info("上传录音成功。oss地址为：" + ossUrl);
            cccTellCall.setRecord_down(true);
            cccTellCallMapper.updateById(cccTellCall);
        } catch (Exception e) {
            log.error("上传录音文件失败", e);
        } finally {
            request.releaseConnection();
           
        }
    }
}
```
初始化
```java
 private static ExecutorService executorService = Executors.newFixedThreadPool(5);
```
方法内调用执行
```java
  executorService.execute(new RecordDownloadThread(cccTellCallMapper, cccTellCall));
```

# 工具类
## 加密解密

XXTEAUtil 工具包
```java
XXTEAUtil.decryptBase64StringToString(str, key)
XXTEAUtil.encryptToBase64String(str, key)

```



# 树形层级结构

map  list为引用传递，修改map中的对象，会修改treeList.add中的对象。
示例：
```java
public class Demo{
    public static List<FilterItemResult> listTreeByPId(List<FilterItemResult> list) {
        if (CollectionUtils.isEmpty(list)) {
            return null;
        }
        List<FilterItemResult> treeList = new ArrayList<>();
        Map<Long, FilterItemResult> map = list.stream().collect(Collectors.toMap(FilterItemResult::getId, item -> item));

        for (FilterItemResult node : list) {
            if (node.getParentId() == 0) {
                treeList.add(node);
            } else {
                FilterItemResult parent = map.get(node.getParentId());
                if (parent != null) {
                    if (null == parent.getChild()) {
                        parent.setChild(new ArrayList<>());
                    }
                    parent.getChild().add(node);
                }
            }
        }

        return treeList;
    }
    
    public static List<FilterItemResult> listTreeByCode(List<FilterItemResult> list) {
        if (CollectionUtils.isEmpty(list)) {
            return null;
        }
        List<FilterItemResult> treeList = new ArrayList<>();
        Map<String, FilterItemResult> map = new HashMap<>();
        for (FilterItemResult node : list) {
            node.setChild(new ArrayList<>());
            map.put(node.getCode(), node);
        }

        for (FilterItemResult node : list) {
            if (!node.getCode().contains("_")) {
                treeList.add(node);
            } else {
                FilterItemResult parent = map.get(StringUtils.substringBeforeLast(node.getCode(), "_"));
                if (parent != null) {
                    parent.getChild().add(node);
                }
            }
        }
        return treeList;
    }
}
```