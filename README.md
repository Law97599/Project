# 选排课系统

本项目基于[Singo](https://github.com/Gourouting/singo)开发 

## 项目模块

系统包括以下模块
- 登录模块
  - 考察登录的设计与实现，对 HTTP 协议的理解。
    - 账密登录
    - Cookie Session
- 成员模块
  - 考察工程实现能力。
    - CURD 及对数据库的操作
    - 参数校验
      - 参数长度
      - 弱密码校验
    - 权限判断
- 排课模块
  - 主要考察算法（二分图匹配）的实现。
- 抢课模块
  - 主要考察简单秒杀场景的设计。

## 项目结构
```
.
├── api # MVC框架的controller，负责协调各部件完成任务
│   ├── auth.go
│   ├── course.go
│   ├── member.go
│   └── student.go
│
├── cache # redis缓存相关的代码
│   └── redis.go
│
├── common # 通用工具、错误状态码、常量等
│   ├── constants
│   │   └── constants.go
│   └── util
│       └── logger.go
│
├── conf # 配置文件
│   └── conf.go
│
├── db # 数据库初始化文件
│   └── mysql.sql
│
├── middleware # gin相关中间件
│   ├── auth.go
│   ├── cors.go
│   └── session.go
│
├── model # 数据库表实体
│   ├── course.go
│   ├── init.go
│   ├── member.go
│   ├── record.go
│   └── student_course.go
├── server
│   └── router.go
│
├── service # MVC框架的Service层，负责处理业务逻辑
│   ├── auth_service.go
│   ├── course_service.go
│   ├── member_service.go
│   └── student_service.go
│
├── vo # 页面输入模型
│   └── types.go
│
├── go.mod
├── go.sum
├── main.go
├── Dockerfile
└── README.md
```

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## 运行
```
go run main.go
```
项目运行后启动在3000端口