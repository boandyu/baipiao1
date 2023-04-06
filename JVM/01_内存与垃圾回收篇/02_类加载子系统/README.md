##                              类加载子系统

### 一、 概述

![img](images/jvm整体结构.png)

![image-20200705080911284](images/image-20200705080911284.png)

### 二、类加载子系统作用：

​	类加载子系统负责从文件系统或者网络中加载 **Class** 文件，**class** 文件在文件开头有特定的文件标识。

**ClassLoader** 只负责 **class** 文件的加载，至于它是否可以运行，则由 **Execution Engine** 决定。

加载的类信息存放于一块称为方法区的内存空间。除了类的信息外，方法区中还会存放运行时常量池信息，可能还包括字符串字面量和数字常量（这部分常量信息是Class文件中常量池部分的内存映射）

![截屏2021-12-03 下午12.18.21](images/21.png)

- class file 存在于本地硬盘上，可以理解为设计师在画在纸上的模板，而最终这个模板在执行的时候是要加载到JVM当中来根据这个文件实例化出 n 个一模一样的实例。
- class file 加载到 JVM 中，被称为 DNA 元数据模板，放在方法区。
- 在.class 文件 -> JVM -> 最终成为元数据模板，此过程就要一个运输工具（类装载器 Class Loader），扮演一个快递员的角色。

![截屏2021-12-03 下午12.25.33](images/33.png)

### 三、类的加载过程

例如下面的一段简单的代码：

```java
public class ClassLoader {
    public static void main(String[] args) {
        System.out.println("我已经被加载了");
    }
}
```

它的加载过程是怎么样的呢？

![截屏2021-12-03 下午2.31.35](images/35.png)

完整的流程如图下所示

![截屏2021-12-03 下午2.33.18](images/18.png)

#### 3.1 加载阶段

通过一个类的全限定名获取定义此类的二进制字节流

将这个字节流所代表的静态存储结构转化为方法区的运行时数据结构

在内存中生成一个代表这个类的java.lang.Class对象，作为方法区这个类的各种数据的访问入口

#### 3.2 加载class文件的方式

- 从本地系统中直接加载
- 通过网络获取，典型场景：Web Applet
- 从zip压缩包中读取，成为日后jar、war格式的基础
- 运行时计算成，使用最多的是：动态代理技术
- 由其他文件生成，典型场景：JSP应用从专有数据库中提取.class文件，比较少见
- 从加密文件中获取，典型的防Class文件被反编译的保护措施

#### 3.3 连接阶段

##### 3.3.1 验证 Verify

目的在于确保Class文件的字节流中包含信息符合当前虚拟机的要求，保证被加载类的正确性，不危害虚拟机自身安全。

主要包括四种验证，文件格式验证，元数据验证，字节码验证，符号引用验证。

> 工具：Binary Viewer 查看

![截屏2021-12-03 下午7.23.17](images/17.png)

如果出现不合法的字节码文件，那么将会验证不通过

同时我们可以通过安装IDEA的插件，来查看我们的Class文件

![截屏2021-12-03 下午7.26.05](images/05.png)

安装完成之后，我们编译完一个class文件后，点解view即可显示我们安装的插件来查看字节码方法了

![截屏2021-12-03 下午7.28.34](images/34.png)

##### 3.3.2 准备 Prepare

为类变量分配内存并且设置该类变量的默认初始值，即零值。

```java
public class ClassLoader {
    private static int a = 1; // 准备阶段为0，在下一个阶段，也就是初始化的时候才为1
    public static void main(String[] args) {
        System.out.println(a);
    }
}
```

上面的变量a在准备阶段会赋初始值，但不是1，而是0。

这里不包含用final修饰的static，因为final在编译的时候就会分配了，准备阶段会显式初始化。

这里不会为实例变量分配初始化，类变量会分配在方法区中，而实例变量是会随着对象一起分配到Java堆中。

例如下面这段代码：

##### 3.3.3 解析 Resolve

将常量池内的符号引用转换为直接引用的过程。

事实上，解析操作往往会伴随着JVM在执行完初始化之后在执行。

