## 加密解密

XXTEAUtil 工具包
```java
XXTEAUtil.decryptBase64StringToString(str, key)
XXTEAUtil.encryptToBase64String(str, key)

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

## 线程池

实现runnable线程类
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


## 函数式编程
```java
   seatResultList = seatResultList.stream().filter(seat -> existAgId.contains(seat.getAg_id())).collect(Collectors.toList());
```   

## 字符串
字符串分割 .split() 
特殊字符如“.”作为分隔符需要添加转义

## spring Boot 配置文件不生效
重新install一下。


##Hashmap  treemap 原理和区别

## 事务回滚
  @Transactional(rollbackFor = Exception.class)


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

## List

Arrays.asList()之后使用remove()
为啥使用了Arrays.asList()之后使用remove是错误用法，我们看一下asList()的源码就能知道了。Arrays.asList()返回的是一个指定数组长度的列表，所以不能做Add、Remove等操作。至于为啥是返回的是固定长度的，看下面源码，asList()函数中调用的new ArrayList<>()并不是我们常用的ArrayList类，而是一个Arrays的内部类，也叫ArrayList，而且这个内部类也是基于数组实现的，但它有一个明显的关键字修饰，那就是final。都用final修饰了，那是肯定不能再对它进行add/remove操作的。


Arrays.asList()之后需要进行add/remove操作，可以使用下面这种方式：
```java
    List list = new ArrayList(Arrays.asList(arr));
```


## 分组 计数
```java
    Map<String, Long> map =
                    qa.stream().collect(Collectors.groupingBy(String::intern, Collectors.counting()));
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

## 字符串分割转Long
```java
List<Long> checkIdList = Arrays.stream(checkIdStr.split(",")).map(checkId -> Long.parseLong(checkId.trim())).collect(Collectors.toList());
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
## 注解

@Deprecated


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