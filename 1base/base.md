# go基础

## 介绍

go是一门越来越流行的静态类型的编译型语言，简洁，清晰，高效著称。go代码编译成机器码非常迅速，同时还具有强大的垃圾收集与运行时反射机制，go语言原生就支持并发编程，有自己的并发编程模型，被称为互联网时代的c语言。
在中国知名互联网公司的核心系统越来越多的迁移到go平台上，很多创业公司从开始就选go语言为团队主力开发语言，并且关键的是go语言的平均工资普遍较高，所以技能之外掌握一门go语言绝对是益处多多。

## 主要特点

1. 自动垃圾回收
2. 更丰富的内置类型
3. 函数多返回值
4. 错误处理
5. 匿名函数和闭包
6. 类型和接口
7. 并发编程
8. 反射
9. 语言交互性

## 安装与设置

科学上网到go语言官网 地址是：https://golang.org/dl/ 根据自己的电脑系统选择对应的go安装包，下载完成，进行一键安装即可（我安装的是最新版go1.12.7）。在命令行 输入
```
go version
```
可以查看安装的go安装包版本信息
ps:通过源码包安装的需要自行设置系统环境变量，推荐使用官网的一键安装包安装。

## go目录结构
 src目录：包含go的源文件，每个目录是一个包
 pkg目录：包含包对象
 bin目录：包含可执行命令
 
 go语言工具构建src目录下的源码包，将生产的二进制文件安装到pkg与bin目录中，src子目录通常存放版本控制的各个项目源代码
 
 举个例子：
 
 ```
 bin/
	streak                         # 可执行命令
	todo                           # 可执行命令
pkg/
	linux_amd64/
		code.google.com/p/goauth2/
			oauth.a                # 包对象
		github.com/nf/todo/
			task.a                 # 包对象
src/
	code.google.com/p/goauth2/
		.hg/                       # mercurial 代码库元数据
		oauth/
			oauth.go               # 包源码
			oauth_test.go          # 测试源码
	github.com/nf/
		streak/
		.git/                      # git 代码库元数据
			oauth.go               # 命令源码
			streak.go              # 命令源码
		todo/
		.git/                      # git 代码库元数据
			task/
				task.go            # 包源码
			todo.go                # 命令源码
 ```
 此工作空间包含三个代码库（goauth2、streak 和 todo），两个命令（streak 和 todo） 以及两个库（oauth 和 task）。

## 变量声明

x int，y int 可以被缩写成 x,y int

 ## 包
 
 1. 单独引入： import "fmt"
 2. 批量引入 
 ```
 import (
        "fmt"
        "time"
        "math/rand"
      m "math"    //别名
 )
 ```
 注意：每个go文件引入的包必须被使用，否则编译报错
 
 # 方法是否公开
 go中，如果一个函数用大写字母开头，那么这个函数就是可以导出的，可以被别的go程序引用，如果是小写开头，则代表是包私有的，是不能导出的