符号引用就是一组符号来描述所引用的目标。符号引用的字面量形式明确定义在《java虚拟机规范》的class文件格式中。直接引用就解析动作主要针对类或接口，字段，类方法，接口方法，方法类型等。对应常量池中CONSTANT Class info、CONSTANT Fieldref info、CONSTANT Methodref info 等。

##### 3.3.4 初始化阶段

初始化阶段就是执行类构造器法 ()的过程。

此方法不需要定义，是javac编译器自动收集类中的所有类变量的赋值动作和静态代码块中的语句合并而来。

- 也就是说，当我们代码中包含static变量的时候，就会有clinit方法

构造器方法中制定按语句在源文件中出现的顺序执行。

（）不同于类的构造器。（关联：构造器是虚拟机视角下的（））若该类具有父类。JVM保证子类的（）执行钱，父类的（）已经执行完毕。

- 任何一个类在声明后，都有生成一个构造器，默认是空参构造器

```java
public class ClassInitTest {
    private static int num = 1;
    static {
        num = 2;
        number = 20;
        System.out.println(num);
        System.out.println(number);  //报错，非法的前向引用
    }

    private static int number = 10;

    public static void main(String[] args) {
        System.out.println(ClassInitTest.num); // 2
        System.out.println(ClassInitTest.number); // 10
    }
}
```

关于涉及到父类时候的变量赋值过程

```java
public class ClinitTest1 {
    static class Father {
        public static int A = 1;
        static {
            A = 2;
        }
    }

    static class Son extends Father {
        public static int b = A;
    }

    public static void main(String[] args) {
        System.out.println(Son.b);
    }
}
```

我们输出结果为2，也就是说首先加载ClinitTest1的时候，会找到main方法，然后执行Son的初始化，但是Son集成了Father，因此还需要执行Father的初始化，同时将A赋值为2。我们通过反编译得到Father的加载过程，首先我们看到原来的值被赋值成1，然后又被复制成2，最后返回

```java
iconst_1
putstatic #2 <com/atguigu/java/chapter02/ClinitTest1$Father.A>
iconst_2
putstatic #2 <com/atguigu/java/chapter02/ClinitTest1$Father.A>
return
```

虚拟机必须保证一个类的（） 方法在多线程下被同步加锁。

```java
public class DeadThreadTest {
    public static void main(String[] args) {
        new Thread(() -> {
            System.out.println(Thread.currentThread().getName() + "\t 线程t1开始");
            new DeadThread();
        }, "t1").start();

        new Thread(() -> {
            System.out.println(Thread.currentThread().getName() + "\t 线程t2开始");
            new DeadThread();
        }, "t2").start();
    }
}
class DeadThread {
    static {
        if (true) {
            System.out.println(Thread.currentThread().getName() + "\t 初始化当前类");
            while(true) {

            }
        }
    }
}
```

上面的代码，输出结果为

```java
t1	 线程t1开始
t2	 线程t2开始
t1	 初始化当前类
```

从上面可以看初始化后，只能够执行一次初始化，这也就是同步加锁的过程

### 四、类加载器的分类

​	JVM 支持两种类型的类加载器。分别为引导类加载器（Bootstrap ClassLoader）和自定义类加载器（User-Defined ClassLoader）。

从概念上来讲，自定义类加载器一般指的是程序中开发人员自定义的一类类加载器，但是Java虚拟机规范却没有这么定义，而是将所有派生抽象类ClassLoader的类加载器划分为自定义类加载器。

无论类加载器的类型如何划分，在程序中我们最常见的类加载器始终只有3个，如下所示：

![截屏2021-12-03 下午11.04.48](images/48.png)

这里的四者之间是包含关系，不是上层和下层，也不是子系统的继承关系。

我们通过一个类，获取它不同的加载器

```java
public class ClassLoaderTest {
    public static void main(String[] args) {
        // 获取系统类加载器
        ClassLoader systemClassLoader = ClassLoader.getSystemClassLoader();
        System.out.println(systemClassLoader);
        // 获取其上层的类加载器
        ClassLoader extClassLoader = systemClassLoader.getParent();
        System.out.println(extClassLoader);
        // 视图获取  根加载器
        ClassLoader bootstrapClassLoader = extClassLoader.getParent();
        System.out.println(bootstrapClassLoader);
        // 获取自定义类加载器
        ClassLoader classLoader = ClassLoaderTest.class.getClassLoader();
        System.out.println(classLoader);
        // 获取string类型的加载器
        ClassLoader strClassLoader = String.class.getClassLoader();
        System.out.println(strClassLoader);
    }
}
```

