# python 规范
***
## 目录
* [`命名规范`](#命名规范)
  * [`文件名规范`](#文件名规范)
  * [`类名规范`](#类名规范)
  * [`方法名规范`](#方法名规范)
  * [`变量名规范`](#变量名规范)
  * [`常量名规范`](#常量名规范)

***
### 命名规范
<a name= "命名规范" />
所有的命名的第一准则：见名知意

#### 文件名规范
<a name="文件名规范" />
小写(字母、数字),两个单词之间使用`_`连接

```
eg.
* hello.py
* hello_test.py
* hell1_1.py
```

#### 类名规范
<a name="类名规范" />
以大写字母开头，两个单词之间使用大写隔开(驼峰式)

```
eg.
* class Print(object):
* class MyPrint(object):
* class Print1(object):
```

#### 方法名规范
<a name="方法名规范" /〉
公有的是以小写字母开头，如果是多个单词链接的使用`_`分割开来
是有的是以`_`开头，若是多个单词以`_`分割开来

```
eg.
* def demo()
* def demo_test()
* def _demo()
* def _demo_test()
```

#### 变量名规范
<a name="变量名规范" />
全部采用小写，以`_`分割

```
eg.
* count = 0
* index = 1
* index_of_text = 1
```

### 常量名规范
〈a name="常量名规范" /〉
全部采用大学，使用`_`分割单词

```
eg.
* INDEX = "Index"
* INDEX_OF_TEXT = "IndexOfText"
```
