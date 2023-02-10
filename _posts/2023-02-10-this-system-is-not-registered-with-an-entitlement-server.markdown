---
layout:     post
title:      "This system is not registered with an entitlement server"
tags:
    - redhat
    - yum
    - makecache
    - register
---

服务器信息：

    Linux version 3.10.0-957.el7.x86_64 (mockbuild@x86-040.build.eng.bos.redhat.com) (gcc version 4.8.5 20150623 (Red Hat 4.8.5-36) (GCC) ) #1 SMP Thu Oct 4 20:48:51 UTC 2018

更新为163的yum源

    cd /etc/yum.repos.d
    mv ./CentOS-Base.repo ./CentOS-Base-repo.bak
    wget http://mirrors.163.com/.help/CentOS7-Base-163.repo
    yum clean all
    mv CentOS7-Base-163.repo CentOS-Base.repo
    yum makecache
 
 报错：
 
    This system is not registered with an entitlement server. You can use subscription-manager to register.

redhat系统需要授权才可以使用外部yum源，所以需要注册一个开发者账号，如果不做授权的话使用yum安装软件会报错的。生产环境建议购买服务在使用。

注册地址：

https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/registrations?client_id=rhd-web&redirect_uri=https%3A%2F%2Fdevelopers.redhat.com%2Fconfirmation&state=37f5326d-63fd-4625-9630-568df03a7eb8&response_mode=fragment&response_type=code&scope=openid&nonce=74e3e940-73e8-4c6a-8cf6-b16bfbaaf7ca

redhat系统中操作：

    subscription-manager register --username=[账号] --password=[密码] --auto-attach
    
操作成功后即可正常使用yum。
