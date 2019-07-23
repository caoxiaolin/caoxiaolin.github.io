---
layout:     post
title:      "Confluence中配置腾讯企业邮箱SMTP邮件服务器"
tags:
    - confluence
    - smtp
    - mail
---

安装完 Confluence 后，在配置 SMTP 邮件服务的时候，遇到几个问题，记录如下：

>com.atlassian.mail.MailException: javax.mail.MessagingException: Exception reading response;  
>nested exception is:
>java.net.SocketTimeoutException: Read timed out

根据 [腾讯企业邮箱配置](https://service.exmail.qq.com/cgi-bin/help?subtype=1&&id=28&&no=1000585)

发送邮件服务器：smtp.exmail.qq.com ，使用SSL，端口号465

配置后报如上错误。

参考了 [Configuring SMTP In Confluence](https://confluence.atlassian.com/fisheye/configuring-smtp-960155604.html)，关于 SMTP PORT 配置中有这么一段：

>Optional. The port to connect to on the SMTP host. Fisheye needs to use port 25 or port 587, because unlike Jira its initial connection doesn't use SSL. Port 25 will be used if no port is specified.

将端口修改为 587 后，连接 OK！

紧接着，再发送测试邮件的时候，又报了第二个问题：

>com.atlassian.mail.MailException: com.sun.mail.smtp.SMTPSendFailedException: 501 ÇëµÇÂ¼exmail.qq.comÐÞ¸ÄÃÜÂë

这个主要是因为用了腾讯企业邮箱管理员开通帐号时分配的原始密码，如果用账号密码登录腾讯企业邮箱，系统会提示需要修改密码。

修改密码后，配置新密码重试，发送 OK！
