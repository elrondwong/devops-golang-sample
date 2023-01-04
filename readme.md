[TOC]

# zeroops-example

该模版演示基于 GoLang + Gin 实现全自动检出代码 -> 单元测试 -> 构建 Docker 镜像 -> 推送到 Docker 制品库 -> 部署到远端服务器

## 文件解释

样例包括:

- README.md - 本文件。项目概述及一些说明
- Dockerfile - 用以自动构建 Docker 镜像的脚本
- go.mod - 依赖包文件
- main.go - golang web 程序入口文件
- hello_test.go - go 测试文件

## 快速开始

如下这些引导，假定你想在自己的本地开发本项目。

1. 启动服务器

```bash
$ go run main.go
```

1. 打开 <http://127.0.0.1:80/> .

## 更多操作

### 执行测试

```bash
go test -v
```

### 构建二进制文件

```bash
go build main.go
```
