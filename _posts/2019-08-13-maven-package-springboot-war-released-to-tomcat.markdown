---
layout:     post
title:      "Maven打包SpringBoot项目到war包发布到tomcat"
tags:
    - spring boot
    - maven
    - war
    - tomcat
---

在 VSCode 下新建了一个 Spring Boot项目，如果想放到 Tomcat 下运行，可以通过 maven 发布成 war 包，直接放到 Tomcat 目录下。

通过 maven 发布，需要做以下几个事情：

* 修改 pom.xml 配置文件，修改或增加外部 Tomcat 部署依赖及 packaging 标签

<pre><code><dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-tomcat</artifactId>
    <scope>provided</scope>
 </dependency>

 <packaging>war</packaging>
</code></pre>

* 修改 DemoApplication.java

<pre><code>import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.web.servlet.support.SpringBootServletInitializer;

@SpringBootApplication
public class DemoApplication extends SpringBootServletInitializer {
    public static void main(String[] args) {
            SpringApplication.run(DemoApplication.class, args);
    }

    @Override
    protected SpringApplicationBuilder configure(SpringApplicationBuilder builder) {
            return builder.sources(DemoApplication.class);
    }
}
</code></pre>

* maven 发布

<pre><code>mvn clean package</code></pre>

* 将 target 目录生成的 war 包拷贝到 Tomcat 的 webapps 目录，重启 Tomcat

最后，访问时要带项目名称，例如 http://localhost:8080/mytest/test