```java
jdk.internal.loader.ClassLoaders$AppClassLoader@2c13da15
jdk.internal.loader.ClassLoaders$PlatformClassLoader@c818063
null
jdk.internal.loader.ClassLoaders$AppClassLoader@2c13da15
null
```

#### 4.1 虚拟机自带的类加载器

#### 4.2 启动类加载器（引导类加载器，Bootstrap ClassLoader）

- 这个类加载器使用C/C++语言实现的，嵌套在JVM内部。
- 它用来加载器Java的核心库（JAVAHOME/jre/1ib/rt.jar、resources.jar或sun.boot.class.path路径下的内容），用于提供JVM自身需要的类
- 并不继承自java.lang.ClassLoader，没有父加载器。
- 加载扩展类和应用类程序类加载器，并指定为他们的父类加载器。
- 处于安全考虑，Bootstrap 启动类加载器只加载名为java、javax、sun等开头的类。

#### 4.3 扩展类加载器（Extension ClassLoader）

- Java 语言编写，由sun.misc.Launcher$ExtClassLoader实现。
- 派生于CLassLoader类。
- 父类加载器为启动类加载器。
- 从java.ext.dirs系统属性所指定的目录中加载类库，或从JDK的安装目录的jre/1ib/ext子目录（扩展目录）下加载类库。如果用户创建的JAR放在此目录下，也会自动由扩展类加载器加载

#### 4.4 应用程序类加载器（系统类加载器，AppClassLoader)

- javl语言编写，由sun.misc.LaunchersAppClassLoader实现。
- 派生于ClassLoader类。
- 父类加载器为扩张类加载器。
- 它负责加载环境变量classpath或系统属性java.class.path指定路径下的类库。
- 该类加载是程序中默认的类加载器，一般来说，Java应用的类都是由它来完成加载的。
- 通过classLoader#getSystemclassLoader() 方法可以获取到该类加载器。

#### 4.5 用户自定义类加载器

在Java的日程应用程序开发中，类的家在几乎是由上述3重类加载器互相配合执行的，在必要时，我们还可以自定义类加载器，来定制类的加载方式。为什么要自定义加载器？

- 隔离加载类
- 修改类加载的方式
- 扩展加载源
- 防止源码泄漏

#### 4.6 用户自定义类加载器实现步骤：

- 开发人员可以通过继承抽象类java.lang.ClassLoader类的方式，实现自己的类加载器，以满足一些特殊的需要
- 在JDK1.2 之前，在自定义类加载器时，总会去集成ClassLoader类并重写LoadClass（）方法，从而实现自定义的类的加载类，但是在啊JDK1.2之后不再建议用户覆盖LoadCladd（）方法，而是建议把自定义的类加载逻辑写在findclass() 方法中。
- 在编写自定义类加载器时，如果没有太过于复杂的需求，可以直接集成URIClassLoader类，这样就可以避免自己去编写findclass() 方法获取字节码流的方式，是自定义类加载器编写更加简洁。

#### 4.7 查看根加载器所能加载的目录

刚刚我们通过概念了解到了，根加载器只能够加载java/lib目录下的class，我们通过下面代码验证一下

```java
public class ClassLoaderTest1 {
    public static void main(String[] args) {
        System.out.println("*********启动类加载器************");
        // 获取BootstrapClassLoader 能够加载的API的路径
        URL[] urls = sun.misc.Launcher.getBootstrapClassPath().getURLs();
        for (URL url : urls) {
            System.out.println(url.toExternalForm());
        }

        // 从上面路径中，随意选择一个类，来看看他的类加载器是什么：得到的是null，说明是  根加载器
        ClassLoader classLoader = Provider.class.getClassLoader();
    }
}
```

#### 4.8 关于ClassLoader

ClassLoader类，他是一个抽象类，其后所有的类加载器都继承自ClassLoader（不包括启动类加载器）

