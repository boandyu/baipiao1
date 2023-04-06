### 一、mac中的IDEA的使用快捷键

1. command+F    在当前文件进行文本查找

2. command+shift+F  进行工程和模块中的文件搜索

3. command+u   找到这个方法的接口

4. ommand+option+commad   找到这个接口的实现类

5. command+/    对代码进行注释，并且自动移动到下一行

6. option+command+L   进行格式化代码

7. command +shift+R   进行整个项目或者指定目录文件进行替换

### 二、Editing(编辑)

command + J 快速查看文档 （按F1也可以）

shift + F1 快速查看外部文档

command + N 生成get、set方法

control + O 重写父类方法

control + I 实现接口方法

command + option + T 包围代码

command + option + / 块注释

option + 向上 选中代码块，向下取消

option + enter 显示意向动作



control + option + I 自动缩进线

command + option + L 格式化代码

command + option + O 优化import

command + shift + V 从最近的缓存区选择粘贴

command + D 复制当前行或选定的块

command + delete 删除当前行或选定的块

shift + enter 开始新的一行

command + shift + U 大小写切换

command + shift + [ /command + shift + ] 选择代码块开始/结束

option + fn + delete 删除到单词末尾

option + delete 删除到单词开始

command + 加号/command + 减号 展开/折叠代码块

command + shift + 加号 展开所有代码块

command + shift + 减号 折叠所有代码块

command + W 关闭活动的编辑选项卡

### 三、查询/替换（search/replace）

double shift 查询任何东西

command + G 向下查找

command + shift + G 向上查找

command + R 文件内替换

command + shift + F 全局查找（根据路径）

command + shift + R 全局替换（根据路径）

### 四、编译和运行（compile and run）

command + F9 编译project

control + option + R 弹出run的可选菜单

control + option + D 弹出debug可选菜单

control + R 运行

control + D 调试

### 五、使用查询（usage search）

option + F7/command + F7 在文件中查找用法/在类中查找用法

command + option + F7 显示用法

### 六、debug调试

F8 进入下一步，不进入方法

F7 进入下一步，进入方法，不进入嵌套方法

shift + F7 智能步入，断点运行的行上如果调用多个行，会弹出进入哪个方法

shift + F8 跳出

option + F9 运行到光标出，如果在光标前面还有断点，则进入到断点

option + F8 计算表达式（可以改变变量值，使其生效）

command + option + R 恢复断点运行，进入到下一个断点（如果还有）

command + F8 切换断点（若光标当前行有断点则取消断点，没有则加上断点）

command + shift + F8 查看断点信息

### 七、Navigation(导航)

command + O 查找类文件

command + shift + O 查找所有类型文件、打开文件、打开目录，打开目录需要在输入的内容前面加上一个反斜杠

command + option + O 前往指定的变量/方法

command + L 在当前文件跳转到指定行位置

command + E 显示最近打开的文件记录

option + 方向键 光标跳转到当前语句的首位或末尾

command + shift + 方向键 退回/前进到上一个操作的地方

command + shift + delete 跳转到最后一个编辑地方

command + Y 快速打开光标所在的方法、定义

control + shift + B 跳转到类型定义处

command + U 跳转到光标所在的方法所在父类的方法/接口定义

control + 方向键 上一个方法/下一个方法

command + F12 在类中找方法

control + H 显示当前类的结构层次

command + shift + H 显示方法的结构层次

control + option + H 显示调用层次结构

F2 跳转到下一个警告或错误处

### 八、Refactoring（重构）

F5 复制文件到指定目录

F6 移动文件到指定目录

Command + Delete 在文件上为安全删除文件，弹出确认框

Shift + F6 重命名文件

Command + F6 更改签名

Command + Option + N 一致性

Command + Option + M 将选中的代码提取为方法

Command + Option + V 提取变量

Command + Option + F 提取字段

Command + Option + C 提取常量

Command + Option + P 提取参数
