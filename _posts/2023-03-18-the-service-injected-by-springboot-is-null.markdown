---
layout:     post
title:      "The service injected by spring boot is null"
tags:
    - spring boot
    - Autowired
    - null
---

先看代码：

    @Autowired
    private UserService userService;

    @RestController
    @RequestMapping("/user")
    public class TestController{
        @PostMapping("/list")
        public String getList(@RequestBody UserQuery userQuery){
            List<User> listUser = userService.getList(userQuery);
            return success(listUser);
        }

        @PostMapping("/detail")
        String getDetail(@RequestBody UserQuery userQuery){
            User user = userService.getDetail(userQuery);
            return success(user);
        }
    }

现象：

    java.lang.NullPointerException: null
        at com.example.project.controller.UserController.getDetail(UserController.java:26)

当请求 /user/list 时，执行正常，返回正常

当请求 /user/detail 时，在 userService.getDetail 处报 null

用idea在本地调试，测试一切正常；加日志放到服务器上部署，发现 userService 是 null，getList 方法里是正常的，说明 service 已经注入成功，但在 getDetail 里却变成了 null，为什么？

仔细观察发现，getDetail 方法没有指定访问修饰符。

在 Java 中，如果一个方法没有指定访问修饰符，那么它默认具有包级私有访问权限，只能在同一个包中被访问。

在本地调试代码时，由于 IDE 或命令行工具在同一个包内运行，因此可以访问方法。而在服务器端调试代码时，由于代码部署到了远程服务器上，远程调试工具无法访问该方法，因此会出现访问权限不足的错误。

为了解决这个问题，你可以将该方法的访问修饰符修改为 public 或 protected，这样无论是本地调试还是远程调试，都可以顺利地访问该方法。如果不希望将方法的访问修饰符设置为 public 或 protected，你可以在方法所在的包中添加一个访问权限更高的类或接口，并让该方法成为该类或接口的成员，然后在远程调试时通过该类或接口访问该方法。
