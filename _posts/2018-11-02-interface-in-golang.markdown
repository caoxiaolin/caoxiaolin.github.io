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

### interface 的具体实现

我们先来看一段代码

	func main(){
		var x *int = nil
		fmt.Printf("x == nil ? %+v\n", x == nil)
		fmt.Printf("x == nil ? %+v\n", func(x interface{}) bool{
			return x == nil
		}(x))
	}

代码很简单，结果很突然。

    x == nil ? true
    x == nil ? false

下面我们从 interface 的底层结构来看这个问题。

根据 interface 是否包含有 method，底层实现上用两种 struct 来表示：iface 和 eface。

eface表示不含 method 的 interface 结构，或者叫 empty interface，下面是其 struct 定义，它包含了两个指针，一个指向值的类型，一个指向具体的值。

	type eface struct {
		_type *_type
		data  unsafe.Pointer
	}

一个接口变量可以存储任意实际值（非接口），只要这个值实现了接口的方法，所以空接口 interface{} 可以存储任意类型。例如上面的例子，我们把 x 赋给了一个空接口 interface{}，通过下面的代码来看一下其具体的 struct：

	type InterfaceStructure struct {
		pt uintptr // 到值类型的指针
		pv uintptr // 到值内容的指针
	}

	// asInterfaceStructure 将一个interface{}转换为InterfaceStructure
	func asInterfaceStructure(i interface{}) InterfaceStructure {
		return *(*InterfaceStructure)(unsafe.Pointer(&i))
	}

	func main(){
		var x *int = nil
		fmt.Printf("x == nil ? %+v\n", x == nil)
		fmt.Printf("x == nil ? %+v\n", func(x interface{}) bool{
			return x == nil
		}(x))
		fmt.Printf("%+v\n", asInterfaceStructure(x))
		fmt.Printf("%+v\n", asInterfaceStructure(nil))
	}

上面代码执行后输出：

	x struct: {pt:4790624 pv:0}
	nil struct: {pt:0 pv:0}

这里我们看到，nil 的类型指针和值指针都是0，而 x 的值是0，但类型不是。

iface 表示 non-empty interface 的底层实现，下面是其 struct 定义，它包含了两个指针，一个指向 interface table，叫 itable，另一个指向具体的值。

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

下面通过一个简单的例子来分析。

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
		fmt.Println(s)
	}

这里的 Binary 实现了 String()，按照 Golang interface 的规则，Binary 实现了 Stringer 接口。这是一种隐式的实现，程序运行时会发现 Binary 有一个 String() 方法，所以它实现了 Stringer，虽然程序本身并不知道 Stringer，也没有打算实现 Stringer。

下面是 interface 的内存组织图
![interface](/img/20181102/gointer2.png)

itable 描绘了实际的类型信息及该接口所需要的方法集。

接口值中包含的指针是灰色的，以强调它们是隐式的，而不是直接暴露给 Golang。



### 参考
 - [https://research.swtch.com/interfaces](https://research.swtch.com/interfaces)
 - [https://www.cnblogs.com/qqmomery/p/6298771.html](https://www.cnblogs.com/qqmomery/p/6298771.html)
 - [https://www.jb51.net/article/128071.htm](https://www.jb51.net/article/128071.htm)
 - [http://legendtkl.com/2017/07/01/golang-interface-implement/](http://legendtkl.com/2017/07/01/golang-interface-implement/)
 - [http://legendtkl.com/2017/06/12/understanding-golang-interface/](http://legendtkl.com/2017/06/12/understanding-golang-interface/)
