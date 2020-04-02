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
