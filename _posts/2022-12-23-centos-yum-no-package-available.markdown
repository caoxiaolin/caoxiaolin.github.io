---
layout:     post
title:      "centos yum no package available"
tags:
    - centos
    - yum
    - no package available
---

使用yum搜索某些rpm包，找不到包是因为CentOS是RedHat企业版编译过来的，去掉了所有关于版权问题的东西。安装EPEL后可以很好的解决这个问题。

EPEL(Extra Packages for Enterprise Linux )即企业版Linux的扩展包，是一个由 Fedora 社区维护的源，为 RHEL (Red Hat Enterprise Linux) / CentOS 等企业级发行版提供了一些额外的软件包，这些软件包通常不包含在官方的发行版源中。

如果你需要安装一些不包含在 CentOS 官方源中的软件包，那么你需要先安装 EPEL 源，然后再使用 yum 安装需要的软件包。提供了很多可共Centos使用的组件，安装完这个以后基本常用的rpm都可以找到。

安装epel

    yum install epel-release

清理本地缓存

    yum clean all

更新

    yum update

生成缓存

    yum makecache
