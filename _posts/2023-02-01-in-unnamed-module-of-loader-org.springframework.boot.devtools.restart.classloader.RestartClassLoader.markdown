基于spring boot开发时，为了方便调试，引入了 spring-boot-devtools，用于热部署。

spring boot集成了shiro用于权限校验，shiro的session是通过redis来管理的，当把某个对象存入session后，取出来做类型转换时抛出ClassCastException

java.lang.ClassCastException: class com.example.model.shiro.SystemUserVO cannot be cast to class com.example.model.shiro.SystemUserVO (com.example.model.shiro.SystemUserVO is in unnamed module of loader 'app'; com.example.shiro.SystemUserVO is in unnamed module of loader org.springframework.boot.devtools.restart.classloader.RestartClassLoader @506d2c49)

