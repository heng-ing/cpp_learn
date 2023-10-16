# 1、从 C 到C++

## 1.1、引用(相当于给变量起别名)

**引用的概念**

- 类型名 & 引用名 = 某变量名;

```c++
int n = 4;

int & r = n;//r引用了n, r的类型是 int &
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313193009784.png" alt="image-20230313193009784" style="zoom: 50%;" />

 

**常引用**

定义引用时，前面加上const关键字

```c++
int n;
const int & r = n;
```

r的类型是const int &

特点：<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313201144930.png" alt="image-20230313201144930" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313201801867.png" alt="image-20230313201801867" style="zoom:50%;" />

## 1.2、const关键字

### 1.2.1、定义常量

```c++
const double Pi = 3.14;
```

### 1.2.2、定义常量指针

- **不可以通过常量指针修改其指向的内容，但是常量指针的指向可以改变**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313202932646.png" alt="image-20230313202932646" style="zoom:50%;" />

- **不能把常量指针赋值给非常量指针，反过来可以**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313203059050.png" alt="image-20230313203059050" style="zoom:50%;" />

### 1.2.3、定义常引用

- **不能通过常引用修改其引用的变量**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313205920996.png" alt="image-20230313205920996" style="zoom:50%;" />

### 1.2.4、char *const

const在指针后：修饰指针-- - 指针常量 指针指向不可以改，指针指向的值可以改

```C++
int main() {

    char ch1[6] = "list1";
    char ch2[6] = "list2";

    char* const pStr2 = ch2;
    pStr2[4] = '3';
    //pStr2 = ch1; 这里会报错
    cout << pStr2;

    return 0;
}
```

<img src="C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230725122704180.png" alt="image-20230725122704180"  />

### 小结：

**const char\* 类型的指针，其所指向的内容是常量，是不可以修改的，但其指针值是可以修改的。但对于char\*const类型的指针来说，它的地址是一个常量，也就是说，它的指针值是常量，不可以修改，但其指向的内容是可以修改的。**

## 1.3、动态内存分配（new运算符实现动态内存分配,delete释放）

- **第一种用法，分配一个变量：**

P = new T;

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313211815923.png" alt="image-20230313211815923" style="zoom:50%;" />

- **第二种用法，分配一个数组：**

P = new  T[N]

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313212850199.png" alt="image-20230313212850199" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313214703482.png" alt="image-20230313214703482" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313214730909.png" alt="image-20230313214730909" style="zoom:50%;" />

**new运算符的返回值类型：**

new T;

new T[n];

这两个表 达式返回值的类型都是 T* 

- **用new实现动态内存分配，一定要用delete运算符进行释放**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313215403747.png" alt="image-20230313215403747" style="zoom:50%;" />

- **用delete释放动态分配的数组，要加[]**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230313215711912.png" alt="image-20230313215711912" style="zoom:50%;" />

## 1.4、内联函数和重载函数

### 1.4.1、内联函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314153738955.png" alt="image-20230314153738955" style="zoom:50%;" />

**在函数定义前面加 inline 关键字，即可定义内联函数**

```c++
inline int Max(int a, int b) {
	if (a > b)return a;
	return b;
}
```

### 1.4.2、函数重载

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314154755060.png" alt="image-20230314154755060" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314155340335.png" alt="image-20230314155340335" style="zoom:50%;" />



<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316111642064.png" alt="image-20230316111642064" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314155957305.png" alt="image-20230314155957305" style="zoom: 50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314160750864.png" alt="image-20230314160750864" style="zoom:50%;" />

## 1.5、作用域符号::

```c++
#include<iostream>
using namespace std;

int atk = 1000;

int main() {

    int atk = 100;
    cout << atk << endl;
    //::代表作用域，如果前面什么都不添加，代表全局作用域
    cout << ::atk << endl;

    return 0;
}
```

![image-20230724164254582](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230724164254582.png)

## 1.6、namespace命名空间

```c++
#include "game1.h"

void KingGlory::goAtk()
{
	cout << "王者荣耀攻击实现" << endl;
}
```

```c++
#include "game2.h"

void  LOL::goAtk()
{
	cout << "LOL攻击实现" << endl;

}
```

```c++
#include<iostream>
using namespace std;
#include "game1.h"
#include "game2.h"

//1、命名空间用途： 解决名称冲突
void test01()
{
	KingGlory::goAtk();//调用的是KingGlory里面的goAtk()

	LOL::goAtk();//调用的是LOL里面的goAtk()
}

//2、命名空间下 可以放  变量、函数、结构体、类...
namespace A
{
	int m_A;
	void func();
	struct Person
	{};
	class Animal
	{};
}

//3、命名空间 必须要声明在全局作用域下
void test02()
{
	//namespace B{}; 不可以命名到局部作用域
}

//4、命名空间可以嵌套命名空间
namespace B
{
	int m_A = 10;
	namespace C
	{
		int m_A = 20;
	}
}
void test03()
{
	cout << "B空间下的m_A = " << B::m_A << endl;
	cout << "C空间下的m_A = " << B::C::m_A << endl;
}

//5、命名空间是开放的，可以随时给命名空间添加新的成员
namespace B
{
	int m_B = 100;
}
void test04()
{
	cout << "B空间下的m_A = " << B::m_A << endl;
	cout << "B空间下的m_B = " << B::m_B << endl;
}

//6、命名空间可以是匿名的
namespace
{
	int m_C = 1000;
	int m_D = 2000; 
	//当写的命名空间的匿名的，相当于写了  static int m_C = 1000; static int m_D = 2000;
}

void test05()
{
	cout << "m_C = " << m_C   << endl;
	cout << "m_D = " << ::m_D << endl;
}

//7、命名空间可以起别名
namespace veryLongName
{
	int m_E = 10000;
	void func()
	{
		cout << "aaa" << endl;
	}
}



void test06()
{
	namespace veryShortName = veryLongName;
	cout << veryShortName::m_E << endl;
	cout << veryLongName::m_E << endl;

}

int main(){
	//test01();
	//test03();
	//test04();
	//test05();
	test06();
	

	system("pause");
	return 0;
}
```

## 1.7、using

```c++
#include<iostream>
using namespace std;

namespace KingGlory
{
	int sunwukongId = 1;
}

namespace LOL
{
	int sunwukongId = 3;
}

void test01()
{
	int sunwukongId = 2;

	//1、using声明
	//using KingGlory::sunwukongId ;  

	//当using声明与 就近原则同时出现，出错，尽量避免
	cout << sunwukongId << endl;

}


void test02()
{
	//int sunwukongId = 2;
	//2、using编译指令
	using namespace KingGlory;
	using namespace LOL;
	//当using编译指令  与  就近原则同时出现，优先使用就近
	//当using编译指令有多个，需要加作用域 区分
	cout << KingGlory::sunwukongId << endl;
	cout << LOL::sunwukongId << endl;
}

int main(){

	//test01();
	test02();
	system("pause");
	return 0;
}
```



# 2、类型和对象基础

## 2.1、类和对象的基本概念

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314185848969.png" alt="image-20230314185848969" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314190236052.png" alt="image-20230314190236052" style="zoom:50%;" />

- **一个类的成员函数在内存中只有一份，被所有对象共享**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314190822279.png" alt="image-20230314190822279" style="zoom:50%;" />

- **类的成员函数和类的定义分开写**

```c++
class CRectangle {

	public:
		int w, h;
		int Area();
		int Perimeter();
		void Init(int w_, int h_);

};

int CRectangle::Area() {
	return w * h;
}

int CRectangle::Perimeter() {
	return 2 * (w + h);
}

