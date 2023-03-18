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

