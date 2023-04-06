### springcloud 详细教程

### 一、前言闲聊和课程说明

教学视频地址 [springcloud](https://www.bilibili.com/video/BV18E411x7eT)

### 二、零基础微服务架构理论入门

#### 2.1 什么是微服务

> In short, the microservice architectural style is an approach to developing a single application as a suite of small services, each running in its own process and communicating with lightweight mechanisms, often an HTTP resource API. These services are built around business capabilities and independently deployable by fully automated deployment machinery. There is a bare minimum of centralized management of these services, which may be written in different programming languages and use different data storage technologies.——James Lewis and Martin Fowler (2014)

- 微服务是一种架构风格
- 一个应用拆分为一组小型服务
- 每个服务运行在自己的进行内，也就是可独立部署和升级
- 服务之间使用轻量级HTTP交互
- 服务围绕业务功能拆分
- 可以由全自动部署机制独立部署
- 去中心化，服务自治。服务可以使用不同的语言，不同的存储技术。

#### 2.2 主题词01：现代数组化生活-落地维度

- 手机
- PC
- 智能家居
- ...

#### 2.3 主题词02：分布式微服务架构-落地维度

满足哪些维度？支撑起这些维度的具体技术？

![img](images/fa69e6841ce850672a3ec9cf8f4acad8.png)

- 服务调用
- 服务降级
- 服务注册与发现
- 服务熔断
- 负载均衡
- 服务消息队列
- 服务网关
- 配置中心管理
- 自动化构建部署
- 服务监控
- 全链路追踪
- 服务定时任务
- 调度操作

#### 2.4 Spring Cloud 简介

是什么？符合微服务技术维度

Spring Cloud=分布式微服务架构的一站式解决方案，是多种微服务架构落地技术的集合体，俗称微服务全家桶

猜猜SpringCloud这个集合里有多少种技术？

![img](images/eeb48f15799b978e45ed980172c9f06e.png)

springcloud 俨然已经成为微服务开发的主流技术栈，在国内开发者社区非常火爆。

威力十足，互联网大厂微服务架构案例

**京东的:**

![京东的](images/2b9f44abea91af3c4b77c1c77eae6eb3.png)



**阿里的：**

![阿里的](images/ef6092b03932cb7f6f4b7e20ff370261.png)



**京东物流：**

![京东物流](images/b7f15e802845e6ecdc4c13e2685789c1.png)

![微服务的简单概括](images/60b96a66ac3b4dceda8f7ac2f8d8d79e.png)

#### 2.5 Spring Cloud 技术栈

![netflix](images/fa347f3da197c3df86bf5d36961c8cde.png)



![img](images/b39a21012bed11a837c1edff840e5024.png)



![img](images/fc8ed10fca5f7cc56e4d4623a39245eb.png)

### 三、第二季Boot和Cloud版本选型

#### 3.1 SpringBoot2.x 版

- Spring Boot 2.X 版
  1. [源码地址](https://github.com/spring-projects/spring-boot/releases/)
  2. Spring Boot 2 的新特性
  3. 通过上面官网发现，Boot官方强烈建议你升级到2.X以上版本
- Spring Cloud H版
  1. [源码地址](https://github.com/spring-projects/spring-boot/releases/)
  2. [官网](https://spring.io/projects/spring-cloud#overview)

- Spring Boot 与 Spring Cloud 兼容性查看
  1. [文档](https://spring.io/projects/spring-cloud#adding-spring-cloud-to-an-existing-spring-boot-application)
  2. [JSON接口](https://start.spring.io/actuator/info)

- 接下来开发用到的组件版本
  1. Cloud - Hoxton.SR1
  2. Boot - 2.2.2.RELEASE
  3. Cloud Alibaba - 2.1.0.RELEASE
  4. Java - Java 8
  5. Maven - 3.5及以上
  6. MySQL - 5.7及以上

### 四、Cloud组件停更说明

- 停更已发的升级惨案
  1. 停更不停用
  2. 被动修复bugs不再接受合并请求
  3. 不再发布新版本

- Cloud升级

  1. 服务注册
     1. ❎Eureka
     2. ✅Zookeeper
     3. ✅Consul
     4. ✅Nacos

  2. 服务调用
     1. ✅Ribbon
     2. ✅LoadBalancer

  3. 服务调用2
     1. ❎Feign
     2. ✅OpenFeign
  4. 服务降级
     1. ❎Hystrix
     2. ✅resilience4j
     3. ✅sentienl

  4. 服务网关
     1. ✅Zuul
     2. ！Zuul2
     3. ✅gateway

  5. 服务配置
     1. ❎Config
     2. ✅Nacos

  6. 服务总线
     1. ❎Bus
     2. ✅Nacos

[Spring Cloud中文文档](https://www.bookstack.cn/read/spring-cloud-docs/docs-index.md)

[Spring Cloud官方文档](https://cloud.spring.io/spring-cloud-static/Hoxton.SR1/reference/htmlsingle/)

[Spring Boot官方文档](https://docs.spring.io/spring-boot/docs/2.2.2.RELEASE/reference/htmlsingle/)

### 五、父工程Project空间新建

约定 > 配置 > 编码

创建微服务cloud整体聚合父工程Project，有8个关键步骤：

1. New Project - maven工程 -create from archetype: maven-archetype-site
2. 聚合总父工程名字
3. Maven选版本
4. 工程名字
5. 字符编码 - Setting -File encoding
6. 注解生效激活 - Setting - Annotation Processors
7. Java编译版本选8
8. File Type 过滤 -Setting -File Type

### 六、父工程pom文件

```xml
<?xml version="1.0" encoding="UTF-8"?>

<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.yooome.springcloud</groupId>
  <artifactId>yooomecloud</artifactId>
  <version>1.0-SNAPSHOT</version>
  <packaging>pom</packaging>

  <name>Maven</name>
  <url>http://maven.apache.org/</url>
  <inceptionYear>2001</inceptionYear>


  <distributionManagement>
    <site>
      <id>website</id>
      <url>scp://webhost.company.com/www/website</url>
    </site>
  </distributionManagement>

  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <maven.compiler.source>1.8</maven.compiler.source>
    <maven.compiler.target>1.8</maven.compiler.target>
    <junit.version>4.12</junit.version>
    <log4j.version>1.2.17</log4j.version>
    <lombok.version>1.16.18</lombok.version>
    <mysql.version>5.1.47</mysql.version>
    <druid.version>1.1.16</druid.version>
    <mybatis.spring.boot.version>1.3.0</mybatis.spring.boot.version>
  </properties>

  <!-- 子模块继承之后，提供作用：锁定版本+子modlue不用写groupId和version -->
  <dependencyManagement>
    <dependencies>
      <!--spring boot 2.2.2-->
      <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-dependencies</artifactId>
        <version>2.2.2.RELEASE</version>
        <type>pom</type>
        <scope>import</scope>
      </dependency>
      <!--spring cloud Hoxton.SR1-->
      <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-dependencies</artifactId>
        <version>Hoxton.SR1</version>
        <type>pom</type>
        <scope>import</scope>
      </dependency>
      <!--spring cloud alibaba 2.1.0.RELEASE-->
      <dependency>
        <groupId>com.alibaba.cloud</groupId>
        <artifactId>spring-cloud-alibaba-dependencies</artifactId>
        <version>2.1.0.RELEASE</version>
        <type>pom</type>
        <scope>import</scope>
      </dependency>
      <dependency>
        <groupId>mysql</groupId>
        <artifactId>mysql-connector-java</artifactId>
        <version>${mysql.version}</version>
      </dependency>
      <dependency>
        <groupId>com.alibaba</groupId>
        <artifactId>druid</artifactId>
        <version>${druid.version}</version>
      </dependency>
      <dependency>
        <groupId>org.mybatis.spring.boot</groupId>
        <artifactId>mybatis-spring-boot-starter</artifactId>
        <version>${mybatis.spring.boot.version}</version>
      </dependency>
      <dependency>
        <groupId>junit</groupId>
        <artifactId>junit</artifactId>
        <version>${junit.version}</version>
      </dependency>
      <dependency>
        <groupId>log4j</groupId>
        <artifactId>log4j</artifactId>
        <version>${log4j.version}</version>
      </dependency>
      <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <version>${lombok.version}</version>
        <optional>true</optional>
      </dependency>
    </dependencies>
  </dependencyManagement>
  
  <build>
    <plugins>
      <plugin>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-maven-plugin</artifactId>
        <configuration>
          <fork>true</fork>
          <addResources>true</addResources>
        </configuration>
      </plugin>
    </plugins>
  </build>
</project>

```

### 七、复习DependencyManagement和Dependencies

#### 7.1 dependencyManagement

 Maven使用dependencyManagement元素提供了一种管理依赖版本号的方式。

通常会在一个组织或者项目的最顶层的父POM中看到dependencyManagement元素。

使用pom.xml中的dependencyManagement元素能让所有在子项目中引入依赖而不用显示的列出版本量。

Maven会沿着父子层次向上走，直到找到一个拥有 dependencyManagement元素的项目，然后它就会使用这个 dependencyManagement元素中指定的版本号。

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
        <groupId>mysq1</groupId>
        <artifactId>mysql-connector-java</artifactId>
        <version>5.1.2</version>
        </dependency>
    <dependencies>
</dependencyManagement>
```

然后在子项目里就可以添加`mysql-connector`时可以不指定版本号，例如：

```xml
<dependencies>
    <dependency>
    <groupId>mysq1</groupId>
    <artifactId>mysql-connector-java</artifactId>
    </dependency>
</dependencies>
```

这样做的**好处**就是：如果有多个子项目都引用同一样依赖，则可以避免在每个使用的子项目里都声明一个版本号，这样当想升级或切换到另一个版本时，只需要在顶层父容器里更新，而不需要一个一个子项目的修改；另外如果某个子项目需要另外的一个版本，只需要声明version就可。

- dependencyManagement里只是声明依赖，并不实现引入，因此子项目需要显示的声明需要用的依赖。

- 如果不在子项目中声明依赖，是不会从父项目中继承下来的；只有在子项目中写了该依赖项,并且没有指定具体版本，才会从父项目

  继承该项，并且version和scope都读取自父pom。

- 如果子项目中指定了版本号，那么会使用子项目中指定的jar版本。

IDEA右侧旁边的Maven插件有 `Toggle 'Skip Test' Mode`按钮，这样maven可以跳过单元测试

父工程创建完成执行 `mvn:install` 将父工程发布到仓库方便子工程继承。

### 八、支付模块构建（上）

#### 8.1 创建微服务模块套路：

- 建Module
- 该POM
- 写YML
- 主启动
- 业务类

创建yooomepaymnet8001微服务提供者支付Module模块：

#### **8.2、创建名为yooomepaymnet8001微服务提供者支付Module模块**

#### **8.3、改POM**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>yooomepaymnet8001</artifactId>

    <dependencies>
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.10</version>
        </dependency>
        <!--mysql-connector-java-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
        </dependency>
        <!--jdbc-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-jdbc</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>1.0-SNAPSHOT</version>
            <scope>compile</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>
</project>
```

#### **8.4、写YML**

```yml
server:
  port: 8001

spring:
  application:
    name: cloud-payment-service
  datasource:
    type: com.alibaba.druid.pool.DruidDataSource
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:3306/springcloud?useUnicode=true&characterEncoding=utf-8&useSSL=false
    username: root
    password: root

mybatis:
  mapper-locations: classpath:mapper/*.xml
  type-aliases-package: com.yooome.springcloud.entities
```

#### 8.5 主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class PaymentApplication {
    public static void main(String[] args) {
        SpringApplication.run(PaymentApplication.class, args);
    }
}
```

#### 8.6 业务类

**SQL：**

```sql
CREATE TABLE `payment`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `serial` varchar(200) DEFAULT '',
	PRIMARY KEY (id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4
```

**Entities：**

实体类Payment

```java
package com.yooome.springcloud.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CommonResult<T> {
    private Integer code;
    private String message;
    private T data;

    public CommonResult(Integer code, String message) {
        this(code, message, null);
    }

}
```

**JSON封装体CommonResult：**

```java
package com.yooome.springcloud.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CommonResult<T> {
    private Integer code;
    private String message;
    private T data;

    public CommonResult(Integer code, String message) {
        this(code, message, null);
    }
}
```

**DAO：**

**接口PaymentDao**

```java
package com.yooome.springcloud.dao;

import com.yooome.springcloud.entities.Payment;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
@Mapper
public interface PaymentDao {

    public int create(Payment payment);

    public Payment getPaymentById(@Param("id") Long id);
}

```

**MyBatis映射文件PaymentMapper.xml，路径：resources/mapper/PaymentMapper.xml**

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >

<mapper namespace="com.yooome.springcloud.dao.PaymentDao">
    <insert id="create" parameterType="Payment" useGeneratedKeys="true" keyProperty="id">
        insert into payment(serial) values (#{serial});
    </insert>
    <resultMap id="BaseResultMap" type="com.yooome.springcloud.entities.Payment">
        <id column="id" property="id" jdbcType="BIGINT"/>
        <id column="serial" property="serial" jdbcType="VARCHAR"/>
    </resultMap>
    <select id="getPaymentById" parameterType="Long" resultMap="BaseResultMap">
        select * from payment where id=#{id};
    </select>
</mapper>
```

**Service**：

**接口PaymentService**

```java
package com.yooome.springcloud.server;

import com.yooome.springcloud.entities.Payment;
import org.apache.ibatis.annotations.Param;

public interface PaymentService {
    public int create(Payment payment);

    public Payment getPaymentById(@Param("id") Long id);
}
```

**实现类PaymentServiceImpl：**

```java
package com.yooome.springcloud.server.impl;

import com.yooome.springcloud.dao.PaymentDao;
import com.yooome.springcloud.entities.Payment;
import com.yooome.springcloud.server.PaymentService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;

@Service
public class PaymentServiceImpl implements PaymentService {
    @Resource
    private PaymentDao paymentDao;
    @Override
    public int create(Payment payment) {
        return paymentDao.create(payment);
    }

    @Override
    public Payment getPaymentById(Long id) {
        return paymentDao.getPaymentById(id);
    }
}
```

**Controller**：

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.entities.CommonResult;
import com.yooome.springcloud.entities.Payment;
import com.yooome.springcloud.server.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

@RestController
@Slf4j
public class PaymentController {
    @Resource
    private PaymentService paymentService;

    @PostMapping(value = "/payment/create")
    public CommonResult create(@RequestBody Payment payment) {
        int result = paymentService.create(payment);
        log.info("payment insert result: " + result);
        if (result > 0) {
            return new CommonResult(200,"payment insert success", result);
        }else {
            return new CommonResult(444,"payment insert fail", null);
        }
    }
    @GetMapping(value = "/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id) {
        Payment payment = paymentService.getPaymentById(id);
        log.info("get payment by id result: " + payment);
        if (payment != null) {
            return new CommonResult(200,"get payment by id success", payment);
        }else {
            return new CommonResult(444,"get payment by id fail", null);
        }
    }
}
```

#### 8.7 测试

1. 浏览器 - http://localhost:8001/payment/get/1
2. Postman - http://localhost:8001/payment/create?serial=lun2

#### **8.8.小总结**

创建微服务模块套路：

1. 建Module
2. 改POM
3. 写YML
4. 主启动
5. 业务类

### 九、热部署Devtools

开发时使用，生产环境关闭

#### 9.1 Adding devtools to your project

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-devtools</artifactId>
    <scope>runtime</scope>
    <optional>true</optional>
</dependency>
```

#### 9.2、Adding plugin to your pom.xml

下段配置复制到聚合父工程的pom.xml文件

```xml
<build>
    <!--
	<finalName>你的工程名</finalName>（单一工程时添加）
    -->
    <plugins>
        <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
            <configuration>
                <fork>true</fork>
                <addResources>true</addResources>
            </configuration>
        </plugin>
    </plugins>
</build>
```

#### 9.3 Enabling automatic build

File -> Settings(New Project Settings->Settings for New Projects) ->Complier

下面项勾选

- Automatically show first error in editor
- Display notification on build completion
- Build project automatically
- Compile independent modules in parallel

#### 9.4 Update the value of

键入Ctrl + Shift + Alt + / ，打开Registry，勾选：

- compiler.automake.allow.when.app.running

- actionSystem.assertFocusAccessFromEdt

#### 9.5 重启IDEA

### 十、消费者订单模块

#### 10.1 建Module

创建名为 cloud-consumer-order80的maven工程。

10.2 改pom

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-consumer-order80</artifactId>

    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>
    <dependencies>
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>1.0-SNAPSHOT</version>
            <scope>compile</scope>
        </dependency>
    </dependencies>
</project>
```

#### 10.3 写YML

```yml
server:
  port: 80

spring:
  application:
    name: cloud-consumer-order80
```

#### 10.4 主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class OrderApplication {
    public static void main(String[] args) {
        SpringApplication.run(OrderApplication.class, args);
    }
}
```

#### 10.5 业务类

**Entities：**

```java
package com.yooome.springcloud.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Payment implements Serializable {
    private Long id;
    private String serial;
}

```

```java
package com.yooome.springcloud.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class CommonResult<T> {
    private Integer code;
    private String message;
    private T data;

    public CommonResult(Integer code, String message) {
        this(code, message, null);
    }
}
```

**控制层：**

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.entities.CommonResult;
import com.yooome.springcloud.entities.Payment;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import javax.annotation.Resource;

@RestController
public class OrderController {
    public static final String PAYMENT_URL = "http://localhost:8001";
    @Resource
    private RestTemplate restTemplate;

    @GetMapping("/consumer/payment/create")
    public CommonResult<Payment> create(Payment payment) {
        return restTemplate.postForObject(PAYMENT_URL + "/payment/create", payment, CommonResult.class);
    }

    @GetMapping("/consumer/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id) {
        return restTemplate.getForObject(PAYMENT_URL + "/payment/get/" + id, CommonResult.class);
    }
}
```

**配置类：**

```java
package com.yooome.springcloud.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;
@Configuration
public class ApplicationContextConfig {
    @Bean
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
    }
}

```

#### 10.6 测试

运行cloud-consumer-order80与cloud-provider-payment8001两工程

浏览器 - http://localhost/consumer/payment/get/1

#### 10.7 RestTemplate

RestTemplate提供了多种便捷访问远程Http服务的方法，是一种简单便捷的访问restful服务模板类，是Spring提供的用于访问Rest服务的客户端模板工具集

[官网地址](https://docs.spring.io/spring-framework/docs/5.2.2.RELEASE/javadoc-api/org/springframework/web/client/RestTemplate.html)

使用：

使用restTemplate访问restful接口非常的简单粗暴无脑。
(url, requestMap, ResponseBean.class)这三个参数分别代表。
REST请求地址、请求参数、HTTP响应转换被转换成的对象类型。

### 11、消费者订单模块（下）

#### 11.1 数据库中添加数据为null

浏览器 - http://localhost/consumer/payment/create?serial=lun3

虽然，返回成功，但是观测数据库中，并没有创建serial为lun3的行。

解决之道：在yooomepaymnet8001工程的PaymentController中添加`@RequestBody`注解。

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.entities.CommonResult;
import com.yooome.springcloud.entities.Payment;
import com.yooome.springcloud.server.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

@RestController
@Slf4j
public class PaymentController {
    @Resource
    private PaymentService paymentService;

    @PostMapping(value = "/payment/create")
    public CommonResult create(@RequestBody Payment payment) {
      ...
    }
    
}

```

#### 11.2 Run Dashboard窗口

通过修改idea的workspace.xml的方式来快速打开Run Dashboard窗口（这个用来显示哪些Spring Boot工程运行，停止等信息。我idea 2020.1版本在名为Services窗口就可以显示哪些Spring Boot工程运行，停止等信息出来，所以这仅作记录参考）。

开启Run DashBoard

打开工程路径下的.idea文件夹的workspace.xml

在 <component name="RunDashboard"> 中修改或添加以下代码：

```xml
<option name="configurationTypes">
	<set>
		<option value="SpringBootApplicationConfigurationType"/>
    </set>
</option>
```

由于idea版本差异，可能需要关闭重启。

### 十二、工程重构

观察cloud-consumer-order80与cloud-provider-payment8001两工程有重复代码（entities包下的实体）（坏味道），重构。

#### 12.1 新建 - cloud-api-commons

#### 12.2 POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-api-common</artifactId>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>cn.hutool</groupId>
            <artifactId>hutool-all</artifactId>
            <version>5.1.0</version>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

#### 12.3 entities

将cloud-consumer-order80与cloud-provider-payment8001两工程的公有entities包移至cloud-api-commons工程下。

#### 12.4 maven clean、install cloud-api-commons工程，以供给cloud-consumer-order80与yooomepaymnet8001两工程调用。

#### 12.5 订单80和支付8001分别改造

将cloud-consumer-order80与yooomepaymnet8001两工程的公有entities包移除
引入cloud-api-common依赖

```xml
<dependency>
    <groupId>com.lun.springcloud</groupId>
    <artifactId>cloud-api-commons</artifactId>
    <version>${project.version}</version>
</dependency>
```

#### 12.6 测试

### 十三、Eureka基础知识

#### 13.1 什么是服务治理

Spring Cloud 封装了Netffix公式开发的Eureka模块实现服务治理

在传统的RPC远程调用框架中，管理每个服务与服务之间依赖关系比较复杂，管理比较复杂，所需要使用服务治理，管理服务与服务之间依赖关系，可以实现服务调用，负载均衡，容错等，实现服务发现与注册。

#### 13.2 什么是服务注册与发现

Eureka采用了CS的设计架构，Eureka Server 作为服务注册功能的服务器，它是服务注册中心。而系统中的其他微服务，使用Eureka的客户端连接到Eureka Server 并维持心跳连接。这样系统的维护人员就可以通过Eureka Server 来监控系统中各个微服务是否正常运行。

在服务注注册与发现中，有一个注册中心。放服务启动的时候，会把当前自己服务器的信息比如服务地址通信地址等别名方式注册到注册中心上。另外一方面（消费者服务提供者），以该别名的方式注册中心上获取到实际的服务通信地址，然后再实现本地RPC调用RPC远程调佣框架设计思想:用于注册中心，因为使用注册中心管理每个服务与服务之间的一个依赖关系（服务治理概念）。在任何RPC远程框架中，都会有一个注册中心存放服务地址相关信息（接口地址）。

![img](images/3956561052b9dc3909f16f1ff26d01bb.png)

#### 13.3 Eureka包含两个组件：Eureka Server 和Eureka Client

1、 Eureka Server 提供服务注册服务

各个微服务节点通过配置情动后，会在Eureka Server 中进行注册，这样Eureka Server 中的福福屋注册表中将会存储所有可用服务节点的信息，服务节点的信息可以在界面中直观看到。

2、 EurekaClient 通过注册中心访问

它是一个Java客户端，用于简化Eureka Server 的交互，客户端同时也具备一个内置的，使用轮训（round-robin）负载算法的负载均衡器。在应用启动后，将会向Eureka Server 发送心跳（默认周期为30秒）。如果Eureka Server 在多个细腻条周期内没有接收到某个节点的心跳，Eureka Server 将会从服务注册表中把这个服务几点溢出（默认为90秒）。

### 十四、EurekaServer服务端安装

IDEA生成eurekaserver端服务注册中心，类似物业公司

#### 14.1 创建cloud-eureka-server7001的Maven工程

#### 14.2 修改pom.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-eureka-server7001</artifactId>
    <dependencies>
        <!--eureka-server-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-server</artifactId>
        </dependency>
        <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <!--boot web actuator-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--一般通用配置-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>
</project>
```

#### 14.3 添加application.yml

```yml
server:
  port: 7001

eureka:
  instance:
    hostname: localhost  #eureka服务端的实例名称
  client:
    register-with-eureka: false  # false表示注册中心不注册自己
    fetch-registry: false        # false表示自己端就是注册中心，我们的指责就是维护服务实例，并不需要检索服务
    service-url:
      # 设置与Eureka server交互的地址查询服务和注册服务都需要依赖这个地址。
      defaultZone: http://${eureka.instance.hostname}:${server.port}/eureka/
```

#### 14.4 主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer;

@SpringBootApplication
@EnableEurekaServer
public class EruekaServerApplication {
    public static void main(String[] args) {
        SpringApplication.run(EruekaServerApplication.class, args);
    }
}

```

### 十五、支付微服务8001入住EurekaServer

EurekaClient端cloud-provider-payment8001将注册进EurekaServer成为服务提供者provider，类似学校对外提供授课服务。

### 15.1 修改cloud-provider-payment8001

#### 15.2 改POM

添加spring-cloud-starter-netflix-eureka-client依赖

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
</dependency>
```

#### 15.3 写yml

```yml
eureka:
  client:
    #表示是否将自己注册进Eurekaserver默认为true。
    register-with-eureka: true
    #是否从EurekaServer抓取已有的注册信息，默认为true。单节点无所谓，集群必须设置为true才能配合ribbon使用负载均衡
    fetchRegistry: true
    service-url:
      defaultZone: http://localhost:7001/eureka
```

#### 15.4 主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableEurekaClient
public class PaymentApplication {
    public static void main(String[] args) {
        SpringApplication.run(PaymentApplication.class, args);
    }
}

```

#### 15.5 测试

启动yooomepaymnet8001和cloud-eureka-server7001工程。

浏览器输入 - http://localhost:7001/ 主页内的Instances currently registered with Eureka会显示cloud-provider-payment8001的配置文件application.yml设置的应用名cloud-payment-service

#### 15.6 自我保护机制

EMERGENCY! EUREKA MAY BE INCORRECTLY CLAIMING INSTANCES ARE UP WHEN THEY’RE NOT. RENEWALS ARELESSER THAN THRESHOLD AND HENCFT ARE NOT BEING EXPIRED JUST TO BE SAFE.

紧急情况！EUREKA可能错误地声称实例在没有启动的情况下启动了。续订小于阈值，因此实例不会为了安全而过期。

### 十六、订单微服务80入住进入EurekaServer

EurekaClient端cloud-consumer-order80将入住进EurekaServer成为微服务消费者consumer，类似来上课消费的同学

#### 16.1 cloud-consumer-order80

#### 16.2 POM

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
</dependency>
```

#### 16.3 YML

```yml
eureka:
  client:
    #表示是否将自己注册进Eurekaserver默认为true。
    register-with-eureka: true
    #是否从EurekaServer抓取已有的注册信息，默认为true。单节点无所谓，集群必须设置为true才能配合ribbon使用负载均衡
    fetchRegistry: true
    service-url:
      defaultZone: http://localhost:7001/eureka
```

#### 16.4 主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableEurekaClient
public class PaymentApplication {
    public static void main(String[] args) {
        SpringApplication.run(PaymentApplication.class, args);
    }
}
```

#### 16.5 测试

- 启动cloud-provider-payment8001，cloud-eureka-server7001和cloud-consumer-order80这三个工程。
- 浏览器输入 http://localhost:7001,在主页的 Instance currently registered with Eureka 将会看到yooomepaymnet8001、cloud-consumer-order80两个工程

#### 16.6 注意

application.yml 配置中层次缩进和空格，两者不能少，否则，会抛出异常 Failed  to  bind properties under 'eureka.client.servoice-url' to java.util.Map <java.lang.String,Java.lang.String>。

### 十七、Eureka集群原理说明

#### 17.1 Eureka集群原理说明

![img](images/14570c4b7c4dd8653be6211da2675e45.png)

服务注册：将服务信息注册进注册中心

服务发现：从注册中心上获取服务信息

实质：存key服务命取value闭用地址

- 县启动eureka主注册中心
- 启动服务提供者payment支付服务
- 支付服务启动会把自身信息（比服务地址L以别名方式注入进eureka）
- 消费者order服务在需要调用接口时，使用服务别名去注册中心获取实际的RPC远程调用地址
- 消费者调用地址后，底层实际是利用HttpClient技术实现远程调用
- 消费者实际的服务地址后缓存在本地jvm内存中，默认每间隔30秒更新一次服务调用地址

**问题**：微服务RPC远程调用最核心的是什么？

**高可用**：试想你的注册中心有一个only one，万一它出故障了，会导致整个服务环境不可用。

解决办法：发件Eureak注册中心集群，实现负载均衡+故障容错。

互相注册，互相守望。

### 十八、Eureka集群环境构建

创建cloud-eureka-server70002工程，过程参考 【十四、EurekaServer服务端安装】

![1](images/1.png)

- 找到本地计算机C:\Windows\System32\drivers\etc路径下的hosts文件，修改映射配置添加进hosts文件

```tex
127.0.0.1 eureka7001.com
127.0.0.1 eureka7002.com
```

- 修改cloud-eureka-server7001配置文件

```yml
server:
  port: 7001

eureka:
  instance:
    hostname: eureka7001.com #eureka服务端的实例名称
  client:
    register-with-eureka: false     #false表示不向注册中心注册自己。
    fetch-registry: false     #false表示自己端就是注册中心，我的职责就是维护服务实例，并不需要去检索服务
    service-url:
    #集群指向其它eureka
      defaultZone: http://eureka7002.com:7002/eureka/
    #单机就是7001自己
      #defaultZone: http://eureka7001.com:7001/eureka/
```

- 修改cloud-eureka-server7002配置文件

```yml
server:
  port: 7002

eureka:
  instance:
    hostname: eureka7002.com #eureka服务端的实例名称
  client:
    register-with-eureka: false     #false表示不向注册中心注册自己。
    fetch-registry: false     #false表示自己端就是注册中心，我的职责就是维护服务实例，并不需要去检索服务
    service-url:
    #集群指向其它eureka
      defaultZone: http://eureka7001.com:7001/eureka/
```

时间的时候，遇到异常情况

在开启cloud-eureka-server7002时，开启失败，说7002端口被占用，然后在cmd中输入`netstat -ano | find "7002"` ,查不到任何动词。

纳闷一阵，重启电脑，问题解决。

### 十九、订单支付两微服务注册进Eureka集群

- 将支付微服务8001微服务，订单服务80 微服务发布到上面2台Eureka集群配置中

将他们的配置文件的eureka.client.service-url.defaultZone进行修改

```yml
eureka:
  client:
    #表示是否将自己注册进Eurekaserver默认为true。
    register-with-eureka: true
    #是否从EurekaServer抓取已有的注册信息，默认为true。单节点无所谓，集群必须设置为true才能配合ribbon使用负载均衡
    fetchRegistry: true
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka, http://eureka7002.com:7002/eureka
```

- 测试01
  1. 先要启动EurekaServer，7001/7002服务
  2. 在启动启动服务提供者provider，8001
  3. 在启动消费者，80
  4. 浏览器输入 -http://localhost/consumer/payment/get/30

### 二十、支付微服务集群配置

#### 20.1 支付服务提供者8001集群环境构建

参考cloud-provicer-payment8001

1. 新建cloud-provider-payment8002
2. 改pom
3. 写YML -端口8002
4. 主启动
5. 业务类
6. 修改8001/8002的Controller，添加serverPort

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.entities.CommonResult;
import com.yooome.springcloud.entities.Payment;
import com.yooome.springcloud.server.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

@RestController
@Slf4j
public class PaymentController {
    @Resource
    private PaymentService paymentService;
    @Value("${server.port}")
    private String serverPort;
    @PostMapping(value = "/payment/create")
    public CommonResult create(@RequestBody Payment payment) {
        int result = paymentService.create(payment);
        log.info("payment insert result: " + result);
        if (result > 0) {
            return new CommonResult(200,"payment insert success, server port " + serverPort, result);
        }else {
            return new CommonResult(444,"payment insert fail", null);
        }
    }
}
```

#### 20.2 负载均衡

cloud-consumer-order80订单服务访问地址不能写死

```java
@Slf4j
@RestController
public class OrderController {

    //public static final String PAYMENT_URL = "http://localhost:8001";
    public static final String PAYMENT_URL = "http://CLOUD-PAYMENT-SERVICE";
    
    ...
}
```

使用`@LoadBalanced`注解赋予RestTemplate负载均衡的能力

```java
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class ApplicationContextConfig {

    @Bean
    @LoadBalanced//使用@LoadBalanced注解赋予RestTemplate负载均衡的能力
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
    }
}
```

ApplicationContextBean - 提前说一下Ribbon的负载均衡功能

#### 20.3 测试

先要启动EurekaServer， 7001/7002服务

再要启动服务提供者provider,8002、8001服务

浏览器输入 -http://localhost/consumer/payment/get/31

结果：负载均衡效果达到，8001/8002端口交替出现

Ribbon和Eureka整合后Consumer可以直接调用服务而不用关心地址和端口号，且该服务还有负载功能

互相注册，互相守望

![img](images/94c4c3eca8c2f9eb7497fe643b9b0622.png)



### 二十一、actuator微服务信息完善

主机名称：服务名称修改（也就是将IP地址，换成可读性高的名字）

修改cloud-provider-payment8001，cloud-provider-payment8002

修改部分 -YML -eureka.instance.instance-id

```yml
eureka:
  ...
  instance:
    instance-id: payment8001 #添加此处

```

```yml
eureka:
  ...
  instance:
    instance-id: payment8002 #添加此处

```

修改之后：

eureka主页将显示payment8001， payment8002代替原来显示的IP地址。

访问信息有IP信息提示，（也就是鼠标指针移动到payment8001，payment8002名下，会有IP地址提示）

修改部分 -YML - eureka.install.prefer-ip-address

```yml
eureka:
  ...
  instance:
    instance-id: payment8001 
    prefer-ip-address: true #添加此处

```

```yml
eureka:
  ...
  instance:
    instance-id: payment8002
    prefer-ip-address: true #添加此处

```

### 二十二、服务发现Discovery

对于注册进eureka里面的微服务，可以通过服务发现来获得该服务的信息

- 修改cloud-provider-payment8001的Controller

```java
@RestController
@Slf4j
public class PaymentController{
	...
    
    @Resource
    private DiscoveryClient discoveryClient;

    ...

    @GetMapping(value = "/payment/discovery")
    public Object discovery()
    {
        List<String> services = discoveryClient.getServices();
        for (String element : services) {
            log.info("*****element: "+element);
        }

        List<ServiceInstance> instances = discoveryClient.getInstances("CLOUD-PAYMENT-SERVICE");
        for (ServiceInstance instance : instances) {
            log.info(instance.getServiceId()+"\t"+instance.getHost()+"\t"+instance.getPort()+"\t"+instance.getUri());
        }

        return this.discoveryClient;
    }
}

```

- 8001主启动类

```java
@SpringBootApplication
@EnableEurekaClient
@EnableDiscoveryClient//添加该注解
public class PaymentMain001 {

    public static void main(String[] args) {
        SpringApplication.run(PaymentMain001.class, args);
    }
}
```

- 自测

先要启动EurekaServer

在启动8001主启动类，需要稍等一会儿

浏览器输入http://localhost:8001/payment/discovery

浏览器输出：

```json
{"services":["cloud-payment-service"],"order":0}
```

后台输出：

```java
*****element: cloud-payment-service
CLOUD-PAYMENT-SERVICE	192.168.199.218	8001	http://192.168.199.218:8001
```

### 二十三、Eureka自我保护理论知识

#### 23.1 **概述：**

保护模式主要用于一组客户端和Eureka Server之间存在网络分区场景下的保护。一旦进入保护模式，Eureka Server将会尝试保护其服务注册表中的信息，不再删除服务注册表中的数据，也就是不会注销任何微服务。

如果在Eureka Server的首页看到一下这段提示，则说明Eureka进入了保护模式：

EMERGENCY! EUREKA MAY BE INCORRECTLY CLAIMING INSTANCES ARE UP WHEN THEY’RE NOT. RENEWALS ARE LESSER THANTHRESHOLD AND HENCE THE INSTANCES ARE NOT BEING EXPIRED JUSTTO BE SAFE

#### 23.2 导致原因

一句话：某事某刻一个微服务不可用了，Eureka不会立刻清理，依旧会对该微服务的信心进行保存。

属于CAP里面的AP分支。

#### 23.3 为什么会产生Eureka自我保护机制

为了EurekaClient可以正常运行，防止与EurekaServer网络不通情况下，EurekaServer不会立刻将EurekaClient服务剔除

#### 23.4 什么是自我保护模式

默认情况下，如果EurekaServer在一定时间内没有接收到某个微服务实例的心跳，EurekaServer将会注销该实例（默认90秒）。但是当网络分区故障发生（延时，卡顿，拥挤）时。微服务与EurekaServer之间无法正常通信贸易商行为可能变的非常危险了---因为微服务本省其实是健康的，此时本不应该注销这个微服务。Eureka通过"自我保护模式"来解决这个问题 ---当EurekaServer节点在短时间内丢失过多客户端时（可能发生了网络分区故障），那么这个节点就会进入自我保护模式。

![img](images/264b66e8099a3761beaea2ba44b8fc5e.png)

自我保护模式机制：默认情况下EurekaClient定时向EurekaServer端发送心跳包

如果Eureka在server端在一定时间内（默认90秒）没有收到EurekaClient发送心跳包，变回直接从服务注册列表中剔除该服务，但是在短时间（90秒中）内丢失了大量的服务实例心跳，这时候EurekaServer会开启自我保护机制，不会剔除该微服务（该现象可能出现在如果网络不同但是EurekaClient为出现宕机，此时如果换做别的注册中心如果一定时间内没有收到心跳将会剔除该服务，这样就出现了严重事务，因为客户端还能正常发送心跳，只是网络延迟问题，二保护保护机制是为了解决此问题而产生的）。

#### 23.5 在自我保护模式中，EurekaServer 会保护服务注册表中的信息，不再注销任何服务实例。

它的设计哲学就是宁可保留错误的服务注册信息，也不盲目注销任何可能将康的服务实例。一句话：好死不如赖活着。

综上，自我保护模式是一种应对网络异常的安全保护措施。它的架构这些是宁可保留所有的微服务（健康的微服务和不健康的微服务都会保留）也不盲目注销任何健康的微服务。使用自我保护模式，可以让Eureka集群更加的健壮，稳定。

### 二十四、怎么禁止自我保护

- 在EurekaServer端7001出设置关闭自我保护机制

出厂默认，自我保护机制是开启的

使用`eureka.server.enable-self-preservation = false` 可以禁用自我保护模式

```yml
eureka:
  ...
  server:
    #关闭自我保护机制，保证不可用服务被及时踢除
    enable-self-preservation: false
    eviction-interval-timer-in-ms: 2000

```

关闭效果：

spring-eureka主页会显示出一句：

**THE SELF PRESERVATION MODE IS TURNED OFF. THIS MAY NOT PROTECT INSTANCE EXPIRY IN CASE OF NETWORK/OTHER PROBLEMS.**

- 生产者客户端eureakeClient端8001

默认：

```properties
eureka.instance.lease-renewal-interval-in-seconds=30
eureka.instance.lease-expiration-duration-in-seconds=90
```

```yml
eureka:
  ...
  instance:
    instance-id: payment8001
    prefer-ip-address: true
    #心跳检测与续约时间
    #开发时没置小些，保证服务关闭后注册中心能即使剔除服务
    #Eureka客户端向服务端发送心跳的时间间隔，单位为秒(默认是30秒)
    lease-renewal-interval-in-seconds: 1
    #Eureka服务端在收到最后一次心跳后等待时间上限，单位为秒(默认是90秒)，超时将剔除服务
    lease-expiration-duration-in-seconds: 2

```

- 测试
  - 7001和8001都配置完成
  - 先启动7001再启动8001

结果：先关闭8001，马上被删除了

### 二十五、Eureka停更说明

https://github.com/Netflix/eureka/wiki

> Eureka 2.0 (Discontinued)
>
> The existing open source work on eureka 2.0 is discontinued. The code base and artifacts that were released as part of the existing repository of work on the 2.x branch is considered use at your own risk.
>
> Eureka 1.x is a core part of Netflix’s service discovery system and is still an active project.

我们用ZooKeeper代替Eureka功能。

### 二十六、支付服务注册进zookeeper

#### 26.1 注册中心Zookeeper

zookeeper是一个分布式协调工具，可以实现注册中心功能

关闭Linux服务器防火墙后，启动zookeeper服务器

用到的Linux命令行：

- systemctl stop firewalld 关闭防火墙
- systemctl status firewalld 查看防火墙的状态
- ipconfig 查看IP地址
- ping 查验结果

zookeeper 服务器取代Eureka服务器，zk作为服务注册中心

视频中的虚拟机CentOS开启Zookeeper，打算在本机启动Zookeeper具体操作参考Zookeeper学习笔记

#### 26.2 服务提供者

1. 新建名为cloud-provider-payment8004的Maven工程
2. POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>LearnCloud</artifactId>
        <groupId>com.lun.springcloud</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-provider-payment8004</artifactId>
    <dependencies>
        <!-- SpringBoot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.lun.springcloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
        <!-- SpringBoot整合zookeeper客户端 -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-zookeeper-discovery</artifactId>
            <!--先排除自带的zookeeper3.5.3 防止与3.4.9起冲突-->
            <exclusions>
                <exclusion>
                    <groupId>org.apache.zookeeper</groupId>
                    <artifactId>zookeeper</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <!--添加zookeeper3.4.9版本-->
        <dependency>
            <groupId>org.apache.zookeeper</groupId>
            <artifactId>zookeeper</artifactId>
            <version>3.4.9</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
</project>
```

3. YML

```yml
#8004表示注册到zookeeper服务器的支付服务提供者端口号
server:
  port: 8004

#服务别名----注册zookeeper到注册中心名称
spring:
  application:
    name: cloud-provider-payment
  cloud:
    zookeeper:
      connect-string: 127.0.0.1:2181 # 192.168.111.144:2181 #

```

4. 主启动类

```java
package com.yooome.springcloud.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
@Slf4j
public class PaymentController {
    @Value("${server.port}")
    private String serverPort;
    @RequestMapping("/payment/zk")
    public String paymentZookeeper() {
        return "spring cloud with zookeeper: " + serverPort + "\t" + UUID.randomUUID().toString();
    }
}
```

5. Controller

```java
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
@Slf4j
public class PaymentController
{
    @Value("${server.port}")
    private String serverPort;

    @RequestMapping(value = "/payment/zk")
    public String paymentzk()
    {
        return "springcloud with zookeeper: "+serverPort+"\t"+ UUID.randomUUID().toString();
    }
}
```

6. 启动8004注册进zookeeper（要先启动zookeeper的server）

- 验证测试： 浏览器 - http://locahost:8004/payment/zk

- 验证测试2：接着用zookeeper客户端操作

```java
[zk: localhost:2181(CONNECTED) 3] ls /services
[cloud-provider-payment]
[zk: localhost:2181(CONNECTED) 4] ls /services/cloud-provider-payment 
[ef010ed8-4b29-4a6f-aaad-f65d463c733b]
[zk: localhost:2181(CONNECTED) 5] get /services/cloud-provider-payment/ef010ed8-4b29-4a6f-aaad-f65d463c733b
{"name":"cloud-provider-payment","id":"ef010ed8-4b29-4a6f-aaad-f65d463c733b","address":"192.168.3.14","port":8004,"sslPort":null,"payload":{"@class":"org.springframework.cloud.zookeeper.discovery.ZookeeperInstance","id":"application-1","name":"cloud-provider-payment","metadata":{}},"registrationTimeUTC":1645081869763,"serviceType":"DYNAMIC","uriSpec":{"parts":[{"value":"scheme","variable":true},{"value":"://","variable":false},{"value":"address","variable":true},{"value":":","variable":false},{"value":"port","variable":true}]}}
[zk: localhost:2181(CONNECTED) 6] get /zookeeper 
```

json格式化 `get /services/cloud-provider-payment/ef010ed8-4b29-4a6f-aaad-f65d463c733b`

```json
{
	"name": "cloud-provider-payment",
	"id": "ef010ed8-4b29-4a6f-aaad-f65d463c733b",
	"address": "192.168.3.14",
	"port": 8004,
	"sslPort": null,
	"payload": {
		"@class": "org.springframework.cloud.zookeeper.discovery.ZookeeperInstance",
		"id": "application-1",
		"name": "cloud-provider-payment",
		"metadata": {}
	},
	"registrationTimeUTC": 1645081869763,
	"serviceType": "DYNAMIC",
	"uriSpec": {
		"parts": [{
			"value": "scheme",
			"variable": true
		}, {
			"value": "://",
			"variable": false
		}, {
			"value": "address",
			"variable": true
		}, {
			"value": ":",
			"variable": false
		}, {
			"value": "port",
			"variable": true
		}]
	}
}
```

### 二十七、临时节点还是持久节点

zookeeper的服务节点是临时节点，没有Eureka那么含情脉脉。

### 二十八、订单服务注册进zookeeper

1. 新建cloud-consumerzk-order80
2. POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-consumerzk-order80</artifactId>
    <dependencies>
        <!-- SpringBoot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <!-- SpringBoot整合zookeeper客户端 -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-zookeeper-discovery</artifactId>
            <!--先排除自带的zookeeper-->
            <exclusions>
                <exclusion>
                    <groupId>org.apache.zookeeper</groupId>
                    <artifactId>zookeeper</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <!--添加zookeeper3.4.9版本-->
        <dependency>
            <groupId>org.apache.zookeeper</groupId>
            <artifactId>zookeeper</artifactId>
            <version>3.4.9</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3. YML

```yml
server:
  port: 80
# 服务别名 ----- 注册zookeeper到注册中心名称
spring:
  cloud:
    zookeeper:
      connect-string: 127.0.0.1:2181 # 192.168.111.144:2181 #
  application:
    name: cloud-consumer-order
```

4. 主启动

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class ConsumerZookeeperApplication {
    public static void main(String[] args) {
        SpringApplication.run(ConsumerZookeeperApplication.class, args);
    }
}

```

5. 业务类

```java
package com.yooome.springcloud.config;

import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class ApplicationContextConfig {
    @Bean
    @LoadBalanced
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }
}
```

```java
package com.yooome.springcloud.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import javax.annotation.Resource;

@RestController
public class OrderZKController {
    private static final  String INVOKE_URL="http://cloud-provider-payment";
    @Resource
    private RestTemplate restTemplate;
    @GetMapping("/consumer/payment/zk")
    public String paymentInfo() {
        String forObject = restTemplate.getForObject(INVOKE_URL + "/payment/zk", String.class);
        return forObject;
    }

}

```

6. 测试验证

运行zookeeper服务端，cloud-consumerzk-order80，cloud-provider-payment8004。

打开zookeeper客户端

```json
[zk: localhost:2181(CONNECTED) 8] ls
ls [-s] [-w] [-R] path
[zk: localhost:2181(CONNECTED) 9] ls /services 
[cloud-consumer-order, cloud-provider-payment]
[zk: localhost:2181(CONNECTED) 10] 
```

7. 访问测试地址 -http://localhost/consumer/payment/zk

### 二十九、Consul简介

[Consul官网](https://www.consul.io/)

[Consule下载地址](https://www.consul.io/downloads)

#### 29.1 What is Consul

```
Consul is a service mesh solution providing a full featured control plane with service discovery, configuration, and segmentation functionality. Each of these features can be used individually as needed, or they can be used together to build a full service mesh. Consul requires a data plane and supports both a proxy and native integration model. Consul ships with a simple built-in proxy so that everything works out of the box, but also supports 3rd party proxy integrations such as Envoy. link

Consul是一个服务网格解决方案，它提供了一个功能齐全的控制平面，具有服务发现、配置和分段功能。这些特性中的每一个都可以根据需要单独使用，也可以一起用于构建全服务网格。Consul需要一个数据平面，并支持代理和本机集成模型。Consul船与一个简单的内置代理，使一切工作的开箱即用，但也支持第三方代理集成，如Envoy。
```

Consul是一套开源的分布式服务发现和配置管理系统，由HashiCorp公司用Go语言开发。

提供了微服务系统中的服务治理，配置中心，控制总线等功能。这些功能中的每一个都可以根据需要单据 使用，也可以一起使用构建全方位的服务网络，总之Consul提供了一种完成的服务网络解决方案。

它具有很多优点。包括：基于raft协议，比较简介；支持健康检查。同时支持HTTP和DNS协议支持跨数据中心的WAN集群提供图形界面跨平台，支持Liunx、Mac、Windows。

> The key features of Consul are:
>
> Service Discovery: Clients of Consul can register a service, such as api or mysql, and other clients can use Consul to discover providers of a given service. Using either DNS or HTTP, applications can easily find the services they depend upon.
> Health Checking: Consul clients can provide any number of health checks, either associated with a given service (“is the webserver returning 200 OK”), or with the local node (“is memory utilization below 90%”). This information can be used by an operator to monitor cluster health, and it is used by the service discovery components to route traffic away from unhealthy hosts.
> KV Store: Applications can make use of Consul’s hierarchical key/value store for any number of purposes, including dynamic configuration, feature flagging, coordination, leader election, and more. The simple HTTP API makes it easy to use.
> Secure Service Communication: Consul can generate and distribute TLS certificates for services to establish mutual TLS connections. Intentions can be used to define which services are allowed to communicate. Service segmentation can be easily managed with intentions that can be changed in real time instead of using complex network topologies and static firewall rules.
> Multi Datacenter: Consul supports multiple datacenters out of the box. This means users of Consul do not have to worry about building additional layers of abstraction to grow to multiple regions.
> 地址：https://www.consul.io/docs/intro#what-is-consul

能干什么？

- 服务发现-提供HTTP和DNS两种发现方式。
- 健康检测-试吃多种方式，HTTP、TCP、Docker、Shell脚本定制化
- KV存储-Key、Value的存储方式
- 多数据中心-Consul支持多数据中心
- 可视化Web界面

### 三十、安装并运行Consul

[官网安装说明](https://learn.hashicorp.com/tutorials/consul/get-started-install)

windows版本解压缩后，的consul.exe,打开cmd

- 查看本本consul -v：

```bash
yooome@192 bin % consul --version
Consul v1.11.3
Revision e319d7ed
Protocol 2 spoken by default, understands 2 to 3 (agent will automatically use protocol >2 when speaking to compatible agents)
```

- 开发模式启动 consul agent -dev :

#### 30.1 Mac consul 安装与启动

1. 下载地址：https://www.consul.io/downloads.html
2. 将解压后的文件consul  拷贝到/usr/local/bin（不要放在中文目录下）

```bash
cp consul /usr/local/bin
```

3. 测试

打开bin文件，执行consul，查看consul命令，如下即表示成功

![img](images/1110742-20210408234658471-1963197479.png)

4. 启动consul

Consul 在启动后分为两种模式：Server 模式； Client 模式；

![img](images/1110742-20210408234758983-216171787.png)

此时在浏览器访问http://localhost:8500，侧可以看到如下页面，说明启动成功；

![img](images/1110742-20210408235008533-51226666.png)

5. 停止consul

可以使用 Ctrl -c (终端信号)正常停止代理。终端代理之后，您应该看到它离开集群并关闭。

### 三十一、服务提供者注册进Consul

1、新建Module支付服务provider8006

2、改POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>
    <artifactId>yooomepayment8006</artifactId>
    <dependencies>
        <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <!--SpringCloud consul-server -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-consul-discovery</artifactId>
        </dependency>
        <!-- SpringBoot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--日常通用jar包配置-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>cn.hutool</groupId>
            <artifactId>hutool-all</artifactId>
            <version>RELEASE</version>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>cn.hutool</groupId>
            <artifactId>hutool-all</artifactId>
            <version>RELEASE</version>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3、YML

```yml
server:
  port: 8006
spring:
  application:
    name: consul-provider-payment
  cloud:
    # consul 注册中心地址
    consul:
      host: localhost
      port: 8500
      discovery:
        #hostname 127.0.0.1
        service-name: ${spring.application.name}

```

4、主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient
public class PaymentApplication8006 {
    public static void main(String[] args) {
        SpringApplication.run(PaymentApplication8006.class, args);
    }
}

```

5、业务类

```java
package com.yooome.springcloud.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class PaymentController {
    @Value("${server.port}")
    private String serverPort;

    @RequestMapping("/payment/consul")
    public String paymentConsul() {
        return "spring cloud consul: " + serverPort + "\t" + UUID.randomUUID().toString();
    }

}

```

6、验证测试

- http://localhost:8006/payment/consul
- http://localhost:8500 - 会显示provider8006

### 三十二、服务消费者注册进Consul

1、新建Module消费服务order80  cloud-consumerconsul-order80

2、POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-consumerconsul-order80</artifactId>
    <dependencies>
        <!--SpringCloud consul-server -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-consul-discovery</artifactId>
        </dependency>
        <!-- SpringBoot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--日常通用jar包配置-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3、YML

```yml
# consul服务端口号
server:
  port: 80

spring:
  application:
    name: clould-consumer-order
  cloud:
    consul:
      port: 8500
      host: localhost
      discovery:
        service-name: ${spring.application.name}

```

4、主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

@SpringBootApplication
@EnableDiscoveryClient //该注解用于向使用consul或者zookeeper作为注册中心时注册服务
public class OrderConsulApplication80 {
    public static void main(String[] args) {
        SpringApplication.run(OrderConsulApplication80.class,args);
    }
}

```

5.配置Bean

```java
package com.yooome.springcloud.config;

import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;


@Configuration
public class ApplicationContextConfig {

    @Bean
    @LoadBalanced
    public RestTemplate getRestTemplate() {
        return new RestTemplate();
    }
}

```

6、controller

```java
package com.yooome.springcloud.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import javax.annotation.Resource;

@RestController
public class OrderConsulController {
    private static final String INVOKE_URL = "http://consul-provider-payment";
    @Resource
    private RestTemplate restTemplate;

    @RequestMapping("/consumer/payment/consul")
    public String paymentInfo() {
        return restTemplate.getForObject(INVOKE_URL + "/payment/consul", String.class);
    }
}

```

7、验证测试

运行consul，cloud-providerconsul-payment8006，cloud-consumerconsul-order80

http://localhost:8500/ 主页会显示出consul，cloud-providerconsul-payment8006，cloud-consumerconsul-order80三服务。

8、访问测试地址 - http://localhost/consumer/payment/consul

### 三十三、三个注册中心异同点

|  组件名   | 语言CAP | 服务健康检查 | 对外暴露接口 | SpringCloud集成 |
| :-------: | :-----: | :----------: | :----------: | :-------------: |
|  Eureka   |  Java   |      AP      |    可支配    |      HTTP       |
|  Consul   |   Go    |      CP      |     支持     |    HTTP/DNS     |
| Zookeeper |  Java   |      CP      |  支持客户端  |     已集成      |

CAP：

- C：Consistency（强一致性）
- A：Availability（可用性）
- P：Partition tolerance（分区容错性）

![img](images/b41e0791c9652955dd3a2bc9d2d60983.png)

最多只能同时较好的满足两个

CAP理论的核心是：一个分布式系统不可能同时很好的满足一致性，可用性和分区容错性这三个需求。

因此，根据CAP原理将NoSql数据库分成了满足CA原则，满足CP原则和满足AP原则的三大类：

- CA - 单点集群，满足一致性，可用性的系统，通常在可扩展性上不太强大。
- CP - 满足一致性，分区容忍性的系统，通常性能不是特别高。
- AP - 满足可用性，分区容忍性的系统，通常可能堆一致性要求低一些。

AP架构（Eureka）

当网络出现后，为了保证可用性，系统B可以返回旧值，保证系统的可用性。

结论：违背了一致性C的要求，只满足可用性和分区容错，即AP

![img](images/2d07748539300b9c466eb1d9bac5cd1b.png)



CP架构（Zookeeper/Consul）

当网络分区出现后，为了保证一致性，就必须拒接请求，否则无法保证一致性。

结论：违背了可用性A的要求，只要满足一致性和分区容错，即CP。

![img](images/c6f2926a97420015fcebc89b094c5598.png)

CP和AP对立统一的矛盾关系。

### 三十四、Ribbon入门介绍

SpringCloud Ribbon 是基于Netflix Ribbon 实现的一套客户端负载均衡的工具。

简单的说，Ribbon是Netfix发布的开源项目，主要功能是提供客户端的软件负载均衡算法和服务调用。Ribbon客户端组件提供一系类晚上的配置项如连接超时，重试等。

简单的说，就是在配置文件中出现Load Balance（简称LB）后面所有的机器，Ribbon会自动的帮助你基于某种规则（如简单轮训，随机连接等）去连接这些机器。我们很容易使用Ribbon实现自定义的负载均衡算法。

[Github - Ribbon](https://github.com/Netflix/ribbon/wiki/Getting-Started)

Ribbon目前也进入维护模式。

Ribbon未来可能被Spring Cloud LoadBalacer替代。

LB负载均衡(Load Balance)是什么

简单的说就是将用户的请求平摊的分配到多个服务上，从而达到系统的HA (高可用)。

常见的负载均衡有软件Nginx，LVS，硬件F5等。

Ribbon本地负载均衡客户端VS Nginx服务端负载均衡区别

Nginx是服务器负载均衡，客户端所有请求都会交给nginx，然后由nginx实现转发请求。即负载均衡是由服务端实现的。
Ribbon本地负载均衡，在调用微服务接口时候，会在注册中心上获取注册信息服务列表之后缓存到JVM本地，从而在本地实现RPC远程服务调用技术。

集中式LB

即在服务的消费方和提供方之间使用独立的LB设施(可以是硬件，如F5, 也可以是软件，如nginx)，由该设施负责把访问请求通过某种策略转发至服务的提供方;

进程内LB

将LB逻辑集成到消费方，消费方从服务注册中心获知有哪些地址可用，然后自己再从这些地址中选择出一个合适的服务器。

Ribbon就属于进程内LB，它只是一个类库，集成于消费方进程，消费方通过它来获取到服务提供方的地址。

一句话

负载均衡 + RestTemplate调用

### 三十五、Ribbon的负载均衡和Rest调用

#### 35.1 架构说明

总结：Ribbon其实就是一个软负载均衡的客户端组件，它乐意和其他所需请求的客户端结合使用，和Eureka结合只是其中的一个实例。

![img](images/145b915e56a85383b3ad40f0bb2256e0.png)

Ribbon在工作时分成两步：

- 第一步先选择EurekaServer，它优先选择在同一个区域负载较少的server。
- 第二步在根据用户指定的策略，在从server渠道的服务注册列表中选择一个地址、
- 其中Ribbon提供了多种策略：比如轮询，随机和根据响应时间加权。

#### 35.2 POM

先前工程项目没有引入spring-cloud-starter-ribbon也可以使用ribbon。

```xml
<dependency>
    <groupld>org.springframework.cloud</groupld>
    <artifactld>spring-cloud-starter-netflix-ribbon</artifactid>
</dependency>
```

这是因为spring-cloud-starter-netflix-eureka-client自带了spring-cloud-starter-ribbon引用

#### 35.3 RestTemplate的使用

RestTemplate Java Doc

getForObject() / getForEntity() - GET请求方法

getForObject()：返回对象为响应体中数据转化成的对象，基本上可以理解为Json。

getForEntity()：返回对象为ResponseEntity对象，包含了响应中的一些重要信息，比如响应头、响应状态码、响应体等。

```java
@GetMapping("/consumer/payment/getForEntity/{id}")
public CommonResult<Payment> getPayment2(@PathVariable("id") Long id)
{
    ResponseEntity<CommonResult> entity = restTemplate.getForEntity(PAYMENT_URL+"/payment/get/"+id,CommonResult.class);

    if(entity.getStatusCode().is2xxSuccessful()){
        return entity.getBody();//getForObject()
    }else{
        return new CommonResult<>(444,"操作失败");
    }
}
```

**postForObject() / postForEntity()** - POST请求方法

### 三十六、Ribbon默认自带的负载规则

IRule：根据特定算法中从服务列表中选取一个要访问的服务

![img](images/87243c00c0aaea211819c0d8fc97e445.png)

- RoundRobinRule轮询
- RandomRule随机
- RetryRule 先按照RoundRobinRule的策略获取服务，如果获取服务失败则在指定时间内会进行重启。
- WeightedResponseTimeRule 对RoundRobinRule的扩展，响应速度越快的实例选择权重越大，越容易被选择
- BestAvailableRule 会先过滤掉由于多次访问故障而处于断路器跳闸状态的服务，然后选择一个并发量最小的服务
- AvailabilityFilteringRule 先过滤掉故障实例，再选择并发较小的实例
- ZoneAvoidanceRule 默认规则,复合判断server所在区域的性能和server的可用性选择服务器

### 三十七、Ribbon负载规则替换

1.修改cloud-consumer-order80

2.注意配置细节

官方文档明确给出了警告:

这个自定义配置类不能放在@ComponentScan所扫描的当前包下以及子包下，

否则我们自定义的这个配置类就会被所有的Ribbon客户端所共享，达不到特殊化定制的目的了。

（也就是说不要将Ribbon配置类与主启动类同包）

3.新建package - com.yooome.springcloud.config

4.在com.yooome.config下新建MySelfRule规则类

```java
import com.netflix.loadbalancer.IRule;
import com.netflix.loadbalancer.RandomRule;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class MySelfRule {

    @Bean
    public IRule myRule(){
        return new RandomRule();
    }
}

```

5. 主启动类添加@RibbonClient

```java
import com.yooome.config.MySelfRule;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.cloud.netflix.ribbon.RibbonClient;

@SpringBootApplication
@EnableEurekaClient
//添加到此处
@RibbonClient(name = "CLOUD-PAYMENT-SERVICE", configuration = MySelfRule.class)
public class OrderApplication
{
    public static void main( String[] args ){
        SpringApplication.run(OrderApplication.class, args);
    }
}
```

6. 测试

开启cloud-eureka-server7001，cloud-consumer-order80，cloud-provider-payment8001，cloud-privider-payment8002

输入浏览器-输入 http://loclahost/consumer/payment/get/30

返回结果中的serverPort在8001与8002两种间反复横跳。

### 三十八、Ribbon默认负载轮询算法原理

默认负载轮训算法: rest接口第几次请求数 % 服务器集群总数量 = 实际调用服务器位置下标，每次服务重启动后rest接口计数从1开始。

List<Servicelnstance> instances = discoveryClient.getInstances("CLOUD-PAYMENT-SERVICE");

如:

- List [0] instances = 127.0.0.1:8002
- List [1] instances = 127.0.0.1:8001
- 8001+ 8002组合成为集群，它们共计2台机器，集群总数为2，按照轮询算法原理：
- 当总请求数为1时:1%2=1对应下标位置为1，则获得服务地址为127.0.0.1:8001
- 当总请求数位2时:2%2=О对应下标位置为0，则获得服务地址为127.0.0.1:8002
- 当总请求数位3时:3%2=1对应下标位置为1，则获得服务地址为127.0.0.1:8001
- 当总请求数位4时:4%2=О对应下标位置为0，则获得服务地址为127.0.0.1:8002
  如此类推…

### 三十九、RoundRobinRule源码分析

```java
public interface IRule{
    /*
     * choose one alive server from lb.allServers or
     * lb.upServers according to key
     * 
     * @return choosen Server object. NULL is returned if none
     *  server is available 
     */

    //重点关注这方法
    public Server choose(Object key);
    
    public void setLoadBalancer(ILoadBalancer lb);
    
    public ILoadBalancer getLoadBalancer();    
}

```

```java
package com.netflix.loadbalancer;

import com.netflix.client.config.IClientConfig;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * The most well known and basic load balancing strategy, i.e. Round Robin Rule.
 *
 * @author stonse
 * @author Nikos Michalakis <nikos@netflix.com>
 *
 */
public class RoundRobinRule extends AbstractLoadBalancerRule {

    private AtomicInteger nextServerCyclicCounter;
    private static final boolean AVAILABLE_ONLY_SERVERS = true;
    private static final boolean ALL_SERVERS = false;

    private static Logger log = LoggerFactory.getLogger(RoundRobinRule.class);

    public RoundRobinRule() {
        nextServerCyclicCounter = new AtomicInteger(0);
    }

    public RoundRobinRule(ILoadBalancer lb) {
        this();
        setLoadBalancer(lb);
    }

    //重点关注这方法。
    public Server choose(ILoadBalancer lb, Object key) {
        if (lb == null) {
            log.warn("no load balancer");
            return null;
        }

        Server server = null;
        int count = 0;
        while (server == null && count++ < 10) {
            List<Server> reachableServers = lb.getReachableServers();
            List<Server> allServers = lb.getAllServers();
            int upCount = reachableServers.size();
            int serverCount = allServers.size();

            if ((upCount == 0) || (serverCount == 0)) {
                log.warn("No up servers available from load balancer: " + lb);
                return null;
            }

            int nextServerIndex = incrementAndGetModulo(serverCount);
            server = allServers.get(nextServerIndex);

            if (server == null) {
                /* Transient. */
                Thread.yield();
                continue;
            }

            if (server.isAlive() && (server.isReadyToServe())) {
                return (server);
            }

            // Next.
            server = null;
        }

        if (count >= 10) {
            log.warn("No available alive servers after 10 tries from load balancer: "
                    + lb);
        }
        return server;
    }

    /**
     * Inspired by the implementation of {@link AtomicInteger#incrementAndGet()}.
     *
     * @param modulo The modulo to bound the value of the counter.
     * @return The next value.
     */
    private int incrementAndGetModulo(int modulo) {
        for (;;) {
            int current = nextServerCyclicCounter.get();
            int next = (current + 1) % modulo;//求余法
            if (nextServerCyclicCounter.compareAndSet(current, next))
                return next;
        }
    }

    @Override
    public Server choose(Object key) {
        return choose(getLoadBalancer(), key);
    }

    @Override
    public void initWithNiwsConfig(IClientConfig clientConfig) {
    }
}

```

### 四十、Ribbon之厚些轮询算法

自己试着写一个类似RoundRobinRule的本地负载均衡器

- 7001/7002集群启动
- 8001/8002微服务改造 -controller

```java
@RestController
@Slf4j
public class PaymentController{

    ...
    
	@GetMapping(value = "/payment/lb")
    public String getPaymentLB() {
        return serverPort;//返回服务接口
    }
    
    ...
}

```

- 80订单微服务改造

1、ApplicationContextConfig去掉注解@LoadBalanced，OrderMain80去掉注解@RibbonClient

```java
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class ApplicationContextConfig {

    @Bean
    //@LoadBalanced
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
    }

}

```

2、创建LoadBalance接口

```java
import org.springframework.cloud.client.ServiceInstance;

import java.util.List;

/**
 */
public interface LoadBalancer
{
    ServiceInstance instances(List<ServiceInstance> serviceInstances);
}

```

3、MyLB

实现LoadBalancer接口

```java
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

/**
 */
@Component//需要跟主启动类同包，或者在其子孙包下。
public class MyLB implements LoadBalancer
{

    private AtomicInteger atomicInteger = new AtomicInteger(0);

    public final int getAndIncrement()
    {
        int current;
        int next;

        do {
            current = this.atomicInteger.get();
            next = current >= 2147483647 ? 0 : current + 1;
        }while(!this.atomicInteger.compareAndSet(current,next));
        System.out.println("*****第几次访问，次数next: "+next);
        return next;
    }

    //负载均衡算法：rest接口第几次请求数 % 服务器集群总数量 = 实际调用服务器位置下标  ，每次服务重启动后rest接口计数从1开始。
    @Override
    public ServiceInstance instances(List<ServiceInstance> serviceInstances)
    {
        int index = getAndIncrement() % serviceInstances.size();

        return serviceInstances.get(index);
    }
}


```

4、OrderController

```java
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import com.lun.springcloud.lb.LoadBalancer;

@Slf4j
@RestController
public class OrderController {

    //public static final String PAYMENT_URL = "http://localhost:8001";
    public static final String PAYMENT_URL = "http://CLOUD-PAYMENT-SERVICE";

	...

    @Resource
    private LoadBalancer loadBalancer;

    @Resource
    private DiscoveryClient discoveryClient;

	...

    @GetMapping(value = "/consumer/payment/lb")
    public String getPaymentLB()
    {
        List<ServiceInstance> instances = discoveryClient.getInstances("CLOUD-PAYMENT-SERVICE");

        if(instances == null || instances.size() <= 0){
            return null;
        }

        ServiceInstance serviceInstance = loadBalancer.instances(instances);
        URI uri = serviceInstance.getUri();

        return restTemplate.getForObject(uri+"/payment/lb",String.class);

    }
}
```

5、测试

不停的刷新http://localhost/consumer/payment/lb，可以看到8001/8002交替出现。

### 四十一、OpenFeign是什么

[官方文档](https://cloud.spring.io/spring-cloud-static/Hoxton.SR1/reference/htmlsingle/#spring-cloud-openfeign)

[Github地址](https://github.com/spring-cloud/spring-cloud-openfeign)

> Feign is a declarative web service client. It makes writing web service clients easier. To use Feign create an interface and annotate it. It has pluggable annotation support including Feign annotations and JAX-RS annotations. Feign also supports pluggable encoders and decoders. Spring Cloud adds support for Spring MVC annotations and for using the same HttpMessageConverters used by default in Spring Web. Spring Cloud integrates Ribbon and Eureka, as well as Spring Cloud LoadBalancer to provide a load-balanced http client when using Feign. link
>
> Feign是一个声明式WebService客户端。使用Feign能让编写Web Service客户端更加简单。它的使用方法是定义一个服务接口然后在上面添加注解。Feign也支持可拔插式的编码器和解码器。Spring Cloud对Feign进行了封装，使其支持了Spring MVC标准注解和HttpMessageConverters。Feign可以与Eureka和Ribbon组合使用以支持负载均衡。

#### 41.1 Fegin能干什么

Fegin旨在使编写Java Http客户端变的更容易。

前面在使用Ribbon + RestTemplate时，利用RestTemplate 对http请求的封装处理，形成了一套模板化的调用方法。但是在实际卡法中，由于对服务依赖的调用可能不止一处，往往一个接口会别多出调用，所以通常都会针对每个微服务自行封装一些客户端类来包装这些依赖服务的调用。所以，Fegin在此基础上做了进一步封装，有他来帮助我们定义和时间依赖服务接口的定义。在Fegin的实现下，我们只需创建一个接口并使用注解的方式来配置它（以前是Dao接口上面标注Mapper注解，现在是一个微服务接口上面标注一个Fegin注解即可），即可完成服务提供方的接口绑定，简化了使用Spring Cloud Ribbon时，自动封装服务调用客户端的开发量。

#### 41.2 Fegin集成了Ribbon

利用Ribbon维护了Payment的服务列表信息，并且通过轮询实现了客户端的负载均衡。而与Ribbon不同的是，通过fegin只需要定义服务绑定接口且以声明式的方法，优雅而简单的实现了服务调用。

#### 41.3 Fegin和OpenFegin两者区别

Feign是Spring Cloud组件中的一个轻量级RESTful的HTTP服务客户端Feign内置了Ribbon，用来做客户点负载均衡，去调用服务注册中心的服务。Feign的使用方式是：使用Feign的注解定义接口，调用这个接口，就可以调用服务注册中心的服务。

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-feign</artifactId>
</dependency>
```

OpenFeign 是Spring Cloud在Feign的基础上支持了SpringMVC的注解，如@RequestMapping等等。OpenFeign的@Feignclient可以解析SpringMVC的@RequestMapping注解下的接口，并通过动态代理的方式产生实现类，实现类中做负载均衡并调用其它服务。

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-openfeign</artifactId>
</dependency>

```

### 四十二、OpenFeign服务调用

接口+注解：微服务调用接口 + @FeignClient

1、新建cloud-consumer-feign-order80

2、POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-consumer-feign-order80</artifactId>
    <dependencies>
        <!--openfeign-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-openfeign</artifactId>
        </dependency>
        <!--eureka client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <!--web-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--一般基础通用配置-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3、YML

```yml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/

```

4、主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class FeignOrderApplication80 {
    public static void main(String[] args) {
        SpringApplication.run(FeignOrderApplication80.class, args);
    }
}

```

5、业务类

业务类逻辑接口 + @FeignClient配置调用provider服务

新建PaymentFeignService接口新增注解@FeignClient

```java
import com.lun.springcloud.entities.CommonResult;
import com.lun.springcloud.entities.Payment;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;


@Component
@FeignClient(value = "CLOUD-PAYMENT-SERVICE")
public interface PaymentFeignService
{
    @GetMapping(value = "/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id);

}

```

控制层Controller

```java
import com.lun.springcloud.entities.CommonResult;
import com.lun.springcloud.entities.Payment;
import com.lun.springcloud.service.PaymentFeignService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import javax.annotation.Resource;

@RestController
@Slf4j
public class OrderFeignController
{
    @Resource
    private PaymentFeignService paymentFeignService;

    @GetMapping(value = "/consumer/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id)
    {
        return paymentFeignService.getPaymentById(id);
    }

}
```

6、测试

县启动2个eureka集群7001/7002

在启动2个而服务8001/8002

在启动OpenFeign启动

http://localhost/consumer/payment/get/30

Feign自带负载均衡配置项

### 四十三、OpenFeign超时控制

#### 43.1 超时设置，故意设置超时演示错误情况

1、服务提供方8001/8002故意写暂停程序

```java
@RestController
@Slf4j
public class PaymentController {
    
    ...
    
    @Value("${server.port}")
    private String serverPort;

    ...
    
    @GetMapping(value = "/payment/feign/timeout")
    public String paymentFeignTimeout()
    {
        // 业务逻辑处理正确，但是需要耗费3秒钟
        try {
            TimeUnit.SECONDS.sleep(3);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        return serverPort;
    }
    
    ...
}

```

2、服务消费方80添加超时方法PaymentService

```java
@Component
@FeignClient(value = "CLOUD-PAYMENT-SERVICE")
public interface PaymentFeignService{

    ...

    @GetMapping(value = "/payment/feign/timeout")
    public String paymentFeignTimeout();
}
```

3、服务消费方80添加超时方法OrderFeignController

```java
@RestController
@Slf4j
public class OrderFeignController
{
    @Resource
    private PaymentFeignService paymentFeignService;

    ...

    @GetMapping(value = "/consumer/payment/feign/timeout")
    public String paymentFeignTimeout()
    {
        // OpenFeign客户端一般默认等待1秒钟
        return paymentFeignService.paymentFeignTimeout();
    }
}

```

4、测试：

多次刷新 http://localhost/consumer/payment/feign/timeout

将会跳出错误Spring Boot默认错误页面，主要异常：`feign.RetryableException:Read timed out executing GET http://CLOUD-PAYMENT-SERVCE/payment/feign/timeout`。

**OpenFeign默认等待1秒钟，超过后报错**

**YML文件里需要开启OpenFeign客户端超时控制**

```yml
#设置feign客户端超时时间(OpenFeign默认支持ribbon)(单位：毫秒)
ribbon:
  #指的是建立连接所用的时间，适用于网络状况正常的情况下,两端连接所用的时间
  ReadTimeout: 5000
  #指的是建立连接后从服务器读取到可用资源所用的时间
  ConnectTimeout: 5000

```

### 四十四、OpenFeign日志增强

**日志打印功能**

Feign提供了日志打印功能，我们可以通过配置来调用日志级别， 从而了解Feign中Http请求的细节。

说白了就是对Feign接口的调用情况进行监控和输出。

**日志级别**

- NONE：默认的，不信啊是任何日志
- BASIC：进记录请求方式，URL、响应状态码及执行时间。
- HEADERS：处理BASIC中定义的信息之外，还有请求和响应的头信息；
- FULL：除了HEADERS中定义的信息之外，还有请求和响应的正文及元数据。

**配置日志bean**

```java
import feign.Logger;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FeignConfig
{
    @Bean
    Logger.Level feignLoggerLevel()
    {
        return Logger.Level.FULL;
    }
}
```

**YML文件里需要开启日志的Feign客户端**

```yml
logging:
  level:
    # feign日志以什么级别监控哪个接口
    com.lun.springcloud.service.PaymentFeignService: debug
```

**后台日志查看**

得到更多日志信息

### 四十五、Hystrix是什么

#### 45.1 概述

**分布式系统面临的问题**

复杂分布式体系结构中的应用程序有数十个依赖关系，每个依赖关系在某些时候将不可避免地失败。

**服务雪崩**

多个微服务之间调用的时候，假设微服务A调用用微服务B和微服务C，微服务B和微服务C又调用其它的微服务，这就是所谓的“扇出”。如果扇出的链路上某个微服务的调用响应时间过程或者不可用，对微服务A的调用就会占用越来越多的系统资源，进而引起系统崩溃，所谓的“雪崩效应”。

对于高流量的应用来说，单一的后避依赖可能会导致所有服务器上的所有资源都在几秒钟内饱和。比失败更糟糕的是，这些应用程序还能导致服务之间的延迟增加，备份队列，线程和其他系统资源紧张，导致整个系统发生更多的级联故障。这些都表示需要对故障和延迟进行隔离和管理，以便单个依赖关系的失败，不能取消整个应用程序或系统。

所以，通常当你发现一个模块下的某个实例失败后，这时候这个模块依赖海辉接受流量，然后这个有问题的模块还调用了其他的模块，这样就会发生级联故障，或者叫雪崩。

#### 45.2 Hystrix停更进维

能干什么？

- 服务降级
- 服务熔断
- 接近实对的监控
- …

**官网资料**

[link](https://github.com/Netflix/Hystrix/wiki/How-To-Use)

**Hystrix官宣，停更进维**

[link](https://github.com/Netflix/Hystrix)

- 被动修bugs
- 不再接受合并请求
- 不再发布新版本

#### 45.3 Hystrix的服务降级熔断限流概念初讲

**服务降级**

服务器忙，请稍后再试，不让客户端等待并立刻返回一个友好提示，fallback

**那些情况会触发降级**

- 程序运行异常
- 超时
- 服务熔断触发服务降级
- 线程池/信号量打满也会导致服务降级

**服务熔断：**

类比保险丝达到最大服务访问后，直接拒绝访问，拉闸限电，然后调用服务降级的方法并返回友好提示。

服务的降级 -> 进而熔断 -> 恢复调用链路

**服务限流**

秒杀高并发等操作，严禁一窝蜂的过来拥挤，大家排队，一秒钟N个，有序进行。

#### 45.4 Hystrix支付微服务构建

将cloud-eureka-server7001该配置单机版

1、新建cloud-provider-hystrix-payment8001

2、POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-provider-hystrix-payment8001</artifactId>
    <dependencies>
        <!--hystrix-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-hystrix</artifactId>
        </dependency>
        <!--eureka client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <!--web-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3、YML

```yml
server:
  port: 8001

spring:
  application:
    name: cloud-provider-hystric-payment

eureka:
  client:
    service-url:
      defalutZone: http://eureka7001.com:7001/eureka
    register-with-eureka: true
    fetch-registry: true
```

4、主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableEurekaClient
public class HystricApplication8001 {
    public static void main(String[] args) {
        SpringApplication.run(HystricApplication8001.class, args);
    }
}

```

5、业务类

service

```java
import org.springframework.stereotype.Service;

import java.util.concurrent.TimeUnit;

/**
 */
@Service
public class PaymentService {
    /**
     */
    public String paymentInfo_OK(Integer id)
    {
        return "线程池:  "+Thread.currentThread().getName()+"  paymentInfo_OK,id:  "+id+"\t"+"O(∩_∩)O哈哈~";
    }

    public String paymentInfo_TimeOut(Integer id)
    {
        try { TimeUnit.MILLISECONDS.sleep(3000); } catch (InterruptedException e) { e.printStackTrace(); }
        return "线程池:  "+Thread.currentThread().getName()+" id:  "+id+"\t"+"O(∩_∩)O哈哈~"+"  耗时(秒): 3";
    }
}

```

controller

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.service.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
public class PaymentController {
    @Resource
    private PaymentService paymentService;
    @Value("${server.port}")
    private String serverPort;

    @GetMapping("/payment/hystrix/ok/{id}")
    public String paymentInfo_OK(@PathVariable("id") Integer id) {
       return paymentService.paymentInfo_OK(id);
    }
    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id) {
        return  paymentService.paymentInfo_TimeOut(id);
    }

}

```

6、正常测试

启动eureka7001

启动cloud-provider-hystrix-payment8001

访问

success的访问 - http://localhost:8001/payment/hystrix/ok/1

每次调用耗费5秒钟 - http://localhost:8001/payment/hystrix/timeout/30

上述module均ok

以上为根基平台，从正确 -> 错误 -> 降级熔断 -> 恢复

### 46、JMeter高并发压测后卡顿

上述在非高并发情形下，还能勉强满足

Jmeter压测测试

[JMeter官网](https://jmeter.apache.org/index.html)

> The Apache JMeter™ application is open source software, a 100% pure Java application designed to load test functional behavior and measure performance. It was originally designed for testing Web Applications but has since expanded to other test functions.

开启Jmeter，来20000个并发压死8001，20000个请求都去访问paymentInfo_TimeOut服务

1.测试计划中右键添加-》线程-》线程组（线程组202102，线程数：200，线程数：100，其他参数默认）

2.刚刚新建线程组202102，右键它-》添加-》取样器-》Http请求-》基本 输入http://localhost:8001/payment/hystrix/ok/1

3.点击绿色三角形图标启动。

看演示结果：拖慢，原因：tomcat的默认的工作线程数被打满了，没有多余的线程来分解压力和处理。

**Jmeter压测结论**

上面还是服务提供者8001自己测试，假如此时外部的消费者80也来访问，那消费者只能干等，最终导致消费端80不满意，服务端8001直接被拖慢。

### 47、订单微服务调用支付服务出现卡顿

**看热闹不限事大，80新建加入**

1、新建 - cloud-consumer-feign-hystrix-order80

2、POM

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>yooomecloud</artifactId>
        <groupId>com.yooome.springcloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-consumer-feign-hystrix-order80</artifactId>
    <dependencies>
        <!--openfeign-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-openfeign</artifactId>
        </dependency>
        <!--hystrix-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-hystrix</artifactId>
        </dependency>
        <!--eureka client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <dependency>
            <groupId>com.yooome.springcloud</groupId>
            <artifactId>cloud-api-common</artifactId>
            <version>${project.version}</version>
        </dependency>
        <!--web-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--一般基础通用配置-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
    </properties>

</project>
```

3、YML

```yml
server:
  port: 80
eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/
```

4、主启动类

```java
package com.yooome.springcloud;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class FeignHystrixOrderApplication80 {
    public static void main(String[] args) {
        SpringApplication.run(FeignHystrixOrderApplication80.class, args);
    }
}
```

5、业务类

```java
package com.yooome.springcloud.controller;

import com.yooome.springcloud.service.PaymentHystrixService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
public class OrderHystirxController {
    @Value("${server.port}")
    private String serverPort;
    @Resource
    private PaymentHystrixService paymentHystrixService;
    @GetMapping("/consumer/payment/hystrix/ok/{id}")
    public String paymentInfo_OK(@PathVariable("id") Integer id)
    {
        String result = paymentHystrixService.paymentInfo_OK(id);
        return result;
    }

    @GetMapping("/consumer/payment/hystrix/timeout/{id}")
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id) {
        String result = paymentHystrixService.paymentInfo_TimeOut(id);
        return result;
    }
}
```

```java
package com.yooome.springcloud.service;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

@Service
@FeignClient(value = "CLOUD-PROVIDER-HYSTRIX-PAYMENT" /*,fallback = PaymentFallbackService.class*/)
public interface PaymentHystrixService {
    @GetMapping("/payment/hystrix/ok/{id}")
    public String paymentInfo_OK(@PathVariable("id") Integer id);

    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id);
}
```

6、正常测试

http://localhost/consumer/payment/hystrix/ok/1

7、高并发测试

2W个线程压8001

消费端80微服务再去访问正常的Ok微服务8001地址

http://localhost/consumer/payment/hystrix/ok/32

消费者80被拖慢

原因：8001同一层次的其它接口服务被困死，因为tomcat线程池里面的工作线程已经被挤占完毕。

正因为有上述故障或不佳表现才有我们的降级/容错/限流等技术诞生。

### 48、降级容错解决的维度要求

超时导致服务器变慢（转圈） - 超时不再等待

出错（宕机或程序运行出错）- 出错要有兜底

解决：

- 对方服务（8001）超时了，调用者（80）不能一直卡死等待，必须有服务降级。
- 对方服务（8001）down机了，调用（80）不能一直卡死等待，必须有服务降级。
- 对方服务（8001）OK，调用者（80）自己出故障或有自我要求（自己的等待时间小于服务提供者），自己处理降级。

### 49、Hystrix之服务降级支付侧fallback

降级配置 - @HystrixCommand

8001先从自身找问题

设置自身调用超时时间的峰值，峰值内可以正常运行，超过了需要有兜底的方法处理，做服务降级fallback。

8001fallback

业务类启动 - @HystrixCommand报异常后如何处理

一旦调用服务方法失败并抛出了错误信息后，会自动调用@HystrixCommand标注好的fallbackMethod调用类中的指定方法

```java
@Service
public class PaymentService{

    @HystrixCommand(fallbackMethod = "paymentInfo_TimeOutHandler"/*指定善后方法名*/,commandProperties = {
            @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",value="3000")
    })
    public String paymentInfo_TimeOut(Integer id)
    {
        //int age = 10/0;
        try { TimeUnit.MILLISECONDS.sleep(5000); } catch (InterruptedException e) { e.printStackTrace(); }
        return "线程池:  "+Thread.currentThread().getName()+" id:  "+id+"\t"+"O(∩_∩)O哈哈~"+"  耗时(秒): ";
    }

    //用来善后的方法
    public String paymentInfo_TimeOutHandler(Integer id)
    {
        return "线程池:  "+Thread.currentThread().getName()+"  8001系统繁忙或者运行报错，请稍后再试,id:  "+id+"\t"+"o(╥﹏╥)o";
    }
    
}

```

上面故意制造两种异常：

1. int age = 10 /0  ，计算异常
2. 我们能接受3秒钟，它运行5秒钟，超时异常。

当前服务不可用了，做服务降级，兜底的方案都是paymentInfo_TimeOutHandler

**主启动类激活**

添加新注解@EnableCircuitBreaker

```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.circuitbreaker.EnableCircuitBreaker;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableEurekaClient
@EnableCircuitBreaker//添加到此处
public class HystricApplication8001{
    public static void main(String[] args) {
            SpringApplication.run(HystricApplication8001.class, args);
    }
}
```

### 50、Hystrix值服务降级订单侧fallback

80订单微服务，也可以更好的保护自己，自己也依样画葫芦进行客户端降级保护

题外话，切记 - 我们自己配置过的热部署方式对java代码的改动明显

但对@HystrixCommand内属性的修改建议重启微服务

YML

```yml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/

#开启
feign:
  hystrix:
    enabled: true

```

主启动类

```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.hystrix.EnableHystrix;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
@EnableHystrix//添加到此处
public class OrderHystrixMain80{
    
    public static void main(String[] args){
        SpringApplication.run(OrderHystrixMain80.class,args);
    }
}


```

业务类

```java
import com.lun.springcloud.service.PaymentHystrixService;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixCommand;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixProperty;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
@Slf4j
public class OrderHystirxController {
    @Resource
    private PaymentHystrixService paymentHystrixService;


    @GetMapping("/consumer/payment/hystrix/timeout/{id}")
    @HystrixCommand(fallbackMethod = "paymentTimeOutFallbackMethod",commandProperties = {
            @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",value="1500")
    })
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id) {
        //int age = 10/0;
        String result = paymentHystrixService.paymentInfo_TimeOut(id);
        return result;
    }
    
    //善后方法
    public String paymentTimeOutFallbackMethod(@PathVariable("id") Integer id){
        return "我是消费者80,对方支付系统繁忙请10秒钟后再试或者自己运行出错请检查自己,o(╥﹏╥)o";
    }
}

```





