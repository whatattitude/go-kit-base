
[英文](README.md) | 中文

## 📖 简介

一个服务后端基础框架，包括线上服务的基础能力（日志切割清理、服务热加载）

## 🚀 功能：

- 日志清理
- 配置热加载
- 热编译
- 服务消耗监控

## 目录结构

- control 启动控制文件
- router 路由跳转注册表
- conf 测试配置目录
- conf_online  线上配置目录
- log 日志

##  使用

### 命令行启动

### 配置文件说明

### 数据可视化
- vm-ui link:
``` go
http://localhost:8428/vmui/?#/?g0.expr=%7B__name__%21%3D%22%22%2Ctype%3D%22now%22%7D&g0.range_input=6h&g0.end_input=2023-02-28T09%3A19%3A46&g0.relative_time=last_6_hours
```
