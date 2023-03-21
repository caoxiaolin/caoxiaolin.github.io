---
layout:     post
title:      "centos7下安装rabbitmq服务"
tags:
    - centos7
    - rabbitmq
---

安装：

    [root@localhost ~]# yum search rabbitmq
    Determining fastest mirrors
    ============================================================================================= N/S matched: rabbitmq ========================
    librabbitmq-devel.i686 : Header files and development libraries for librabbitmq
    librabbitmq-devel.x86_64 : Header files and development libraries for librabbitmq
    librabbitmq-examples.x86_64 : Examples built using the librabbitmq
    opensips-event_rabbitmq.x86_64 : Event RabbitMQ module
    rabbitmq-java-client-doc.noarch : Documentation for rabbitmq-java-client
    rabbitmq-java-client-javadoc.noarch : Javadoc for rabbitmq-java-client
    rabbitmq-server.noarch : The RabbitMQ server
    golang-github-michaelklishin-rabbit-hole-devel.noarch : RabbitMQ HTTP API client in Go
    golang-github-streadway-amqp-devel.noarch : Go client for AMQP 0.9.1 with RabbitMQ extensions
    librabbitmq.i686 : C-language AMQP client library
    librabbitmq.x86_64 : C-language AMQP client library
    rabbitmq-java-client.noarch : Java Advanced Message Queue Protocol client library

    [root@localhost ~]# yum install rabbitmq-server.noarch
    
    等待安装。。。
    
    已安装:
    rabbitmq-server.noarch 0:3.3.5-34.el7
    
    作为依赖被安装:
    erlang-asn1.x86_64 0:R16B-03.18.el7            erlang-compiler.x86_64 0:R16B-03.18.el7              erlang-crypto.x86_64 0:R16B-03.18.el7
    erlang-hipe.x86_64 0:R16B-03.18.el7            erlang-inets.x86_64 0:R16B-03.18.el7                 erlang-kernel.x86_64 0:R16B-03.18.el7
    erlang-os_mon.x86_64 0:R16B-03.18.el7          erlang-otp_mibs.x86_64 0:R16B-03.18.el7              erlang-public_key.x86_64 0:R16B-03.18.
    erlang-sasl.x86_64 0:R16B-03.18.el7            erlang-sd_notify.x86_64 0:0.1-1.el7                  erlang-snmp.x86_64 0:R16B-03.18.el7   
    erlang-stdlib.x86_64 0:R16B-03.18.el7          erlang-syntax_tools.x86_64 0:R16B-03.18.el7          erlang-tools.x86_64 0:R16B-03.18.el7  
    lksctp-tools.x86_64 0:1.0.17-2.el7

    完毕！

    ### 启动管理插件
    [root@localhost ~]# rabbitmq-plugins enable rabbitmq_management
    Plugin configuration unchanged.

    ### 创建账号
    [root@localhost rabbitmq]# rabbitmqctl add_user admin pass123
    Creating user "admin" ...
    ...done.

    ### 设为管理员
    [root@localhost rabbitmq]# rabbitmqctl set_user_tags admin administrator
    Setting tags for user "admin" to [administrator] ...
    ...done.

    ### 设置权限
    [root@localhost rabbitmq]# rabbitmqctl set_permissions -p "/" admin ".*" ".*" ".*"
    Setting permissions for user "admin" in vhost "/" ...
    ...done.

    ### 启动服务
    service rabbitmq-server start

访问：http://[IP]:15672 即可登录后台进行管理

rabbitmq常用命令：



以下是 RabbitMQ 管理常用的命令（均需要管理员权限）：

* 启动 RabbitMQ 服务：systemctl start rabbitmq-server
* 停止 RabbitMQ 服务：systemctl stop rabbitmq-server
* 重启 RabbitMQ 服务：systemctl restart rabbitmq-server
* 查看 RabbitMQ 服务状态：systemctl status rabbitmq-server
* 添加用户：rabbitmqctl add_user [username] [password]
* 设置用户角色：rabbitmqctl set_user_tags [username] [role] （role 可以是多个，用逗号分隔）
* 授权用户访问虚拟主机：rabbitmqctl set_permissions -p [vhost] [username] ".*" ".*" ".*"
* 查看 RabbitMQ 队列列表：rabbitmqctl list_queues
* 查看 RabbitMQ 交换器列表：rabbitmqctl list_exchanges
* 查看 RabbitMQ 绑定列表：rabbitmqctl list_bindings
* 清空队列：rabbitmqctl purge_queue [queue name]
* 查看 RabbitMQ 配置信息：rabbitmqctl environment
* 查看 RabbitMQ 用户列表：rabbitmqctl list_users
* 删除用户：rabbitmqctl delete_user [username]
* 删除队列：rabbitmqctl delete_queue [queue name]
* 删除交换器：rabbitmqctl delete_exchange [exchange name]


