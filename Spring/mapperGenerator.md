# Spring代码生成器
```
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE generatorConfiguration PUBLIC "-//mybatis.org//DTD MyBatis Generator Configuration 1.0//EN"
        "http://mybatis.org/dtd/mybatis-generator-config_1_0.dtd" >
<generatorConfiguration>

    <properties resource="application-dev.yml"/>

    <!-- 数据库驱动包位置 -->
    <classPathEntry
            location="/Users/yiche/Documents/mysql-connector-java-5.1.44.jar" />

    <context id="Mysql" targetRuntime="MyBatis3Simple" defaultModelType="flat">
        <!-- TKmybatis配置 -->
        <property name="javaFileEncoding" value="UTF-8"/>
        <property name="beginningDelimiter" value="`"/>
        <property name="endingDelimiter" value="`"/>

        <!--<plugin type="tk.mybatis.mapper.generator.MapperPlugin">
            <property name="mappers" value="tk.mybatis.mapper.common.Mapper"/>
        </plugin>-->
        <!--<plugin type="tk.mybatis.mapper.generator.MapperPlugin">-->
            <!--<property name="mappers" value="operation.util.MyMapper"/>-->
        <!--</plugin>-->

        <commentGenerator>
            <!-- 是否去除自动生成的注释 true：是 ： false:否 -->
            <property name="suppressAllComments" value="true"/>
        </commentGenerator>
        <!-- 数据库链接URL、用户名、密码 -->
        <jdbcConnection driverClass="com.mysql.jdbc.Driver"
                        connectionURL="jdbc:mysql://192.168.15.55:3306/email"
                        userId="root"
                        password="123456">
        </jdbcConnection>
        <!-- 生成模型的包名和位置 -->
        <!--<javaModelGenerator targetPackage="${targetEntityPackage}" targetProject="${targetJavaProject}" />-->
        <javaModelGenerator targetPackage="com.yiche.bdc.dataexport.entity" targetProject="src/main/java"/>
        <!-- 生成的映射文件包名和位置 -->
        <!--<sqlMapGenerator targetPackage="${targetMapperPackage}" targetProject="${targetJavaProject}" />-->
        <sqlMapGenerator targetPackage="mapper" targetProject="src/main/resources"/>
        <!-- 生成service的包名和位置 -->
        <!--<javaClientGenerator targetPackage="${targetMapperPackage}" targetProject="${targetJavaProject}" type="XMLMAPPER" />-->
        <javaClientGenerator targetPackage="com.yiche.bdc.dataexport.dao" targetProject="src/main/java"
                             type="XMLMAPPER"/>

        <!-- 要生成的那些表 -->
        <table tableName="email_info" domainObjectName="EmailInfoEntity">
        </table>
    </context>
</generatorConfiguration>
```
