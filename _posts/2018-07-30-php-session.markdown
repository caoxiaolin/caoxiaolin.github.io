---
layout:     post
title:      "PHP中session的使用"
tags:
    - php
    - session
---

## 什么是 Session？

如今的互联网应用，基本上都是基于 [HTTP协议](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE) 的，HTTP是建立在有状态的 TCP 协议上的无状态连接。

HTTP 的无状态使得协议具有可伸缩性，可以处理海量请求，但同时也带来更多的麻烦，因为几乎所有应用都是需要上下文的。于是，两种用于保持 HTTP 连接状态的技术就应运而生了，一个是 [Cookie](https://en.wikipedia.org/wiki/HTTP_cookie)，而另一个则是 [Session](https://en.wikipedia.org/wiki/Session_(computer_science))。

Session，在计算机中，尤其是在网络应用中，称为“会话控制”。Session 对象存储特定用户会话所需的属性及配置信息。这样，当用户在应用程序的 Web 页之间跳转时，存储在 Session 对象中的变量将不会丢失，而是在整个用户会话中一直存在下去。

## Session 和 Cookie

Session 存储在服务器端，Cookie 存储在客户端。

客户端发出 HTTP 请求时，Cookie 会被放在 HEADER 中一起发给服务器，服务器拿到 Cookie 中的数据，在服务器上找到对应的 Session 数据，从而有了上下文。

PHP 中，默认会生成一个 name 为 PHPSESSID 的 Cookie，你可以通过 php.ini 中的 session.name 来修改它。

	; Name of the session (used as cookie name).
	; http://php.net/session.name
	session.name = PHPSESSID

当发起一个请求时：

![http请求中的cookie](/img/20180730/php-session-1.png)

## PHP中 Session 的使用

PHP 中针对 Session，提供了[一系列的操作函数](https://secure.php.net/manual/zh/book.session.php)，你可以很方便的使用 Session。

你可以通过 php.ini 配置自动开启 Session，或者显式的用 session_start() 来开启。

	; Initialize session on request startup.
	; http://php.net/session.auto-start
	session.auto_start = 0

手工开启：

	session_start();

设置 Session：

	$_SESSION['name'] = 'caoxiaolin';

销毁 Session:

	session_destroy();

## PHP中 Session 的存储

PHP 中的 Session 默认以文件形式存储，在php.ini中配置如下：

	session.save_handler = files
	session.save_path = "/tmp"

文件存储存在一个非常严重的问题，[官方文档](https://secure.php.net/manual/zh/session.examples.basic.php)是这样说的：

>无论是通过调用函数 session_start() 手动开启会话， 还是使用配置项 session.auto_start 自动开启会话， 对于基于文件的会话数据保存（PHP 的默认行为）而言， 在会话开始的时候都会给会话数据文件加锁， 直到 PHP 脚本执行完毕或者显式调用 session_write_close() 来保存会话数据。 在此期间，其他脚本不可以访问同一个会话数据文件。

>对于大量使用 Ajax 或者并发请求的网站而言，这可能是一个严重的问题。 解决这个问题最简单的做法是如果修改了会话中的变量， 那么应该尽快调用 session_write_close() 来保存会话数据并释放文件锁。 还有一种选择就是使用支持并发操作的会话保存管理器来替代文件会话保存管理器。

另外，对于分布式系统而言，文件存储在 Session 共享方面有诸多不便。因此，当前的互联网应用，Session 一般不会以文件形式来存储。

除了默认的文件存储，PHP Session 还支持很多的存储方式，其中比较主流的就是 Redis。

安装 redis 扩展，同时修改 php.ini 配置如下：

	session.save_handler = redis
	session.save_path = "tcp://127.0.0.1:6379"

这时，再请求页面的时候，Session 就会存储在 Redis 中了

	telnet localhost 6379
	Trying ::1...
	Connected to localhost.
	Escape character is '^]'.
	keys *
	*1
	$43
	PHPREDIS_SESSION:6fn70hr00mghe8vibvgc7q1sdj
	get PHPREDIS_SESSION:6fn70hr00mghe8vibvgc7q1sdj
	$23
	name|s:10:"caoxiaolin";

## PHP 中 Session 的过期及回收

Session 都是有生命周期的，这个时间配置在：

	; After this number of seconds, stored data will be seen as 'garbage' and
	; cleaned up by the garbage collection process.
	; http://php.net/session.gc-maxlifetime
	session.gc_maxlifetime = 1440

默认24分钟，从 redis 可以很直观的看到过期时间

	telnet localhost 6379
	Trying ::1...
	Connected to localhost.
	Escape character is '^]'.
	keys *
	*1
	$43
	PHPREDIS_SESSION:ek8bskgq32as8518j8g0ufc2oa
	ttl PHPREDIS_SESSION:ek8bskgq32as8518j8g0ufc2oa
	:1430

如果 Session 存储在 redis 里，会通过 redis 的 GC 回收，这时就没 PHP 什么事了。如果是存储在文件里，则 PHP 提供了一套简单的回收机制。

这里先介绍两个配置项：

	; Defines the probability that the 'garbage collection' process is started
	; on every session initialization. The probability is calculated by using
	; gc_probability/gc_divisor. Where session.gc_probability is the numerator
	; and gc_divisor is the denominator in the equation. Setting this value to 1
	; when the session.gc_divisor value is 100 will give you approximately a 1% chance
	; the gc will run on any give request.
	; Default Value: 1
	; Development Value: 1
	; Production Value: 1
	; http://php.net/session.gc-probability
	session.gc_probability = 1

	; Defines the probability that the 'garbage collection' process is started on every
	; session initialization. The probability is calculated by using the following equation:
	; gc_probability/gc_divisor. Where session.gc_probability is the numerator and
	; session.gc_divisor is the denominator in the equation. Setting this value to 1
	; when the session.gc_divisor value is 100 will give you approximately a 1% chance
	; the gc will run on any give request. Increasing this value to 1000 will give you
	; a 0.1% chance the gc will run on any give request. For high volume production servers,
	; this is a more efficient approach.
	; Default Value: 100
	; Development Value: 1000
	; Production Value: 1000
	; http://php.net/session.gc-divisor
	session.gc_divisor = 1000

这两个配置综合起来定义了在每次 session_start() 时启动垃圾回收进程的概率，此概率用 gc_probability/gc_divisor 计算得来。例如 1/1000 意味着在每1000个请求中会启动1次 gc 进程。

PHP 源码：

	static zend_long php_session_gc(zend_bool immediate)
	{
	    int nrand;
	    zend_long num = -1;

	    /* GC must be done before reading session data. */
	    if ((PS(mod_data) || PS(mod_user_implemented))) {
	        if (immediate) {
	            PS(mod)->s_gc(&PS(mod_data), PS(gc_maxlifetime), &num);
	            return num;
	        }
	        nrand = (zend_long) ((float) PS(gc_divisor) * php_combined_lcg());
	        if (PS(gc_probability) > 0 && nrand < PS(gc_probability)) {
	            PS(mod)->s_gc(&PS(mod_data), PS(gc_maxlifetime), &num);
	        }
	    }
	    return num;
	}

这里 php_combined_lcg() 是一个随机数发生器, 生成0到1范围的随机数
