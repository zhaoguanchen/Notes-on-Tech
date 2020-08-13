# 阿里云


## 应用配置管理 ACM


## 访问控制 RAM


## 数据管理 DMS

## 对象存储 OSS

```java
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.CopyObjectRequest;
import com.aliyun.oss.model.ObjectMetadata;
import com.netease.nsip.component.common.utils.StringUtils;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Date;
import java.util.UUID;

 
@Slf4j
@Component
public class OssServiceUtil {

    // 访问域名
    private static String endpoint;
    // 用户
    private static String accessKeyId;
    // 密钥
    private static String accessKeySecret;
    //存储空间
    private static String bucketName;

    private static String url;

    private static String BASE_PATH = "";

    @Value("${oss.endpoint}")
    public void setEndpoint(String endpoint) {
        OssServiceUtil.endpoint = endpoint;
    }

    @Value("${oss.accesskeyid}")
    public void setAccessKeyId(String accessKeyId) {
        OssServiceUtil.accessKeyId = accessKeyId;
    }

    @Value("${oss.accesskeysecret}")
    public void setAccessKeySecret(String accessKeySecret) {
        OssServiceUtil.accessKeySecret = accessKeySecret;
    }

    @Value("${oss.bucketname}")
    public void setBucketName(String bucketName) {
        OssServiceUtil.bucketName = bucketName;
    }

    @Value("${oss.url}")
    public void setUrl(String url) {
        OssServiceUtil.url = url;
    }

    /**
     * 文件上传
     *
     * @param fileInputStream
     * @param fileName
     * @return
     */
    public static String uploadSftpFileFolder(InputStream fileInputStream, String fileName) {
        OSS ossClient = null;
        try {
            ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret);
            ossClient.putObject(bucketName, BASE_PATH + fileName, fileInputStream);
        } catch (Exception e) {
            log.error("[OssService] uploadSftpSettleFile error: {}", fileName, e);
            return null;
        } finally {
            if (null == ossClient) {
                log.error("[OssService] 上传文件时ossClient新建失败 {} ", fileName);
            } else {
                ossClient.shutdown();
            }
        }
        log.info("地址：" + url + BASE_PATH + fileName);
        return url + BASE_PATH + fileName;
    }

    /**
     * @Author zhaoguanchen
     * @Description 上传mp3录音文件
     * @Date 2019/12/30 16:57
     **/
    public static String uploadRecordFile(InputStream fileInputStream, String fileName) {
        OSS ossClient = null;
        try {
            ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret);
            ObjectMetadata metadata = new ObjectMetadata();
            metadata.setContentDisposition("attachment");
            ossClient.putObject(bucketName, BASE_PATH + fileName, fileInputStream, metadata);
        } catch (Exception e) {
            log.error("[OssService] uploadSftpSettleFile error: {}", fileName, e);
            return null;
        } finally {
            if (null == ossClient) {
                log.error("[OssService] 上传录音文件时ossClient新建失败 {} ", fileName);
            } else {
                ossClient.shutdown();
            }
        }
        log.info("地址：" + url + BASE_PATH + fileName);
        return url + BASE_PATH + fileName;
    }

    /**
     * @Author zhaoguanchen
     * @Description 生成下载链接
     * @Date 2019/12/31 14:37
     **/
    public static String generateDownloadUrl(String sourceBucketName, String sourceObjectName) {
        OSS ossClient = null;
        try {
            ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret);
            Date expiration = new Date(System.currentTimeMillis() + 3600L * 1000 * 24 * 180);
            URL url = ossClient.generatePresignedUrl(sourceBucketName, sourceObjectName, expiration);
            log.info("[OssService] 获取文件下载链接：{},{}", url, sourceObjectName);
            return url.toString();
        } catch (Exception e) {
            log.error("[OssService] 生成下载链接异常：{},{}", e, sourceObjectName);
            return null;
        } finally {
            if (null == ossClient) {
                log.error("[OssService] 生成下载链接时ossClient新建失败 {} ", sourceObjectName);
            } else {
                ossClient.shutdown();
            }
        }
    }

    /**
     * @Author zhaoguanchen
     * @Description 设置文件http元信息   添加 ContentDisposition：attachment
     * @Date 2019/12/31 14:33
     **/
    public static void updateObjectMetadata(String sourceBucketName, String sourceObjectName) {
        OSS ossClient = null;
        try {
            ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret);
            ObjectMetadata metadata = ossClient.getObjectMetadata(sourceBucketName, sourceObjectName);
            if (!"attachment".equals(metadata.getContentDisposition())) {
                CopyObjectRequest request = new CopyObjectRequest(sourceBucketName, sourceObjectName, sourceBucketName, sourceObjectName);
                ObjectMetadata meta = new ObjectMetadata();
                meta.setContentDisposition("attachment");
                request.setNewObjectMetadata(meta);
                ossClient.copyObject(request);
                log.info("[OssService] 设置文件http元信息： {}", sourceObjectName);
            } else {
                log.info("可直接下载");
            }
        } catch (Exception e) {
            log.error("[OssService] 设置文件http元信息异常： {} ", e.toString());
        } finally {
            if (null == ossClient) {
                log.error("[OssService] 设置文件http元信息时ossClient新建失败 {} ", sourceObjectName);
            } else {
                ossClient.shutdown();
            }

        }
    }


    /**
     * 下载网络图片，并将图片上传至sftp上
     *
     * @param img
     * @return
     * @throws Exception
     */
    public String uploadHttpImg(String img) throws Exception {
        log.info("img:{}", img);
        String imgUrl;
        if (StringUtils.isEmpty(img)) {
            return null;
        }
        URL url = new URL(img);
        HttpURLConnection http = (HttpURLConnection) url.openConnection();
        http.setRequestMethod("GET");
        http.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");
        http.setDoOutput(true);
        http.setDoInput(true);
        http.connect();
        InputStream is = http.getInputStream(); //网络返回的输入流
        try {
            String[] arr = img.split("/");
            String fileName = arr[arr.length - 1];
            //如果给的图片不是.结尾的，那么默认为jpg形式
            fileName = fileName.indexOf('.') > 0 ? fileName.substring(img.indexOf('.')) : UUID.randomUUID().toString() + ".jpg";
            imgUrl = uploadSftpFileFolder(is, fileName);
            if (StringUtils.isEmpty(imgUrl)) {
                log.error("上传图片失败, oss返回结果为空: {}", imgUrl);
            }
        } catch (Exception e) {
            log.error("上传图片失败", e);
            imgUrl = null;
        } finally {
            http.disconnect();
        }
        return imgUrl;
    }
}

```