void CRectangle::Init(int w_, int h_) {
	w = w_;
	h = h_;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314192710180.png" alt="image-20230314192710180" style="zoom: 50%;" />

- **类成员的可访问范围**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314192913381.png" alt="image-20230314192913381" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314193220929.png" alt="image-20230314193220929" style="zoom:50%;" />

​		**私有成员变量或函数**在类的外部是不可访问的，甚至是不可查看的。只有类和友元函数可以访问私有成员。

- **成员函数的重载及参数缺省**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314211347547.png" alt="image-20230314211347547" style="zoom:50%;" />

## 2.2、构造函数

- **基本概念**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314212204707.png" alt="image-20230314212204707" style="zoom: 50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314213219251.png" alt="image-20230314213219251" style="zoom: 50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314220614554.png" alt="image-20230314220614554" style="zoom:50%;" />

在C++中调用构造函数的方式有如下几种：

> **A a = A();//调用默认构造函数**
>
> **A a = A(xx);//调用带参的构造函数**
>
> **A a(xx);//调用带参的构造函数的简写形式**
>
> **A a;//调用默认构造函数的简写形式**

- **构造函数在数组中的使用**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314221053283.png" alt="image-20230314221053283" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314221421399.png" alt="image-20230314221421399" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230314221956340.png" alt="image-20230314221956340" style="zoom:50%;" />

**最后一条只初始化两个对象**

## 2.3、复制构造函数

**基本概念：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316113907295.png" alt="image-20230316113907295" style="zoom: 67%;" />

```c++
#include <iostream>
using namespace std;

class Complex {
private :
	double real, image;

public: 
	Complex(int real_, int image_) {
		real = real_;
		image = image_;
	}

	void print() {
		cout << real << endl;
		cout << image << endl;
	}

};
int main() {
	Complex c1(1, 5);
	Complex c2(c1);//调用缺省的复制构造函数，将c2初始化成c1一样

	c1.print();
	c2.print();
	return 0;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316121021623.png" alt="image-20230316121021623" style="zoom:80%;" />

**参数必须是引用，不能是对象**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316123226983.png" alt="image-20230316123226983" style="zoom: 67%;" />

**复制构造函数起作用的三种情况**

1、当用一个对象去初始化同类的另一个对象时。

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316134405409.png" alt="image-20230316134405409" style="zoom: 67%;" />

2、如果某函数有一个参数是类A的对象，那么该函数被调用的时，类A的复制构造函数将被调用。

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316135823005.png" alt="image-20230316135823005" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316135855600.png" alt="image-20230316135855600" style="zoom:67%;" />

3、如果函数的返回值是类A的对象时，则函数返回时，A的复制构造函数被调用：

**注意：对象间的赋值并不会导致复制构造函数被调用**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316141943576.png" alt="image-20230316141943576" style="zoom: 80%;" />

**常量引用参数的使用**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316142558048.png" alt="image-20230316142558048" style="zoom:80%;" />

## 2.4、类型转换构造函数

**什么是类型转换构造函数**：类型转换构造函数也是一个**带有一个参数的普通构造函数**。

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316143920645.png" alt="image-20230316143920645" style="zoom:67%;" />

举例：

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316144840549.png" alt="image-20230316144840549" style="zoom: 80%;" />

![image-20230316144953482](C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316144953482.png)

## 2.5、析构函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316150041566.png" alt="image-20230316150041566" style="zoom: 80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316150914188.png" alt="image-20230316150914188" style="zoom:80%;" />

**析构函数和运算符delete**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316154428334.png" alt="image-20230316154428334" style="zoom: 80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316162412894.png" alt="image-20230316162412894" style="zoom:67%;" />

## 2.6、构造函数和析构函数调用时机

```c++
class Demo {
	int id;
public:
	Demo(int i) {
		id = i;
		cout << "id=" << id << "构造函数被调用" << endl;
	}
	~Demo() {
		cout << "id=" << id << "析构函数被调用" << endl;
	}

};
Demo d1(1);
void Func() {
	static Demo d2(2);
	Demo d3(3);
	cout << "func" << endl;
}

int main() {
	Demo d4(4);
	d4 = 6;
	cout << "main" << endl;
	{
		Demo d5(5);
	}
	Func();
	cout << "main ends" << endl;
	return 0;
}
```

![image-20230316174915813](C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230316174915813.png)

​	**全局对象**在main()函数之前初始化，即Demo d1(1);

​	d4 = 6;类型不一样，调用类型转换构造函数，将6转换成一个临时对象，此时调用一次**构造构造函数**。这语句执行完后，临时对象销毁，执行**析构函数**。

​	全局对象Demo d1(1);静态对象static Demo d2(2);**先构造的后析构**

# 3、类和对象提高

## 3.1、this指针

​	**C++程序到C程序的翻译**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317165058505.png" alt="image-20230317165058505" style="zoom:67%;" />

**this指针作用：其作用就是指向成员函数所作用的对象**

**非静态成员函数**中可以直接使用this来代表指向该函数作用的对象的指针，即指向调用该函数的当前对象的指针。

```c++
class Complex {
private:
	double real, image;

public:
	Complex(double real_, double image_) {
		real = real_;
		image = image_;
	}

	void print() {
		cout << real <<" ";
		cout << image << endl;
	}
	
	Complex AddOne() {
		this->real++;//等价于 real ++
		this->print//等价于 print()
		return *this;
	}

};

int main() {
	Complex c1(1, 1), c2(0, 0);
	c2 = c1.AddOne();
	return 0;
}
```

![image-20230317163440743](C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317163440743.png)

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317164619028.png" alt="image-20230317164619028" style="zoom:67%;" />

**这里并不会出错**，因为翻译成c语言后，并没有用到空指针p

以下是会出错的例子：

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317165643399.png" alt="image-20230317165643399" style="zoom: 67%;" />

**this指针和静态成员函数**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317165900661.png" alt="image-20230317165900661" style="zoom:67%;" />

## 3.2、静态成员变量（static关键字）

- ​	普通成员变量每个对象有各自的一份，而静态成员变量**一共就一份**，**为所有对象共享**

 <img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317170828169.png" alt="image-20230317170828169" style="zoom:67%;" />

- 普通成员函数必须具体作用于某个对象，而静态成员函数**并不具体作用于某个对象**

**如何访问静态成员？**

1、类名::成员名

2、对象名.成员名

3、指针->成员名

4、引用.成员名

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317172125778.png" alt="image-20230317172125778" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317172955560.png" alt="image-20230317172955560" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317173009749.png" alt="image-20230317173009749" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317173331680.png" alt="image-20230317173331680" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317173351993.png" alt="image-20230317173351993" style="zoom:67%;" />

**注意事项：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317173802745.png" alt="image-20230317173802745" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317173817844.png" alt="image-20230317173817844" style="zoom:67%;" />



以上写法有严重的缺陷

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317174613903.png" alt="image-20230317174613903" style="zoom: 67%;" />

​	临时的对象并不会时nTotalNumber和nTotalArea增加1，但是临时对象消亡时，会调用析构函数，即nTotalNumber和nTotalArea会减少1

解决方法：自己写一个复制构造函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230317175020162.png" alt="image-20230317175020162" style="zoom:50%;" />

## 3.3、成员对象和封闭类

**有成员对象的类叫封闭类：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230320143819265.png" alt="image-20230320143819265" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230320145018635.png" alt="image-20230320145018635" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321143636569.png" alt="image-20230321143636569" style="zoom: 67%;" />

**封闭类构造函数和析构类构造函数的执行顺序**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321144426700.png" alt="image-20230321144426700" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321150038709.png" alt="image-20230321150038709" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321164355164.png" alt="image-20230321164355164" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321151027483.png" alt="image-20230321151027483" style="zoom:80%;" />

**封闭类的复制构造函数**

```c++
class A {
public:
	A() { cout << "A的无参构造函数" << endl; }
	A(A & a) { cout << "A的复制构造函数" << endl; }
};

class B { 
public:
	A a; 

	B() { cout << "B的无参构造函数" << endl; }
	B(B & b):a(b.a) { 
		cout << "B的复制构造函数" << endl;
	}

};
int main() {

	B b1;
	cout << "***********"<<endl;
	B b2(b1);
	return 0;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321170057385.png" alt="image-20230321170057385" style="zoom: 80%;" />

​	B b2(b1);语句表示，b2对象**会**调用**类B的复制构造函数**（用初始化列表的方式去初始化类B的成员对象a，B类对象只有一个成员对象A a），但是因为**有成员对象A**，所以此时b2.a会**先调用类A的复制构造函数**，对b2.a进行初始化，并将b1.a的值复制给b2.1。当去掉类B的构造函数B(B & b):a(b.a) { }这条语句时，此时调用类B的默认构造函数，结果不变。

**注意以下代码，与上面的区别：**

当构造函数写成以下时

```c++
B(B & b) {
		a = b.a;
		cout << "B的复制构造函数" << endl;
	}
```

输出结果如下：

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321171138417.png" alt="image-20230321171138417" style="zoom:80%;" />

**此时并不会调用A的复制构造函数**

## 3.4、常量对象、常量成员函数和常引用（const关键字）

- 常量对象：不希望某个对象的值被改变

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230321172337018.png" alt="image-20230321172337018" style="zoom:67%;" />

- 常量成员函数：常量成员函数执行期间，不应该修改其作用的对象


<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322135807791.png" alt="image-20230322135807791" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322141302705.png" alt="image-20230322141302705" style="zoom: 80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322141643051.png" alt="image-20230322141643051" style="zoom:67%;" />

**两个成员函数，名字和参数列表都一样，但是一个是const，一个不是，算重载而不是重复定义**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322144558439.png" alt="image-20230322144558439" style="zoom:67%;" />

objTest1是常量对象，会调用常量成员函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322145513991.png" alt="image-20230322145513991" style="zoom:80%;" />

- 常引用：不能通过常引用，修改其作用的变量

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322150107206.png" alt="image-20230322150107206" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322150205987.png" alt="image-20230322150205987" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322150244018.png" alt="image-20230322150244018" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322150340377.png" alt="image-20230322150340377" style="zoom:67%;" />

## 3.5、友元（分为友元函数和友元类两种）

**1）、友元函数：一个类的友元函数可以访问该类的私有成员**

```c++
class CCar;
class CDriver {
public:
	void ModifyCar(CCar * pCar);//改装汽车
};

class CCar {
private:
	int price;
	friend int MostExpensiveCar(CCar cars[], int total);//声明友元
	friend void CDriver::ModifyCar(CCar * pCar);//将CDriver类的ModifyCar函数声明成友元，
												//所以访问CCar的私有成员price
};

void CDriver::ModifyCar(CCar * pCar) {
	pCar->price += 1000;//汽车改装后价值增加
}

//全局函数
int MostExpensiveCar(CCar cars[], int total) {
	//求最贵汽车的价格
	int tmpMax = -1;
	for (int i = 0; i < total; ++i) {
		if (cars[i].price > tmpMax)
			tmpMax = cars[i].price;
	}
}
```

**2)、友元类：如果A是B的友元类，那么A的成员函数可以访问B的私有成员**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322162547385.png" alt="image-20230322162547385" style="zoom: 67%;" />

**友元类之间的关系，不能传递，不能继承。**

# 4、运算符重载

## 4.1、运算符重载的基本概念

**运算符重载的需求**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322163325943.png" alt="image-20230322163325943" style="zoom:67%;" />

**运算符重载的形式**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322164421621.png" alt="image-20230322164421621" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322164438048.png" alt="image-20230322164438048" style="zoom:80%;" />

```c++
#include <iostream>
using namespace std;

class Complex {
public:
	double real, imag;
	Complex(double r = 0.0, double i = 0.0) :real(r), imag(i) {}
	Complex operator-(const Complex & c);
};

Complex operator+(const Complex & a, const Complex & b) {
	return Complex(a.real + b.real, a.imag + b.imag);//返回一个临时对象
}

Complex Complex::operator-(const Complex & c) {
	return Complex(real - c.real, imag - c.imag);//返回一个临时对象
}

int main() {
	Complex a(4, 4), b(1, 1), c;
	c = a + b;//等价于c = operator+(a,b)
	cout << c.real << "," << c.imag << endl;
	cout << (a-b).real << "," << (a - b).imag << endl;
	//a-b等价于a.operator-(b)
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322190724589.png" alt="image-20230322190724589" style="zoom:80%;" />

**重载为成员函数时，参数个数为运算符目数减一。**

**重载为普通函数时，参数个数为运算符目数。**

## 4.2、复制运算符的重载（=）

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230322192105115.png" alt="image-20230322192105115" style="zoom: 67%;" />

```c++
#include <iostream>
#include <cstring>
using namespace std;

class String {
private:
	char * str;
public:
	String() :str(new char[1]) { str[0] = 0; }
	const char * c_str() { return str; }
	String & operator = (const char * s);
	String & operator = (const String & s);

	~String() { delete[] str; }
	
	String(const char * s) {
		delete[] str;
		str = new char[strlen(s) + 1];
		strcpy(str, s);
	}
    
    String(String & s) {
		str = new char[strlen(s.str) + 1];
		strcpy(str, s.str);
	}

};

String & String::operator = (const char * s) {
    //重载“=”以使得obj = "hello"能够成立
	delete[] str;
	str = new char[strlen(s) + 1];//strlen(s)并不会算上\0
	strcpy(str, s);//字符串复制，遇到‘\0’时停止，还会复制字符串的结束符'\0'
	return *this;
}

String & String::operator = (const String & s) {
    //避免s1 = s2时，s1和s2指向同一个内存空间
	if (this == &s)//排除s = s;这种操作
		return *this;
	delete[] str;
	str = new char[strlen(s.str) + 1];//strlen(s)并不会算上\0
	strcpy(str, s.str);//字符串复制，遇到‘\0’时停止，还会复制字符串的结束符'\0'
	return *this;
}


int main() {
	String s;
	s = "Good Luck,";
	cout << s.c_str() << endl;
	s = "123456798";
	cout << s.c_str() << endl;
	String s2 = "nihai a a";
	cout << s2.c_str() << endl;
	return 0;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325120655589.png" alt="image-20230325120655589" style="zoom:80%;" />

如果不重载，当S1 = S2这样的语句会出现问题

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325121920908.png" alt="image-20230325121920908" style="zoom:67%;" />

​	原生的 = 的含义是将等号右边的全部复制给左边的，这样s1.str与s2.,str会指向同一个内存空间，原来的s1.str指向的空间无法释放, s1对象消亡，析构函数将释放s1.str指向的空间，**则s2消亡时还要释放一次,**不妥。

**对operator = 返回值的讨论**

对运算符进行重载的时候，好的风格是应该尽量保持运算符原本的特性，例如赋值运算符的特性，**= 原本就是返回的等号左边的引用**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325124823341.png" alt="image-20230325124823341" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325124847836.png" alt="image-20230325124847836" style="zoom: 67%;" />

## 4.3、运算符重载为友元函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325151827100.png" alt="image-20230325151827100" style="zoom:80%;" />

```c++
class Complex {
private:
	double real, imag;
public:
	Complex(double r, double i) :real(r), imag(i) {}
	Complex operator-(const Complex & c);
	Complex operator+(double r);

};

Complex Complex::operator+(double r) {
	//能解释c+5
	return Complex(real + r, imag);
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325152056647.png" alt="image-20230325152056647" style="zoom:80%;" />

所以，为了使得上述的表达式能成立，需要将+重载为普通函数。

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325161721832.png" alt="image-20230325161721832" style="zoom:67%;" />

但是普通函数又不能访问私有成员，所以，需要将运算符+重载为友元。

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325163757624.png" alt="image-20230325163757624" style="zoom: 80%;" />

## 4.4、可变长数组类的实现

**可变长整型数组**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325214247050.png" alt="image-20230325214247050" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230325214301723.png" alt="image-20230325214301723" style="zoom:80%;" />

**非引用的函数返回值，不可以作为 = 的左值使用**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328142556700.png" alt="image-20230328142556700" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328152416322.png" alt="image-20230328152416322" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328152457161.png" alt="image-20230328152457161" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328152531439.png" alt="image-20230328152531439" style="zoom:67%;" />

## 4.5、流插入运算符和流提取运算符的重载

**问题：**

**cout << 5 << "this"; 为什么能够成立？**

**cout 是什么？"<<"为什么能够用在cout上？**

- cout是在iostream头文件中定义的ostream类的对象。

- "<<"能用在cout上是因为，在iostream里对"<<"进行了重载。

**考虑怎么重载才能使得cout<<5;和cout<<"this";都能成立？**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328162455139.png" alt="image-20230328162455139" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328162525623.png" alt="image-20230328162525623" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328162553114.png" alt="image-20230328162553114" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328163103567.png" alt="image-20230328163103567" style="zoom:67%;" />

**只能重载为全局函数：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328163245685.png" alt="image-20230328163245685" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328163929876.png" alt="image-20230328163929876" style="zoom:67%;" />

```c++
#include <iostream>
#include <string>
#include <cstdlib>
using namespace std;

class Complex {
public:
	double real, imag;
	Complex(double r = 0.0, double i = 0.0) :real(r), imag(i) {}
	Complex operator-(const Complex & c);
};

Complex operator+(const Complex & a, const Complex & b) {
	return Complex(a.real + b.real, a.imag + b.imag);//返回一个临时对象
}

Complex Complex::operator-(const Complex & c) {
	return Complex(real - c.real, imag - c.imag);//返回一个临时对象
}

ostream & operator<<(ostream &os,const Complex & c) {
	os << c.real << "+" << c.imag << "i";
	return os;
}

istream & operator>>(istream &is, Complex & c) {
	string s;
	is >> s;//将"a+bi"作为字符串读入，"a+bi"中间不能有空格
	int pos = s.find("+", 0);
	string sTmp = s.substr(0, pos);//分离出代表实部的字符串.截取下表为0,长度为pos个字符串
	c.real = atof(sTmp.c_str());//atof库函数能将const char*指针指向的内容转换成float
	sTmp = s.substr(pos + 1, s.length() - pos - 2);//分离出代表的虚部的字符串
	c.imag = atof(sTmp.c_str());
	return is;
}

int main() {
	
	Complex c;
	int n;
	cin >> c >> n;
	cout << c << "," << n;

	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328170724699.png" alt="image-20230328170724699" style="zoom:80%;" />

## 4.6、类型转换运算符的重载

```c++
#include <iostream>

using namespace std;

class Complex {
	double real, imag;
public:
	Complex(double r = 0.0, double i = 0.0) :real(r), imag(i) {}
	operator double() { return real; }
    //重载强制类型转换运算符double，不写返回值类型
};

int main() {

	Complex c(1.2, 3.4);
	cout << (double)c << endl;
	double n = 2 + c;
	cout << n;
	return 0;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328183039333.png" alt="image-20230328183039333" style="zoom:80%;" />

## 4.7、自增自减运算符的重载

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328184823982.png" alt="image-20230328184823982" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328184910146.png" alt="image-20230328184910146" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328184931016.png" alt="image-20230328184931016" style="zoom:80%;" />

  **原生的++a;返回的是a的引用。原生的a++;返回的是一个临时变量，就是+操作之前的值。**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328201214190.png" alt="image-20230328201214190" style="zoom: 67%;" />

# 5、继承

## 5.1、继承和派生的基本概念

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328201511905.png" alt="image-20230328201511905" style="zoom:67%;" />

**派生类的写法：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328202727035.png" alt="image-20230328202727035" style="zoom:50%;" />

**派生类对象的内存空间：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328204642701.png" alt="image-20230328204642701" style="zoom:80%;" />

## 5.2、继承关系和复合关系

**类之间的两种关系：**

 <img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328210546527.png" alt="image-20230328210546527" style="zoom:67%;" />

例如**基类是学生类**，**派生类是中学生**，从逻辑上，一个中学生也是一个学生

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328211724707.png" alt="image-20230328211724707" style="zoom:67%;" />

## 5.3、覆盖和保护成员

- **覆盖：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328213834268.png" alt="image-20230328213834268" style="zoom:67%;" />

**基类和派生类有同名成员的情况：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328214204512.png" alt="image-20230328214204512" style="zoom:80%;" />

**但是一般不要在派生类与基类中写同名的成员变量！！！！！**

- **类的保护成员**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230328215201074.png" alt="image-20230328215201074" style="zoom:80%;" />

## 5.4、派生类的构造函数

```c++
class Bug {
private:
	int nLegs; int nColor;
public:
	int nType;
	Bug(int legs, int colotr);
	void PrintBug() {};

};

Bug::Bug(int legs, int colotr):nLegs(legs),nColor(colotr) {
	
}

class FlyBug :public Bug {//FlyBug是Bug的派生类
	int nWings;
public:
	FlyBug(int legs, int color, int wings);
};

//错误的FlyBug构造函数
FlyBug::FlyBug(int legs, int color, int wings) {
	nLegs = legs;  //不能访问
	nColor = color;//不能访问
	nType = 1;//OK
	nWings = wings;
}

//正确的FlyBug构造函数:使用初始化列表的方式
FlyBug::FlyBug(int legs, int color, int wings):Bug(legs,color) {
	nWings = wings;
}
```

- 在创建派生类的对象时，需要调用基类的构造函数: 初始化派生类对象中从基类继承的成员。在执行一个派生类的构造函数之前，总是先执行基类的构造函数


- 调用基类构造函数的两种方式

  ​    显式方式（初始化列表的方式）：在派生类的构造函数中，为基类的构造函数提供参数.derived::derived(arg derived-list):base(arg base-list)

  ​	隐式方式: 在派生类的构造函数中，**省略基类构造函数时**，派生类的构造函数则自动调用**基类的默认构造函数**

- 派生类的析构函数被执行时，**先执行完派生类的析构函数后，自动调用基类的析构函数。**



```c++
#include <iostream>
using namespace std;

class Base {
public:
	int n;
	Base(int i) :n(i) {
		cout << "Base" << n << "constructed" << endl;
	}
	~Base() {
		cout << "Base" << n << "destructed" << endl;
	}
};

class Derived :public Base {
public:
	Derived(int i) :Base(i) {
		cout << "Derived constructed" << endl;
	}
	~Derived(){
		cout << "Derived destructed" << endl;
	}
};

int main() {
	Derived derived(12);
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416203837442.png" alt="image-20230416203837442" style="zoom:80%;" />

**封闭派生类对象的构造函数执行顺序:**

在创建派生类的对象时: 

1. 先执行基类的构造函数，用以初始化派生类对象中从基类继承的成员
2. 再执行成员对象类的构造函数，用以初始化派生类对象中成员对象。
3. 最后执行派生类自己的构造函数



在派生类对象消亡时:

1. 先执行派生类自己的析构函数
2. 依次执行各成员对象类的析构函数
3. 最后执行基类的析构函数

**析构函数的调用顺序与构造函数的调用顺序相反**。

## 5.5、公有继承的赋值兼容规则

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416210726483.png" alt="image-20230416210726483" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416210853770.png" alt="image-20230416210853770" style="zoom:80%;" />

**直接基类和间接基类**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416211342882.png" alt="image-20230416211342882" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416211418156.png" alt="image-20230416211418156" style="zoom:80%;" />

# 6、虚函数

## 6.1、虚函数和多态的基本概念

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416212320940.png" alt="image-20230416212320940" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416212845207.png" alt="image-20230416212845207" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416213952866.png" alt="image-20230416213952866" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416213331080.png" alt="image-20230416213331080" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230416214051805.png" alt="image-20230416214051805" style="zoom:80%;" />

```c++
#include <iostream>
using namespace std;


class Base {
public:
	void fun1() { fun2(); }
	virtual void fun2() {
		cout << "Base::fun2()" << endl;
	}
};

class Derived :public Base {
public:
	virtual void fun2() {
		cout << "Derived:fun2()" << endl;
	}
};

int main() {
	Derived d;
	Base * pBase = &d;
	pBase->fun1();
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230418220356803.png" alt="image-20230418220356803" style="zoom:80%;" />

**在非构造函数，非析构函数的成员函数中调用虚函数，是多态!!!**

Base中的fun1()调用fun2()，

```c++
void fun1() { fun2(); }
```

等价于

```c++
void fun1() { this->fun2(); }
```

this的含义是指向当前对象，

```c++
Base * pBase = &d;
pBase->fun1();
```

此时this指向的是d，所以**this->fun2();**调用的是**Derived中的fun2()**

​		在**构造函数和析构函数**中调用虚函数，不是多态。编译时即可确定，调用的函数是自己的类或基类中定义的函数，不会等到运行时才决定调用自己的还是派生类的函数。

## 6.2、多态的实现原理

​		“多态”的关键在于通过基类指针或引用调用一个虚函数时，编译时不确定到底调用的是基类还是派生类的函数，运行时才确定 ---- 这叫“动态联编”。“动态联编” 底是怎么实现的呢?

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230420122245727.png" alt="image-20230420122245727" style="zoom:80%;" />

多态实现的关键——**虚函数表**
		每一个有虚函数的类 (或有虚函数的类的派生类)都有一个虚函数表，该类的任何对象中都放着虚函数表的指针。虚函数表中列出了该类的虚函数地址。**多出来的4个字节就是用来放虚函数表的地址的。**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230420122947111.png" alt="image-20230420122947111" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230619161857796.png" alt="image-20230619161857796" style="zoom:67%;" />

```c++
#include <iostream>
using namespace std;

class A {
public:virtual void Func() {
	cout << "A::Func" << endl;
}
};

class B :public A {
public:virtual void Func() {
	cout << "B::Func" << endl;
}
};

int main() {
	A a;
	A * pa = new B();
	pa->Func();
	long long * p1 = (long long *)& a;
	long long * p2 = (long long *)pa;
	*p2 = *p1;
	pa->Func();
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230421220421512.png" alt="image-20230421220421512" style="zoom:80%;" />

## 6.3、虚析构函数、纯虚函数和抽象类

### 6.3.1、虚析构函数

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422141143997.png" alt="image-20230422141143997" style="zoom:80%;" />

### 6.3.2、纯虚函数和抽象类

**没有函数体的虚函数**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422142829987.png" alt="image-20230422142829987" style="zoom:80%;" />

- **在抽象类的成员函数内可以调用纯虚函数，但是在构造函数或析构函数内部不能调用纯虚函数。**

- **如果一个类从抽象类派生而来，那么当且仅当它实现了基类中的所有纯虚函数，它才能成为非抽象类。**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422145546264.png" alt="image-20230422145546264" style="zoom:80%;" />

# 7、输入输出和模板

## 7.1、输入输出流相关的类

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422150101633.png" alt="image-20230422150101633" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422150238028.png" alt="image-20230422150238028" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422150422490.png" alt="image-20230422150422490" style="zoom:80%;" />

在VS2017中要在开头加入

**define _CRT_SECURE_NO_WARNINGS**

才能用文件的操作  

```c++
  #include<iostream>
using namespace std;

int main() {
	int x, y;
	cin >> x >> y;
	freopen("E:\\C++\\test.txt", "w", stdout);
	if (y == 0) {
		cerr << "error" << endl;
	}
	else {
		cout << x / y;
	}
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422161517818.png" alt="image-20230422161517818" style="zoom:80%;" />

```c++
#include<iostream>
using namespace std;

int main() {
	double f;
	int n;
	freopen("E:\\C++\\test.txt", "r", stdin);
	cin >> f >> n;
	cout << f << "," << n << endl;
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422162622624.png" alt="image-20230422162622624" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422164658310.png" alt="image-20230422164658310" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422171639294.png" alt="image-20230422171639294" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422172011326.png" alt="image-20230422172011326" style="zoom:67%;" />

**EOF的值一般为-1**

```c++
#include<iostream>
using namespace std;

int main() {
	int x;
	char buf[100];
	cin >> x;
	cin.getline(buf, 90);
	cout << buf << endl;
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422175318331.png" alt="image-20230422175318331" style="zoom:67%;" />

## 7.2、流操作算子

真的需要用就查就行

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422224403536.png" alt="image-20230422224403536" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230422224436800.png" alt="image-20230422224436800" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230423000035659.png" alt="image-20230423000035659" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230423000103339.png" alt="image-20230423000103339" style="zoom:80%;" />

## 7.3、文件的读写

- **创建文件**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424190515019.png" alt="image-20230424190515019" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424191133489.png" alt="image-20230424191133489" style="zoom:80%;" />

- **文件名的绝对路径与相对路径**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424191337707.png" alt="image-20230424191337707" style="zoom:80%;" />

- **文件的读写指针**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424192336666.png" alt="image-20230424192336666" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424192956613.png" alt="image-20230424192956613" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230424193030946.png" alt="image-20230424193030946" style="zoom:67%;" />

```C++
#include <fstream>
#include <iostream>
using namespace std;

int main() {
	ofstream fout("E:\\C++\\test.txt", ios::out|ios::in| ios::ate);
	//long location = fout.tellp();
	//cout << location;
	//location = 6;
	//fout.seekp(location);
	//fout.seekp(-9, ios::cur);
	fout << "2";
	fout.seekp(1);
	fout << "8";

	//location = fout.tellp();
	//cout << location;
}
```



| 模式标志   | 描述                                                         |
| :--------- | :----------------------------------------------------------- |
| ios::app   | 追加模式。所有写入都追加到文件末尾。                         |
| ios::ate   | 文件打开后定位到文件末尾。                                   |
| ios::in    | 打开文件用于读取。                                           |
| ios::out   | 打开文件用于覆盖写入。打开之前会清空文件                     |
| ios::trunc | 如果该文件已经存在，其内容将在打开文件之前被截断，即把文件长度设为 0。其实就是先清空 |

**注：**

1. **ios::app追加模式会将写入位置设置为文件的尾部，因此即使调用了 `fout.seekp(location)` 设置了写入位置为第 6 个字节，但由于文件以追加模式打开，写入操作仍然会在文件的尾部进行。**
2. **ios::out|ios::in模式不会清空文件，此时文件指针在文件首部位置，可读可覆盖写入**
3. **ofstream fout("E:\\C++\\test.txt", ios::out|ios::in| ios::ate);首次为追加，接着也可以使用fout.seekp(1);改变文件读写指针的位置**

```c++
#include<iostream>
#include<fstream>
#include<vector>
#include<algorithm>
using namespace std;

int main() {
	vector<int> v;
	ifstream srcFile("E:\\C++\\in.txt", ios::in);
	ofstream destFile("E:\\C++\\out.txt", ios::out);
	int x;
	while (srcFile >> x)
		v.push_back(x);
	sort(v.begin(), v.end());
	for (int i = 0; i < v.size(); i++)
		destFile << v[i] << " ";
	destFile.close();
	srcFile.close();
	return 0;
}
```

**二进制文件读写：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510120023145.png" alt="image-20230510120023145" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510120050174.png" alt="image-20230510120050174" style="zoom:80%;" />

```c++
#include<iostream>
#include<fstream>
using namespace std;

int main() {
	ofstream fout("some.dat", ios::out | ios::binary);
	int x = 120;
	fout.write((const char *)(&x), sizeof(int));
	fout.close();
	ifstream fin("some.dat", ios::in | ios::binary);
	int y;
	fin.read((char *)(&y), sizeof(int));
	fin.close();
	cout << y << endl;
	return 0;
}
```

**文件拷贝：**

```c++
#include<iostream>
#include<fstream>

using namespace std;

int main(int argc, char * argv[]) {
	if (argc != 3) {
		cout << "File name missing!" << endl;
		return 0;
	}
	
	ifstream inFile(argv[1], ios::binary | ios::in);
	if (!inFile) {
		cout << "Source file open error." << endl;
		return 0;
	}

	ofstream outFile(argv[2], ios::binary | ios::out);
	if (!outFile) {
		cout << "New file open error." << endl;
		inFile.close();
		return 0;
	}
	char c;
	while(inFile.get(c)) {
		outFile.put(c);
	}
	outFile.close();
	inFile.close();
	return 0;
}
```

**二进制文件和文本文件的区别：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510171453744.png" alt="image-20230510171453744" style="zoom: 67%;" />

- Unix/Linux下打开文件，用不用 ios::binary 没区别

- Windows下打开文件，如果不用 ios::binary，则:
  **读取文件时，所有的"\r\n" 会被当做一个字符 "\n"处理，即少读了一个字符"\r"。**

  **写入文件时，写入单独的"\n"时，系统自动在前面加一个 "\r"，即多写了一个"\r"**

## 7.4、函数模板

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510203729253.png" alt="image-20230510203729253" style="zoom:67%;" />

**可以根据参数类型推导T的类型：**

```c++
#include<iostream>
using namespace std;

template <class T>

void Swap(T &x, T &y) {
	T tmp = x;
	x = y;
	y = tmp;
}

int main() {
	int n = 1, m = 2;
	Swap(n, m);
	double f = 1.2, g = 2.3;
	Swap(f, g);

	cout << "n的值是：" << n << "，" << "m的值是：" << m << endl;
	cout << "f的值是：" << f << "，" << "g的值是：" << g << endl;

	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510204618035.png" alt="image-20230510204618035" style="zoom:80%;" />

**不通过参数实例化函数模板，直接指定类型是double**

```c++
#include<iostream>
using namespace std;
template <class T>

T Inc(T n) {
	return n + 1;
}

int main() {

	cout << Inc<double>(4) / 2;
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230510210636013.png" alt="image-20230510210636013" style="zoom:80%;" />

 **函数模板可以重载，只要它们的形参表或类型参数表不同即可**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230511112115876.png" alt="image-20230511112115876" style="zoom:67%;" />

**函数模板和函数的次序：**

​	在有多个函数和函数模板名字相同的情况下，编译器如下处理条函数调用语句

1. 先找参数完全匹配的普通函数(非由模板实例化而得的函数)
2. 再找参数完全匹配的模板函数。
3. 再找实参数经过自动类型转换后能够匹配的普通函数
4. 上面的都找不到，则报错。



 **匹配模板函数时，不进行类型自动转换：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230511114440912.png" alt="image-20230511114440912" style="zoom:67%;" />

**myFunction()函数有两个相同的参数T，myFunction(5, 8.4)是两个不同的参数，故会报错**



## 7.5、类模板

**类模板的定义：**

```c++
template <class 类型参数1, class 类型参数2, ...... >//类型参数表
class 类模板名 {
	成员函数和成员变量
};
```

**类模板示例：**

```c++
#include<iostream>
#include <string>
using namespace std;

template <class T1, class T2>
class Pair {
public:
	T1 key;
	T2 value;
	Pair(T1 k, T2 v) :key(k), value(v) {};
	bool operator < (const Pair<T1, T2> & p) const;
};
template  <class T1, class T2>
bool Pair<T1, T2>::operator < (const Pair<T1, T2> & p) const {
	return key < p.key;
}

int main() {

	Pair<string, int> student1("Tom", 19);
	cout << student1.key << " " << student1.value << endl;
	Pair<string, int> student2("U", 23);
	cout << student2.key << " " << student2.value << endl;
	bool z = student1 < student2;
	cout << z << endl;
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230513170119274.png" alt="image-20230513170119274" style="zoom:80%;" />

**类模板与非类型参数：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230513170553719.png" alt="image-20230513170553719" style="zoom:67%;" />

**int size**是非类型参数，**不是代表类型**，是**代表的某个值**

## 7.6、类模板与派生、友元和静态成员变量

**类模板与继承（四种情况）**

- **类模板从类模板派生**

- **类模板从模板类派生**

- **类模板从普通类派生**

- **普通类从模板类派生**



**类模板与友元**

- **函数、类、类的成员函数作为类模板的友元**

- **函数模板作为类模板的友元**

- **函数模板作为类的友元**

- **类模板作为类模板的友元**



**类模板与static成员(注意写法)：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230513184846495.png" alt="image-20230513184846495" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230513184913356.png" alt="image-20230513184913356" style="zoom:67%;" />

# 8、标准模板库STL(一)

## 8.1、string类

string 类是模板类:	**typedef basic_string <char> string;**

使用string类要 包含头文件<string>

string对象的初始化：

```C++
#include<iostream>
#include <string>
using namespace std;

int main() {
	string s1("Hello");
	string s2 = "March";
	string s3(8, 'a');
	
	cout << s1 <<"的长度："<< s1.length() << endl;
	cout << s2 << endl;
	cout << s3 << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230515220716807.png" alt="image-20230515220716807"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230515220806685.png" alt="image-20230515220806685" style="zoom:50%;" />

**string 的赋值和拼接**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230515221607971.png" alt="image-20230515221607971" style="zoom:50%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517110813538.png" alt="image-20230517110813538" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517114644042.png" alt="image-20230517114644042" style="zoom:67%;" />

```c++
#include<iostream>
#include <string>
using namespace std;

int main() {
	string s1("Hello");
	string s2 = "March";
	//string s3(8, 'a');
	cout << s1 + s2 << endl;
	s1.append(s2);
	cout << s1 << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517114738690.png" alt="image-20230517114738690"  />

**比较string（字典序大小）**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517115752608.png" alt="image-20230517115752608" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517115806731.png" alt="image-20230517115806731" style="zoom:67%;" />

**子串：**

```c++
#include<iostream>
#include <string>
using namespace std;

int main() {
	string s1("Hello world");
	string s2 = s1.substr(4, 5);
	cout << s1 << "的长度：" << s1.length() << endl;
	cout << s2 << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517121414101.png" alt="image-20230517121414101"  />

**交换string:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517132743744.png" alt="image-20230517132743744" style="zoom:67%;" />

**寻找string中的字符**

```c++
#include<iostream>
#include <string>
using namespace std;

int main() {
	string s1("Hello worldlo");
	cout << s1 << "的长度：" << s1.length() << endl;
	cout << s1.find("lo") << endl;
	cout << s1.rfind("lo") << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517134726793.png" alt="image-20230517134726793"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517135614526.png" alt="image-20230517135614526" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517135818586.png" alt="image-20230517135818586" style="zoom:67%;" />

**转换成C语言形式的char* 字符串：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517140748470.png" alt="image-20230517140748470" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517140818941.png" alt="image-20230517140818941" style="zoom:67%;" />

**字符串流处理：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517141305607.png" alt="image-20230517141305607" style="zoom:67%;" />



**字符串流处理- 字符串输入流 istringstream**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517141413051.png" alt="image-20230517141413051" style="zoom:67%;" />

**字符串流处理- 字符串输出流 ostringstream**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517141720170.png" alt="image-20230517141720170" style="zoom:67%;" />

**替换：string& replace(int pos, int n, const string& str); //替换从pos开始n个字符为字符串str**

```C++
#include<iostream>
#include <string>

using namespace std;

int main() {
    
    string s1 = "abc123456789";
    string s2 = "abc123456789";
    s2.replace(1, 2, "BCDE");
    cout << s1 << endl;
    cout << s2 << endl;
    
}
```

![image-20230729121934210](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230729121934210.png)

把**s2索引为1**的位置用**"BCDE"**替换掉**2个字符**，也就是把**bc**替换成**BCDE**

**string& replace (const_iterator i1, const_iterator i2, const string& str);**

**用str替换 迭代器起始位置和结束位置的字符**

## 8.2、标准模板库STL概述（一）

**STL中的基本概念：**

**容器:** 可容纳各种数据类型的通用数据结构,是**类模板**

**迭代器:** 可用于依次存取容器中元素，**类似于指针**

**算法:** 用来操作容器中的元素的**函数模板**

- ​	sort()来对一个vector中的数据进行排序

- ​	find()来搜索一个list中的对象

 

举一个简单的例子：

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517181523003.png" alt="image-20230517181523003" style="zoom:67%;" />



**容器概述：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517182336446.png" alt="image-20230517182336446" style="zoom:50%;" />

​		对象被插入容器中时，被插入的是对象的一个复制品。许多算法，比如排序，查找，要求对容器中的元素进行比较，有的容器本身就是排序的，所以，放入容器的对象所属的类，**往往还应该重载 ==和< 运算符。**



**顺序容器简介**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517183350440.png" alt="image-20230517183350440" style="zoom: 67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517184639392.png" alt="image-20230517184639392" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517190310556.png" alt="image-20230517190310556" style="zoom:67%;" />



**关联容器简介：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230517191051487.png" alt="image-20230517191051487" style="zoom:67%;" />

**容器适配器简介：**

 <img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604121239702.png" alt="image-20230604121239702" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604121319238.png" alt="image-20230604121319238" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604121520837.png" alt="image-20230604121520837" style="zoom:80%;" />

- **顺序容器和关联容器中都有的成员函数**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604122725967.png" alt="image-20230604122725967" style="zoom:80%;" />

- **顺序容器的常用成员函数**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604123343453.png" alt="image-20230604123343453" style="zoom:80%;" />

## 8.3、标准模板库STL概述（二）

**迭代器**

- 用于指向顺序容器和关联容器中的元素

- 迭代器用法和指针类似

- 有const 和非 const两种

- 通过迭代器可以读取它指向的元素

- 通过非const迭代器还能修改其指向的元素

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604152024398.png" alt="image-20230604152024398" style="zoom:80%;" />

​		迭代器上可以执行 ++ 操作，以使其指向容器中的下一个元素。如果迭代器到达了容器中的最后一个元素的后面，此时再使用它，就会出错，类似于使用NULL或未初始化的指针一样。

**迭代器实例：**

```c++
#include<vector>
#include<iostream>
using namespace std;

int main() {

	vector<int> v;//一个存放int元素的数组，一开始里面没有元素
	v.push_back(1);
	v.push_back(2);
	v.push_back(3);
	v.push_back(4);

	vector<int> ::const_iterator i;
	for (i = v.begin(); i != v.end(); ++i)
		cout << *i << ",";
	cout << endl;
    
	vector<int>::reverse_iterator r;//反向迭代器
	for (r = v.rbegin(); r != v.rend(); r++) 
		cout << *r << ",";
	cout << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604155328488.png" alt="image-20230604155328488" style="zoom:80%;" />

**迭代器类别**

**双向迭代器：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604155732012.png" alt="image-20230604155732012" style="zoom:80%;" />

**随机访问迭代器：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604160021180.png" alt="image-20230604160021180" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604160856846.png" alt="image-20230604160856846" style="zoom:80%;" />

**随机访问迭代器举例：**

```c++
#include<vector>
#include<iostream>
using namespace std;

int main() {

	vector<int> v;
	v.push_back(1);
	v.push_back(2);
	v.push_back(3);
	v.push_back(4);
    
	int j;
	for (j = 0; j < v.size(); j++) {
		cout << v[j] << ",";
	}
	cout << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604163528222.png" alt="image-20230604163528222" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604163653816.png" alt="image-20230604163653816" style="zoom:80%;" />

**算法简介：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604164341852.png" alt="image-20230604164341852" style="zoom:80%;" />

**算法示例：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604170611698.png" alt="image-20230604170611698" style="zoom:80%;" />

```c++
#include<vector>
#include<iostream>
#include<algorithm>
using namespace std;

int main() {

	vector<int> v;
	v.push_back(1);
	v.push_back(2);
	v.push_back(3);
	v.push_back(4);

	vector<int> ::iterator p;
	p = find(v.begin(), v.end(), 3);
	cout << *p << endl;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604170712842.png" alt="image-20230604170712842" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604171718569.png" alt="image-20230604171718569" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604171739683.png" alt="image-20230604171739683" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604172521836.png" alt="image-20230604172521836" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604172540405.png" alt="image-20230604172540405" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230604172603869.png" alt="image-20230604172603869" style="zoom:80%;" />

## 8.4、vector,depue和list

**vector 示例程序**  

```c++
#include<iostream>
#include<vector>
using namespace std;

template<class T>
void PrintVector(T s, T e) {
	for (; s != e; ++s)
		cout << *s << "";
	cout << endl;
}

int main() {
	int a[5] = { 1,2,3,4,5 };
	vector<int> v(a, a + 5);
	cout << "1) " << v.end() - v.begin() << endl;
	cout << "2) "; PrintVector(v.begin(), v.end());
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230605153407297.png" alt="image-20230605153407297" style="zoom:80%;" />

**depue:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230605154849115.png" alt="image-20230605154849115" style="zoom: 67%;" />

**双向链表list:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230607120634621.png" alt="image-20230607120634621" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230607120659275.png" alt="image-20230607120659275" style="zoom:80%;" />

## 8.5、函数对象

**若一个类重载了运算符“()”,则该类的对象就成为函数对象**

```c++
#include<iostream>
using namespace std;

class CMyAverge {
public:
	double operator()(int a1, int a2 ,int a3) {
		//重载()运算符
		return (double)(a1 + a2 + a3)/3;
	}
};

int main() {
	CMyAverge average;//函数对象
	cout << average(3, 2, 3);//average.operator()(3, 2, 3)
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609135559210.png" alt="image-20230609135559210" style="zoom:80%;" />

**函数对象的应用：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163015973.png" alt="image-20230609163015973" style="zoom:80%;" />



<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163036917.png" alt="image-20230609163036917" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163056648.png" alt="image-20230609163056648" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163116490.png" alt="image-20230609163116490" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163136490.png" alt="image-20230609163136490" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163158710.png" alt="image-20230609163158710" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163222088.png" alt="image-20230609163222088" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163249312.png" alt="image-20230609163249312" style="zoom:80%;" />

**STL的函数对象类模板：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609163450243.png" alt="image-20230609163450243"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609165323251.png" alt="image-20230609165323251"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609165345623.png" alt="image-20230609165345623"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609165408667.png" alt="image-20230609165408667" style="zoom: 80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609165434270.png" alt="image-20230609165434270"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609170832376.png" alt="image-20230609170832376"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230609170903805.png" alt="image-20230609170903805"  />

# 9、标准模板库STL(二)

## 9.1、set和multiset

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610101517999.png" alt="image-20230610101517999"  />



**pair模板：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610103722285.png" alt="image-20230610103722285"  />



**multiset：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610105845059.png" alt="image-20230610105845059"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610113534252.png" alt="image-20230610113534252"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610113600209.png" alt="image-20230610113600209" style="zoom: 67%;" />

```c++
#include<set>
using namespace std;
class A{};
int main() {
	multiset<A> a;
	a.insert(A());//error
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230610115544212.png" alt="image-20230610115544212"  />

```c++
#include<iostream>
#include<set>
using namespace std;

template <class T>
void Print(T first, T last) {
	for (; first != last; ++first)
		cout << *first << " ";
	cout << endl;
}

class A {
private:
	int n;
public:
	A(int n_) {n = n_;}
	friend bool operator < (const A & a1, const A & a2) {
		return a1.n < a2.n;
	}
	friend ostream & operator<<(ostream & o, const A & a2) {
		o << a2.n;
		return o;
	}
	friend class MyLess;
};

struct MyLess {
	bool operator()(const A & a1, const A & a2) {
		//按个位数比较大小
		return (a1.n % 10) < (a2.n % 10);
	}
};

typedef multiset<A> MSET1;//MSET1用"<"比较大小
typedef multiset<A, MyLess> MSET2;//MSET2用MyLess::operator()比较大小
int main() {
	const int SIZE = 6;
	A a[SIZE] = { 4,22,19,8,33,40 };
	MSET1 m1;
	m1.insert(a, a + SIZE);
	m1.insert(22);
	cout << "1)" << m1.count(22) << endl;
	cout << "2)";
	Print(m1.begin(), m1.end());
	//m1的元素：4 8 19 22 22 33 40
	MSET1::iterator pp = m1.find(19);
	if (pp != m1.end()) {//条件为真说明找到
		cout << "found" << endl;
	}
	cout << "3)" << *m1.lower_bound(22) << "," << *m1.upper_bound(22) << endl;
	//3)22,33
	pp = m1.erase(m1.lower_bound(22), m1.upper_bound(22));
	// 删除一个区间内的所有元素
	// st.erase(frist, last)即删除[frist, last)内的元素
	// pp指向被删元素的下一个元素
	cout << "4)";Print(m1.begin(), m1.end());
	//4)4 8 19 33 40
	cout << "5)"; cout << *pp << endl;
	//5)33

	MSET2 m2;
	m2.insert(a, a + SIZE);
	cout << "6) "; Print(m2.begin(), m2.end());
	return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230612164502024.png" alt="image-20230612164502024"  />



**set:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230612171020463.png" alt="image-20230612171020463"  />

```c++
#include<iostream>
#include<set>
using namespace std;

int main() {
	typedef set<int>::iterator IT;
	int a[5] = { 1,2,3,4,6 };
	set<int> st(a, a + 5);
	pair<IT, bool> result;
	result = st.insert(5);//用pair去接收insert后的返回值，用来判断插入是否成功
	cout << result.second << endl;
	
	result = st.insert(5);
	cout << result.second << endl;
	pair<IT, IT> bounds = st.equal_range(4);
	cout << *bounds.first << "," << *bounds.second;

}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230612173308078.png" alt="image-20230612173308078"  />



## 9.2、map和multimap

**multimap：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613154812382.png" alt="image-20230613154812382"  />

```c++
#include <iostream>
#include <map>
using namespace std;

int main()
{
    typedef multimap<int, double> mmid;
    mmid pairs;
    cout << "1)" << pairs.count(15) << endl;
    pairs.insert(mmid::value_type(15, 2.7));
    pairs.insert(mmid::value_type(15, 99.3));
    cout << "2)" << pairs.count(15) << endl;
    pairs.insert(mmid::value_type(30, 111.11));
    pairs.insert(mmid::value_type(10, 22.22));
    pairs.insert(mmid::value_type(25, 33.333));
    pairs.insert(mmid::value_type(20, 9.3));
    
    for (mmid::const_iterator i = pairs.begin(); i != pairs.end(); i++)
    {
        cout << "(" << (*i).first << "," << (*i).second << ")"
             << ",";
    }
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613164314097.png" alt="image-20230613164314097"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613175706739.png" alt="image-20230613175706739"  />

```c++
#include <iostream>
#include <map>
#include <string>
using namespace std;

class CStudent
{
public:
    struct CInfo
    {
        int id;
        string name;
    };
    int score;
    CInfo info; // 学生的其他信息
};

typedef multimap<int, CStudent::CInfo> MAP_STD;
int main()
{
    MAP_STD mp;
    CStudent st;
    string cmd;
    while (cin >> cmd)
    {
        if (cmd == "Add")
        {
            cin >> st.info.name >> st.info.id >> st.score;
            mp.insert(MAP_STD::value_type(st.score, st.info));
        }
        else if (cmd == "Query")
        {
            int score;
            cin >> score;
            MAP_STD::iterator p = mp.lower_bound(score);
            if (p != mp.begin())
            {
                --p;
                score = p->first;
                MAP_STD::iterator maxp = p;
                int maxId = p->second.id;
                for (; p != mp.begin() && p->first == score; --p)
                {
                    if (p->second.id > maxId)
                    {
                        maxp = p;
                        maxId = p->second.id;
                    }
                }
                if (p->first == score)
                {
                    if (p->second.id > maxId)
                    {
                        maxp = p;
                        maxId = p->second.id;
                    }
                }
                cout << maxp->second.name << " " << maxp->second.id << " " << maxp->first << endl;
            }
            else
            {
                cout << "Nobody" << endl;
            }
            /* code */
        }
        
    }
    return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613175804446.png" alt="image-20230613175804446" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613180203768.png" alt="image-20230613180203768" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230613180229363.png" alt="image-20230613180229363" style="zoom:67%;" />

## 9.3、容器适配器

**Stack:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614112605549.png" alt="image-20230614112605549" style="zoom: 67%;" />



**queue:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614113305162.png" alt="image-20230614113305162" style="zoom:67%;" />



**priority_queue:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614113908122.png" alt="image-20230614113908122" style="zoom:67%;" />

```c++
#include <queue>
#include <iostream>
using namespace std;

int main()
{
    priority_queue<double> pq1;
    pq1.push(3.2);
    pq1.push(9.8);
    pq1.push(9.8);
    pq1.push(5.4);
    while (!pq1.empty())
    {
        cout << pq1.top() << " ";
        pq1.pop();
        /* code */
    }
    cout << endl;

    priority_queue<double, vector<double>, greater<double>> pq2;
    pq2.push(3.2);
    pq2.push(9.8);
    pq2.push(9.8);
    pq2.push(5.4);
    while (!pq2.empty())
    {
        cout << pq2.top() << " ";
        pq2.pop();
        /* code */
    }
    return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614143719740.png" alt="image-20230614143719740"  />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614143744612.png" alt="image-20230614143744612" style="zoom:67%;" />

## 9.4、适配器

### 9.4.1、函数对象适配器

**把两个参数绑定在一起**

```C++
#include<iostream>
#include "vector"
#include "functional"

using namespace std;

class MyPrint : public binary_function<int, int, void> {
public:
    void operator()(int val, int start) const {
        cout << val + start << endl;
    }
};

void test01() {
    vector<int> v;
    for (int i = 0; i < 10; i++) {
        v.push_back(i);
    }
    cout << "请输入起始累加值：" << endl;
    int num;
    cin >> num;

    //函数对象适配器
    for_each(v.begin(), v.end(), bind2nd(MyPrint(),num));

}


int main() {

    test01();

    return 0;
}
```

![image-20230729160421696](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230729160421696.png)

bind2nd(MyPrint(),num)将**两个参数绑定在一起**

**//1、利用bind2nd 进行绑定**

**//2、继承 public binary_function<参数1 类型,参数2类型,返回值类型>**

**//3、加const**



### **9.4.2、取反适配器（把大于取反为小于）**

```C++
class GreaterThan5: public unary_function<int, bool>{
public:
    bool operator()(int val) const {
        return val > 5;
    }
};

void test02() {
    vector<int> v;
    for (int i = 0; i < 10; i++) {
        v.push_back(i);
    }
    //一元取反
    auto pos = find_if(v.begin(), v.end(), not1(GreaterThan5()));
    if (pos != v.end()) {
        cout << "找到小于5的值为：" << *pos << endl;
    } else {
        cout << "没有找到小于5的值" << endl;
    }
};


int main() {

    test02();
    return 0;
}
```



**//1、利用not1进行取反**

**//2、继承 public unary_function<int,bool>**

**//3、加const**

![image-20230729162745434](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230729162745434.png)



另一种写法：

```c++
void test02() {
    vector<int> v;
    for (int i = 0; i < 10; i++) {
        v.push_back(i);
    }
    int num;
    cout << "输入你想要查找的小于的数" << endl;
    cin >> num;
    //一元取反
//    auto pos = find_if(v.begin(), v.end(), not1(GreaterThan5()));
    auto pos = find_if(v.begin(), v.end(), not1(bind2nd(greater<int>(), num)));
    if (pos != v.end()) {
        cout << "找到小于5的值为：" << *pos << endl;
    } else {
        cout << "没有找到小于5的值" << endl;
    }
};
```

![image-20230729164156116](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230729164156116.png)

**3、二元取反**

```c++
class Print  {
public:
    void operator()(int val) {
        cout << val << endl;
    }
};

void  test03(){
    vector<int> v;
    v.push_back(10);
    v.push_back(30);
    v.push_back(20);
    v.push_back(50);
    v.push_back(40);

    sort(v.begin(),v.end(),not2(less<int>()));

    for_each(v.begin(),v.end(),Print());

}


int main() {
    test03();
    return 0;
}
```

![image-20230802114945712](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230802114945712.png)

原本是从小到大排序，取反后从大到小排序

### 9.4.3、函数适配器

**将函数指针 适配成函数对象 ptr_fun**

```C++
void myPrint3(int val, int start) {
    cout << val + start << endl;
}

void test03() {
    vector<int> v;
    v.push_back(10);
    v.push_back(30);
    v.push_back(20);
    v.push_back(50);
    v.push_back(40);

    //将函数指针 适配成函数对象 ptr_fun
    for_each(v.begin(), v.end(), bind2nd(ptr_fun(myPrint3), 1000));

}


int main() {
    test03();
    return 0;
}
```

### 9.4.4、成员函数适配器

**如果存放的是对象实体   mem_fun_ref**
**如果存放的是对象指针   mem_fun**

```C++
class Person {
public:
    string name;
    int age;

    Person(string name, int age) : name(name), age(age) {}
    void showPerson(){
        cout << "成员函数打印--姓名：" << this->name << "年龄：" << this->age << endl;
    }


};

void myPrint4(Person &person) {

    cout << "姓名：" << person.name << "年龄：" << person.age << endl;
}

void test04() {
    vector<Person> v;
    Person p1("aaa", 10);
    Person p2("bbb", 20);
    Person p3("ccc", 30);
    Person p4("ddd", 40);
    v.push_back(p1);
    v.push_back(p2);
    v.push_back(p3);
    v.push_back(p4);

    //利用mem_fun_ref
    for_each(v.begin(), v.end(), mem_fun_ref(&Person::showPerson));

}

int main() {
    test04();
    return 0;
}
```

将成员函数适配成全局函数，这里就是将showPerson()适配成myPrint4(Person &person)

# 10、C++11特性

**统一的初始化方法：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614151712295.png" alt="image-20230614151712295" style="zoom:67%;" />

**成员变量默认初始值：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614151815168.png" alt="image-20230614151815168" style="zoom:67%;" />

**auto关键字：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614151925584.png" alt="image-20230614151925584" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614153249941.png" alt="image-20230614153249941" style="zoom:67%;" />

**decltype关键字：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614153558844.png" alt="image-20230614153558844" style="zoom:67%;" />

用括号()括起来后，返回值变成**引用**



**智能指针shared_ptr:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614163259827.png" alt="image-20230614163259827" style="zoom:67%;" />

```c++
#include <memory>
#include <iostream>
using namespace std;

struct A
{
    int n;
    A(int v = 0) : n(v) {}
    ~A()
    {
        cout << n << " destructor" << endl;
    }
};
int main()
{
    shared_ptr<A> sp1(new A(2));//sp1托管A(2)
    shared_ptr<A> sp2(sp1);//sp2也托管A(2)
    shared_ptr<A> sp3(new A(3));//sp3托管A(3)

    sp1.reset();
    sp2.reset();

    cout << "aaaa" << endl;
}
```



**空指针nullptr:**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614190505529.png" alt="image-20230614190505529" style="zoom:67%;" />



**右值引用和move语义：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230614194517513.png" alt="image-20230614194517513" style="zoom:67%;" />

```C++
#include <iostream>
#include <string>
#include <cstring>

using namespace std;
class String
{
public:
    char *str;
    String() : str(new char[1]) { str[0] = 0; }
    String(const char *s)
    {
        str = new char[strlen(s) + 1];
        strcpy(str, s);
    }

    

    String(const String &s)
    {
        cout << "copy constructor called" << endl;
        str = new char[strlen(s.str) + 1];
        strcpy(str, s.str);
    }

    String &operator=(const String &s)
    {
        cout << "copy operator= called" << endl;
        if (str != s.str)
        {
            delete[] str;
            str = new char[strlen(s.str) + 1];
            strcpy(str, s.str);
        }
        return *this;
    }

    String(String &&s) : str(s.str)
    {
        cout << "move constructor called" << endl;
        s.str = new char[1];
        s.str[0] = 0;
    }

    String &operator=(String &&s)
    {
        cout << "move operator = called" << endl;
        if (str != s.str)
        {
            delete[] str;
            str = s.str;
            s.str = new char[1];
            s.str[0] = 0;
        }
        return *this;
    }

    ~String() { delete[] str; }
};

template <class T>
void MoveSwap(T &a, T &b)
{
    T tmp(move(a));//std::move(a)为右值，这里会调用move constructor
    a = move(b);//move(b)为右值，因此这里会调用move assigment
    b = move(tmp);//move(tmp)为右值，因此这里会调用move assigment
}

int main()
{
    String s;
    s = String("ok");
    cout << "***********" << endl;
    String &&r = String("this");
    cout << r.str << endl;
    String s1 = "hello", s2 = "world";
    MoveSwap(s1, s2);
    cout << s2.str << endl;
    return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615165109565.png" alt="image-20230615165109565" style="zoom:80%;" />

**哈希表：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615170247086.png" alt="image-20230615170247086" style="zoom:67%;" />

**正则表达式：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615172242036.png" alt="image-20230615172242036" style="zoom:67%;" />



**Lambda表达式：**

只使用一次的函数对象，能否不要专门为其编写一个类?
只调用一次的简单函数，能否在调用时才写出其函数体?

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615180546495.png" alt="image-20230615180546495" style="zoom:67%;" />

```c++
#include <iostream>

using namespace std;

int main()
{
    int x = 100, y = 200, z = 300;
    cout << [](double a, double b)
    { return a + b; }(1.2, 2.5)
         << endl;
    auto ff = [=, &y, &z](int n)
    {
        cout << x << endl;
        y++;
        z++;
        return n * n;
    };

    cout << ff(15) << endl;
    cout << y << "," << z << endl;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615181531081.png" alt="image-20230615181531081" style="zoom:80%;" />

```c++
#include <iostream>
#include <algorithm>

using namespace std;

int main() {
    int a[4] = {4,2,11,33};
    sort(a, a+4, [](int x, int y){return  x%10 < y%10;});
    for_each(a,a+4,[](int x){cout << x << " ";});
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615203359319.png" alt="image-20230615203359319" style="zoom:80%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230615204643438.png" alt="image-20230615204643438" style="zoom: 67%;" />



**强制类型转换：**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616134206898.png" alt="image-20230616134206898" style="zoom:67%;" />

1. **static_cast**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616135229342.png" alt="image-20230616135229342" style="zoom: 67%;" />

```c++
#include <iostream>
using namespace std;
class A{
public:
    operator int(){return 1;}
    operator char *(){return NULL;}
};

int main() {
    A a;
    int n;
    char * p = "New Dragob Inn";
    n = static_cast<int>(3.14);//n的值变为3
    n = static_cast<int>(a);//调用a.operator int, n的值变为 1
    p = static_cast<char*>(a);//调用a.operator int *,p的值变为 NULL
    return 0;
}
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616152226552.png" alt="image-20230616152226552" style="zoom:67%;" />

2. **reinterpret_cast**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616153132981.png" alt="image-20230616153132981" style="zoom:67%;" />

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616161507207.png" alt="image-20230616161507207" style="zoom: 80%;" />

3. **const_cast**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616162346684.png" alt="image-20230616162346684" style="zoom:67%;" />

4. **dynamic_cast**

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616164536284.png" alt="image-20230616164536284" style="zoom:67%;" />

假设我们有一个简单的类层次结构，由基类 `Shape` 和两个派生类 `Circle` 和 `Rectangle` 组成。每个类都有一个公共的成员函数 `Draw()`。

```c++
cppCopy codeclass Shape {
public:
    virtual void Draw() {
        cout << "Drawing a shape." << endl;
    }
};

class Circle : public Shape {
public:
    void Draw() {
        cout << "Drawing a circle." << endl;
    }
    void Special(){
        cout << "Circle Special";
    }
};

class Rectangle : public Shape {
public:
    void Draw() {
        cout << "Drawing a rectangle." << endl;
    }
};
```

现在，我们创建了一个指向基类的指针，并将其指向一个派生类的对象：

```c++
Shape* shapePtr = new Circle();
```

我们可以使用 `dynamic_cast` 将基类指针转换为派生类指针。如果转换成功，我们就可以调用派生类特有的函数或访问派生类的成员：

```c++
    shapePtr->Draw();
    Circle* circlePtr = dynamic_cast<Circle*>(shapePtr);
    if (circlePtr) {
        // 转换成功，可以访问 Circle 类的成员函数
        circlePtr->Special();
    } else {
        // 转换失败，shapePtr 并不指向 Circle 类的对象
        cout << "Failed to cast to Circle." << endl;
    }
```

<img src="C++%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20230616165607790.png" alt="image-20230616165607790"  />

# 11、异常

## 11.1、异常的使用方式

```c++
#include<iostream>
#include "exception"

using namespace std;

class MyException1 : public exception {

public:
    const char *what() const throw() {
        return "自定义的异常";
    }

};

int main() {
    try{
        throw MyException1();
    }catch (MyException1 &e){
        cout << "输出:";
        cout << e.what()<<endl;
    }
    return 0;
}
```

<img src="C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230727133841258.png" alt="image-20230727133841258"  />

```c++
try {
    throw "发生了异常";
} catch (const char * i) {
    cout << i << endl;
}
```

![image-20230727133909230](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230727133909230.png)

## 11.2、栈解旋

​		从try代码块开始，到throw抛出异常之前，所有栈上的数据都会被释放掉，释放的顺序和创建顺序相反的，这个过程我们称为**栈解旋**

```c++
#include<iostream>
#include "exception"
#include "string"
using namespace std;

class MyException1 : public exception {
public:
    const char *what() const throw() {
        return "自定义的异常";
    }

};

class Person{
public:
    int i;

    Person(int a){
        i=a;
        cout<<"Person的默认构造函数调用："<< i <<endl;
    }
    ~Person(){
        cout<<"Person的析构函数调用："<< i <<endl;
    }
};

int MyDivision(int a, int b){
    if(b == 0){
        Person p1(1);
        Person p2(2);

        throw MyException1();
    } else{
        return a/b;
    }
}

int main() {
    try{
        MyDivision (1,0);

    }catch (MyException1 &e){
        cout << "输出:";
        cout << e.what()<<endl;
    }
    return 0;
}
```

![image-20230727151627340](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230727151627340.png)

## 11.3、异常的抛出与接收

```c++
class MyException
{
public:
	MyException()
	{
		cout << "MyException默认构造函数调用" << endl;
	}
	MyException(const MyException& e)
	{
		cout << "MyException拷贝构造函数调用" << endl;
	}
	~MyException()
	{
		cout << "MyException析构函数调用" << endl;
	}
};

void doWork()
{
	throw  new MyException();
}

void test01()
{
	try
	{
		doWork();
	}
	//抛出的是 throw MyException();  catch (MyException e) 调用拷贝构造函数 效率低
	//抛出的是 throw MyException();  catch (MyException &e)  引用方式接收，只调用默认构造函数 效率高 推荐
	//抛出的是 throw &MyException(); catch (MyException *e) 对象方式接收，对象会提前释放掉，不能在非法操作
	//抛出的是 new MyException();   catch (MyException *e) 只调用默认构造函数 自己要管理释放
	catch (MyException* e)  //这里根据上面的可选不同方法，这里是最后一种
	{
		cout << "自定义异常捕获" << endl;
		delete e;
	}
}
```

**推荐：抛出的是 throw MyException();  catch (MyException &e)**  

## 11.4、异常的多态使用

```c++
#include<iostream>
#include "exception"
#include "string"

using namespace std;

class BaseException {
public:
    virtual void printError() = 0;//纯虚函数
};

//空指针异常
class NULLPointerException : public BaseException {
public:
    virtual void printError() {
        cout << "空指针异常" << endl;
    }
};

//越界异常
class OutOfRangeException : public BaseException {
public:
    virtual void printError() {
        cout << "越界异常" << endl;
    }
};

void doWork() {
//    throw NULLPointerException();
    throw OutOfRangeException();
}

void test01() {
    try {
        doWork();
    } catch (BaseException &e) {
        e.printError();
    }
}

int main() {
    test01();
    return 0;
}
```

![image-20230727165513492](C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230727165513492.png)

抛出的异常对象不同，但是都用**父类BaseException进行接收**，体现了多态的作用

## 11.5、系统标准异常

C++ 提供了一系列标准的异常，定义在 **<exception>** 中，我们可以在程序中使用这些标准的异常。它们是以父子类层次结构组织起来的，如下所示：

<img src="C:\Users\13771\AppData\Roaming\Typora\typora-user-images\image-20230727171700381.png" alt="image-20230727171700381" style="zoom:80%;" />

下表是对上面层次结构中出现的每个异常的说明：

| 异常                   | 描述                                                         |
| :--------------------- | :----------------------------------------------------------- |
| **std::exception**     | 该异常是所有标准 C++ 异常的父类。                            |
| std::bad_alloc         | 该异常可以通过 **new** 抛出。                                |
| std::bad_cast          | 该异常可以通过 **dynamic_cast** 抛出。                       |
| std::bad_typeid        | 该异常可以通过 **typeid** 抛出。                             |
| std::bad_exception     | 这在处理 C++ 程序中无法预期的异常时非常有用。                |
| **std::logic_error**   | 理论上可以通过读取代码来检测到的异常。                       |
| std::domain_error      | 当使用了一个无效的数学域时，会抛出该异常。                   |
| std::invalid_argument  | 当使用了无效的参数时，会抛出该异常。                         |
| std::length_error      | 当创建了太长的 std::string 时，会抛出该异常。                |
| std::out_of_range      | 该异常可以通过方法抛出，例如 std::vector 和 std::bitset<>::operator[]()。 |
| **std::runtime_error** | 理论上不可以通过读取代码来检测到的异常。                     |
| std::overflow_error    | 当发生数学上溢时，会抛出该异常。                             |
| std::range_error       | 当尝试存储超出范围的值时，会抛出该异常。                     |
| std::underflow_error   | 当发生数学下溢时，会抛出该异常。                             |











