![截屏2021-12-04 下午12.31.54](images/54.png)

sun.misc.Launcher 它是一个java虚拟机的入口应用

![截屏2021-12-04 下午12.34.01](images/01.png)

获取ClassLoader的途径

- 获取当前ClassLoader：clazz.getClassLoader()
- 获取当前线程上下文的ClassLoader：Thread.currentThread().getContextClassLoader()
- 获取系统的ClassLoader：ClassLoader.getSystemClassLoader()
- 获取调用者的ClassLoader：DriverManager.getCallerClassLoader()

### 五、双亲委派机制

​	Java虚拟机对class文件采用的是按需加载的方式，也就是说当需要使用该类时才会将它的class文件加载到内存生成class对象。而且加载某个类的class文件时，Java虚拟机采用的是双亲委派模式，即把请求交由父类处理，它是一种任务委派模式。

#### 5.1 工作原理

- 如果一个类加载器收到了类加载请求，它并不会自己先去加载，而是把这个请求委托给父类的加载器去执行；
- 如果父类加载器还存在其父类加载器，则进一步向上委托，依次递归，请求最终将到达顶层的启动类加载器；
- 如果父类加载器可以完成类加载任务，就成功返回，倘若父类加载器无法完成此加载任务，子加载器才会尝试自己去加载，这就是双亲委派模式。

![截屏2021-12-04 下午12.37.27](images/27.png)

#### 5.2  双亲委派机制举例

当我们加载jdbc.jar 用于实现数据库连接的时候，首先我们需要知道的是 jdbc.jar是基于SPI接口进行实现的，所以在加载的时候，会进行双亲委派，最终从根加载器中加载 SPI核心类，然后在加载SPI接口类，接着在进行反向委派，通过线程上下文类加载器进行实现类 jdbc.jar的加载。

![截屏2021-12-04 下午12.38.25](images/25.png)



#### 5.3 沙箱安全机制

自定义string类，但是在加载自定义String类的时候会率先使用引导类加载器加载，而引导类加载器在加载的过程中会先加载jdk自带的文件（rt.jar包中java\lang\String.class），报错信息说没有main方法，就是因为加载的是rt.jar包中的string类。这样可以保证对java核心源代码的保护，这就是沙箱安全机制。

#### 5.4 双亲委派机制的优势

通过上面的例子，我们可以知道，双亲机制可以

- 避免类的重复加载
- 保护程序安全，防止核心API被随意篡改
  - 自定义类：java.lang.String
  - 自定义类：java.lang.ShkStart（报错：阻止创建 java.lang开头的类）

### 六、其它

#### 6.1 如何判断两个class对象是否相同

在JVM中表示两个class对象是否为同一个类存在两个必要条件：

- 类的完整类名必须一致，包括包名。
- 加载这个类的ClassLoader（指ClassLoader实例对象）必须相同。

换句话说，在JvM中，即使这两个类对象（class对象）来源同一个Class文件，被同一个虚拟机所加载，但只要加载它们的ClassLoader实例对象不同，那么这两个类对象也是不相等的。

JVM必须知道一个类型是由启动加载器加载的还是由用户类加载器加载的。如果一个类型是由用户类加载器加载的，那么JVM会将这个类加载器的一个引用作为类型信息的一部分保存在方法区中。当解析一个类型到另一个类型的引用的时候，JVM需要保证这两个类型的类加载器是相同的。

#### 6.2 类的主动使用和被动使用

Java程序对类的使用方式分为：王动使用和被动使用。 主动使用，又分为七种情况：

- 创建类的实例
- 访问某个类或接口的静态变量，或者对该静态变量赋值
- 调用类的静态方法I
- 反射（比如：Class.forName（"com.atguigu.Test"））
- 初始化一个类的子类
- Java虚拟机启动时被标明为启动类的类
- JDK7开始提供的动态语言支持：
- java.lang.invoke.MethodHandle实例的解析结果REF getStatic、REF putStatic、REF invokeStatic句柄对应的类没有初始化，则初始化

除了以上七种情况，其他使用Java类的方式都被看作是对类的被动使用，都不会导致类的初始化。















































































