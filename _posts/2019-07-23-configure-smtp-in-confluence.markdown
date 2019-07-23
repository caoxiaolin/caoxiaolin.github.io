---
layout:     post
title:      "Confluence中配置腾讯企业邮箱SMTP邮件服务器"
tags:
    - confluence
    - smtp
    - mail
---

## com.atlassian.mail.MailException: javax.mail.MessagingException: Exception reading response;
## nested exception is:
## java.net.SocketTimeoutException: Read timed out

SMTP Port Optional. The port to connect to on the SMTP host. FishEye needs to use port 25 or 
port 587, because unlike JIRA its initial connection doesn't use SSL. Port 25 will be used if no 
port is specified.

开始配置了465端口，报read timed out，后修改为587，OK！

## com.atlassian.mail.MailException: com.sun.mail.smtp.SMTPSendFailedException: 501 ÇëµÇÂ¼exmail.qq.comÐÞ¸ÄÃÜÂë

这个主要是因为用了管理员开通帐号分配的原始密码，如果用账号密码登录腾讯企业邮箱，会提示修改密码。

修改密码后，配置新密码重试，OK！
