# GqGo Engine Tools

## This project is a fork from Go Zero. It contains lots of features for GqGo Engine.
## 这是一个基于 Go Zero 跨境的云生态项目，提供了很多针对 GqGo Engine 的优化方案

## 特性

- **最新技术栈**：使用 ent, casbin, kafka 等前沿技术开发
- **完全支持go-swagger**: 直接在api文件内编写注释即可直接生成swagger文档
- **统一的错误处理**: 整个系统拥有国际化的统一错误处理
- **国际化**：内置完善的国际化方案
- **服务注册发现**: 完善的服务注册发现机制，原生支持K8s
- **权限**: 内置完善的动态路由权限生成方案, 集成RBAC权限控制
- **代码生成**: 内置三端 Web, API, RPC 代码生成
- **多种扩展**: 提供多种扩展，同时具有非常简单的接入功能
- **其他**: 流量控制， ES服务
- 引入Template 模版
- go-swagger : 基于go-swagger而不是官方的@doc注解
- 多国语言支持
- 优化错误信息处理,支持多语言错误
- 简单易用的校验器
- 支持代码生成，生成API,RPC 和 web 端的CRUD代码
- 支持多种额外插件如GORM, RocketMQ
- 对 GqGo Engine 的针对性优化
- rpc logic group分组
- GqGo Cloud 生态

## 支持功能

- 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
- 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
- 岗位管理：配置系统用户所属担任职务。
- 菜单管理：配置系统菜单，操作权限，按钮权限标识，接口权限等。
- 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
- 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
- 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
- 会员管理：管理注册会员信息
- 接口文档：根据业务代码自动生成相关的api接口文档。
- 代码生成：根据数据表结构生成对应的增删改查相对应业务
- 服务监控：查看一些服务器的基本信息

## 准备

- [Golang](http://go.dev/) and [git](https://git-scm.com/) - Go 语言
- [Ent](https://entgo.io/docs/getting-started) - Ent
- [Mysql](https://www.mysql.com/) - Mysql数据库
- [GORM](https://gorm.io/) - GORM 数据库ORM组件
- [Casbin](https://casbin.org/) - 权限管理
- [Go-swagger](https://goswagger.io/) - Go-swagger 文档生成调试

## 快速开始

[快速开始文档](https://doc.ryansu.pro/zh/guide/basic-config/env_setting.html)

## 更新日志

[CHANGELOG](./CHANGELOG.md)

## 相关项目

- [Simple Admin](https://github.com/gmiddlecloud/gqgo-engine-core)
- [Simple Admin 后端界面](https://github.com/gmiddlecloud/gqgo-engine-manage)

## 可选组件

- [文件管理](https://github.com/gmiddlecloud/gqgo-engine-file)
- [定时任务](https://github.com/gmiddlecloud/gqgo-engine-job)
- [会员管理](https://github.com/gmiddlecloud/gqgo-engine-member-api)
