---
layout:     post
title:      "Golang中interface"
tags:
    - golang
    - interface
---

### Interface 是什么
Go 语言不是一种 “传统” 的面向对象的编程语言，它没有类和继承的概念。但是 Go 语言里有非常灵活的 interface，通过它可以实现很多面向对象的特性。

在 C++、Java 中，如果要实现一个接口，需要通过关键字 implements 进行显式的申明，像下面这样：

    public class MyInterface implements A, B, C {
        //code here
    }

而 Golang 的 interface 是隐式的、非侵入式的，Golang 中没有 implements 类似的关键字。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现了这个接口。这种做法就像Python这种纯动态语言中使用的 [Duck Typing](https://en.wikipedia.org/wiki/Duck_typing)，它是静态的。

>If it walks like a duck and it quacks like a duck, then it must be a duck.

如果它像鸭子一样走路，像鸭子一样呱呱叫，那它就是鸭子。

interface 定义了一组方法（方法集），这些方法是抽象的，没有被实现的，接口里也不能包含变量。

如果一个变量含有了多个 interface 类型的方法，那么这个变量就实现了多个接口；如果一个变量只含有了1个 interface 的部分方法，那么这个变量就没有实现这个接口。

空接口 Interface{} 没有任何方法，所以所有类型都实现了空接口。

因此，interface 是什么呢？具有什么特点？

+ interface 是没有变量，只有一堆抽象方法声明的集合
+ 类型不需要显式声明它实现了某个接口，接口被隐式地实现
+ 任何类型包含了在 interface 中声明的全部方法，则表明该类型实现了该接口
+ 实现某个接口的类型（除了实现接口方法外）可以有其他的方法
+ 一个类型可以实现多个接口，多个类型可以实现同一个接口
+ interface 可以作为一种数据类型，实现了该接口的任何对象都可以给对应的接口类型变量赋值

### Golang 中的 interface

作为一门编程语言，对方法的处理一般有两种方式：一种是将所有方法放在一个 table 里，静态的调用（C++, java）；第二种是在调用时动态查找方法(python, smalltalk, javascript)。


而 Go 语言介于两者中间，它虽然也有 table，但又在运行时计算 table。

为什么要这样呢？

我们知道 Golang 是没有严格意义上的继承的，Golang 的 interface 也不存在继承关系。一个类型可以实现多个接口，多个类型可以实现同一个接口，只要你实现了接口定义的方法都可以成为接口类型，这就给 Golang 的 table 初始化带来很大麻烦。到底有多少类型实现了这个接口，一个类型到底实现了多少接口，编译器就懵逼了。但如果我们在运行时实时计算呢，只需要分析一下类型是否实现了接口的所有方法就行了，so easy!

### interface 的底层实现

我们先来看一段代码

	func main() {
		var x *int = nil
		var y interface{} = x
		fmt.Printf("x == nil ? %+v\n", x == nil)
		fmt.Printf("y == nil ? %+v\n", y == nil)
	}

代码很简单，结果很突然。

	x == nil ? true
	y == nil ? false

下面我们从 interface 的底层结构来看这个问题。

根据 interface 是否包含有 method，底层实现上用两种 struct 来表示：eface 和 iface。

eface表示不含 method 的 interface 结构，或者叫 empty interface，下面是其 struct 定义，它包含了两个指针，一个指向值的类型，一个指向实际的数据。

	type eface struct {
		_type *_type
		data  unsafe.Pointer
	}

一个接口变量可以存储任意实际值（非接口），只要这个值实现了接口的方法，所以空接口 interface{} 可以存储任意类型。在上面的例子里，我们把 x 赋给了一个空接口 interface{}，下面我们通过 gdb 来看这两个值有什么不同。

	[root@localhost ~]# go build -gcflags '-l -N' -o test test-interface1.go 
	[root@localhost ~]# gdb test
	GNU gdb (GDB) Red Hat Enterprise Linux 7.6.1-110.el7
	Copyright (C) 2013 Free Software Foundation, Inc.
	License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
	This is free software: you are free to change and redistribute it.
	There is NO WARRANTY, to the extent permitted by law.  Type "show copying"
	and "show warranty" for details.
	This GDB was configured as "x86_64-redhat-linux-gnu".
	For bug reporting instructions, please see:
	<http://www.gnu.org/software/gdb/bugs/>...
	Reading symbols from /root/test...done.
	(gdb) l main.main
	1	package main
	2	
	3	import "fmt"
	4	
	5	func main() {
	6		var x *int = nil
	7		var y interface{} = x
	8		fmt.Printf("x == nil ? %+v\n", x == nil)
	9		fmt.Printf("y == nil ? %+v\n", y == nil)
	(gdb) b test-interface.go:8
	Breakpoint 1 at 0x4897b6: file /root/test-interface1.go, line 8.
	(gdb) r
	Starting program: /root/test 

	Breakpoint 1, main.main () at /root/test-interface1.go:8
	8		fmt.Printf("x == nil ? %+v\n", x == nil)
	(gdb) info locals
	x = 0x0
	y = {_type = 0x494c60, data = 0x0}
	(gdb) 

这里我们看到，y 是一个 eface 的 struct，它的 data 是 nil，但 _type 是有值的。而 Golang 中的 nil，_type 和 data 都是 nil。因此判断 interface 和 nil 的时候一定要注意这一点。

eface 运行时通过 iface.convT2E 系列方法来转换。

	[root@localhost ~]# cat test-interface2.go 
	package main

	import "fmt"

	func main() {
		x := "test"
		var y interface{} = x
		fmt.Printf("%+v\n", y)
	}
	[root@localhost ~]# go build -gcflags '-l -N' -o test test-interface2.go
	[root@localhost ~]# go tool objdump -s "main\.main" test
	TEXT main.main(SB) /root/test-interface2.go
	  test-interface2.go:5	0x489760		64488b0c25f8ffffff		MOVQ FS:0xfffffff8, CX			
	  test-interface2.go:5	0x489769		488d4424d8			LEAQ -0x28(SP), AX			
	  test-interface2.go:5	0x48976e		483b4110			CMPQ 0x10(CX), AX			
	  test-interface2.go:5	0x489772		0f8616010000			JBE 0x48988e				
	  test-interface2.go:5	0x489778		4881eca8000000			SUBQ $0xa8, SP				
	  test-interface2.go:5	0x48977f		4889ac24a0000000		MOVQ BP, 0xa0(SP)			
	  test-interface2.go:5	0x489787		488dac24a0000000		LEAQ 0xa0(SP), BP			
	  test-interface2.go:6	0x48978f		488d05dcf40200			LEAQ 0x2f4dc(IP), AX			
	  test-interface2.go:6	0x489796		4889442458			MOVQ AX, 0x58(SP)			
	  test-interface2.go:6	0x48979b		48c744246004000000		MOVQ $0x4, 0x60(SP)			
	  test-interface2.go:7	0x4897a4		4889442478			MOVQ AX, 0x78(SP)			
	  test-interface2.go:7	0x4897a9		48c784248000000004000000	MOVQ $0x4, 0x80(SP)			
	  test-interface2.go:7	0x4897b5		488d0564060100			LEAQ 0x10664(IP), AX			
	  test-interface2.go:7	0x4897bc		48890424			MOVQ AX, 0(SP)				
	  test-interface2.go:7	0x4897c0		488d442478			LEAQ 0x78(SP), AX			
	  test-interface2.go:7	0x4897c5		4889442408			MOVQ AX, 0x8(SP)			
	  test-interface2.go:7	0x4897ca		e89130f8ff			CALL runtime.convT2Estring(SB)		//在这里调用convT2Estring
	  test-interface2.go:7	0x4897cf		488b442418			MOVQ 0x18(SP), AX			
	  test-interface2.go:7	0x4897d4		488b4c2410			MOVQ 0x10(SP), CX			
	  test-interface2.go:7	0x4897d9		48894c2448			MOVQ CX, 0x48(SP)			
	  test-interface2.go:7	0x4897de		4889442450			MOVQ AX, 0x50(SP)			

iface 表示 non-empty interface 的底层实现，下面是其 struct 定义，它包含了两个指针，一个指向 interface table，叫 itable，另一个指向实际的数据。

	type iface struct {
		tab  *itab
		data unsafe.Pointer
	}

	type itab struct {
		inter *interfacetype
		_type *_type
		hash  uint32 // copy of _type.hash. Used for type switches.
		_     [4]byte
		fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
	}

下面再通过一段简单的代码来分析。

	type Stringer interface {
		 String() string
	}

	type Binary uint64

	func (i Binary) String() string {
		return strconv.FormatUint(i.Get(), 2)
	}

	func (i Binary) Get() uint64 {
		return uint64(i)
	}

	func main(){
		var b Binary = 200
		s := Stringer(b)
		fmt.Println(s.String())
	}

这里的 Binary 实现了 String()，按照 Golang interface 的规则，Binary 实现了 Stringer 接口。这是一种隐式的实现，程序运行时会发现 Binary 有一个 String() 方法，所以它实现了 Stringer，虽然程序本身并不知道 Stringer，也没有打算实现 Stringer。

下面是 iface 的内存组织图
![iface](/img/20181102/gointer2.png)

itable 开头是一些描述类型的元字段，后面是一串方法。注意这个方法是 interface 本身的方法，并非其 dynamic value（Binary）的方法。所以这里只有 String()方法，而没有Get()方法。但这个方法的实现肯定是具体类的方法，这里就是 Binary 的方法。

另一个指针指向实际的数据，在这里是一个 b 的拷贝。


### 参考
 - [https://research.swtch.com/interfaces](https://research.swtch.com/interfaces)
 - [https://www.cnblogs.com/qqmomery/p/6298771.html](https://www.cnblogs.com/qqmomery/p/6298771.html)
 - [https://www.jb51.net/article/128071.htm](https://www.jb51.net/article/128071.htm)
 - [http://legendtkl.com/2017/07/01/golang-interface-implement/](http://legendtkl.com/2017/07/01/golang-interface-implement/)
 - [http://legendtkl.com/2017/06/12/understanding-golang-interface/](http://legendtkl.com/2017/06/12/understanding-golang-interface/)
